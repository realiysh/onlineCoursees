CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author_id INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE
);

INSERT INTO courses (title, author_id, price) VALUES
('Backend дамыту — Go тілімен', 1, 150000),
('Frontend негіздері — React.js', 2, 180000),
('DevOps практикумы: Docker & CI/CD', 3, 220000),
('Киберқауіпсіздікке кіріспе', 4, 160000),
('Машиналық оқыту негіздері', 5, 275000);
