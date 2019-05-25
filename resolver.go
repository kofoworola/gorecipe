package gorecipe

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateRecipe(ctx context.Context, input *NewRecipe, ingredients []*NewIngredient) (*Recipe, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateRecipe(ctx context.Context, id *int, input *NewRecipe, ingredients []*NewIngredient) (*Recipe, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteRecipe(ctx context.Context, id *int) ([]*Recipe, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Recipes(ctx context.Context) ([]*Recipe, error) {
	panic("not implemented")
}
