package mockdb

import (
	"github.com/go-pg/pg/v9/orm"
)

// User database mock
type User struct {
	CreateFn         func(orm.DB, andro.User) (andro.User, error)
	ViewFn           func(orm.DB, int) (andro.User, error)
	FindByUsernameFn func(orm.DB, string) (andro.User, error)
	FindByTokenFn    func(orm.DB, string) (andro.User, error)
	ListFn           func(orm.DB, *andro.ListQuery, andro.Pagination) ([]andro.User, error)
	DeleteFn         func(orm.DB, andro.User) error
	UpdateFn         func(orm.DB, andro.User) error
}

// Create mock
func (u *User) Create(db orm.DB, usr andro.User) (andro.User, error) {
	return u.CreateFn(db, usr)
}

// View mock
func (u *User) View(db orm.DB, id int) (andro.User, error) {
	return u.ViewFn(db, id)
}

// FindByUsername mock
func (u *User) FindByUsername(db orm.DB, uname string) (andro.User, error) {
	return u.FindByUsernameFn(db, uname)
}

// FindByToken mock
func (u *User) FindByToken(db orm.DB, token string) (andro.User, error) {
	return u.FindByTokenFn(db, token)
}

// List mock
func (u *User) List(db orm.DB, lq *andro.ListQuery, p andro.Pagination) ([]andro.User, error) {
	return u.ListFn(db, lq, p)
}

// Delete mock
func (u *User) Delete(db orm.DB, usr andro.User) error {
	return u.DeleteFn(db, usr)
}

// Update mock
func (u *User) Update(db orm.DB, usr andro.User) error {
	return u.UpdateFn(db, usr)
}
