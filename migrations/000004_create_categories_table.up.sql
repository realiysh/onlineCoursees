CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT
);

INSERT INTO categories (name, description) VALUES ('Books', 'All kinds of books');
INSERT INTO categories (name, description) VALUES ('Tech', 'Technology and gadgets');

UPDATE categories
SET name = 'Programmer', description = 'nothing'
WHERE name = 'Books';

UPDATE categories
SET name = 'IT', description = 'nothing'
WHERE name = 'Tech';
