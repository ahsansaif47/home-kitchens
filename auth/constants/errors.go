package constants

import "errors"

var CustomValidationErrors = map[string]string{
	"Email":    "Email is incorrect",
	"Password": "Passowrd is incorrect",
	"PhoneNo":  "Phone No. is incorrect",
	"DOB":      "Date of birth is incorrect",
}

var ErrUserAlreadyExists = errors.New("user already exists")
