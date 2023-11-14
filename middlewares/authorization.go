package middlewares

import (
	"net/http"
	"strconv"

	"github.com/Kozzen890/project2-group2-glng-ks-08/databases"
	"github.com/Kozzen890/project2-group2-glng-ks-08/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func MediaAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := databases.GetDB()
		socialMediaId, err := strconv.Atoi(ctx.Param("socialMediaId"))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error" : "Bad Request",
				"message" : "Invalid Parameter",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		Media := models.Media{}

		if err := db.Preload("User").First(&Media, socialMediaId).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Data Not Found",
				"message" : "Data doesn't exist",
			})
			return
		}

		if Media.UserId != userId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You're not allowed to access this data",
			})
		}

		ctx.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := databases.GetDB()
		photoId, err := strconv.Atoi(ctx.Param("photoId"))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error" : "Bad Request",
				"message" : "Invalid Parameter",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		Photos := models.Photo{}

		if err := db.Preload("User").Preload("Comments").First(&Photos, photoId).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Data Not Found",
				"message" : "Data doesn't exist",
			})
			return
		}

		if Photos.UserId != userId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You're not allowed to access this data",
			})
		}

		ctx.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := databases.GetDB()
		commentId, err := strconv.Atoi(ctx.Param("commentId"))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error" : "Bad Request",
				"message" : "Invalid Parameter",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		Comment := models.Comment{}

		if err := db.Preload("User").Preload("Photo").First(&Comment, commentId).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error" : "Data Not Found",
				"message" : "Data doesn't exist",
			})
			return
		}

		if Comment.UserId != userId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You're not allowed to access this data",
			})
		}

		ctx.Next()
	}
}