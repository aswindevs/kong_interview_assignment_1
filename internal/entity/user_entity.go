package entity

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Roles    []Role `json:"roles" gorm:"many2many:user_roles;"`
}

// Implementing custom JSON marshaling
func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID        uint      `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email     string    `json:"email"`
		Roles     []Role    `json:"roles"`
		Name      string    `json:"name"`
	}{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Email:     u.Email,
		Roles:     u.Roles,
		Name:      u.Name,
	})
}
