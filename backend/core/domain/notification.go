package domain

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gameiro/gila-test/core/dto"
)

type Notification struct {
	Message  string    `json:"message"`
	Category *Category `json:"category"`
	Channel  *Channel  `json:"-"`
	User     *User     `json:"user"`
}

type NotificationUseCaseContract interface {
	SendNotification(*dto.MessagePosted) []*Notification
}

type NotificationServiceContract interface {
	Send(c *gin.Context) []*Notification
}
