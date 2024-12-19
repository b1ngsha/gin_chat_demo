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

func GetLimitedRoomMsg(roomId int, offset int) []map[string]interface{} {
	var results []map[string]interface{}
	ChatDB.Model(&Message{}).
		Select("messages.*, users.username, users.avatar_id").
		Joins("INNER join users on users.id = messages.from_user_id").
		Where("messages.room_id = ?", roomId).
		Where("messages.to_user_id = -1").
		Order("messages.id desc").
		Offset(offset).
		Limit(DefaultPageLimit).
		Scan(&results)
	return results
}

func GetLimitedPrivateMsg(fromUserId, toUserId string, offset int) []map[string]interface{} {
	var results []map[string]interface{}
	ChatDB.Model(&Message{}).
		Select("messages.*, users.username, users.avatar_id").
		Joins("INNER join users on users.id = messages.from_user_id").
		Where("messages.from_user_id = ? and messages.to_user_id = ?", fromUserId, toUserId).
		Or("messages.from_user_id = ? and messages.to_user_id = ?", toUserId, fromUserId).
		Order("messages.id desc").
		Offset(offset).
		Limit(DefaultPageLimit).
		Scan(&results)
	return results
}
