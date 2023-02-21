package domain

import (
	"github.com/gin-gonic/gin"
)

type CategoryContract interface {
	GetID() string
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CategoryRepositoryContract interface {
	FetchAll() []*Category
	Get(string) *Category
	GetSubscribers(*Category) []*User
}

type CategoryUseCaseContract interface {
	FetchAll() []*Category
}

type CategoryServiceContract interface {
	FetchAll(c *gin.Context) []*Category
}

func (c Category) GetID() string {
	return c.ID
}
