package controllers

import (
	"net/http"
	"strconv"

	"github.com/Kozzen890/project2-group2-glng-ks-08/databases"
	"github.com/Kozzen890/project2-group2-glng-ks-08/dto"
	"github.com/Kozzen890/project2-group2-glng-ks-08/helper"
	"github.com/Kozzen890/project2-group2-glng-ks-08/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetAllPhotos(ctx *gin.Context) {
	var Photos []models.Photo
	if err := databases.DB.Debug().Preload("Comments").Preload("User").Find(&Photos).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}

	var res []dto.GetPhotosWithUser

	// Iterasi melalui setiap foto dan buat objek DTO
	for _, photo := range Photos {
		user := dto.GetUsersPhoto{
			Email: photo.User.Email,
			Name:  photo.User.Name,
		}

		dtoPhoto := dto.GetPhotosWithUser{
			Id:        photo.Id,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoUrl,
			UserId:    photo.UserId,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User:      user,
		}

		// Tambahkan objek DTO ke slice
		res = append(res, dtoPhoto)
	}
	ctx.JSON(http.StatusOK, res)
}

func GetPhotoById(ctx *gin.Context) {
	var photo models.Photo

	// Mendapatkan ID dari parameter URL
	photoId, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid photo ID",
		})
		return
	}

	// Menggunakan Preload untuk memuat informasi pengguna terkait
	err = databases.DB.Debug().Preload("User").First(&photo, "id = ?", photoId).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Photo not found",
			"message": err.Error(),
		})
		return
	}

	// Memeriksa apakah ada pengguna terkait dengan foto
	var user dto.GetUsersPhoto
	if photo.User != nil {
		user = dto.GetUsersPhoto{
			Email: photo.User.Email,
			Name:  photo.User.Name,
		}
	}

	// Membuat objek DTO untuk respons
	res := dto.GetPhotosWithUser{
		Id:        photo.Id,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoUrl,
		UserId:    photo.UserId,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
		User:      user,
	}
	ctx.JSON(http.StatusOK, res)
}

func UploadPhoto(ctx *gin.Context) {
	Photos := models.Photo{}
	contextType := helper.GetContentType(ctx)

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	if contextType == "application/json" {
		ctx.ShouldBindJSON(&Photos)
	} else {
		ctx.ShouldBind(&Photos)
	}

	Photos.UserId = uint(userId)
	if err := databases.DB.Debug().Create(&Photos).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to updload photo",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":         Photos.Id,
		"title":      Photos.Title,
		"caption":    Photos.Caption,
		"photo_url":  Photos.PhotoUrl,
		"user_id":    Photos.UserId,
		"created_at": Photos.CreatedAt,
	})
}

func UpdatePhoto(ctx *gin.Context) {
	Photos := models.Photo{}
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"]

	contextType := helper.GetContentType(ctx)

	if contextType == "application/json" {
		ctx.ShouldBindJSON(&Photos)
	} else {
		ctx.ShouldBind(&Photos)
	}

	Photos.UserId = uint(userId.(float64))
	Photos.Id = uint(photoId)

	err := databases.DB.Model(&Photos).Where("id = ?", photoId).Updates(models.Photo{Title: Photos.Title, Caption: Photos.Caption, PhotoUrl: Photos.PhotoUrl}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	res := dto.PhotoUpdateRes{
		Id:        photoId,
		Title:     Photos.Title,
		Caption:   Photos.Caption,
		PhotoUrl:  Photos.PhotoUrl,
		UserId:    Photos.UserId,
		UpdatedAt: Photos.UpdatedAt,
	}
	
	ctx.JSON(http.StatusOK, res)
}

func DeletePhoto(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))
	Photos := models.Photo{}

	if err := databases.DB.Debug().First(&Photos, photoId).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "photo not found",
			"message": err.Error(),
		})
		return
	}

	if err := databases.DB.Debug().Delete(&Photos).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to delete photo",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  "Delete Success",
		"Message": "The photo has been successfully deleted",
	})
}