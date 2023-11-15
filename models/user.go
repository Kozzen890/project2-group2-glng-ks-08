package models

import (
	"github.com/Kozzen890/project2-group2-glng-ks-08/helper"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Name      string    `gorm:"not null;uniqueIndex;constraint" json:"name" form:"name" valid:"required~Your username is required"`
	Email     string    `gorm:"not null;uniqueIndex;constraint" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password  string     `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age       int      `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required,range(9|100)~Age has to be above 8 years old"`
	Photos    []Photo   `gorm:"constraint:OnUpdate:CASCADE,onDelete:SET NULL;" json:"photos"`
	Comments  []Comment `gorm:"constraint:OnUpdate:CASCADE,onDelete:SET NULL;" json:"comments"`
	Medias    []Media   `gorm:"constraint:OnUpdate:CASCADE,onDelete:SET NULL;" json:"social_medias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helper.HashThePass(u.Password)
	err = nil
	return
}