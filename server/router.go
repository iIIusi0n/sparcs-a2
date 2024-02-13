package server

import (
	"api-server/config"
	"api-server/middlewares"
	"github.com/gin-gonic/gin"

	cGathering "api-server/controllers/gathering"
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
				user.GET("/:id", cUser.GetUserRouter)
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

				post.GET("/:id/like", cPost.GetPostLikesRouter)

				post.POST("/:id/like", cPost.LikePostRouter)

				post.DELETE("/:id/like", cPost.UnlikePostRouter)
			}

			gathering := v1.Group("/gathering")
			{
				gathering.Use(middlewares.JwtAuthMiddleware)

				gathering.GET("/", cGathering.GetGatheringsRouter)
				gathering.GET("/:id", cGathering.GetGatheringRouter)
				gathering.GET("/upcoming", cGathering.GetUpcomingGatheringsRouter)

				gathering.POST("/", cGathering.CreateGatheringRouter)

				gathering.PATCH("/:id", cGathering.UpdateGatheringRouter)

				gathering.GET("/:id/participant", cGathering.GetGatheringParticipantsRouter)

				gathering.POST("/:id/participant", cGathering.ParticipateGatheringRouter)

				gathering.DELETE("/:id/participant", cGathering.CancelParticipateGatheringRouter)
			}
		}
	}

	return r
}
