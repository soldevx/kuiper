package mock

// JWT mock
type JWT struct {
	GenerateTokenFn func(andro.User) (string, error)
}

// GenerateToken mock
func (j JWT) GenerateToken(u andro.User) (string, error) {
	return j.GenerateTokenFn(u)
}
