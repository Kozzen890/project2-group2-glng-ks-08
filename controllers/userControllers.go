package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kozzen890/project2-group2-glng-ks-08/databases"
	"github.com/Kozzen890/project2-group2-glng-ks-08/dto"
	"github.com/Kozzen890/project2-group2-glng-ks-08/helper"
	"github.com/Kozzen890/project2-group2-glng-ks-08/models"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	db := databases.GetDB()
	contentType := helper.GetContentType(ctx)
	_, _ = db, contentType
	User := models.User{}

	if contentType == "application/json" {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	if err := databases.GetDB().Create(&User).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to create user data",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User Created",
		"id": User.Id,
		"email": User.Email,
		"name":  User.Name,
		"age": User.Age,
	})

	// db := databases.GetDB()
	// contentType := helper.GetContentType(ctx)
	// _, _ = db, contentType
	// User := models.User{}

	// if contentType == "application/json" {
	// 	ctx.ShouldBindJSON(&User)
	// } else {
	// 	ctx.ShouldBind(&User)
	// }

	// err := db.Create(&User).Error

	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error":   "Bad Request",
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusCreated, gin.H{
	// 	"id":       User.Id,
	// 	"name": User.Name,
	// 	"email":    User.Email,
	// 	"age":      User.Age,
	// })
}

func UserLogin(ctx *gin.Context){
	User := models.User{}
	password := ""

	if err := ctx.ShouldBindJSON(&User); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Input must be in JSON format"})
		return
	}

	password = User.Password

	if err := databases.DB.Where("email = ?", User.Email).Take(&User).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unathorized",
			"message": "Email not registered",
		})
		return
	}

	comparePass := helper.CompareThePass([]byte(User.Password), []byte(password))

	if !comparePass{
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	token := helper.GenerateToken(User.Id, User.Name, User.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserUpdate(ctx *gin.Context) {
	UserData := ctx.MustGet("userData").(jwt.MapClaims)
	request := dto.UpdateUserReq{}
	// contentType := helper.GetContentType(ctx)
	userEntity := models.User{}
	getUserID, _ := strconv.Atoi(ctx.Param("userId"))
	userId := uint(UserData["id"].(float64))

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	userEntity.Id = userId
	
	if userId != uint(getUserID) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not authorized to edit this user.",
		})
		return
	}

	if err := databases.DB.Model(&userEntity).Where("id = ?", getUserID).Updates(request).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update user data"})
		return
	}

	res := dto.UpdateUserRes{
		Id:        userEntity.Id,
		Email:     userEntity.Email,
		Username:  userEntity.Name,
		Age:       userEntity.Age,
		UpdatedAt: userEntity.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, res)
}

func UserDelete(ctx *gin.Context) {
	UserData := ctx.MustGet("userData").(jwt.MapClaims)
	UserId := int(UserData["id"].(float64))
	userEntity := models.User{}

	if err := databases.DB.First(&userEntity, UserId).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "data not found",
			"message": err.Error(),
		})
		return
	}

	if err := databases.DB.Preload("Photos").Preload("Comments").Preload("Medias").Model(&userEntity).Delete(&userEntity).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Failed to delete data",
			"message": err.Error(),
		})
		return
	}

	// if err := databases.DB.Delete(&userEntity).Error; err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{
	// 		"error":   "Unauthorizated",
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	ctx.JSON(http.StatusOK, gin.H{
		// "message": "your account has been succesfully deleted",
		"message": fmt.Sprintf(" Your Account with Username %s and deleted successfully", userEntity.Name),
	})
}