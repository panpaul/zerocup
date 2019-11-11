package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{},&Article{},&Comment{})
	DB.Model(&Comment{}).AddForeignKey("user_id", "users(id)",
		"CASCADE", "CASCADE")
	DB.Model(&Comment{}).AddForeignKey("reply_to_id", "users(id)",
		"CASCADE", "CASCADE")
	DB.Model(&Comment{}).AddForeignKey("Article_id", "Articles(id)",
		"CASCADE", "CASCADE")
	DB.Model(&Comment{}).AddForeignKey("parent_id", "comments(id)",
		"CASCADE", "CASCADE")
	DB.Model(&Comment{}).AddForeignKey("root_id", "comments(id)",
		"CASCADE", "CASCADE")
}
