package authutils

import (
	"minecv/internal/domain/entities"
)

// BuildUserResponse generates the user response data
func BuildUserResponse(user *entities.UserEntity) map[string]interface{} {
	return map[string]interface{}{
		"user_id":    user.UserID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"username":   user.Username,
		"email":      user.Email,
		"role":       user.Role,
	}
}
