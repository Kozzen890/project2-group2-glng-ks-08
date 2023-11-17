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

func GetAllMedia(ctx *gin.Context) {
	var Media []models.Media
	if err := databases.DB.Debug().Preload("User").Find(&Media).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "can't find media",
			"message": err.Error(),
		})
	}

	var res []dto.GetSocialMediaRes

	for _, media := range Media {
		user := dto.GetUserSocialMediaRes{
			Id: uint(media.User.Id),
			Username:  media.User.Name,
			ProfileImageUrl: "",
		}

		dtoMedia := dto.GetSocialMediaRes{
			Id:        media.Id,
			Name:     media.Name,
			SocialMediaUrl:   media.SocialMediaUrl,
			UserId:    media.UserId,
			CreatedAt: media.CreatedAt,
			UpdatedAt: media.UpdatedAt,
			User:      user,
		}

		// Tambahkan objek DTO ke slice
		res = append(res, dtoMedia)
	}
	ctx.JSON(http.StatusOK, res)
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "Get All Media has been successful",
	// 	"media": Media,
	// })
}

func GetSocialMediaById(ctx *gin.Context) {
	socialMedia := models.Media{}
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	if err := databases.DB.Preload("User").First(&socialMedia, socialMediaId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Media has been successful",
		"media": socialMedia,
	})
}

func CreateMedia(ctx *gin.Context) {
	contentType := helper.GetContentType(ctx)
	Media := models.Media{}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if contentType == "application/json" {
		ctx.ShouldBindJSON(&Media)
	} else {
		ctx.ShouldBind(&Media)
	}

	Media.UserId = userId

	if err := databases.DB.Debug().Create(&Media).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to upload social media",
			"message": err.Error(),
		})
		return
	}

	responses := dto.UploadMediaRes{
		Message:				"Upload Social Media has been successfully",
		Id:             Media.Id,
		Name:           Media.Name,
		SocialMediaUrl: Media.SocialMediaUrl,
		UserId:         Media.UserId,
		CreatedAt:      Media.CreatedAt,
	}

	ctx.JSON(http.StatusCreated, responses)

}

func UpdateMedia(ctx *gin.Context){
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)
	Media := models.Media{}

	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		ctx.ShouldBindJSON(&Media)
	} else {
		ctx.ShouldBind(&Media)
	}

	Media.UserId = userID
	Media.Id = uint(socialMediaId)

	err := databases.DB.Model(&Media).Where("id = ?", socialMediaId).Updates(models.Media{Name: Media.Name, SocialMediaUrl: Media.SocialMediaUrl}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	response := dto.UpdateMediaRes{
		Message:				"Update Social Media has been successfully",
		Id:             socialMediaId,
		Name:           Media.Name,
		SocialMediaUrl: Media.SocialMediaUrl,
		UserId:         Media.UserId,
		UpdatedAt:      Media.UpdatedAt,
	}
	ctx.JSON(http.StatusOK, response)
}

func DeleteMedia(ctx *gin.Context) {
	getId, _ := strconv.Atoi(ctx.Param("socialMediaId"))
	Media := models.Media{}

	if err := databases.DB.Debug().First(&Media, getId).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "media not found",
			"message": err.Error(),
		})
		return
	}

	if err := databases.DB.Debug().Delete(&Media).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to delete media",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your social media has been successfully deleted",
	})
}

