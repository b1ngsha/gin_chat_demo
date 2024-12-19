package controller

import (
	"gin_chat_demo/models"
	"gin_chat_demo/service/session"
	"gin_chat_demo/service/user"
	"gin_chat_demo/websocket"
	"net/http"
	"strconv"

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

func Room(c *gin.Context) {
	roomIdStr := c.Param("room_id")
	roomId, _ := strconv.Atoi(roomIdStr)
	userInfo := session.GetUserInfo(c)
	msgList := models.GetLimitedRoomMsg(roomId, 0)
	c.HTML(http.StatusOK, "room.html", gin.H{
		"room_id":   roomId,
		"msg_list":  msgList,
		"msg_count": len(msgList),
		"user_info": userInfo,
	})
}

func PrivateChat(c *gin.Context) {
	roomId := c.Query("room_id")
	toUserId := c.Query("user_id")
	userInfo := session.GetUserInfo(c)
	fromUserId := strconv.Itoa(int(userInfo["user_id"].(uint)))
	msgList := models.GetLimitedPrivateMsg(fromUserId, toUserId, 0)
	c.HTML(http.StatusOK, "private_chat.html", gin.H{
		"room_id":   roomId,
		"user_info": userInfo,
		"msg_list":  msgList,
	})
}

func Pagination(c *gin.Context) {
	roomId := c.Query("room_id")
	toUserId := c.Query("to_user_id")
	offset := c.Query("offset")

	offsetInt, err := strconv.Atoi(offset)
	if err != nil || offsetInt <= 0 {
		offsetInt = 0
	}

	msgList := []map[string]interface{}{}
	if toUserId != "" {
		userInfo := session.GetUserInfo(c)
		userIdStr := strconv.Itoa(int(userInfo["user_id"].(uint)))
		msgList = models.GetLimitedPrivateMsg(userIdStr, toUserId, offsetInt)
	} else {
		roomIdInt, _ := strconv.Atoi(roomId)
		msgList = models.GetLimitedRoomMsg(roomIdInt, offsetInt)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": map[string]interface{}{
			"list": msgList,
		},
	})
}
