package websocket

var (
	rooms = make(map[int][]interface{})
)

type Server struct {
}

func (server *Server) GetOnlineUserCountByRoom(roomId int) int {
	return len(rooms[roomId])
}
