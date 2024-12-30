package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/paundraP/be-mcs/user-service/config"
	"github.com/paundraP/be-mcs/user-service/directives"
	"github.com/paundraP/be-mcs/user-service/graphql/generated"
	"github.com/paundraP/be-mcs/user-service/graphql/resolvers"
	"github.com/paundraP/be-mcs/user-service/middleware"
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

	if err := db.AutoMigrate(&generated.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware)

	userRepo := repository.NewUserRepo(db)

	resolver := &resolvers.Resolver{
		UserRepo: userRepo,
	}
	c := generated.Config{Resolvers: resolver}
	c.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	// Route the Playground to "/"
	router.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	// Route GraphQL queries/mutations to "/query"
	router.Handle("/query", srv)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
