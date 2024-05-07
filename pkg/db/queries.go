package db

const scheme = `
CREATE TABLE IF NOT EXISTS url (
    name CHAR(4),
    link TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    times INT DEFAULT 0,
    PRIMARY KEY (name),
    CHECK (LENGTH(name) = 4)
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (id),
    UNIQUE (username)
);
`
