package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/paundraP/be-mcs/user-service/config"
	"github.com/paundraP/be-mcs/user-service/graphql/generated"
	"github.com/paundraP/be-mcs/user-service/graphql/resolvers"
	"github.com/paundraP/be-mcs/user-service/models"
	"github.com/paundraP/be-mcs/user-service/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := config.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	userRepo := repository.NewUserRepo(db)

	resolver := &resolvers.Resolver{
		UserRepo: userRepo,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	// Route the Playground to "/"
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	// Route GraphQL queries/mutations to "/query"
	http.Handle("/query", srv)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
