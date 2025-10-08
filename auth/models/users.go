package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName         string
	FirstName        string
	Email            string
	PasswordHash     *string
	AuthProviderType string
	RoleID           string
	DOB              *string
	PhoneNumber      *string
}
