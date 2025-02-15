package dtos

type DeletePostRequestDTO struct {
	ID string `json:"id" validate:"required"`
}
