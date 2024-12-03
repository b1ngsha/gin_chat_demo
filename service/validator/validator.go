package validator

type User struct {
	Username string `binding:"required,min=2"`
	Password string `binding:"required,min=6"`
	AvatarId string `binding:"required,numeric"`
}
