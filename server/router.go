package server

import (
	"api-server/config"
	"api-server/middlewares"
	"github.com/gin-gonic/gin"

	cImage "api-server/controllers/image"
	cPost "api-server/controllers/post"
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
					debug.GET("/token/:id/:name", cUser.TemporaryTokenRouter)
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
				user.GET("/stats", cUser.GetStatsRouter)

				user.PATCH("/", cUser.UpdateUserRouter)
			}

			post := v1.Group("/post")
			{
				post.Use(middlewares.JwtAuthMiddleware)

				post.GET("/", cPost.GetPostsRouter)
				post.GET("/:id", cPost.GetPostRouter)

				post.POST("/", cPost.CreatePostRouter)

				post.PATCH("/:id", cPost.UpdatePostRouter)

				post.DELETE("/:id", cPost.DeletePostRouter)
			}
		}
	}

	return r
}
