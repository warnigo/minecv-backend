package schemas

// CreateUserSchemas represents the input data for user sign up
type CreateUserSchemas struct {
	FirstName string `json:"first_name" binding:"required,min=3,max=100" example:"John"`
	LastName  string `json:"last_name" binding:"required,min=3,max=100" example:"Doe"`
	Username  string `json:"username" binding:"required,min=3,max=22" example:"warnigo"`
	Email     string `json:"email" binding:"required,email" example:"john@doe.com"`
	Password  string `json:"password" binding:"required,min=8,max=20"`
}

// LoginUserSchemas represents the input data for user log in
type LoginUserSchemas struct {
	Email    string `json:"email,omitempty" binding:"omitempty,email" example:"john@doe.com"`
	Username string `json:"username,omitempty" binding:"omitempty,min=3,max=22" example:"warnigo"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}
