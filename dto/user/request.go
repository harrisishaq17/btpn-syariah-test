package user

type (
	CreateUserRequest struct {
		UserID string `json:"user_id" validate:"required,max=15"`
		Name   string `json:"name" validate:"required,max=100"`
		Phone  string `json:"phone" validate:"required,max=15"`
	}
)

