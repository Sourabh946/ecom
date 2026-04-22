-- VENDORS TABLE
CREATE TABLE IF NOT EXISTS vendors (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36),
    store_name TEXT,
    is_verified BOOLEAN,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);