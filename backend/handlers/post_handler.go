package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/TimiBolu/lema-ai-users-service/dtos"
	"github.com/TimiBolu/lema-ai-users-service/services"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PostHandler struct {
	service  services.PostService
	logger   *logrus.Logger
	validate *validator.Validate
}

func NewPostHandler(service services.PostService, logger *logrus.Logger) *PostHandler {
	return &PostHandler{
		service:  service,
		logger:   logger,
		validate: validator.New(),
	}
}

func (h *PostHandler) GetPostsByUser(w http.ResponseWriter, r *http.Request) {
	var reqDTO dtos.GetPostsByUserRequestDTO
	reqDTO.UserId = r.URL.Query().Get("userId")

	if err := h.validate.Struct(reqDTO); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid request data", h.logger, err)
		return
	}

	posts, err := h.service.GetPostsByUser(r.Context(), reqDTO.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			SendErrorResponse(w, http.StatusNotFound, "Posts not found", h.logger, err)
		} else {
			SendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve posts", h.logger, err)
		}
		return
	}

	respDTO := dtos.GetPostsByUserResponseDTO{
		Posts: posts,
	}

	SendSuccessResponse(w, http.StatusOK, "posts fetched successfully", respDTO)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var reqDTO dtos.CreatePostRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&reqDTO); err != nil {
		if err.Error() == "http: request body too large" {
			SendErrorResponse(w, http.StatusRequestEntityTooLarge, "Request body size exceeded", h.logger, err)
		} else {
			SendErrorResponse(w, http.StatusBadRequest, "Invalid request body", h.logger, err)
		}
		return
	}

	if err := h.validate.Struct(reqDTO); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid request data", h.logger, err)
		return
	}

	reqDTO.Title = strings.TrimSpace(reqDTO.Title)
	reqDTO.Body = strings.TrimSpace(reqDTO.Body)

	newPost, err := h.service.CreatePost(r.Context(), reqDTO.Title, reqDTO.Body, reqDTO.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			SendErrorResponse(w, http.StatusNotFound, "User does not exist", h.logger, err)
		} else {
			SendErrorResponse(w, http.StatusInternalServerError, "Failed to create post", h.logger, err)
		}
		return
	}

	respDTO := dtos.CreatePostResponseDTO{
		ID:        newPost.ID,
		Title:     newPost.Title,
		Body:      newPost.Body,
		UserID:    newPost.UserID,
		CreatedAt: newPost.CreatedAt,
	}

	SendSuccessResponse(w, http.StatusCreated, "post created successfully", respDTO)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	var reqDTO dtos.DeletePostRequestDTO
	vars := mux.Vars(r)
	reqDTO.ID = vars["id"]

	if err := h.validate.Struct(reqDTO); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid request data", h.logger, err)
		return
	}

	if err := h.service.DeletePost(r.Context(), reqDTO.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			SendErrorResponse(w, http.StatusNotFound, "Post not found", h.logger, err)
		} else {
			SendErrorResponse(w, http.StatusInternalServerError, "Failed to delete post", h.logger, err)
		}
		return
	}

	SendSuccessResponse[any](w, http.StatusNoContent, "post deleted successfully", nil)
}
