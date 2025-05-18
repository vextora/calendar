package dto

type UserResponse struct {
	ID       uint                `json:"id"`
	Username string              `json:"username"`
	Email    string              `json:"email"`
	Password string              `json:"password"`
	Detail   *UserDetailResponse `json:"detail,omitempty"`
}

type UserDetailResponse struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}
