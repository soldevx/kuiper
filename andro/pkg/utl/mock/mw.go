package mock

import (
	andro "github.com/soldevx/kuiper/andro/pkg/utl/model"
)

// JWT mock
type JWT struct {
	GenerateTokenFn func(andro.User) (string, error)
}

// GenerateToken mock
func (j JWT) GenerateToken(u andro.User) (string, error) {
	return j.GenerateTokenFn(u)
}
