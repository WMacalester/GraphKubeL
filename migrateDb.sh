#!/bin/bash
source ./product.env 
migrate -database "postgres://${PGUSER}:${POSTGRES_PASSWORD}@localhost:${PRODUCT_DB_PORT}/${POSTGRES_DB}?sslmode=${PRODUCT_PG_SSLMODE}" -path ./services/product/database/migrations down
migrate -database "postgres://${PGUSER}:${POSTGRES_PASSWORD}@localhost:${PRODUCT_DB_PORT}/${POSTGRES_DB}?sslmode=${PRODUCT_PG_SSLMODE}" -path ./services/product/database/migrations up