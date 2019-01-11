package todos_db

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/deslee/gqlgen_todos/models"
	"log"
)

func CreateUser(db *sql.DB, name string) (models.User, error) {
	newUuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	id := newUuid.String()

	_, err = db.Exec(`INSERT INTO Users(ID, Name) VALUES (?, ?)`, id, name)
	if err != nil {
		panic(err)
	}

	return models.User{ID: id, Name: name}, nil
}


func CreateTodo(db *sql.DB, text string, userID string) (models.Todo, error) {
	newUuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	id := newUuid.String()

	_, err = db.Exec(`INSERT INTO Todos(ID, Text, UserID) VALUES (?, ?, ?)`, id, text, userID)
	if err != nil {
		panic(err)
	}

	return models.Todo{ID: id, Text: text, UserID: userID}, nil
}

func GetUsers(db *sql.DB) ([]models.User, error) {
	var (
		id string
		name string
		users []models.User
	)

	rows, err := db.Query(`SELECT ID, Name from Users`)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		users = append(users, models.User{ID: id, Name: name})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return users, nil
}

func GetUser(db *sql.DB, id string) (models.User, error) {
	var (
		name string
	)

	err := db.QueryRow(`SELECT Name from Users where ID = ?`, id).Scan(&name)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return models.User{ID: id, Name: name}, nil
}

func GetTodos(db *sql.DB) ([]models.Todo, error) {
	var (
		id string
		text string
		userID string
		todos []models.Todo
	)

	rows, err := db.Query(`SELECT ID, Text, UserID from Todos`)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &text, &userID)
		if err != nil {
			panic(err)
		}
		todos = append(todos, models.Todo{ID: id, Text: text, UserID: userID})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return todos, nil
}

func GetTodosForUser(db *sql.DB, userID string) ([]models.Todo, error) {
	var (
		id string
		text string
		todos []models.Todo
	)

	rows, err := db.Query(`SELECT ID, Text FROM Todos WHERE UserID = ?`, userID)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &text)
		if err != nil {
			panic(err)
		}
		todos = append(todos, models.Todo{ID: id, Text: text, UserID: userID})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return todos, nil
}

func GetTodo(db *sql.DB, id string) (models.Todo, error) {
	var (
		text string
		userID string
	)

	err := db.QueryRow(`SELECT Text, UserID from Todos where ID = ?`, id).Scan(&text, &userID)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return models.Todo{ID: id, UserID: userID}, nil
}

func DeleteTodo(db *sql.DB, id string) error {
	_, err := db.Exec(`DELETE FROM Todos WHERE ID=?`, id)
	if err != nil {
		panic(err)
	}
	return nil
}

func DeleteUser(db *sql.DB, id string) error {
	_, err := db.Exec(`DELETE FROM Users WHERE ID=?`, id)
	if err != nil {
		panic(err)
	}
	return nil
}