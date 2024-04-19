package db

const scheme = `

CREATE TABLE IF NOT EXISTS url (
    name CHAR(5) PRIMARY KEY,
    link TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    times INT DEFAULT 0
);

`
