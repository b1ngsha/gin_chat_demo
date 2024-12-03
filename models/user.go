package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	AvatarId string `json:"avatar_id"`
}

func CreateUser(params interface{}) User {
	var user User
	user.Username = params.(map[string]interface{})["username"].(string)
	user.Password = params.(map[string]interface{})["password"].(string)
	user.AvatarId = params.(map[string]interface{})["avatar_id"].(string)
	ChatDB.Create(&user)
	return user
}

func FindUserById(id uint) User {
	var user User
	ChatDB.First(&user, id)
	return user
}

func FindUserByName(name string) User {
	var user User
	ChatDB.Where("username = ?", name).First(&user)
	return user
}

func SaveAvatarId(AvatarId string, u User) User {
	u.AvatarId = AvatarId
	ChatDB.Save(&u)
	return u
}