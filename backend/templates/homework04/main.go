package main

import (
	"homework04/database"
	"homework04/handlers"
	"homework04/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()
	// 配置信任的代理（开发环境可以为空，生产环境需配置实际代理IP）
	r.SetTrustedProxies(nil)

	// Welcome route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Blog API",
			"version": "1.0",
			"endpoints": gin.H{
				"auth":     "/register, /login",
				"posts":    "/posts",
				"comments": "/comments",
			},
		})
	})

	// Public routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Post routes
	posts := r.Group("/posts")
	{
		posts.GET("", handlers.GetPosts)
		posts.GET("/:id", handlers.GetPost)

		// Protected routes
		protected := posts.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("", handlers.CreatePost)
			protected.PUT("/:id", handlers.UpdatePost)
			protected.DELETE("/:id", handlers.DeletePost)
		}
	}

	// Comment routes
	comments := r.Group("/comments")
	comments.Use(middleware.AuthMiddleware())
	{
		comments.POST("", handlers.CreateComment)
		comments.GET("/post/:post_id", handlers.GetCommentsByPost)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
