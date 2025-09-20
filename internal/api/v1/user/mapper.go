package user

import "calendarapi/internal/api/v1/user/dto"

func DetailResponse(data *User) dto.UserResponse {
	res := dto.UserResponse{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
	}

	if data.Detail.ID != 0 {
		res.Detail = &dto.UserDetailResponse{
			FullName: data.Detail.FullName,
			Phone:    data.Detail.Phone,
			Address:  data.Detail.Address,
		}
	}

	return res
}
