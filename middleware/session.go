package middleware

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/sessions"
import "github.com/gin-contrib/sessions/cookie"

func Session(secret string) gin.HandlerFunc{

	store := cookie.NewStore([]byte(secret))

	store.Options(sessions.Options{
		HttpOnly:true,
		MaxAge: 7 * 86400,
		Path: "/",
	})

	return sessions.Sessions("gin-session", store)
}
