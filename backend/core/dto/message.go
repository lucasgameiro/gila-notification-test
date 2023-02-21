package dto

import (
	"github.com/gin-gonic/gin"
)

type MessagePosted struct {
	CategoryId string `form:"category_id" binding:"required"`
	Message    string `form:"message" binding:"required"`
}

func GetMessageFromRequest(c *gin.Context) (*MessagePosted, error) {
	var message MessagePosted
	err := c.Bind(&message)
	return &message, err
}
