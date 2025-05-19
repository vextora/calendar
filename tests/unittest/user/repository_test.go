package user_test

import (
	"errors"
	"oncomapi/internal/api/v1/user/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserRepository(t *testing.T) {

	testFindID := "FindByID"
	t.Run(testFindID, func(t *testing.T) {
		tests := []struct {
			name       string
			inputID    int
			setupMock  func(*MockUserRepo)
			expectUser *domain.User
			expectErr  bool
		}{
			{
				name:    "Success",
				inputID: 1,
				setupMock: func(m *MockUserRepo) {
					m.On(testFindID, 1).Return(&domain.User{ID: 1, Username: "oncom"}, nil)
				},
				expectUser: &domain.User{ID: 1, Username: "oncom"},
				expectErr:  false,
			},
			{
				name:    "UserNotFound",
				inputID: 2,
				setupMock: func(m *MockUserRepo) {
					m.On(testFindID, 2).Return((*domain.User)(nil), errors.New("not found"))
				},
				expectUser: nil,
				expectErr:  true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockRepo := &MockUserRepo{}
				tt.setupMock(mockRepo)

				u, err := mockRepo.FindByID(tt.inputID)
				if tt.expectErr {
					assert.Error(t, err)
					assert.Nil(t, u)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.expectUser, u)
				}
				mockRepo.AssertExpectations(t)
			})
		}
	})

	testFindEmail := "FindByEmail"
	t.Run(testFindEmail, func(t *testing.T) {
		tests := []struct {
			name        string
			email       string
			setupMock   func(*MockUserRepo)
			expected    *domain.User
			expectError bool
		}{
			{
				name:  "Success",
				email: "oncom@mail.com",
				setupMock: func(m *MockUserRepo) {
					m.On(testFindEmail, "oncom@mail.com").Return(&domain.User{Email: "oncom@mail.com"}, nil)
				},
				expected:    &domain.User{Email: "oncom@mail.com"},
				expectError: false,
			},
			{
				name:  "NotFound",
				email: "x@mail.com",
				setupMock: func(m *MockUserRepo) {
					m.On(testFindEmail, "x@mail.com").Return((*domain.User)(nil), errors.New("not found"))
				},
				expected:    nil,
				expectError: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockRepo := &MockUserRepo{}
				tt.setupMock(mockRepo)

				res, err := mockRepo.FindByEmail(tt.email)
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
	})

	testSave := "Save"
	t.Run(testSave, func(t *testing.T) {
		tests := []struct {
			name        string
			input       *domain.User
			setupMock   func(*MockUserRepo)
			expected    *domain.User
			expectError bool
		}{
			{
				name:  "Success",
				input: &domain.User{Username: "newuser", Email: "new@example.com", Password: "password123"},
				setupMock: func(m *MockUserRepo) {
					m.On(testSave, mock.AnythingOfType("*domain.User")).
						Return(&domain.User{ID: 1, Username: "oncom", Email: "oncom@mail.com"}, nil)
				},
				expected:    &domain.User{ID: 1, Username: "oncom", Email: "oncom@mail.com"},
				expectError: false,
			},
			{
				name:  "FailedToSave",
				input: &domain.User{Username: "", Email: "", Password: ""},
				setupMock: func(m *MockUserRepo) {
					m.On(testSave, mock.AnythingOfType("*domain.User")).Return(nil, errors.New("failed to save"))
				},
				expected:    nil,
				expectError: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockRepo := &MockUserRepo{}
				tt.setupMock(mockRepo)

				res, err := mockRepo.Save(tt.input)
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
	})

	testCreateDetail := "CreateDetail"
	t.Run(testCreateDetail, func(t *testing.T) {
		tests := []struct {
			name        string
			input       *domain.UserDetail
			setupMock   func(*MockUserRepo)
			expected    *domain.UserDetail
			expectError bool
		}{
			{
				name:  "Success",
				input: &domain.UserDetail{UserID: 1, FullName: "John Doe", Phone: "08123456789", Address: "Jakarta"},
				setupMock: func(m *MockUserRepo) {
					m.On(testCreateDetail, mock.AnythingOfType("*domain.UserDetail")).
						Return(&domain.UserDetail{ID: 1, FullName: "John Doe", Phone: "08123456789", Address: "Jakarta"}, nil)
				},
				expected:    &domain.UserDetail{ID: 1, FullName: "John Doe", Phone: "08123456789", Address: "Jakarta"},
				expectError: false,
			},
			{
				name:  "FailedToCreate",
				input: &domain.UserDetail{},
				setupMock: func(m *MockUserRepo) {
					m.On(testCreateDetail, mock.AnythingOfType("*domain.UserDetail")).Return(nil, errors.New("failed to create"))
				},
				expected:    nil,
				expectError: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockRepo := &MockUserRepo{}
				tt.setupMock(mockRepo)

				res, err := mockRepo.CreateDetail(tt.input)
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
	})
}
