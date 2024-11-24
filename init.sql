CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    customer VARCHAR(100),
    total NUMERIC(10, 2)
);

INSERT INTO orders (customer, total) VALUES ('John Doe', 123.45);
