package models

import "time"

type GormModel struct {
	Id         int       `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}