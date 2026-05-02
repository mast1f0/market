CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price_snapshot NUMERIC(10,2) NOT NULL,
    name_snapshot TEXT NOT NULL,
    image_snapshot TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_order
        FOREIGN KEY(order_id)
            REFERENCES orders(id)
                ON DELETE CASCADE
);