package oncom

import "oncomapi/internal/api/v1/oncom/dto"

func OncomResponse(data *Oncom) dto.OncomResponse {
	return dto.OncomResponse{
		ID:      data.ID,
	}
}
