-- name: GetProducts :many
SELECT *
FROM products
ORDER BY name;

-- name: GetProductCategoryById :one
SELECT *
FROM product_categories
WHERE id = $1;

-- name: GetProductCategories :many
SELECT *
FROM product_categories
ORDER BY name;

-- name: InsertProductCategory :one
INSERT INTO product_categories (name)
VALUES ($1)
ON CONFLICT (name) DO NOTHING
RETURNING id, name;

-- name: InsertProduct :one
INSERT INTO products (name, category_id, description)
VALUES ($1, $2, $3)
RETURNING id, name, category_id, description;