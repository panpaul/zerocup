package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)
const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

type User struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Nickname       string
	Status         string
	Avatar         string `gorm:"size:1000"`
}

type Article struct {
	ID        uint
	CreAt time.Time
	UpdAt time.Time
	Art   string
	ArtID uint  
}
// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
// User 用户模型
type Comment struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	User      User       `gorm:"foreignkey:UserID;association_foreignkey:ID"`
	UserID    uint       `gorm:"not null"`
	ReplyTo   User       `gorm:"foreignkey:ReplyToID;association_foreignkey:ID"`
	ReplyToID uint       `sql:"default:null"`
	Article   string     `gorm:"foreignkey:ArticleID;association_foreignkey:ID"`
	//Article   Article    `gorm:"foreignkey:ArticleID;association_foreignkey:ID"`
	ArticleID uint      `gorm:"not null"`
	Parent    *Comment  `gorm:"foreignkey:ParentID;association_foreignkey:ID;"`
	ParentID  uint      `sql:"default:null"`
	Root      *Comment  `gorm:"foreignkey:RootID;association_foreignkey:ID;"`
	RootID    uint      `sql:"default:null"`
	Content   string    `gorm:"not null"`
	Replys    []Comment `sql:"default:null"`
}
