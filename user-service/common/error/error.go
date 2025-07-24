package error

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ValidateResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"messagge,omitempty"`
}

var ErrValidator = map[string]string{}

func ErrValidatorResponse(err error) (validationResponse []ValidateResponse) {
	var fieldErrors validator.ValidationErrors

	// cek apakah error termasuk dalam error validasi
	if errors.As(err, &fieldErrors) {
		for _, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				validationResponse = append(validationResponse, ValidateResponse{
					Field:   err.Field(),
					Message: fmt.Sprintf("%s is required", err.Field()),
				})
			case "email":
				validationResponse = append(validationResponse, ValidateResponse{
					Field:   err.Field(),
					Message: fmt.Sprintf("%s is not validate email addres", err.Field()),
				})
			case "phonenumber":
				validationResponse = append(validationResponse, ValidateResponse{
					Field:   err.Field(),
					Message: fmt.Sprintf("%s is not validate phone number", err.Field()),
				})
			default:
				errValidator, ok := ErrValidator[err.Tag()]
				if ok {
					count := strings.Count(errValidator, "%s")
					if count == 1 {
						validationResponse = append(validationResponse, ValidateResponse{
							Field:   err.Field(),
							Message: fmt.Sprintf(errValidator, err.Field()),
						})
					} else {
						validationResponse = append(validationResponse, ValidateResponse{
							Field:   err.Field(),
							Message: fmt.Sprintf(errValidator, err.Field(), err.Param()),
						})
					}
				} else {
					validationResponse = append(validationResponse, ValidateResponse{
						Field:   err.Field(),
						Message: fmt.Sprintf("something wrong on %s; %s", err.Field(), err.Tag()),
					})
				}
			}

		}
	}
	return validationResponse
}

// untuk log error
func WrapError(err error) error {
	logrus.Error("error %v", err)
	return err
}
