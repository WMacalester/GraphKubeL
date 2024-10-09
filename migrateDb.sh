#!/bin/bash
source ./product.env 
migrate -database "postgres://${PRODUCT_PG_USER}:${PRODUCT_PG_PASSWORD}@${PRODUCT_PG_HOST}:${PRODUCT_PG_PORT}/${PRODUCT_PG_DATABASE}?sslmode=${PRODUCT_PG_SSLMODE}" -path ./services/product/database/migrations up