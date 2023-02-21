package usecase

import (
	"gitlab.com/gameiro/gila-test/core/domain"
	"gitlab.com/gameiro/gila-test/core/repository"
)

type CategoryUseCase struct {
}

func (usecase CategoryUseCase) FetchAll() []*domain.Category {
	var c repository.CategoryRepository
	return c.FetchAll()
}
