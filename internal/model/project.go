package model

import "gorm.io/gorm"

type ProjectStatus string

const (
	ProjectStatusActive   ProjectStatus = "active"
	ProjectStatusPaused   ProjectStatus = "paused"
	ProjectStatusArchived ProjectStatus = "archived"
)

type Project struct {
	gorm.Model
	OwnerID     uint          `gorm:"not null;index"`
	Owner       User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name        string        `gorm:"not null"`
	Key         string        `gorm:"uniqueIndex;not null"`
	Description string        `gorm:"type:text"`
	Status      ProjectStatus `gorm:"type:text;not null;default:active"`
}
