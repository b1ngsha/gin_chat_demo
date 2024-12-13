package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	wsUpgrader     = websocket.Upgrader{}
	rooms          = make(map[int][]interface{})
)

type Server struct {
}

type Client struct {
	Conn       *websocket.Conn `json:"conn"`
	RemoteAddr string          `json:"remote_addr"`
	UserId     uint            `json:"user_id"`
	Username   string          `json:"username"`
	RoomId     int             `json:"room_id"`
	AvatarId   string          `json:"avatar_id"`
}

type Message struct {
	FromUserId uint   `json:"from_user_id"`
	Username   string `json:"username"`
	AvatarId   string `json:"avatar_id"`
	ToUserId   uint   `json:"to_user_id"`
	Content    string `json:"content"`
	ImageUrl   string `json:"image_url"`
	RoomId     int    `json:"room_id"`
	Time       int64  `json:"time"`

	Conn   *websocket.Conn `json:"conn"`
	Status int             `json:"status"`
}

func (server *Server) RunWebSocket(c *gin.Context) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, _ := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	defer conn.Close()

	channel := make(chan struct{}) // channel between client and server to communicate
	go read(conn, channel)
	go write(channel)

	select {}
}

func read(conn *websocket.Conn, channel chan<- struct{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("read error:", err)
		}
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			// TODO 离线通知
			log.Println("read message error:", err)
			conn.Close()
			close(channel)
			return
		}
		log.Println("client message:", string(message), conn.RemoteAddr())

		// TODO 处理心跳响应

		// TODO maybe 加锁

		// TODO 服务端返回响应
	}
}

func write(channel <-chan struct{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("write error:", err)
		}
	}()

	for {
		select {
		case <-channel:
			return
			// TODO handle conn clients

			// TODO handle clients messages

			// TODO handle offline
		}
	}
}

func (server *Server) GetOnlineUserCountByRoom(roomId int) int {
	return len(rooms[roomId])
}
