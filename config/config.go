package config

var AppConfig = []byte(`
{
	"app": {
		"port": "8080"
	},
	"mysql": {
		"dsn": "root:@tcp(127.0.0.1:3306)/gin_chat_demo?charset=utf8mb4&parseTime=True&loc=Local"
	}
}
`)