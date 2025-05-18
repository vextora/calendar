package {{.EntityLower}}

import "oncomapi/internal/{{.ApiPrefix}}/{{.Version}}/{{.EntityLower}}/dto"

func {{.Entity}}Response(data *{{.Entity}}) dto.{{.Entity}}Response {
	return dto.{{.Entity}}Response{
		ID:      data.ID,
	}
}
