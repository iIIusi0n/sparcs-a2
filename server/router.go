package server

import (
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
			user := v1.Group("/debug") // TODO: Remove this
			{
				user.GET("/token/:id/:name", cUser.TemporaryTokenRouter)
			}

			image := v1.Group("/image")
			{
				// image.Use(middlewares.JwtAuthMiddleware)

				image.GET("/:id", cImage.GetImageRouter)
				image.POST("/", cImage.UploadImageRouter)
				image.DELETE("/:id", cImage.DeleteImageRouter)
			}
		}
	}

	return r
}
