package senders

import "log"

type SenderContract interface {
	Dispatch(string, string) int
}

type SMSSender struct {
	Name string
}

func (s SMSSender) Dispatch(phone string, message string) int {
	log.Output(1, "[SMS][NUMBER:"+phone+"][MESSAGE:"+message+"]")
	return 0
}

type EmailSender struct {
	Name string
}

func (s EmailSender) Dispatch(email string, message string) int {
	log.Output(1, "[EMAIL][EMAIL:"+email+"][MESSAGE:"+message+"]")
	return 0
}

type PushNotificationSender struct {
	Name string
}

func (s PushNotificationSender) Dispatch(phone string, message string) int {
	log.Output(1, "[PUSH_NOTIFICATION][NUMBER:"+phone+"][MESSAGE:"+message+"]")
	return 0
}
