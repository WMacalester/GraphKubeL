-- name: GetProductById :one
SELECT
    p.id AS id,
    p.name AS name,
    p.description AS description,
    pc.id AS category_id,
    pc.name AS category_name
FROM products p
JOIN product_categories pc ON p.category_id = pc.id
WHERE p.id = $1;

-- name: GetProducts :many
SELECT
    p.id AS id,
    p.name AS name,
    p.description AS description,
    pc.id AS category_id,
    pc.name AS category_name
FROM products p
JOIN product_categories pc ON p.category_id = pc.id
ORDER BY p.name;

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