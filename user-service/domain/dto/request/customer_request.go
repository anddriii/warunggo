package request

type CustomerRequest struct {
	Name                 string  `json:"name" validate:"required"`
	Email                string  `json:"email" validate:"email, required"`
	Phone                string  `json:"phone" validate:"required"`
	Password             string  `json:"password" validate:"required"`
	PasswordConfirmation string  `json:"password_confirmation" validate:"required"`
	Address              string  `json:"address"`
	Photo                string  `json:"photo"`
	Lat                  float64 `json:"lat"`
	Lgn                  float64 `json:"lgn"`
	RoleId               int     `json:"role_id"`
}
