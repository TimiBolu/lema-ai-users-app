package dtos

type GetUserRequestDTO struct {
	ID string `json:"id" validate:"required,uuid"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}

type GetUserResponseDTO struct {
	ID        string  `json:"id"`
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	Email     string  `json:"email"`
	Address   Address `json:"address"`
}
