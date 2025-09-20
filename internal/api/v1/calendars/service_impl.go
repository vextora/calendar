package calendars

type calendarsService struct {
	repo Repository
}

func NewCalendarsService(repo Repository) Service {
	return &calendarsService{repo: repo}
}
