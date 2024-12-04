package session

import (
	"gin_chat_demo/models"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(viper.GetString("app.cookie_key")))
	return sessions.Sessions("gin_chat_demo", store)
}

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userId := session.Get("user_id")
		if userId == nil {
			ctx.Redirect(http.StatusFound, "/")
			return
		}
		ctx.Next()
		return
	}
}

func SaveAuthSession(c *gin.Context, userId string) {
	session := sessions.Default(c)
	session.Set("user_id", userId)
	session.Save()
}

func GetUserInfo(c *gin.Context) map[string]interface{} {
	session := sessions.Default(c)
	userId := session.Get("user_id")
	userInfo := make(map[string]interface{})
	if userId != nil {
		userIdStr := userId.(string)
		userId, _ := strconv.Atoi(userIdStr)
		user := models.FindUserById(uint(userId))
		userInfo["user_id"] = user.ID
		userInfo["username"] = user.Username
		userInfo["avatar_id"] = user.AvatarId
	}
	return userInfo
}
