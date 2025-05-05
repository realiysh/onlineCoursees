CREATE TABLE IF NOT EXISTS courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    category_id INT NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

INSERT INTO courses (title, description, price, category_id) VALUES
('Backend дамыту — Go тілімен', 'Изучите основы разработки на Go', 150000, 1),
('Frontend негіздері — React.js', 'Научитесь создавать современные веб-приложения', 180000, 1),
('DevOps практикумы: Docker & CI/CD', 'Освойте инструменты DevOps', 220000, 2),
('Киберқауіпсіздікке кіріспе', 'Основы информационной безопасности', 160000, 2),
('Машиналық оқыту негіздері', 'Введение в машинное обучение', 275000, 1); 