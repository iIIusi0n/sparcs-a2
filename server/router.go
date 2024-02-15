package server

import (
	"api-server/config"
	"api-server/middlewares"
	"github.com/gin-gonic/gin"

	cChat "api-server/controllers/chat"
	cData "api-server/controllers/data"
	cHospital "api-server/controllers/hospital"
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
					debug.GET("/dropalldata", cData.DropAllDataRouter)
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

				user.POST("/", cUser.CreateUserRouter)

				user.PATCH("/", cUser.UpdateUserRouter)
			}

			hospital := v1.Group("/hospital")
			{
				hospital.Use(middlewares.JwtAuthMiddleware)

				hospital.GET("/", cHospital.GetHospitalsRouter)
				hospital.GET("/:id", cHospital.GetHospitalRouter)
				hospital.GET("/:id/waiting", cHospital.GetHospitalWaitingNumberRouter)

				hospital.POST("/", cHospital.CreateHospitalRouter)
				hospital.POST("/:id/waiting", cHospital.CreateHospitalWaitingNumberRouter)

				hospital.PATCH("/:id", cHospital.UpdateHospitalRouter)
			}

			chat := v1.Group("/chat")
			{
				chat.Use(middlewares.JwtAuthMiddleware)

				chat.GET("/room", cChat.GetChatRoomsRouter)
				chat.GET("/room/:id", cChat.GetChatRoomRouter)

				chat.POST("/room", cChat.CreateChatRoomRouter)

				chat.GET("/:room_id", cChat.GetChatMessagesRouter)

				chat.POST("/:room_id", cChat.CreateChatMessageRouter)
			}
		}
	}

	return r
}
