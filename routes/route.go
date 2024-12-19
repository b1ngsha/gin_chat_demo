package routes

import (
	"gin_chat_demo/controller"
	"gin_chat_demo/service/session"
	"gin_chat_demo/views"
	"gin_chat_demo/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.New()
	router.SetHTMLTemplate(views.ViewsTemplate)

	router.StaticFS("/static", http.Dir("static"))
	r := router.Group("/", session.EnableCookieSession())
	{
		r.GET("/", controller.Index)
		r.POST("/login", controller.Login)
		r.GET("/ws", websocket.Start)

		authorized := r.Group("/", session.SessionAuthMiddleware())
		{
			authorized.GET("/home", controller.Home)
			authorized.GET("/room/:room_id", controller.Room)
			authorized.GET("/private_chat", controller.PrivateChat)
			authorized.GET("/pagination", controller.Pagination)
		}
	}
	return router
}