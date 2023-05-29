CREATE DATABASE smart_it;

\c smart_it

CREATE TYPE role_t AS ENUM ('user', 'admin', '');

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    role role_t NOT NULL,
    age INTEGER DEFAULT 0,
    email TEXT UNIQUE NOT NULL
);