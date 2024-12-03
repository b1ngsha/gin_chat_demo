package routes

import (
	"gin_chat_demo/controller"
	"gin_chat_demo/service/session"
	"gin_chat_demo/views"
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
	}
	return router
}
