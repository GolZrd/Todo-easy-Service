package handler

import (
	"todoservice/internal/model"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateTask(c *gin.Context) {

	// Структура задачи
	var input model.Task

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "internal server error",
		})
	}

	// Создаем задачу

	c.JSON(200, gin.H{
		"message": "Create task",
	})

}

func (h *Handler) Tasks(c *gin.Context) {

	// Получаем список задач
	c.JSON(200, gin.H{
		"message": "Get tasks",
	})
}

func (h *Handler) DeleteTask(c *gin.Context) {

	// Удаляем задачу
	c.JSON(200, gin.H{
		"message": "Delete tasks",
	})
}

func (h *Handler) UpdateTask(c *gin.Context) {

	// Обновляем статус задачи
	c.JSON(200, gin.H{
		"message": "Update task",
	})
}
