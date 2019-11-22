package config

import (
	"github.com/gin-gonic/gin"
	//"comment/cache"
	"comment/models"
	"comment/util"
	"os"

	"github.com/joho/godotenv"
)

var CurrentTime string

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	CurrentTime = "2006-01-02 15:04:05"
	godotenv.Load()
	gin.SetMode(os.Getenv("GIN_MODE"))

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	models.Database(os.Getenv("MYSQL_DSN"))
	//cache.Redis()
}
