package user_test

import (
	"oncomapi/internal/api/v1/user/domain"

	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Save(user *domain.User) (*domain.User, error) {
	args := m.Called(user)

	var result *domain.User
	if args.Get(0) != nil {
		result = args.Get(0).(*domain.User)
	}

	return result, args.Error(1)
}

func (m *MockUserRepo) FindByID(id int) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepo) FindByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepo) CreateDetail(detail *domain.UserDetail) (*domain.UserDetail, error) {
	args := m.Called(detail)

	var result *domain.UserDetail
	if args.Get(0) != nil {
		result = args.Get(0).(*domain.UserDetail)
	}
	return result, args.Error(1)
}
