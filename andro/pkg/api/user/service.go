package user

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/labstack/echo"

	"github.com/soldevx/kuiper/andro/pkg/api/user/platform/pgsql"

	andro "github.com/soldevx/kuiper/andro/pkg/utl/model"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, andro.User) (andro.User, error)
	List(echo.Context, andro.Pagination) ([]andro.User, error)
	View(echo.Context, int) (andro.User, error)
	Delete(echo.Context, int) error
	Update(echo.Context, Update) (andro.User, error)
}

// New creates new user application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *User {
	return &User{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes User application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *User {
	return New(db, pgsql.User{}, rbac, sec)
}

// User represents user application service
type User struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents user repository interface
type UDB interface {
	Create(orm.DB, andro.User) (andro.User, error)
	View(orm.DB, int) (andro.User, error)
	List(orm.DB, *andro.ListQuery, andro.Pagination) ([]andro.User, error)
	Update(orm.DB, andro.User) error
	Delete(orm.DB, andro.User) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) andro.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, andro.AccessRole, int, int) error
	IsLowerRole(echo.Context, andro.AccessRole) error
}
