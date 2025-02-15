package dtos

import "time"

type CreatePostRequestDTO struct {
	Title  string `json:"title" validate:"required,min=1,max=100"`
	Body   string `json:"body" validate:"required,min=1,max=500"`
	UserID string `json:"userId" validate:"required"`
}

type CreatePostResponseDTO struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}
