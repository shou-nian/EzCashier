package models

import (
	"gorm.io/gorm"
)

type (
	UserStatus string
	UserRole   string
)

const (
	StatusActive   UserStatus = "active"
	StatusInactive UserStatus = "inactive"

	RoleAdmin  UserRole = "admin"
	RoleUser   UserRole = "user"
	RoleViewer UserRole = "viewer"
)

// User represents a user in the system.
// It contains basic user information and is used for authentication and authorization.
type User struct {
	// gorm.Model embeds fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model

	// Email is the unique identifier for the user
	Email string `gorm:"unique;not null" json:"email"`

	// Password is the hashed password of the user
	// It is not included in JSON output for security reasons
	Password string `gorm:"not null" json:"-"`

	// Name is the full name of the user
	Name string `gorm:"not null" json:"name"`

	// UserStatus represents the current status of the user (e.g., active, inactive)
	UserStatus UserStatus `gorm:"not null" json:"status"`

	// Role defines the user's role in the system, determining their permissions
	Role UserRole `gorm:"not null" json:"role"`
}

func (User) TableName() string { return "user" }
