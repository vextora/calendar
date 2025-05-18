package {{.EntityLower}}

import "oncomapi/internal/{{.ApiPrefix}}/{{.Version}}/{{.EntityLower}}/domain"

type (
	{{.Entity}} = domain.{{.Entity}}
	Repository = domain.Repository
	Service    = domain.Service
)
