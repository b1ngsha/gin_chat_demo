package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(viper.GetString("app.cookie_key")))
	return sessions.Sessions("gin_chat_demo", store)
}

func SaveAuthSession(c *gin.Context, userId string) {
	session := sessions.Default(c)
	session.Set("user_id", userId)
	session.Save()
}