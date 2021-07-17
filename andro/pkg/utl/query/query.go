package query

import (
	"github.com/labstack/echo"

	andro "github.com/soldevx/kuiper/andro/pkg/utl/model"
)

// List prepares data for list queries
func List(u andro.AuthUser) (*andro.ListQuery, error) {
	switch true {
	case u.Role <= andro.AdminRole: // user is SuperAdmin or Admin
		return nil, nil
	case u.Role == andro.CompanyAdminRole:
		return &andro.ListQuery{Query: "company_id = ?", ID: u.CompanyID}, nil
	case u.Role == andro.LocationAdminRole:
		return &andro.ListQuery{Query: "location_id = ?", ID: u.LocationID}, nil
	default:
		return nil, echo.ErrForbidden
	}
}
