package database

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type InventoryRepository struct {
	Db *redis.Client
}

func NewInventoryRepository() *InventoryRepository{
	addr, err := buildAddress()

	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
        Addr: addr, 
        Password: "", 
        DB: 0, 
    })

	return &InventoryRepository{Db: rdb}
}

func (r *InventoryRepository) GetProductInventory(ctx context.Context, key string) (int, error){
	val, err := r.Db.Get(ctx, key).Result();
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(val)
}

func (r *InventoryRepository) SetProductInventory(ctx context.Context, key string, value int) (int, error){
	val := strconv.Itoa(value);

	err := r.Db.Set(ctx, key, val, 0).Err();
	if err != nil {
		return 0, err
	}
	return value, nil
}

func buildAddress() (string, error) {
	db := os.Getenv("INV_HOST_NAME")
	port := os.Getenv("INV_DB_PORT")

	if (db == "" || port == ""){
		return "", fmt.Errorf("one or more required environment variables are missing. DB hostname: %v, DB port: %v", db, port)
	}

	return fmt.Sprintf("%s:%s", db, port), nil
}