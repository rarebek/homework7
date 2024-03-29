CREATE TABLE IF NOT EXISTS users_products (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    product_id UUID NOT NULL,
    amount INT NOT NULL DEFAULT 1
);