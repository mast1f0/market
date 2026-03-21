CREATE TABLE cart_items (
    id SERIAL PRIMARY KEY,
    cart_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL DEFAULT 1,

    CONSTRAINT fk_cart
        FOREIGN KEY(cart_id)
            REFERENCES carts(id)
                 ON DELETE CASCADE,

    CONSTRAINT fk_product
        FOREIGN KEY(product_id)
            REFERENCES products(id)
                ON DELETE CASCADE
);