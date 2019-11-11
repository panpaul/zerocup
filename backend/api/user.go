package api

import (
	"github.com/gin-gonic/gin"
	"comment/models"
	"comment/serializers"
	"comment/services"
)

func Register(c *gin.Context)  {
	var service services.RegisterService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200,"ERR" + err.Error())
	} else {
		res := service.Register()
		c.JSON(200, res)
	}
}

func Login(c *gin.Context)  {
	var service services.LoginService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, err.Error())
	} else {
		res := service.Login()
		c.JSON(200, res)
	}
}

func CurrentUserInfo(c *gin.Context)  {
	user := models.CurrentUser(c)
	c.JSON(200, serializers.BuildUser(*user))
}
