CREATE TABLE IF NOT EXISTS roles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    role TEXT UNIQUE
);

CREATE TABLE IF NOT EXISTS permissions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    role_id INTEGER,
    name TEXT,
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

-- CREATE TABLE IF NOT EXISTS users (
--     id TEXT PRIMARY KEY,
--     name TEXT UNIQUE,
--     email TEXT UNIQUE,
--     password TEXT,
--     role_id INTEGER,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
-- 	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     FOREIGN KEY (role_id) REFERENCES roles(id)
-- );
