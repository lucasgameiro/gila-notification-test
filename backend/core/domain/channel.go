package domain

import "gitlab.com/gameiro/gila-test/senders"

type Channel struct {
	ID     string                 `json:"id"`
	Name   string                 `json:"name"`
	SendTo senders.SenderContract `json:"-"`
}

type ChannelRepositoryContract interface {
	Get(string) *Channel
}
