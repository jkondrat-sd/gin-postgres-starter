// Why: user models define the database shape that GORM persists for application users.
// What to do: add persistent user fields here, then mirror public fields in DTO responses when needed.
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Name         string `gorm:"not null"`
	Role         string `gorm:"not null;default:user"`
	Projects     []Project
}
