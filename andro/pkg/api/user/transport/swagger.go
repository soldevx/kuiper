package transport

import kuiper "github.com/soldevx/kuiper/kuipersrv"

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*kuiper.User
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Users []kuiper.User `json:"users"`
		Page  int           `json:"page"`
	}
}
