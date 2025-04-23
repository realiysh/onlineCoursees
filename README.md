# API для курсов и авторов

REST API на Go для управления курсами и авторами, включая систему аутентификации и авторизации.

## Запуск проекта

1. Установите зависимости:
   ```
   go mod tidy
   ```

2. Создайте файл `.env` с настройками БД:
   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=postgres
   ```

3. Запустите проект:
   ```
   go run main.go
   ```

## Тестирование

В проекте имеется 12 тестов в следующих файлах:
- `tests/auth_controller_test.go` (3 теста)
- `tests/author_controller_test.go` (6 тестов)
- `tests/course_controller_test.go` (3 теста)

Для запуска тестов используйте:
```
go test -v ./tests
```

Примечание: для запуска тестов необходимо настроить соединение с PostgreSQL. Если не хотите настраивать базу данных, тесты будут пропущены (Skip).

## API Endpoints

API поддерживает следующие операции:

### Публичные эндпоинты

- `POST /api/register` - Регистрация пользователя
- `POST /api/login` - Аутентификация и получение JWT-токена
- `GET /api/courses` - Получение списка курсов
- `GET /api/courses/:id` - Получение курса по ID
- `GET /api/authors` - Получение списка авторов
- `GET /api/authors/:id` - Получение автора по ID
- `GET /api/categories` - Получение списка категорий
- `GET /api/categories/:id` - Получение категории по ID

### Защищенные эндпоинты (требуют JWT токен)

- `POST /api/courses` - Создание нового курса
- `PUT /api/courses/:id` - Обновление курса
- `DELETE /api/courses/:id` - Удаление курса
- `POST /api/authors` - Создание нового автора
- `PUT /api/authors/:id` - Обновление автора
- `DELETE /api/authors/:id` - Удаление автора

## Формат запросов и ответов

Для пагинации используйте параметры `page` и `limit`:
```
/api/courses?page=1&limit=10
```

Для фильтрации курсов используйте параметры:
```
/api/courses?title=Go&min_price=10&max_price=50
``` 