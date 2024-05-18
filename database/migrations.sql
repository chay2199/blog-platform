-- Create Users table
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    role TEXT NOT NULL
);

-- Create Posts table
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    author TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Insert an admin user
INSERT INTO users (username, password, role) VALUES ('admin', '$2a$10$miq1qfmdIwnEZeTO35K96.iuIMEH3p1Y4JyqgLZ0WsZ9zhce6jLw2', 'admin');
