package websocket

import "github.com/gin-gonic/gin"

var WsServer = &Server{}

func GetOnlineUserCountByRoom(roomId int) int {
	return WsServer.GetOnlineUserCountByRoom(roomId)
}

func Start(c *gin.Context) {
	WsServer.RunWebSocket(c)
}