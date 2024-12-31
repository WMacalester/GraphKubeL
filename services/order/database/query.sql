-- name: GetOrders :many
SELECT * FROM orders;

-- name: InsertOrder :one
INSERT INTO orders (transaction_id, product_id, number_of_items) 
VALUES ($1, $2, $3)
RETURNING transaction_id, product_id, number_of_items;
