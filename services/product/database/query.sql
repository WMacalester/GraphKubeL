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