package service

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gameiro/gila-test/core/domain"
	"gitlab.com/gameiro/gila-test/core/domain/usecase"
)

type CategoryService struct {
}

func (service CategoryService) FetchAll(c *gin.Context) []*domain.Category {
	var usecase usecase.CategoryUseCase
	return usecase.FetchAll()
}
