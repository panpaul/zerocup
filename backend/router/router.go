package router

import (
	"github.com/gin-gonic/gin"
	"comment/api"
	"comment/middleware"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Cors())

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)

		v1.POST("register", api.Register)
		v1.POST("login", api.Login)


		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("comment", api.CreateComment)
			authed.GET("comment/:article_id", api.CommentList)

			authed.GET("me/info", api.CurrentUserInfo)
		}
	}

	r.StaticFS("./static", http.Dir("./static"))
	return r
}
