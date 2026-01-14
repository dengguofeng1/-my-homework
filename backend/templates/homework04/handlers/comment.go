package handlers

import (
	"homework04/database"
	"homework04/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var input struct {
		Content string `json:"content" binding:"required"`
		PostID  uint   `json:"post_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")

	// Check if post exists
	var post models.Post
	if err := database.DB.First(&post, input.PostID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	comment := models.Comment{
		Content: input.Content,
		UserID:  userID.(uint),
		PostID:  input.PostID,
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

func GetCommentsByPost(c *gin.Context) {
	postID := c.Param("post_id")

	var comments []models.Comment
	if err := database.DB.Preload("User").Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
