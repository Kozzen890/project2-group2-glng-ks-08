package models

import (
	"time"

	"github.com/Kozzen890/project2-group2-glng-ks-08/helper"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	Id          uint    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null;uniqueIndex" json:"name" form:"name" valid:"required~Username is required"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required, email~Invalid email formal"`
	Password  string    `gorm:"not null" json:"password" form:"password" valid:"required,minstringlength(6)~Password has to have minimum length of 6 characters"`
	Age       int      `gorm:"not null" json:"age" form:"age" valid:"required~Age is required,range(8|70)~Minimum age is 8 years old"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Photos    []Photo   
	Comments  []Comment 
	Medias    []Media   
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

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helper.HashThePass(u.Password)
	err = nil
	return
}