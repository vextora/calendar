package shared

type TokenValidator interface {
	Validate(tokenString string) (any, error)
}
