package repository

import (
	"log"

	"gitlab.com/gameiro/gila-test/core/domain"
	"golang.org/x/exp/slices"
)

type CategoryRepository struct {
}

func (c CategoryRepository) FetchAll() []*domain.Category {
	cat1 := domain.Category{ID: "1", Name: "Sports"}
	cat2 := domain.Category{ID: "2", Name: "Finance"}
	cat3 := domain.Category{ID: "3", Name: "Movies"}
	return []*domain.Category{&cat1, &cat2, &cat3}
}

func (c CategoryRepository) Get(id string) *domain.Category {
	categories := c.FetchAll()
	idx := slices.IndexFunc(categories, func(c *domain.Category) bool { return c.ID == id })
	if idx == -1 {
		log.Println("Get category: could not find category")
		return nil
	}
	return categories[idx]
}

func (c CategoryRepository) GetSubscribers(category *domain.Category) []*domain.User {
	var u UserRepository
	var users []*domain.User
	for _, user := range u.FetchAll() {
		for _, userCat := range user.Categories {
			if userCat.GetID() == category.GetID() {
				users = append(users, user)
			}
		}
	}
	return users
}
