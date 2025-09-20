package domain

type Calendars struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name" validate:"required"`
	Color     string `json:"color"`
	IsDefault bool   `json:"is_default"`
}
