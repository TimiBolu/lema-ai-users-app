package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/TimiBolu/lema-ai-users-service/dtos"
	"github.com/TimiBolu/lema-ai-users-service/services"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
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
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to get posts by user ID", h.logger, err)
		return
	}

	respDTO := dtos.GetPostsByUserResponseDTO{
		Posts: posts,
	}

	SendSuccessResponse(w, http.StatusOK, "posts fetched successfully", respDTO)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	// Set request body size limit
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MB limit

	var reqDTO dtos.CreatePostRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&reqDTO); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid request body", h.logger, err)
		return
	}

	if err := h.validate.Struct(reqDTO); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Invalid request data", h.logger, err)
		return
	}

	// Sanitize the input fields
	reqDTO.Title = strings.TrimSpace(reqDTO.Title)
	reqDTO.Body = strings.TrimSpace(reqDTO.Body)

	newPost, err := h.service.CreatePost(r.Context(), reqDTO.Title, reqDTO.Body, reqDTO.UserID)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to create post", h.logger, err)
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
		SendErrorResponse(w, http.StatusInternalServerError, "Failed to delete post", h.logger, err)
		return
	}

	SendSuccessResponse[any](w, http.StatusNoContent, "post deleted successfully", nil)
}
