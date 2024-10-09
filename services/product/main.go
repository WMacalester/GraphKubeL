//go:generate echo -e "\033[1;34mGenerating sqlc...\033[0m"
//go:generate sqlc generate
//go:generate echo -e "\033[1;32mGenerating graphql...\033[0m"
//go:generate gqlgen generate
package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/WMacalester/GraphKubeL/services/product/graph"
	"github.com/jackc/pgx/v5/pgxpool"
)

const defaultPort = "8080"

type AppConfig struct {
	DB ProductRepository
}

func main() {
	connString, err := CreateConnString()
	if err != nil {
		log.Fatal(err)
	}

	connPool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	defer connPool.Close()

	// appConfig := AppConfig{DB: *NewProductRepository(connPool)}
	
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
