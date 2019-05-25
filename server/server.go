package main

import (
	"github.com/kofoworola/gorecipe/models"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/kofoworola/gorecipe"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//Migrate Db
	db := models.FetchConnection()
	db.AutoMigrate(&models.Recipe{},&models.Ingredient{})
	db.Close()

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gorecipe.NewExecutableSchema(gorecipe.Config{Resolvers: &gorecipe.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
