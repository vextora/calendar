package domain

type Repository interface {
	Save(user *User) (*User, error)
	FindByID(id int) (*User, error)
	FindByEmail(email string) (*User, error)

	CreateDetail(detail *UserDetail) (*UserDetail, error)
}
