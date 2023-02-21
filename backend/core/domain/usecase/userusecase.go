package usecase

import (
	"gitlab.com/gameiro/gila-test/core/domain"
	"gitlab.com/gameiro/gila-test/core/dto"
	"gitlab.com/gameiro/gila-test/core/repository"
)

type UserUseCase struct {
}

func (usecase UserUseCase) Notify(user *domain.User, message *dto.MessagePosted) []*domain.Notification {
	var categoriesRepo repository.CategoryRepository
	var notifications []*domain.Notification
	for _, channel := range user.Channels {
		att := user.Phone
		if channel.Name == "E-Mail" {
			att = user.Email
		}
		go channel.SendTo.Dispatch(att, message.Message)
		notifications = append(notifications, &domain.Notification{
			User: user, Channel: channel, Category: categoriesRepo.Get(message.CategoryId), Message: message.Message,
		})
	}
	return notifications
}
