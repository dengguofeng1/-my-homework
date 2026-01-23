package database

import (
	"fmt"
	"homework04/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	// 优先尝试连接 MySQL
	dsn := "root:password@tcp(127.0.0.1:3306)/blog_system?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("无法连接到 MySQL: %v\n正在回退到 SQLite...\n", err)
		DB, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect SQLite: %v", err)
		}
		fmt.Println("✓ 已使用 SQLite (blog.db)")
	} else {
		fmt.Println("✓ 已成功连接到 MySQL")
	}

	// 自动迁移模型
	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
