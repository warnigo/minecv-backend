package valueobjects

// Role defines the available user roles
type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)
