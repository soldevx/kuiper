package user

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/labstack/echo"

	kuiper "github.com/soldevx/kuiper/kuipersrv"
	"github.com/soldevx/kuiper/kuipersrv/pkg/api/user/platform/pgsql"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, kuiper.User) (kuiper.User, error)
	List(echo.Context, kuiper.Pagination) ([]kuiper.User, error)
	View(echo.Context, int) (kuiper.User, error)
	Delete(echo.Context, int) error
	Update(echo.Context, Update) (kuiper.User, error)
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
	Create(orm.DB, kuiper.User) (kuiper.User, error)
	View(orm.DB, int) (kuiper.User, error)
	List(orm.DB, *kuiper.ListQuery, kuiper.Pagination) ([]kuiper.User, error)
	Update(orm.DB, kuiper.User) error
	Delete(orm.DB, kuiper.User) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) kuiper.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, kuiper.AccessRole, int, int) error
	IsLowerRole(echo.Context, kuiper.AccessRole) error
}
