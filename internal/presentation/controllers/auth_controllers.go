package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"minecv/internal/domain/schemas"
	"minecv/internal/domain/services"
	"minecv/internal/infrastructure/localization"
	responseutils "minecv/internal/presentation/controllers/utils"
	"minecv/internal/presentation/controllers/validation"
	"minecv/pkg/utils"
)

// Register godoc
// @Summary Register a new user
// @Description Create a new user by providing email, username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body schemas.CreateUserSchemas true "User Data"
// @Success 201 {object} entities.UserEntity
// @Failure 400 {string} string "Bad request"
// @Router /register [post]
func Register(c *gin.Context) {
	var input schemas.CreateUserSchemas
	lang := c.GetString("lang")
	localizer, _ := c.Get("localizer")
	i18n := localizer.(*localization.I18n)
	translate := utils.GetTranslator(i18n, lang)

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondValidationError(c, map[string]string{"error": "validations.invalid_input"}, "validations.validation_failed")
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		validationErrors := validation.ParseValidationErrors(err, i18n, lang)
		utils.RespondValidationError(c, validationErrors, translate("validations.validation_failed", nil))
		return
	}

	user, accessToken, refreshToken, err := services.CreateUser(input)
	if err != nil {
		utils.Respond(c, http.StatusBadRequest, true, false, nil, nil, err.Error())
		return
	}

	responseData := map[string]interface{}{
		"user":          responseutils.BuildUserResponse(user),
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	utils.RespondSuccess(c, http.StatusCreated, responseData, translate("auth.registered_successfully", nil))
}

// LoginUser godoc
// @Summary Log in a user
// @Description Log in by providing email and password to receive a JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body schemas.UserLoginSchemaIn true "User Login Data"
// @Success 200 {string} string "JWT Token"
// @Failure 401 {string} string "Unauthorized"
// @Router /login [post]
func Login(c *gin.Context) {
	var input schemas.LoginUserSchemas

	lang := c.GetString("lang")
	localizer, _ := c.Get("localizer")
	i18n := localizer.(*localization.I18n)
	translate := utils.GetTranslator(i18n, lang)

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondValidationError(c, map[string]string{"error": "validations.invalid_input"}, "validations.validation_failed")
		return
	}

	user, accessToken, refreshToken, err := services.AuthenticateUser(input)
	if err != nil {
		utils.Respond(c, http.StatusUnauthorized, true, false, nil, nil, translate("auth.invalid_credentials", nil))
		return
	}

	responseData := map[string]interface{}{
		"user":          responseutils.BuildUserResponse(user),
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	utils.RespondSuccess(c, http.StatusOK, responseData, translate("auth.login_successful", nil))
}
