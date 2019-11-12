package middleware

import (
	"github.com/gin-gonic/gin"
	"comment/models"
	"comment/serializers"
	"comment/util"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := http.StatusOK

		token := c.GetHeader("token")
		if token == "" {
			code = http.StatusBadRequest
		}

		claims, err := util.ParseToken(token)
		if err != nil {
			code = http.StatusForbidden
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = http.StatusForbidden
		} else {
			var user models.User
			if err := models.DB.Find(&user, claims.UserId).Error; err != nil {
				code = http.StatusForbidden
			} else {
				c.Set("CurrentUser", &user)
			}
		}

		if code != http.StatusOK {
			c.JSON(200, serializers.Response{
				Status:code,
				Message:"forbidden",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

