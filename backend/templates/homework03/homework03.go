package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;uniqueIndex;not null"`
	PostCount int    `gorm:"default:0"` // 文章数量统计
	Posts     []Post `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Post 文章模型
type Post struct {
	ID            uint      `gorm:"primaryKey"`
	Title         string    `gorm:"size:200;not null"`
	Content       string    `gorm:"type:text"`
	UserID        uint      `gorm:"not null;index"`
	User          User      `gorm:"foreignKey:UserID"`
	Comments      []Comment `gorm:"foreignKey:PostID"`
	CommentStatus string    `gorm:"size:50;default:'有评论'"` // 评论状态
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// AfterCreate 钩子函数：文章创建后自动更新用户文章数量
func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// Comment 评论模型
type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	PostID    uint   `gorm:"not null;index"`
	Post      Post   `gorm:"foreignKey:PostID"`
	UserID    uint   `gorm:"not null;index"`
	User      User   `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// AfterDelete 钩子函数：评论删除后检查文章评论数量
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var count int64
	// 注意：在删除钩子中，如果需要统计剩余数量，通常需要检查数据库
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论").Error
	}
	return nil
}

func main() {
	// 数据库连接配置 (请确保本地 MySQL 已启动并创建了 blog_system 数据库)
	// 用户名: root, 密码: password (请根据实际情况修改)
	dsn := "root:password@tcp(127.0.0.1:3306)/blog_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 题目 1：自动迁移表结构
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal("迁移表结构失败:", err)
	}
	fmt.Println("✓ 题目1：数据库表创建成功")

	// 清理旧数据（可选，方便重复运行测试）
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Comment{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Post{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})

	// 插入测试数据
	user1 := User{Name: "张三", Email: "zhangsan@example.com"}
	db.Create(&user1)

	post1 := Post{Title: "Go语言入门", Content: "这是一篇关于Go语言的文章", UserID: user1.ID}
	db.Create(&post1)

	comment1 := Comment{Content: "写得很好", PostID: post1.ID, UserID: user1.ID}
	db.Create(&comment1)

	fmt.Println("✓ 测试数据插入成功")

	// 题目 2：关联查询
	fmt.Println("\n--- 题目2：关联查询测试 ---")
	var user User
	db.Preload("Posts.Comments").First(&user, user1.ID)
	fmt.Printf("用户 %s 的文章及评论已查询完成\n", user.Name)

	// 题目 3：钩子函数测试
	fmt.Println("\n--- 题目3：钩子函数测试 ---")
	fmt.Printf("创建文章前用户文章数: %d\n", user.PostCount)

	// 再次触发 AfterCreate
	newPost := Post{Title: "新文章", UserID: user1.ID}
	db.Create(&newPost)

	var updatedUser User
	db.First(&updatedUser, user1.ID)
	fmt.Printf("创建新文章后用户文章数: %d\n", updatedUser.PostCount)

	fmt.Println("\n所有任务演示完成！")
}
