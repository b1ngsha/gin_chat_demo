package websocket

var WsServer = &Server{}

func GetOnlineUserCountByRoom(roomId int) int {
	return WsServer.GetOnlineUserCountByRoom(roomId)
}