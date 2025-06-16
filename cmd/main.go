package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Gin Framework",
		})
	})

	// POST /create - создаем заметку
	router.POST("/create", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Create Node",
		})
	})

	// GET /notes - получаем список заметок
	router.GET("/notes", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get notes",
		})
	})

	// DELETE /notes/:id - удаляем заметку по id
	router.DELETE("/notes/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Delete note",
		})
	})
	// PUT /notes/{id} - обновляем статус заметки
	router.PUT("/notes/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update note",
		})
	})

	router.Run(":8080")
}
