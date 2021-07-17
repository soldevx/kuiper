package transport

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*andro.User
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Users []andro.User `json:"users"`
		Page  int          `json:"page"`
	}
}
