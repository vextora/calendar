package accesscontrol

type AccessRequest struct {
	Subject  string
	Object   string
	Action   string
	Resource string
	Context  map[string]any
}

type Evaluator interface {
	Evaluate(req AccessRequest) bool
}
