package user

import (
	"calendarapi/internal/api/v1/user/dto"
	"calendarapi/pkg/hash"
	"calendarapi/pkg/jwt"
	"errors"
)

type userService struct {
	repo Repository
}

func NewUserService(repo Repository) Service {
	return &userService{repo: repo}
}

func (serv *userService) Register(input *dto.RegisterRequest) (*User, error) {
	existingUser, err := serv.repo.FindByEmail(input.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already in use")
	}

	hashedPassword, err := hash.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
	}

	createdUser, err := serv.repo.Save(user)
	if err != nil {
		return nil, err
	}

	detail := &UserDetail{
		UserID:   createdUser.ID,
		FullName: input.Detail.FullName,
		Phone:    input.Detail.Phone,
		Address:  input.Detail.Address,
	}

	createdDetail, err := serv.repo.CreateDetail(detail)
	if err != nil {
		return nil, err
	}

	createdUser.Detail = *createdDetail

	return createdUser, nil
}

func (serv *userService) Login(input *dto.LoginRequest) (string, error) {
	user, err := serv.repo.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("email not exist")
	}

	if !hash.CheckPassword(input.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := jwt.GenerateToken(int(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (serv *userService) FindByID(id int) (*User, error) {
	user, err := serv.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("email not exist")
	}

	return user, nil
}
