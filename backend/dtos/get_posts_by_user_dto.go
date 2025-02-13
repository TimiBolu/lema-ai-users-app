package dtos

import "github.com/TimiBolu/lema-ai-users-service/models"

type GetPostsByUserRequestDTO struct {
	UserId string `query:"userId" validate:"required,uuid"`
}

type GetPostsByUserResponseDTO struct {
	Posts []models.Post `json:"posts"`
}
