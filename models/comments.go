package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `json:"user_id"`
	PhotoId   uint      `json:"photo_id"`
	Message   string    `json:"message" valid:"required~Your message is required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User
	Photo     *Photo
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(c)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}