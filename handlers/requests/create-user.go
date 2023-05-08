package requests

type CreateUserRequest struct {

	// Name of the user
	// required: true
	Name string `json:"name" validate:"required"`
	// Username of the user
	// required: true
	Username string `json:"username" validate:"required"`
	// Password of the user
	// required: true
	Password string `json:"password" validate:"required"`
}
