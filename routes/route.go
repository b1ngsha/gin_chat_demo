package routes

import (
	"gin_chat_demo/controller"
	"gin_chat_demo/service/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.New()

	router.StaticFS("/static", http.Dir("static"))
	r := router.Group("/", session.EnableCookieSession())
	{
		r.GET("/", controller.Index)
	}
	return router
}
