package services

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"minecv/internal/domain/entities"
	"minecv/internal/domain/schemas"
	valueobjects "minecv/internal/domain/value_objects"
	"minecv/internal/infrastructure/database"
	"minecv/pkg/lib"
)

func CreateUser(input schemas.CreateUserSchemas) (*entities.UserEntity, string, string, error) {
	// Check if the email is already registered
	var existingUser entities.UserEntity
	if err := database.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return nil, "", "", fmt.Errorf("email already exists")
	}

	// Check if the username is already registered
	if err := database.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		return nil, "", "", fmt.Errorf("username already exists")
	}

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", "", fmt.Errorf("error hashing password: %v", err)
	}

	// Create a new User model instance and set fields
	newUser := entities.UserEntity{
		UserID:       uuid.New().String(),
		Email:        input.Email,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Username:     input.Username,
		PasswordHash: string(hashedPassword),
		Role:         valueobjects.Role("user"),
	}

	// Save new user to the database
	if err := database.DB.Create(&newUser).Error; err != nil {
		return nil, "", "", fmt.Errorf("error creating user: %v", err)
	}

	// Generate tokens
	accessToken, err := lib.GenerateToken(newUser.UserID)
	if err != nil {
		return nil, "", "", fmt.Errorf("error generating access token: %v", err)
	}

	refreshToken, err := lib.GenerateRefreshToken(newUser.UserID)
	if err != nil {
		return nil, "", "", fmt.Errorf("error generating refresh token: %v", err)
	}

	return &newUser, accessToken, refreshToken, nil
}

func AuthenticateUser(input schemas.LoginUserSchemas) (*entities.UserEntity, string, string, error) {
	var user entities.UserEntity

	// Check if user provided email or username
	if input.Email != "" {
		if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
			return nil, "", "", fmt.Errorf("email not found")
		}
	} else if input.Username != "" {
		if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
			return nil, "", "", fmt.Errorf("user not found")
		}
	} else {
		return nil, "", "", fmt.Errorf("email or username must be provided")
	}

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return nil, "", "", fmt.Errorf("invalid password")
	}

	// Generate tokens
	accessToken, err := lib.GenerateToken(user.UserID)
	if err != nil {
		return nil, "", "", fmt.Errorf("error generating access token: %v", err)
	}

	refreshToken, err := lib.GenerateRefreshToken(user.UserID)
	if err != nil {
		return nil, "", "", fmt.Errorf("error generating refresh token: %v", err)
	}

	return &user, accessToken, refreshToken, nil
}
