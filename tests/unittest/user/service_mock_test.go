package user_test

import (
	"oncomapi/internal/api/v1/user"
	"oncomapi/internal/api/v1/user/dto"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Register(req *dto.RegisterRequest) (*user.User, error) {
	args := m.Called(req)
	if args.Get(0) != nil {
		return args.Get(0).(*user.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserService) Login(req *dto.LoginRequest) (string, error) {
	args := m.Called(req)
	if args.Get(0) != nil {
		return args.String(0), args.Error(1)
	}
	return "", args.Error(1)
}
