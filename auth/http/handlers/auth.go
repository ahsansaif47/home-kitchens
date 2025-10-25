package handlers

import (
	"errors"
	"time"

	"github.com/ahsansaif47/home-kitchens/auth/config"
	"github.com/ahsansaif47/home-kitchens/auth/constants"
	"github.com/ahsansaif47/home-kitchens/auth/http/controllers"
	"github.com/ahsansaif47/home-kitchens/auth/http/dto"
	"github.com/ahsansaif47/home-kitchens/auth/models"
	"github.com/ahsansaif47/home-kitchens/auth/utils"
	"github.com/ahsansaif47/home-kitchens/auth/utils/jwt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service   controllers.IUserService
	validator *validator.Validate
}

func NewAuthHandler(service controllers.IUserService) *AuthHandler {
	return &AuthHandler{
		service:   service,
		validator: validator.New(),
	}
}

// SignUp
//
//	@Summary		Create User
//	@Description	Create a New User
//	@Tags			User
//	@Accept			json
//	@Param			body	body	dto.UserSignupRequest	true	"Signup User Request"
//
//	@Produce		json
//	@Body			user  dto.UserSignupRequest true "User Signup Request"
//	@Success		200	{object}	dto.UserSignupResponse
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/users/signup [post]
func (h *AuthHandler) CreateUser(ctx *fiber.Ctx) error {
	userReq := dto.UserSignupRequest{}

	err := ctx.BodyParser(&userReq)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	h.validator.RegisterValidation("password", utils.PasswordValidator)
	// Validate the request
	if err = h.validator.Struct(userReq); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			fieldName := err.Field()
			if customMsg, exists := constants.CustomValidationErrors[fieldName]; exists {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": customMsg,
				})
			}
		}
	}

	passwordHash, err := utils.GeneratePasswordHash(userReq.Password)
	if passwordHash == "" || err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	now := time.Now().UTC()
	newUser := &models.User{
		UserName:         userReq.UserName,
		FirstName:        userReq.FirstName,
		Email:            userReq.Email,
		PasswordHash:     &passwordHash,
		CreatedAt:        &now,
		UpdatedAt:        &now,
		AuthProviderType: "local",
		DOB:              userReq.DOB,
		PhoneNumber:      userReq.PhoneNumber,
		RoleID:           userReq.RoleID,
	}

	err = h.service.CreateUser(newUser)
	if err == constants.ErrUserAlreadyExists {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User with this email already exists"})
	}

	tokenStr, err := jwt.GenerateJWT(newUser.Email, userReq.UserName, userReq.RoleID, config.GetConfig().JWTSecret)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate access token",
		})
	}

	response := dto.UserSignupResponse{
		AccessToken: tokenStr,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// SignIn
//
//	@Summary		Authenticate User
//	@Description	Login User into the System
//	@Tags			User
//	@Accept			json
//	@Param			body	body	dto.UserLoginRequest	true	"Signin User Request"
//
//	@Produce		json
//	@Success		200	{object}	dto.UserLoginResponse
//	@Failure		400	{object}	dto.ErrorResponse
//	@Failure		500	{object}	dto.ErrorResponse
//	@Router			/users/signin [post]
func (h *AuthHandler) Signin(ctx *fiber.Ctx) error {
	signinReq := dto.UserLoginRequest{}

	err := ctx.BodyParser(&signinReq)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	h.validator.RegisterValidation("password", utils.PasswordValidator)

	if err := h.validator.Struct(&signinReq); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			fieldName := err.Field()
			if msg, exist := constants.CustomValidationErrors[fieldName]; exist {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": msg,
				})
			}
		}
	}

	passwordHash, err := utils.GeneratePasswordHash(signinReq.Password)
	if passwordHash == "" || err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Get the user details from DB
	user, err := h.service.FindUserByEmail(signinReq.Email)
	if err != nil {
		errResponse := dto.ErrorResponse{
			Message: err.Error(),
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(errResponse)

		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(errResponse)
	}

	accessToken, err := jwt.GenerateJWT(signinReq.Email, user.UserName, user.RoleID, config.GetConfig().JWTSecret)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Message: err.Error(),
		})
	}

	response := dto.UserLoginResponse{
		AccessToken: accessToken,
		UserName:    user.UserName,
		RoleID:      user.RoleID,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
