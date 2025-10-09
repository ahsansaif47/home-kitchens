package dto

type UserSignupRequest struct {
	UserName    string  `json:"user_name" validate:"required,min=3,max=50"`
	FirstName   string  `json:"first_name" validate:"required, min=3,max=50"`
	Email       string  `json:"email" validate:"required,email"`
	Password    string  `json:"password" validate:"required,min=6"`
	DOB         *string `json:"dob,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	RoleID      uint    `json:"role_id" validate:"required,oneof=1 2 3"`
}

type UserSignupResponse struct {
	AccessToken string `json:"access_token"`
}
