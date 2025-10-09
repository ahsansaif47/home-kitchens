package handlers

import (
	"time"

	"github.com/ahsansaif47/home-kitchens/auth/config"
	"github.com/ahsansaif47/home-kitchens/auth/constants"
	"github.com/ahsansaif47/home-kitchens/auth/http/controllers"
	"github.com/ahsansaif47/home-kitchens/auth/http/dto"
	"github.com/ahsansaif47/home-kitchens/auth/models"
	"github.com/ahsansaif47/home-kitchens/auth/utils"
	"github.com/ahsansaif47/home-kitchens/auth/utils/jwt"
	"github.com/go-playground/validator/v10"

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

// CreateUser
//
//	@Summary		Create user
//	@Description	Create a new user
//	@Tags			User
//	@Accept			json
//	@Param			body	body	dto.UserSignupRequest	true	"Signup User Request"
//
//	@Produce		json
//	@Body			user  dto.UserSignupRequest true "User Signup Request"
//	@Success		200	{object}	dto.UserSignupResponse
//	@Failure		400	{object}	fiber.Error
//	@Failure		500	{object}	fiber.Error
//	@Router			/signup [post]
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

	passwordHash, err := utils.HashPassword(userReq.Password)
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
