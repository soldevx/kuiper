package query

import (
	"github.com/labstack/echo"

	kuiper "github.com/soldevx/kuiper/kuipersrv"
)

// List prepares data for list queries
func List(u kuiper.AuthUser) (*kuiper.ListQuery, error) {
	switch true {
	case u.Role <= kuiper.AdminRole: // user is SuperAdmin or Admin
		return nil, nil
	case u.Role == kuiper.CompanyAdminRole:
		return &kuiper.ListQuery{Query: "company_id = ?", ID: u.CompanyID}, nil
	case u.Role == kuiper.LocationAdminRole:
		return &kuiper.ListQuery{Query: "location_id = ?", ID: u.LocationID}, nil
	default:
		return nil, echo.ErrForbidden
	}
}
