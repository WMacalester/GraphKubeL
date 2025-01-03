#!/bin/bash

source .env
source ./product.env 
PRODUCT_DB_PORT=5434
migrate -database "postgres://${PGUSER}:${POSTGRES_PASSWORD}@localhost:${PRODUCT_DB_PORT}/${POSTGRES_DB}?sslmode=${PRODUCT_PG_SSLMODE}" -path ./services/product/database/migrations down
migrate -database "postgres://${PGUSER}:${POSTGRES_PASSWORD}@localhost:${PRODUCT_DB_PORT}/${POSTGRES_DB}?sslmode=${PRODUCT_PG_SSLMODE}" -path ./services/product/database/migrations up

source ./order.env 
ORDER_DB_PORT=5433
migrate -database "postgres://${PGUSER}:${POSTGRES_PASSWORD}@localhost:${ORDER_DB_PORT}/${POSTGRES_DB}?sslmode=${ORDER_PG_SSLMODE}" -path ./services/order/database/migrations down
migrate -database "postgres://${PGUSER}:${POSTGRES_PASSWORD}@localhost:${ORDER_DB_PORT}/${POSTGRES_DB}?sslmode=${ORDER_PG_SSLMODE}" -path ./services/order/database/migrations up
