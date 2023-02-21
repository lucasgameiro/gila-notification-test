package repository

import (
	"log"

	"gitlab.com/gameiro/gila-test/core/domain"
	"gitlab.com/gameiro/gila-test/senders"
	"golang.org/x/exp/slices"
)

type ChannelRepository struct {
}

func (c ChannelRepository) FetchAll() []*domain.Channel {
	var mSMSSender senders.SMSSender
	var mEmailSender senders.EmailSender
	var mPushNotificationSender senders.PushNotificationSender
	c1 := domain.Channel{ID: "1", Name: "SMS", SendTo: &mSMSSender}
	c2 := domain.Channel{ID: "2", Name: "E-Mail", SendTo: &mEmailSender}
	c3 := domain.Channel{ID: "3", Name: "Push Notification", SendTo: &mPushNotificationSender}
	return []*domain.Channel{&c1, &c2, &c3}
}

func (c ChannelRepository) Get(id string) *domain.Channel {
	channels := c.FetchAll()
	idx := slices.IndexFunc(channels, func(c *domain.Channel) bool { return c.ID == id })
	if idx == -1 {
		log.Println("Get Channel: could not find channel")
		return nil
	}
	return channels[idx]
}
