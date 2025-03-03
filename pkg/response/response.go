package response

import "github.com/gin-gonic/gin"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(200, Response{Message: message, Data: data})
}

func Error(c *gin.Context, message string, err error) {
	c.JSON(400, Response{Message: message, Error: err.Error()})
}
