package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Kozzen890/project2-group2-glng-ks-08/databases"
	"github.com/Kozzen890/project2-group2-glng-ks-08/helper"
	"github.com/Kozzen890/project2-group2-glng-ks-08/models"
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

	token := helper.GenerateToken(uint(User.GormModel.Id), User.Name, User.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserUpdate(ctx *gin.Context) {
	GetId, _ := strconv.Atoi(ctx.Param("userId"))
	UserData := ctx.MustGet("userData").(jwt.MapClaims)
	UserId := UserData["id"].(float64)

	contextType := helper.GetContentType(ctx)
	_, _ = databases.DB, contextType

	User := models.User{}
	OldUser := models.User{}

	if contextType == "application/json" {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	User.UpdatedAt = time.Now()
	User.Id = int(UserId)

	if err := databases.DB.Where("id=?", GetId).Take(&OldUser).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "data not found",
			"message": err.Error(),
		})
		return
	}
	if err := databases.DB.Preload("Photos").Preload("Comments").Preload("Medias").Model(&OldUser).Updates(&User).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to update data",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update User Success",
		"id":         OldUser.Id,
		"email":      OldUser.Email,
		"username":   OldUser.Name,
		"age":        OldUser.Age,
		"updated_at": OldUser.UpdatedAt,
	})
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

	ctx.JSON(http.StatusOK, gin.H{
		// "message": "your account has been succesfully deleted",
		"message": fmt.Sprintf(" Your Account with Username %s and deleted successfully", userEntity.Name),
	})
}