package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/eveld/graphql/auth"
	"github.com/eveld/graphql/models"
	"github.com/eveld/graphql/resolver"
	"github.com/eveld/graphql/server"
	"github.com/eveld/graphql/service"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()

	// Config.
	config := service.NewConfig()

	// Database.
	database, err := service.NewDatabase(
		config.DBhost,
		config.DBport,
		config.DBuser,
		config.DBpassword,
		config.DBname,
		config.DBmaxopenconns,
		config.DBconnmaxlifetime,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Database migrations.
	service.Migrate(database)

	// Services.
	trackService := service.NewTrackService(database)
	challengeService := service.NewChallengeService(database)

	// Authentication.
	firebaseApp, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	firebaseClient, err := firebaseApp.Auth(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Server.
	r := mux.NewRouter()
	r.Use(auth.Middleware(firebaseClient))

	// Graphql.
	c := server.Config{
		Resolvers: &resolver.Resolver{
			TrackService:     trackService,
			ChallengeService: challengeService,
		},
	}

	// Check if the user is authenticated.
	c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		isAuthenticated := auth.IsAuthenticated(ctx)
		if !isAuthenticated {
			return nil, fmt.Errorf("Access denied")
		}

		return next(ctx)
	}

	// Check if the user has a role.
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role models.Role) (interface{}, error) {
		log.Printf("%#v", role)
		return next(ctx)
	}

	// Handlers.
	r.Handle("/", handler.Playground("Playground", "/api"))
	r.Handle("/api", handler.GraphQL(server.NewExecutableSchema(c)))

	log.Println("Listing on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
