-- PRODUCTS TABLE
CREATE TABLE IF NOT EXISTS products (
    id CHAR(36) PRIMARY KEY,
    vendor_id CHAR(36),
    name TEXT,
    price DECIMAL(10,2),
    stock INT,
    FOREIGN KEY (vendor_id) REFERENCES vendors(id) ON DELETE CASCADE
);