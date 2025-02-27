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
	"github.com/WMacalester/GraphKubeL/internal/common"
	"github.com/WMacalester/GraphKubeL/services/product/database"
	"github.com/WMacalester/GraphKubeL/services/product/graph"
)

const defaultPort = "8080"

type AppConfig struct {
	DB database.ProductRepository
}

func main() {
	connPool := common.ConnectToPostgresDb(context.Background(), database.CreateConnString)
	defer connPool.Close()

	appConfig := AppConfig{DB: *database.NewProductRepository(connPool)}
	
	exposedPort := os.Getenv("PRODUCT_SERVICE_PORT")
	if exposedPort == "" {
		exposedPort = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Repository: &appConfig.DB}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.HandleFunc("/health", common.HealthCheck())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", exposedPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}

