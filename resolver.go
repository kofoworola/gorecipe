package gorecipe

import (
	"context"
	"github.com/kofoworola/gorecipe/models"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

//Resolver for mutations
type mutationResolver struct{ *Resolver }

//Create recipe mutation
func (r *mutationResolver) CreateRecipe(ctx context.Context, input *NewRecipe, ingredients []*NewIngredient) (*models.Recipe, error) {
	//Fetch Connection and close db
	db := models.FetchConnection()
	defer db.Close()

	//Create the recipe using the input structs
	recipe := models.Recipe{Name: input.Name, Procedure: *input.Procedure}

	//initialize the ingredients with the length of the input for ingredients
	recipe.Ingredients = make([]models.Ingredient,len(ingredients))
	//Loop and add all items
	for index,item := range ingredients{
		recipe.Ingredients[index] = models.Ingredient{Name: item.Name}
	}
	//Create by passing the pointer to the recipe
	db.Create(&recipe)
	return &recipe, nil
}

//Update recipe mutation
func (r *mutationResolver) UpdateRecipe(ctx context.Context, id *int, input *NewRecipe, ingredients []*NewIngredient) (*models.Recipe, error) {
	//Fetch Connection and close db
	db := models.FetchConnection()
	defer db.Close()

	var recipe models.Recipe
	//Find recipe based on ID and update
	db = db.Preload("Ingredients").Where("id = ?",*id).First(&recipe).Update("name",input.Name)
	if input.Procedure != nil{
		db.Update("procedure",*input.Procedure)
	}

	//Update Ingredients
	recipe.Ingredients = make([]models.Ingredient,len(ingredients))
	for index,item := range ingredients{
		recipe.Ingredients[index] = models.Ingredient{Name:item.Name}
	}
	db.Save(&recipe)
	return &recipe,nil
}

//Delete recipe mutation
func (r *mutationResolver) DeleteRecipe(ctx context.Context, id *int) ([]*models.Recipe, error) {
	//Fetch connection
	db := models.FetchConnection()
	defer db.Close()
	var recipe models.Recipe

	//Fetch based on ID and delete
	db.Where("id = ?",*id).First(&recipe).Delete(&recipe)

	//Preload and fetch all recipes
	var recipes []*models.Recipe
	db.Preload("Ingredients").Find(&recipes)
	return recipes,nil
}


//Query resolver
type queryResolver struct{ *Resolver }

//Get all recipes
func (r *queryResolver) Recipes(ctx context.Context) ([]*models.Recipe, error) {
	//Fetch a connection
	db := models.FetchConnection()
	//Defer closing the database
	defer db.Close()
	//Create an array of recipes to populate
	var recipes []*models.Recipe
	// .Preload loads the Ingredients relationship into each recipe
	db.Preload("Ingredients").Find(&recipes)
	return recipes,nil
}
