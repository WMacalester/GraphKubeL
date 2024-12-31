CREATE TABLE orders (
    transaction_id INT NOT NULL,
    product_id INT NOT NULL,
    number_of_items INT NOT NULL,
    PRIMARY KEY (transaction_id, product_id)
);

