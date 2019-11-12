package serializers

import (
	"comment/models"
)

type UserSerializer struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	HeadImg  string `json:"head_img"`
}

func BuildUser(item models.User) UserSerializer {
	return UserSerializer{
		ID:       item.ID,
		UserName: item.UserName,
		HeadImg:  item.HeadImg,
	}
}

func BuildUsers(items []models.User) []UserSerializer {
	var users []UserSerializer
	for _, item := range items {
		users = append(users, BuildUser(item))
	}
	return users
}
