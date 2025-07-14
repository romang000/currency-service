CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) unique,
    password_hash TEXT
);