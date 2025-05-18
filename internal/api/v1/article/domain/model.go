package domain

type Article struct {
	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID  int    `json:"user_id"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required,minWord=3"`
	Slug    string `json:"slug"`
}
