CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    transaction_id INT NOT NULL,
    product_id INT NOT NULL UNIQUE,
    number_of_items INT NOT NULL
);
