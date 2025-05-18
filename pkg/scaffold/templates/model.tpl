package domain

type {{.Entity}} struct {
	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement"`
}
