package domain

type User struct {
	ID       uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Detail   UserDetail `gorm:"foreignKey:UserID"`
}

type UserDetail struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint
	FullName string
	Phone    string
	Address  string
}
