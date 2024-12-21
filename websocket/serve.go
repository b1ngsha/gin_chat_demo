package websocket

import (
	"encoding/json"
	"gin_chat_demo/models"
	"gin_chat_demo/utils"
	"log"
	"net/http"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jianfengye/collection"
)

var (
	wsUpgrader = websocket.Upgrader{}
	rooms      = make(map[int][]interface{})
	clientMsg  = Message{}
	enterRooms = make(chan Client)
	chNotify   = make(chan int, 1)
	chMsg      = make(chan Message)
)

const (
	msgTypeOnline      = 1
	msgTypeSend        = 3
	msgTypeGetUserList = 4
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
	FromUserId uint          `json:"from_user_id"`
	Username   string        `json:"username"`
	AvatarId   string        `json:"avatar_id"`
	ToUserId   uint          `json:"to_user_id"`
	Content    string        `json:"content"`
	ImageUrl   string        `json:"image_url"`
	RoomId     int           `json:"room_id"`
	Time       int64         `json:"time"`
	Count      int           `json:"count"`
	List       []interface{} `json:"list"`

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
		json.Unmarshal(message, &clientMsg)

		// TODO 处理心跳响应

		// TODO maybe 加锁

		// Handle client online msg
		if clientMsg.Status == msgTypeOnline {
			enterRooms <- Client{
				Conn:       conn,
				RemoteAddr: conn.RemoteAddr().String(),
				UserId:     clientMsg.FromUserId,
				Username:   clientMsg.Username,
				RoomId:     clientMsg.RoomId,
				AvatarId:   clientMsg.AvatarId,
			}
		}
		serverMsgObject := getServerMsgObject(clientMsg.Status, conn)
		chMsg <- serverMsgObject
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
		case client := <-enterRooms:
			handleClientConn(client.Conn)
		case msg := <-chMsg:
			serverMsgStr, _ := json.Marshal(msg)
			switch msg.Status {
			case msgTypeOnline, msgTypeSend:
				notify(msg.Conn, string(serverMsgStr))
			case msgTypeGetUserList:
				chNotify <- 1
				msg.Conn.WriteMessage(websocket.TextMessage, serverMsgStr)
				<-chNotify
			}

			// TODO handle offline
		}
	}
}

/** Remove old client object in rooms slice and add new one **/
func handleClientConn(conn *websocket.Conn) {
	objCollection := collection.NewObjCollection(rooms[clientMsg.RoomId])
	retCollection := utils.Safety.ThreadSafetyDo(func() interface{} {
		return objCollection.Reject(func(item interface{}, key int) bool {
			// remove former client object in rooms slice
			if item.(Client).UserId == clientMsg.FromUserId {
				// use channel to sync, ensure the former client object has been removed
				chNotify <- 1
				item.(Client).Conn.WriteMessage(websocket.TextMessage, []byte(`{"status": -1, "data": []}`))
				<-chNotify
				return true
			}
			return false
		})
	}).(collection.ICollection)

	retCollection = utils.Safety.ThreadSafetyDo(func() interface{} {
		// add a new one
		return retCollection.Append(Client{
			Conn:       conn,
			RemoteAddr: conn.RemoteAddr().String(),
			UserId:     clientMsg.FromUserId,
			Username:   clientMsg.Username,
			RoomId:     clientMsg.RoomId,
			AvatarId:   clientMsg.AvatarId,
		})
	}).(collection.ICollection)

	interfaces, _ := retCollection.ToInterfaces()
	rooms[clientMsg.RoomId] = interfaces
}

func (server *Server) GetOnlineUserCountByRoom(roomId int) int {
	return len(rooms[roomId])
}

func getServerMsgObject(status int, conn *websocket.Conn) Message {
	message := Message{
		Status:     status,
		Username:   clientMsg.Username,
		Conn:       conn,
		RoomId:     clientMsg.RoomId,
		FromUserId: clientMsg.FromUserId,
		Time:       time.Now().UnixNano() / 1e6,
	}

	if status == msgTypeSend {
		message.AvatarId = clientMsg.AvatarId
		message.Content = clientMsg.Content

		// save message to database
		messageEntity := models.Message{}
		content := clientMsg.Content
		if utf8.RuneCountInString(content) > 800 {
			messageEntity.Content = string([]rune(content)[:800])
		}
		messageEntity.FromUserId = clientMsg.FromUserId
		messageEntity.RoomId = clientMsg.RoomId
		messageEntity.ToUserId = clientMsg.ToUserId
		models.ChatDB.Create(&messageEntity)
	}

	if status == msgTypeGetUserList {
		userList := rooms[clientMsg.RoomId]
		message.Count = len(userList)
		message.List = userList
	}

	return message
}

func notify(conn *websocket.Conn, msg string) {
	chNotify <- 1
	room := rooms[clientMsg.RoomId]
	for _, client := range room {
		if client.(Client).RemoteAddr == conn.RemoteAddr().String() {
			client.(Client).Conn.WriteMessage(websocket.TextMessage, []byte(msg))
		}
	}
	<-chNotify
}
