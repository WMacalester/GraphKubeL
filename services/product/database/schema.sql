CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    category INTEGER,
    description VARCHAR(250)
);
