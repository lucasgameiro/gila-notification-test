package usecase

import (
	"gitlab.com/gameiro/gila-test/core/domain"
	"gitlab.com/gameiro/gila-test/core/dto"
	"gitlab.com/gameiro/gila-test/core/repository"
)

type NotificationUseCase struct {
}

func (usecase NotificationUseCase) SendNotification(m *dto.MessagePosted) []*domain.Notification {
	var categoriesRepo repository.CategoryRepository
	category := categoriesRepo.Get(m.CategoryId)
	var userUC UserUseCase
	var r []*domain.Notification

	if category != nil {
		users := categoriesRepo.GetSubscribers(category)
		for _, user := range users {
			notifications := userUC.Notify(user, m)
			r = append(r, notifications...)
		}
		return r
	}
	return nil
}
