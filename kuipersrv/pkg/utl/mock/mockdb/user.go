package mockdb

import (
	"github.com/go-pg/pg/v9/orm"

	kuiper "github.com/soldevx/kuiper/kuipersrv"
)

// User database mock
type User struct {
	CreateFn         func(orm.DB, kuiper.User) (kuiper.User, error)
	ViewFn           func(orm.DB, int) (kuiper.User, error)
	FindByUsernameFn func(orm.DB, string) (kuiper.User, error)
	FindByTokenFn    func(orm.DB, string) (kuiper.User, error)
	ListFn           func(orm.DB, *kuiper.ListQuery, kuiper.Pagination) ([]kuiper.User, error)
	DeleteFn         func(orm.DB, kuiper.User) error
	UpdateFn         func(orm.DB, kuiper.User) error
}

// Create mock
func (u *User) Create(db orm.DB, usr kuiper.User) (kuiper.User, error) {
	return u.CreateFn(db, usr)
}

// View mock
func (u *User) View(db orm.DB, id int) (kuiper.User, error) {
	return u.ViewFn(db, id)
}

// FindByUsername mock
func (u *User) FindByUsername(db orm.DB, uname string) (kuiper.User, error) {
	return u.FindByUsernameFn(db, uname)
}

// FindByToken mock
func (u *User) FindByToken(db orm.DB, token string) (kuiper.User, error) {
	return u.FindByTokenFn(db, token)
}

// List mock
func (u *User) List(db orm.DB, lq *kuiper.ListQuery, p kuiper.Pagination) ([]kuiper.User, error) {
	return u.ListFn(db, lq, p)
}

// Delete mock
func (u *User) Delete(db orm.DB, usr kuiper.User) error {
	return u.DeleteFn(db, usr)
}

// Update mock
func (u *User) Update(db orm.DB, usr kuiper.User) error {
	return u.UpdateFn(db, usr)
}
