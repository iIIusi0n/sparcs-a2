package config

const (
	ServerName = "localhost"
	ServerHost = "0.0.0.0"
	ServerPort = "8080"

	ServerDebug = true // TODO: Change this to false when deploying

	ServerSecret = "REDACTED" // TODO: Change this to a more secure secret

	ServerLog = "server.log"
	GinLog    = "gin.log"
)
