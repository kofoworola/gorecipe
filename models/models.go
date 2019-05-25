package models

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Ingredient struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	//Foreign key
	RecipeID int    `json:"recipeId"`
}

type Recipe struct {
	ID          int           `json:"id" gorm:"primary_key"`
	Name        string        `json:"name"`
	Procedure   string        `json:"procedure"`
	//Ingredients owned by a recipe
	Ingredients []Ingredient `json:"ingredients"`
}

func FetchConnection() *gorm.DB{
	db,err := gorm.Open("mysql","user:password@/golang")
	if err != nil{
		panic(err)
	}
	return db
}