package query_test

import (
	"testing"

	"github.com/labstack/echo"

	kuiper "github.com/soldevx/kuiper/kuipersrv"

	"github.com/stretchr/testify/assert"

	"github.com/soldevx/kuiper/kuipersrv/pkg/utl/query"
)

func TestList(t *testing.T) {
	type args struct {
		user kuiper.AuthUser
	}
	cases := []struct {
		name     string
		args     args
		wantData *kuiper.ListQuery
		wantErr  error
	}{
		{
			name: "Super admin user",
			args: args{user: kuiper.AuthUser{
				Role: kuiper.SuperAdminRole,
			}},
		},
		{
			name: "Company admin user",
			args: args{user: kuiper.AuthUser{
				Role:      kuiper.CompanyAdminRole,
				CompanyID: 1,
			}},
			wantData: &kuiper.ListQuery{
				Query: "company_id = ?",
				ID:    1},
		},
		{
			name: "Location admin user",
			args: args{user: kuiper.AuthUser{
				Role:       kuiper.LocationAdminRole,
				CompanyID:  1,
				LocationID: 2,
			}},
			wantData: &kuiper.ListQuery{
				Query: "location_id = ?",
				ID:    2},
		},
		{
			name: "Normal user",
			args: args{user: kuiper.AuthUser{
				Role: kuiper.UserRole,
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
