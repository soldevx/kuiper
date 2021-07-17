package query_test

import (
	"testing"

	"github.com/labstack/echo"

	"github.com/stretchr/testify/assert"

	"github.com/soldevx/kuiper/andro/pkg/utl/query"

	andro "github.com/soldevx/kuiper/andro/pkg/utl/model"
)

func TestList(t *testing.T) {
	type args struct {
		user andro.AuthUser
	}
	cases := []struct {
		name     string
		args     args
		wantData *andro.ListQuery
		wantErr  error
	}{
		{
			name: "Super admin user",
			args: args{user: andro.AuthUser{
				Role: andro.SuperAdminRole,
			}},
		},
		{
			name: "Company admin user",
			args: args{user: andro.AuthUser{
				Role:      andro.CompanyAdminRole,
				CompanyID: 1,
			}},
			wantData: &andro.ListQuery{
				Query: "company_id = ?",
				ID:    1},
		},
		{
			name: "Location admin user",
			args: args{user: andro.AuthUser{
				Role:       andro.LocationAdminRole,
				CompanyID:  1,
				LocationID: 2,
			}},
			wantData: &andro.ListQuery{
				Query: "location_id = ?",
				ID:    2},
		},
		{
			name: "Normal user",
			args: args{user: andro.AuthUser{
				Role: andro.UserRole,
			}},
			wantErr: echo.ErrForbidden,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q, err := query.List(tt.args.user)
			assert.Equal(t, tt.wantData, q)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
