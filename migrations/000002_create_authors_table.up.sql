CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

INSERT INTO authors (name) VALUES
('Алтынбек Нурлы'),
('Айкерим Тулепберген'),
('Нурбол Абдрахман'),
('Жанель Сапаркызы'),
('Рамазан Ержанулы');
