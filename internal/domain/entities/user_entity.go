package entities

import valueobjects "minecv/internal/domain/value_objects"

// UserEntity represents a user in the system
type UserEntity struct {
	BaseEntity
	UserID    	 string            `gorm:"unique;not null" json:"user_id"`
	Username     string            `gorm:"unique;not null" json:"username"`
	FirstName    string            `gorm:"not null" json:"first_name" example:"John"`
	LastName     string            `gorm:"not null" json:"last_name" example:"Deo"`
	Patronymic   *string           `json:"patronymic,omitempty" example:"Williamson"`
	Email        string            `gorm:"unique;not null" json:"email" example:"email@example.com"`
	PasswordHash string            `gorm:"column:password_hash;not null" json:"-"`
	PhoneNumber  string            `gorm:"unique;not null" json:"phone_number" example:"+1234567890"`
	DateOfBirth  *string           `json:"date_of_birth,omitempty" example:"01-01-2006" format:"day-month-year"`
	Gender       *string           `json:"gender,omitempty" example:"male"`
	Role         valueobjects.Role `gorm:"not null;default:'user'" json:"role" example:"user"`
}
