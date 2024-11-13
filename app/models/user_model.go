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
	PhoneNum string `gorm:"unique;not null;size:11" json:"phone_num" binding:"required"`

	// Password is the hashed password of the user
	// It is not included in JSON output for security reasons
	Password string `gorm:"not null;size:128" json:"-" binding:"required"`

	// Name is the full name of the user
	Name string `gorm:"not null;size:128" json:"name" binding:"required"`

	// UserStatus represents the current status of the user (e.g., active, inactive)
	UserStatus UserStatus `gorm:"not null;default:active" json:"status"`

	// Role defines the user's role in the system, determining their permissions
	Role UserRole `gorm:"not null" json:"role"`
}

func (User) TableName() string { return "user" }

type CreateUserRequest struct {
	Name     string   `json:"name" binding:"required,min=1,max=100"`
	PhoneNum string   `json:"phone_num" binding:"required,len=11"`
	Password string   `json:"password" binding:"required,min=8,max=18"`
	Role     UserRole `json:"role" binding:"required,oneof=admin user viewer"`
}

type UpdateUserRoleRequest struct {
	Role UserRole `json:"role" binding:"required,oneof=admin user viewer"`
}

type UpdateUserInfoRequest struct {
	Name     string `json:"name" binding:"min=1,max=100"`
	PhoneNum string `json:"phone_num" binding:"len=11"`
	Password string `json:"password" binding:"min=8,max=18"`
}
