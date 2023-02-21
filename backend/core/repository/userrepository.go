package repository

import (
	"gitlab.com/gameiro/gila-test/core/domain"
)

type UserRepository struct {
}

func (u UserRepository) FetchAll() []*domain.User {
	var catRepo CategoryRepository
	var chanRepo ChannelRepository
	cat1 := catRepo.Get("1")
	cat2 := catRepo.Get("2")
	cat3 := catRepo.Get("3")
	chan1 := chanRepo.Get("1")
	chan2 := chanRepo.Get("2")
	chan3 := chanRepo.Get("3")
	user1 := domain.User{
		ID: "1", Name: "John Doe", Email: "john.doe@example.org", Phone: "+1 (555)999-9999",
		Categories: []*domain.Category{cat1, cat2},
		Channels:   []*domain.Channel{chan1},
	}
	user2 := domain.User{
		ID: "2", Name: "Martha S Miller", Email: "miller_martha@example.org", Phone: "+1 (555)999-9999",
		Categories: []*domain.Category{cat2, cat3},
		Channels:   []*domain.Channel{chan2},
	}
	user3 := domain.User{
		ID: "3", Name: "Mark Moses", Email: "iammark@example.org", Phone: "+1 (555)999-9999",
		Categories: []*domain.Category{cat1, cat3},
		Channels:   []*domain.Channel{chan3},
	}
	return []*domain.User{&user1, &user2, &user3}
}
