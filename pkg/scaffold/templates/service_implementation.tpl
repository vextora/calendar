package {{.EntityLower}}

type {{.EntityLower}}Service struct {
	repo Repository
}

func New{{.Entity}}Service(repo Repository) Service {
	return &{{.EntityLower}}Service{repo: repo}
}
