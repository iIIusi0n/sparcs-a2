package server

import (
	"api-server/config"
	"api-server/middlewares"
	"github.com/gin-gonic/gin"

	cImage "api-server/controllers/image"
	cUser "api-server/controllers/user"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			if config.ServerDebug {
				debug := v1.Group("/debug")
				{
					debug.POST("/token", cUser.TemporaryTokenRouter)
				}
			}

			image := v1.Group("/image")
			{
				image.GET("/:id", cImage.GetImageRouter)

				image.POST("/", cImage.UploadImageRouter)

				image.DELETE("/:id", cImage.DeleteImageRouter)
			}

			user := v1.Group("/user")
			{
				user.Use(middlewares.JwtAuthMiddleware)

				user.GET("/", cUser.GetLoggedInUserRouter)
				user.GET("/:id", cUser.GetUserRouter)

				user.PATCH("/", cUser.UpdateUserRouter)
			}
		}
	}

	return r
}
