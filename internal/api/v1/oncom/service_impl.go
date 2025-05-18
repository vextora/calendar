package oncom

type oncomService struct {
	repo Repository
}

func NewOncomService(repo Repository) Service {
	return &oncomService{repo: repo}
}
