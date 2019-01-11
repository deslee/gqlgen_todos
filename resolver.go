package gqlgen_todos

import (
	"context"
	"database/sql"
	"github.com/deslee/gqlgen_todos/database"
	"github.com/deslee/gqlgen_todos/models"
)

type Resolver struct{
	Db *sql.DB
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (models.User, error) {
	return todos_db.CreateUser(r.Db, input.Name)
}
func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (models.Todo, error) {
	return todos_db.CreateTodo(r.Db, input.Text, input.UserID)
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]models.User, error) {
	return todos_db.GetUsers(r.Db)
}
func (r *queryResolver) User(ctx context.Context, id string) (models.User, error) {
	return todos_db.GetUser(r.Db, id)
}

func (r *queryResolver) Todos(ctx context.Context) ([]models.Todo, error) {
	return todos_db.GetTodos(r.Db)
}
func (r *queryResolver) Todo(ctx context.Context, id string) (models.Todo, error) {
	return todos_db.GetTodo(r.Db, id)
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (models.User, error) {
	return todos_db.GetUser(r.Db, obj.UserID)
}

type userResolver struct{ *Resolver }

func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]models.Todo, error) {
	return todos_db.GetTodosForUser(r.Db, obj.ID)
}
