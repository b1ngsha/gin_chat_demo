package models

import "gorm.io/gorm"

const DefaultPageLimit = 100

type Message struct {
	gorm.Model
	FromUserId uint
	ToUserId   uint
	RoomId     int
	Content    string
	ImageUrl   string
}

func GetRoomMsg(roomId int) []map[string]interface{} {
	var results []map[string]interface{}
	ChatDB.Model(&Message{}).
		Select("messages.*, users.username, users.avatar_id").
		Joins("INNER join users on users.id = messages.from_user_id").
		Where("messages.room_id = ?", roomId).
		Where("messages.to_user_id = -1").
		Order("messages.id desc").
		Limit(DefaultPageLimit).
		Scan(&results)
	return results
}
