package dto

type UpdateUserReq struct {
	Email string `json:"email" form:"name" valid:"required~Your email is required,email~Invalid email format"`
	Name  string `json:"name" form:"name" valid:"required~Your username is required"`
	Age   int    `json:"age" form:"age" valid:"required~Your age is required,range(8|99)~Minimum age 8 years"`
}