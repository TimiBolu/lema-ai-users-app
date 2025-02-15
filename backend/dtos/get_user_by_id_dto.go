package dtos

type GetUserRequestDTO struct {
	ID string `json:"id" validate:"required"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipcode"`
}

type GetUserResponseDTO struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Address  Address `json:"address"`
}
