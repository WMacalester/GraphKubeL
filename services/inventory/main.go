package main

import (
	"context"
	"log"

	"github.com/WMacalester/GraphKubeL/services/inventory/database"
)

func main() {
    inventoryRepository := database.NewInventoryRepository()
	
	ctx := context.Background()
    _, err := inventoryRepository.Db.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    }

    log.Println("Connected to Redis!")
}
