package enum

import "github.com/traPtitech/game3-back/openapi/models"

type UserRole int

const (
	Guest UserRole = iota // 0
	User                  // 1
	Admin                 // 2
)

func (r UserRole) IsAdmin() bool {
	return r == Admin
}
func (r UserRole) IsUser() bool {
	return r == User
}
func (r UserRole) IsGuest() bool {
	return r == Guest
}
func (r UserRole) IsUserOrAbove() bool {
	return r >= User
}

func (r UserRole) String() string {
	switch r {
	case Guest:
		return "guest"
	case User:
		return "user"
	case Admin:
		return "admin"
	default:
		return "guest"
	}
}

func (r UserRole) ToModelsUserRole() models.UserRole {
	switch r {
	case Guest:
		return models.UserRoleGuest
	case User:
		return models.UserRoleUser
	case Admin:
		return models.UserRoleAdmin
	default:
		return models.UserRoleGuest
	}
}
