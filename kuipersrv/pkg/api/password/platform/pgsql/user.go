package pgsql

import (
	"github.com/go-pg/pg/v9/orm"

	kuiper "github.com/soldevx/kuiper/kuipersrv"
)

// User represents the client for user table
type User struct{}

// View returns single user by ID
func (u User) View(db orm.DB, id int) (kuiper.User, error) {
	user := kuiper.User{Base: kuiper.Base{ID: id}}
	err := db.Select(&user)
	return user, err
}

// Update updates user's info
func (u User) Update(db orm.DB, user kuiper.User) error {
	return db.Update(&user)
}