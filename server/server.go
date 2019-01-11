package main

import (
	"database/sql"
	"github.com/99designs/gqlgen/handler"
	"github.com/deslee/gqlgen_todos"
	"github.com/deslee/gqlgen_todos/database"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("sqlite3", "file:database.db")
	if err != nil {
		panic(err)
	}

	todos_db.CreateTablesIfNotExist(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gqlgen_todos.NewExecutableSchema(gqlgen_todos.Config{Resolvers: &gqlgen_todos.Resolver{Db: db}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}