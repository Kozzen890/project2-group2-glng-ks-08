package controllers

import (
	"net/http"
	"strconv"

	"github.com/Kozzen890/project2-group2-glng-ks-08/databases"
	"github.com/Kozzen890/project2-group2-glng-ks-08/helper"
	"github.com/Kozzen890/project2-group2-glng-ks-08/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetAllComment(ctx *gin.Context) {
	Comment := []models.Comment{}
	if err := databases.DB.Debug().Preload("User").Preload("Photo").Find(&Comment).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "comment not found",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Comment,
	})
}

func GetCommentById(ctx *gin.Context) {
	comment := models.Comment{}
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))

	if err := databases.DB.Debug().Preload("User").Preload("Photo").Find(&comment, commentId).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "comment not found",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func UploadComment(ctx *gin.Context) {
	contentType := helper.GetContentType(ctx)
	Comment := models.Comment{}
	Photo := models.Photo{}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if contentType == "application/json" {
		ctx.ShouldBindJSON(&Comment)
	} else {
		ctx.ShouldBind(&Comment)
	}

	Comment.UserId = userId

	if err := databases.DB.Select("id").First(&Photo, Comment.PhotoId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Photo not found",
		})
		return
	}

	if err := databases.DB.Debug().Create(&Comment).Error; err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error":   "failed to upload comment",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"result": "Upload Comment has been successfully",
		"id":         Comment.Id,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoId,
		"user_id":    Comment.UserId,
		"created_at": Comment.CreatedAt,
	})
}

func EditComment(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(ctx)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(ctx.Param("commentId"))
	userId := uint(userData["id"].(float64))

	if contentType == "application/json" {
		ctx.ShouldBindJSON(&Comment)
	} else {
		ctx.ShouldBind(&Comment)
	}

	Comment.UserId = userId
	Comment.Id = uint(commentId)

	if err := databases.DB.Select("user_id").First(&Comment, commentId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Data not found",
		})
		return
	}

	if err := databases.DB.Debug().Model(&Comment).Where("id = ?", commentId).Updates(&Comment).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to update comment",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         Comment.Id,
		"message":    Comment.Message,
		"user_id":    Comment.UserId,
		"updated_at": Comment.UpdatedAt,
	})
}

func DeleteComment(ctx *gin.Context) {
	GetCommentId, _ := strconv.Atoi(ctx.Param("commentId"))

	Comment := models.Comment{}

	if err := databases.DB.Debug().First(&Comment, GetCommentId).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "comment not found",
			"message": err.Error(),
		})
		return
	}

	if err := databases.DB.Debug().Delete(&Comment).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed to delete comment",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your comment has been successfully deleted",
	})

}