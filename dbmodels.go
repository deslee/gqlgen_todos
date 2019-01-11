package gqlgen_todos

type Todo struct {
	ID string
	Text string
	Done bool
	UserID string
}

type User struct {
	ID    string
	Name  string
}
