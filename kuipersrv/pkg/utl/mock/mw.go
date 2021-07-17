package mock

import (
	kuiper "github.com/soldevx/kuiper/kuipersrv"
)

// JWT mock
type JWT struct {
	GenerateTokenFn func(kuiper.User) (string, error)
}

// GenerateToken mock
func (j JWT) GenerateToken(u kuiper.User) (string, error) {
	return j.GenerateTokenFn(u)
}
