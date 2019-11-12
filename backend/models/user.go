package models

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	UserName string `gorm:"unique"`
	PassWord string
	// TODO: get user head_img
	HeadImg string `sql:"default:'http://localhost:8888/static/base/gopher01.png'"`
}

const (
	passwordCost = 10
)

func (user *User) SetPassWord(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
	if err != nil {
		return err
	}
	user.PassWord = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password))
	return err == nil
}

func CurrentUser(c *gin.Context) *User {
	if value, exists := c.Get("CurrentUser"); exists == false {
		return nil
	} else {
		if user, ok := value.(*User); !ok {
			return nil
		} else {
			return user
		}
	}
}
