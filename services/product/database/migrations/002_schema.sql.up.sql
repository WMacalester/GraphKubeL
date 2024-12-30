WITH inserted_categories AS (
    INSERT INTO product_categories (name)
    VALUES 
        ('Electronics'),
        ('Furniture'),
        ('Clothing')
    RETURNING id, name
)
INSERT INTO products (name, category_id, description)
VALUES
    ('Smartphone', (SELECT id FROM inserted_categories WHERE name = 'Electronics'), 'A high-tech smartphone.'),
    ('Table', (SELECT id FROM inserted_categories WHERE name = 'Furniture'), 'A sturdy wooden table.'),
    ('Jacket', (SELECT id FROM inserted_categories WHERE name = 'Clothing'), 'A warm winter jacket.');