package controller

import (
	"gin_chat_demo/service/session"
	"gin_chat_demo/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Login(c *gin.Context) {
	user.Login(c)
}

func Home(c *gin.Context) {
	userInfo := session.GetUserInfo(c)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"user_info": userInfo,
	})
}