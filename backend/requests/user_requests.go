package requests

type (
	UserRequest struct {
		Username  string `json:"username" validate:"required"`
	}
)
