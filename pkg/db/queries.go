package db

const scheme = `

CREATE TABLE IF NOT EXISTS url (
    name CHAR(5),
    link TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    times INT DEFAULT 0,
    PRIMARY KEY (name)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    isAdmin BOOLEAN DEFAULT False,
    PRIMARY KEY (id)
);

`
