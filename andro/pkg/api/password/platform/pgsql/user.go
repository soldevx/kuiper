package pgsql

import (
	"github.com/go-pg/pg/v9/orm"

	andro "github.com/soldevx/kuiper/andro/pkg/utl/model"
)

// User represents the client for user table
type User struct{}

// View returns single user by ID
func (u User) View(db orm.DB, id int) (andro.User, error) {
	user := andro.User{Base: andro.Base{ID: id}}
	err := db.Select(&user)
	return user, err
}

// Update updates user's info
func (u User) Update(db orm.DB, user andro.User) error {
	return db.Update(&user)
}
