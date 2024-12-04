package controller

import (
	"gin_chat_demo/service/session"
	"gin_chat_demo/service/user"
	"gin_chat_demo/websocket"
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
	rooms := []map[string]interface{}{
		{"id": 1, "num": websocket.GetOnlineUserCountByRoom(1)},
		{"id": 2, "num": websocket.GetOnlineUserCountByRoom(2)},
		{"id": 3, "num": websocket.GetOnlineUserCountByRoom(3)},
		{"id": 4, "num": websocket.GetOnlineUserCountByRoom(4)},
		{"id": 5, "num": websocket.GetOnlineUserCountByRoom(5)},
		{"id": 6, "num": websocket.GetOnlineUserCountByRoom(6)},
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"rooms":     rooms,
		"user_info": userInfo,
	})
}
