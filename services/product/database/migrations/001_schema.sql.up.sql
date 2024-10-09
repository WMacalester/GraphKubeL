CREATE TABLE product_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(31) NOT NULL UNIQUE
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    category_id INT,
    description VARCHAR(250),
    Foreign Key (category_id) REFERENCES product_categories(id)
);
