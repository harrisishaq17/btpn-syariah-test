package user

type (
    UserResponse struct {
        UserID string `json:"user_id"`
        Name   string `json:"name"`
        Phone  string `json:"phone"`
    }
)
