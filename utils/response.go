package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Time    string      `json:"time,omitempty"`
}

func RespondJSON(c *gin.Context, statusCode int, status string, message string, data interface{}) {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
		Time:    time.Now().Format(time.RFC3339),
	}
	c.JSON(statusCode, response)
}

func RespondError(c *gin.Context, statusCode int, message string, err error) {
	if err != nil {
		c.JSON(statusCode, gin.H{"status": "error", "message": message, "error": err.Error()})
	} else {
		c.JSON(statusCode, gin.H{"status": "error", "message": message})
	}
}
