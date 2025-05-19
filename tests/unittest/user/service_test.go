package user_test

import (
	"errors"
	"oncomapi/internal/api/v1/user"
	"oncomapi/internal/api/v1/user/domain"
	"oncomapi/internal/api/v1/user/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_Register(t *testing.T) {

	tests := []struct {
		name        string
		input       *dto.RegisterRequest
		setupMock   func(*MockUserRepo)
		expected    *user.User
		expectError bool
	}{
		{
			name: "Success",
			input: &dto.RegisterRequest{
				Username: "newuser",
				Email:    "newuser@mail.com",
				Password: "password123",
				Detail:   &dto.UserDetailRequest{FullName: "New User", Phone: "08123456789", Address: "Jakarta"},
			},
			setupMock: func(m *MockUserRepo) {
				m.On("FindByEmail", "newuser@mail.com").Return((*domain.User)(nil), nil)
				m.On("Save", mock.AnythingOfType("*domain.User")).Return(&user.User{ID: 1, Username: "newuser", Email: "newuser@mail.com"}, nil)
				m.On("CreateDetail", mock.AnythingOfType("*domain.UserDetail")).Return(&user.UserDetail{ID: 1, FullName: "New User", Phone: "08123456789", Address: "Jakarta"}, nil)
			},
			expected:    &user.User{ID: 1, Username: "newuser", Email: "newuser@mail.com", Detail: user.UserDetail{ID: 1, FullName: "New User", Phone: "08123456789", Address: "Jakarta"}},
			expectError: false,
		},
		{
			name: "EmailAlreadyInUse",
			input: &dto.RegisterRequest{
				Username: "newuser",
				Email:    "usedemail@mail.com",
				Password: "password123",
			},
			setupMock: func(m *MockUserRepo) {
				m.On("FindByEmail", "usedemail@mail.com").Return(&domain.User{Email: "usedemail@mail.com"}, nil)
			},
			expected:    nil,
			expectError: true,
		},
		{
			name:  "FailedToSaveUser",
			input: &dto.RegisterRequest{Username: "newuser", Email: "newuser@mail.com", Password: "password123"},
			setupMock: func(m *MockUserRepo) {
				m.On("FindByEmail", "newuser@mail.com").Return((*domain.User)(nil), nil)
				m.On("Save", mock.AnythingOfType("*domain.User")).Return(nil, errors.New("failed to save"))
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "FailedToCreateDetail",
			input: &dto.RegisterRequest{
				Username: "newuser",
				Email:    "newuser@mail.com",
				Password: "password123",
				Detail:   &dto.UserDetailRequest{FullName: "New User", Phone: "08123456789", Address: "Jakarta"},
			},
			setupMock: func(m *MockUserRepo) {
				m.On("FindByEmail", "newuser@mail.com").Return((*domain.User)(nil), nil)
				m.On("Save", mock.AnythingOfType("*domain.User")).Return(&user.User{ID: 1, Username: "newuser", Email: "newuser@mail.com"}, nil)
				m.On("CreateDetail", mock.AnythingOfType("*domain.UserDetail")).Return(nil, errors.New("failed to create user detail"))
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockUserRepo{}
			tt.setupMock(mockRepo)

			service := user.NewUserService(mockRepo)
			res, err := service.Register(tt.input)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, res)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, res)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
