CREATE TABLE products (
    id SERIAL,
    name VARCHAR(128) NOT NULL,
    price NUMERIC(10, 2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)