package dto

type CreateRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required,minWords=3"`
}

type UpdateRequest struct {
	ID      uint   `json:"id"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required,minWords=3"`
}
