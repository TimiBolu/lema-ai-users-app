package dtos

type GetUsersRequestDTO struct {
	PageNumber int `query:"pageNumber" validate:"required,min=1"`
	PageSize   int `query:"pageSize" validate:"required,min=1,max=20"`
}

type Pagination struct {
	CurrentPage int   `json:"currentPage"`
	PageSize    int   `json:"pageSize"`
	TotalPages  int   `json:"totalPages"`
	TotalItems  int64 `json:"totalItems"`
	HasNext     bool  `json:"hasNext"`
	HasPrev     bool  `json:"hasPrev"`
}

type GetUsersResponseDTO struct {
	Users      []GetUserResponseDTO `json:"users"`
	Pagination Pagination           `json:"pagination"`
}
