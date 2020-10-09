package middleware


import "github.com/gin-contrib/cors"
import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc{
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin","Content-Length", "Content-Type", "Cookie"}
	config.AllowOrigins = []string{"http://localhost:8000"}
	config.AllowCredentials = true


	return cors.New(config)


}
