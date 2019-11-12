# comment

comment for the website.

> based on Singo
> (https://github.com/bydmm/singo)

## godotenv
项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)
```
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local"# Mysql连接地址
REDIS_ADDR="127.0.0.1:6379"# Redis端口和地址
REDIS_PW=""# Redis连接密码
REDIS_DB=""# Redis库从0到10
GIN_MODE="debug"
LOG_LEVEL="debug"
JWT_SECRET_KEY="secretkey"# JWT密钥，必须设置而且不要泄露
```
## run server
本项目使用Go Mod管理依赖。
```bash
go mod init comment
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装
```
