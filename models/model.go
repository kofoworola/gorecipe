package models

type Ingredient struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	RecipeID int    `json:"recipeId"`
}

type Recipe struct {
	ID          int           `json:"id" gorm:"primary_key"`
	Name        string        `json:"name"`
	Procedure   string        `json:"procedure"`
	Ingredients []Ingredient `json:"ingredients"`
}