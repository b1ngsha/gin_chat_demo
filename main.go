package main

import (
	"bytes"
	"gin_chat_demo/config"
	"gin_chat_demo/models"
	"gin_chat_demo/routes"
	"gin_chat_demo/views"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("json")
	if err := viper.ReadConfig(bytes.NewBuffer(config.AppConfig)); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("config file not found")
		} else {
			log.Println("read config error")
		}
		log.Fatal(err)
	}

	models.InitDB()
}

func main() {
	port := viper.GetString("app.port")
	router := routes.InitRoute()
	router.SetHTMLTemplate(views.ViewsTemplate)
	log.Println("start server ", "localhost:"+port)
	http.ListenAndServe("127.0.0.1:" + port, router)
}
