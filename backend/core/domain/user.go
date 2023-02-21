package domain

type User struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	Phone      string      `json:"phone"`
	Categories []*Category `json:"-"`
	Channels   []*Channel  `json:"-"`
}

type UserUseCaseContract interface {
	Notify(*User, string)
}
