package api

import (
	"comment/service"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
    var service service.CreateCommentService
    if err := c.ShouldBind(&service); err != nil {
        c.JSON(400, err.Error())
    } else {
        // 从前端传过来的headers的token解析出当前用户
        user := CurrentUser(c)
        res := service.Create(user)
        c.JSON(200, res)
    }
}
