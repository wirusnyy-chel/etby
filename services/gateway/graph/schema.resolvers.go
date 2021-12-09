package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"etby/services/gateway/graph/generated"
	"etby/services/gateway/graph/model"
	"fmt"
)

func (r *mutationResolver) CreateEvent(ctx context.Context, newEvent model.NewEvent) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTask(ctx context.Context, newTask model.NewTask) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditEvent(ctx context.Context, id string, event model.Edit) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditTask(ctx context.Context, id string, task model.Edit) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Event(ctx context.Context, id string) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Task(ctx context.Context, id string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
