package gqlgen_todos

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
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

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (User, error) {
	tx, err := r.Db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	newUuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	stmt, err := tx.Prepare("INSERT INTO Users (ID, Name) VALUES (? ?)")
	defer stmt.Close()

	id := newUuid.String()
	name := input.Name

	_, err = stmt.Exec(id, name)
	if err != nil {
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	return User{ID: id, Name: input.Name}, nil
}
func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (Todo, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id string) (User, error) {
	panic("not implemented")
}
func (r *queryResolver) Todos(ctx context.Context) ([]Todo, error) {
	panic("not implemented")
}
func (r *queryResolver) Todo(ctx context.Context, id string) (Todo, error) {
	panic("not implemented")
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *Todo) (User, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) Todos(ctx context.Context, obj *User) ([]Todo, error) {
	panic("not implemented")
}
