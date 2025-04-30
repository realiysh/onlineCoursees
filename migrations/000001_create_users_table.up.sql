CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

INSERT INTO users (username, password) VALUES
('altynbek_dev', 'securepass123'),
('aikerim_code', 'qwerty987'),
('nurbol_ds', 'pass456'),
('zhanel_js', 'abc123def'),
('ramazan_root', 'rootpass321');
