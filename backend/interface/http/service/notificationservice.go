package service

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/gameiro/gila-test/core/domain"
	"gitlab.com/gameiro/gila-test/core/domain/usecase"
	"gitlab.com/gameiro/gila-test/core/dto"
	"gitlab.com/gameiro/gila-test/logging"
)

type NotificationService struct {
}

func (service NotificationService) Send(c *gin.Context) []*domain.Notification {
	message, err := dto.GetMessageFromRequest(c)
	if err != nil {
		log.Println("Notification Send: could not get message from request")
		return nil
	}
	var usecase usecase.NotificationUseCase
	return usecase.SendNotification(message)
}

func (service NotificationService) Log(logger *logging.Logger, notifications []*domain.Notification) {
	for _, n := range notifications {
		logJson, _ := json.Marshal(n)
		logger.Push("[" + time.Now().Format("2006-01-02 15:04:05") + "][CATEGORY: " + n.Category.Name + "][USER: " + n.User.ID + "][CHANNEL:" + n.Channel.Name + "]: " + string(logJson))
	}
}
