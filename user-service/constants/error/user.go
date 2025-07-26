package error

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordIncorrect    = errors.New("password incorrect")
	ErrUsernameExist        = errors.New("username already exists")
	ErrEmailExist           = errors.New("email already exists")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
	ErrPhoneNumberExist     = errors.New("phone number already exists")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordIncorrect,
	ErrUsernameExist,
	ErrPasswordDoesNotMatch,
}
