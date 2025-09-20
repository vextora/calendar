package {{.EntityLower}}

import "calendarapi/internal/{{.ApiPrefix}}/{{.Version}}/{{.EntityLower}}/domain"

type (
	{{.Entity}} = domain.{{.Entity}}
	Repository = domain.Repository
	Service    = domain.Service
)
