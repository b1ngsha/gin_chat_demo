package main

import (
	"bytes"
	"gin_chat_demo/config"
	"gin_chat_demo/routes"
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
}

func main() {
	port := viper.GetString("app.port")
	router := routes.InitRoute()
	log.Println("start server ", "localhost:"+port)
	http.ListenAndServe("127.0.0.1:" + port, router)
}
