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

// RefreshToken godoc
// @Summary Refresh access and refresh tokens
// @Description Refresh JWT tokens using a valid refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param token body struct { RefreshToken string `json:"refresh_token"` } true "Refresh Token"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Router /refresh [post]
func RefreshToken(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}

	lang := c.GetString("lang")
	localizer, _ := c.Get("localizer")
	i18n := localizer.(*localization.I18n)
	translate := utils.GetTranslator(i18n, lang)

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondError(c, http.StatusBadRequest, translate("validations.invalid_payload", nil))
		return
	}

	newAccessToken, newRefreshToken, err := services.RefreshTokenService(request.RefreshToken)
	if err != nil {
		utils.RespondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.RespondSuccess(c, http.StatusOK, map[string]string{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	}, translate("auth.token_refreshed_successfully", nil))
}