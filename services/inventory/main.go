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
	"github.com/WMacalester/GraphKubeL/services/inventory/database"
	"github.com/WMacalester/GraphKubeL/services/inventory/graph"
)

const defaultPort = "8080"

func main() {
    inventoryRepository := database.NewInventoryRepository()
	
	ctx := context.Background()
    _, err := inventoryRepository.Db.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    }

    log.Println("Connected to Redis!")

	exposedPort := os.Getenv("INV_SERVICE_PORT")
	if exposedPort == "" {
		exposedPort = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Repository: inventoryRepository}}))

	http.Handle("/", playground.Handler("Inventory GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.HandleFunc("/health", HealthCheck())

	log.Printf("connect to http://localhost:%s/ for inventory GraphQL playground", exposedPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}

func HealthCheck()(http.HandlerFunc){
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
