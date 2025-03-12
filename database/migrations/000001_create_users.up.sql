CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    uuid varchar(36) unique NOT NULL,
    first_name varchar(250) NOT NULL,
    last_name varchar(250) NOT NULL,
    email varchar(250) unique NOT NULL,
    password varchar(100) NOT NULL,
    token varchar(250) NOT NULL,
    verified_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);
