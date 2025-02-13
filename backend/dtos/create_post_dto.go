package dtos

import "time"

type CreatePostRequestDTO struct {
	Title  string `json:"title" validate:"required,min=1,max=255"`
	Body   string `json:"body" validate:"required,min=1,max=1000"`
	UserID string `json:"userId" validate:"required,uuid"`
}

type CreatePostResponseDTO struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}
