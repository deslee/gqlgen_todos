package models

type Todo struct {
	ID string
	Text string
	UserID string
}

type User struct {
	ID    string
	Name  string
}
