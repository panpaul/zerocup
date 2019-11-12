package api

import (
	"github.com/gin-gonic/gin"
	"comment/models"
	"comment/services"
	"strconv"
)

func CreateComment(c *gin.Context) {
	var service services.CreateCommentService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(400, err.Error())
	} else {
		user := models.CurrentUser(c)
		res := service.Create(user)
		c.JSON(200, res)
	}
}

func CommentList(c *gin.Context) {
	var service services.CommentListService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(400, err.Error())
	} else {
		articleID, _ := strconv.Atoi(c.Param("article_id"))
		user := models.CurrentUser(c)
		p := c.DefaultQuery("p", "1")
		res := service.List(uint(articleID), user, p)
		c.JSON(200, res)
	}
}
