package dto

type UpdateUserReq struct {
	Email    string `json:"email" valid:"required~Your email is required,email~Invalid email format"`
	Username string `json:"username" valid:"required~Your username is required"`
	Age      int    `json:"age" valid:"required~Your age is required,range(8|99)~Minimum age 8 years"`
}