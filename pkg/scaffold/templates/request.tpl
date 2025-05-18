package dto

type {{.Entity}}Request struct {
	ID   uint `json:"id" validate:"required"`
}