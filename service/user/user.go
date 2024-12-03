package user

import (
	"gin_chat_demo/models"
	"gin_chat_demo/service/session"
	"gin_chat_demo/service/validator"
	"gin_chat_demo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	avatarId := c.PostForm("avatar_id")

	var u validator.User
	u.Username = username
	u.Password = password
	u.AvatarId = avatarId
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	user := models.FindUserByName(username)
	encryptPwd := utils.Md5Encrypt(password)
	if user.ID == 0 {
		// user not exists
		user = models.CreateUser(map[string]interface{}{
			"username":  username,
			"password":  encryptPwd,
			"avatar_id": avatarId,
		})
	} else {
		// user exists
		// check password
		if user.Password != encryptPwd {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "password error",
			})
			return
		}

		// update avatar id
		models.SaveAvatarId(avatarId, user)
	}

	// save auth info into session
	session.SaveAuthSession(c, string(strconv.Itoa(int(user.ID))))
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "login success",
	})
}
