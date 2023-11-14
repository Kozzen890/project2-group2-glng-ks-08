package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Media struct {
	Id             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(155)" json:"name" valid:"required~Your name is required"`
	SocialMediaUrl string    `json:"social_media_url" valid:"required~Your social media url is required"`
	UserId         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           *User
}

func (m *Media) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(m)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (m *Media) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(m)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}