package request

type SignInRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"min=8,required"`
}

type SignUpRequest struct {
	Email                string `json:"email" validate:"email,required"`
	Name                 string `json:"name" validate:"required"`
	Password             string `json:"password" validate:"min=8,required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"min=8,required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"email,required"`
}

type UpdatePasswordRequest struct {
	CurrentPassword string `json:"current_password,omitempty"`
	NewPassword     string `json:"password_new" validate:"required, min=8"`
	ConfirmPassword string `json:"password_confirmation" validate:"required"`
}

type UpdateDataUserRequest struct {
	Name    string  `json:"name" validate:"required"`
	Email   string  `json:"email" validate:"email, required"`
	Phone   string  `json:"phone" validate:"required"`
	Address string  `json:"address" validate:"required"`
	Lat     float64 `json:"lat" validate:"required"`
	Lgn     float64 `json:"lgn" validate:"required"`
	Photo   string  `json:"photo" validate:"required"`
}
