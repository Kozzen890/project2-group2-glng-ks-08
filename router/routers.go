package routers

import (
	"net/http"

	"github.com/Kozzen890/project2-group2-glng-ks-08/controllers"
	"github.com/Kozzen890/project2-group2-glng-ks-08/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		// Mengirimkan kata-kata sebagai respons
		c.String(http.StatusOK, "Hello welcome to MyGram-API")
	})

	userRouter := router.Group("/users")
	{
		userRouter.GET("/", controllers.GetUsers)
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userId", middlewares.Authentication(), controllers.UserUpdate)
		userRouter.DELETE("/", middlewares.Authentication(), controllers.UserDelete)
	}

	mediaRouter := router.Group("/socialmedias")
	{
		mediaRouter.Use(middlewares.Authentication())
		mediaRouter.GET("/", controllers.GetAllMedia)
		mediaRouter.GET("/:socialMediaId", controllers.GetSocialMediaById)
		mediaRouter.POST("/", controllers.CreateMedia)
		mediaRouter.PUT("/:socialMediaId", middlewares.MediaAuthorization(),controllers.UpdateMedia)
		mediaRouter.DELETE("/:socialMediaId", middlewares.MediaAuthorization(), controllers.DeleteMedia)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", controllers.GetAllPhotos)
		photoRouter.GET("/:photoId", controllers.GetPhotoById)
		photoRouter.POST("/", controllers.UploadPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", controllers.GetAllComment)
		commentRouter.GET("/:commentId", controllers.GetCommentById)
		commentRouter.POST("/", controllers.UploadComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.EditComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	return router
}