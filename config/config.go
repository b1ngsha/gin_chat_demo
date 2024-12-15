package config

var AppConfig = []byte(`
{
	"app": {
		"port": "8080",
		"cookie_key": "Q9nImSw20ecdeoGbTm64y+8Cex95RHprhPgwlQkS2MU="
	},
	"mysql": {
		"dsn": "root:root@tcp(127.0.0.1:3306)/gin_chat_demo?charset=utf8mb4&parseTime=True&loc=Local"
	}
}
`)