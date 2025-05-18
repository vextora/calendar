package user_test

/*func RouterUserTest(service user.Service) *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	api := r.Group("/api")
	handler := user.NewUserHandler(service)
	router.SetupUserRoutes(api, handler)
	return r
}

func TestUserHandler_Register(t *testing.T) {
	tests := []struct {
		name           string
		inputBody      *dto.RegisterRequest
		setupMock      func(*MockUserService)
		expectedCode   int
		expectedStatus string
	}{
		{
			name: "Success",
			inputBody: &dto.RegisterRequest{
				Username: "newuser",
				Email:    "new@mail.com",
				Password: "secret123",
				Detail: &dto.UserDetailRequest{
					FullName: "New User",
					Phone:    "08123456789",
					Address:  "Jakarta",
				},
			},
			setupMock: func(m *MockUserService) {
				m.On("Register", mock.AnythingOfType("*dto.RegisterRequest")).
					Return(&user.User{
						ID:       1,
						Username: "newuser",
						Email:    "new@mail.com",
						Detail: user.UserDetail{
							ID:       1,
							FullName: "New User",
							Phone:    "08123456789",
							Address:  "Jakarta",
						},
					}, nil)
			},
			expectedCode:   http.StatusOK,
			expectedStatus: "OK",
		},
		{
			name: "BadRequest",
			inputBody: &dto.RegisterRequest{
				Username: "erroruser",
				Email:    "error@mail.com",
				Password: "sec123",
				Detail: &dto.UserDetailRequest{
					FullName: "Error User",
					Phone:    "08123456789",
					Address:  "Jakarta",
				},
			},
			setupMock: func(m *MockUserService) {
				m.On("Register", mock.AnythingOfType("*dto.RegisterRequest")).Return(nil, errors.New("internal server error"))
			},
			expectedCode:   http.StatusBadRequest,
			expectedStatus: "Bad Request",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := new(MockUserService)
			if tt.setupMock != nil {
				tt.setupMock(mockSvc)
			}
			router := RouterUserTest(mockSvc)

			body, _ := json.Marshal(tt.inputBody)
			req, _ := http.NewRequest(http.MethodPost, "/api/users/register", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			var response map[string]interface{}
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.Equal(t, tt.expectedStatus, response["status"])

			mockSvc.AssertExpectations(t)
		})
	}
}
*/
