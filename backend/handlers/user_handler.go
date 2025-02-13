package handlers

import (
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/TimiBolu/lema-ai-users-service/dtos"
	"github.com/TimiBolu/lema-ai-users-service/services"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserHandler struct {
	service  services.UserService
	logger   *logrus.Logger
	validate *validator.Validate
}

func NewUserHandler(service services.UserService, logger *logrus.Logger) *UserHandler {
	return &UserHandler{
		service:  service,
		logger:   logger,
		validate: validator.New(),
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	pageNumber, err := strconv.Atoi(r.URL.Query().Get("pageNumber"))
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid pageNumber parameter", h.logger, err)
		return
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid pageSize parameter", h.logger, err)
		return
	}

	reqDTO := dtos.GetUsersRequestDTO{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}

	if err := h.validate.Struct(reqDTO); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid pagination parameters", h.logger, err)
		return
	}

	users, totalUsers, err := h.service.GetUsers(r.Context(), reqDTO.PageNumber, reqDTO.PageSize)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve users", h.logger, err)
		return
	}

	userDTOs := make([]dtos.GetUserResponseDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dtos.GetUserResponseDTO{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Address: dtos.Address{
				Street:  user.Address.Street,
				City:    user.Address.City,
				State:   user.Address.State,
				ZipCode: user.Address.ZipCode,
			},
		}
	}

	totalPages := int(math.Ceil(float64(totalUsers) / float64(reqDTO.PageSize)))
	respDTO := dtos.GetUsersResponseDTO{
		Users: userDTOs,
		Pagination: dtos.Pagination{
			CurrentPage: reqDTO.PageNumber,
			TotalItems:  totalUsers,
			PageSize:    reqDTO.PageSize,
			TotalPages:  totalPages,
			HasNext:     reqDTO.PageNumber < totalPages,
			HasPrev:     reqDTO.PageNumber > 1,
		},
	}

	SendSuccessResponse(w, http.StatusOK, "users fetched successfully", respDTO)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	var reqDTO dtos.GetUserRequestDTO
	reqDTO.ID = mux.Vars(r)["id"]

	if err := h.validate.Struct(reqDTO); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid request data", h.logger, err)
		return
	}

	user, err := h.service.GetUserByID(r.Context(), reqDTO.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			SendErrorResponse(w, http.StatusNotFound, "User not found", h.logger, err)
		} else {
			SendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve user", h.logger, err)
		}
		return
	}

	respDTO := dtos.GetUserResponseDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Address: dtos.Address{
			Street:  user.Address.Street,
			City:    user.Address.City,
			State:   user.Address.State,
			ZipCode: user.Address.ZipCode,
		},
	}
	SendSuccessResponse(w, http.StatusOK, "user fetched successfully", respDTO)
}

func (h *UserHandler) GetUsersCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.service.GetUserCount(r.Context())
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve user count", h.logger, err)
		return
	}

	respDTO := dtos.GetUserCountResponseDTO{
		Count: count,
	}
	SendSuccessResponse(w, http.StatusOK, "user count fetched successfully", respDTO)
}
