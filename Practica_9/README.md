# Практика 9 - Регистрация и аутентификация пользователей с bcrypt

## Структура проекта
```
Practica_9/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── core/
│   │   └── user.go
│   ├── http/
│   │   └── handlers/
│   │       └── auth.go
│   ├── platform/
│   │   └── config/
│   │       └── config.go
│   └── repo/
│       ├── postgres.go
│       └── user_repo.go
├── go.mod
├── go.sum
└── README.md
```

## Описание файлов
- **cmd/api/main.go** - точка входа приложения, настройка сервера и маршрутов
- **internal/core/user.go** - модель пользователя для GORM
- **internal/http/handlers/auth.go** - обработчики регистрации и входа
- **internal/platform/config/config.go** - загрузка конфигурации из переменных окружения
- **internal/repo/postgres.go** - подключение к PostgreSQL через GORM
- **internal/repo/user_repo.go** - репозиторий для работы с пользователями
- **go.mod** - файл зависимостей Go
- **go.sum** - контрольные суммы зависимостей

## Настройка базы данных

### 1. Создание базы данных PostgreSQL
```sql
-- Подключиться к PostgreSQL
psql -U postgres

-- Создать базу данных
CREATE DATABASE pz9_auth;

-- Подключиться к базе
\c pz9_auth
```

### 2. Настройка переменных окружения

Создайте файл `.env` в корне проекта:
```env
DB_DSN=postgres://postgres:ваш_пароль@localhost:5432/pz9_auth?sslmode=disable
BCRYPT_COST=12
APP_ADDR=:8080
```

**Windows PowerShell:**
```powershell
$env:DB_DSN="postgres://postgres:ваш_пароль@localhost:5432/pz9_auth?sslmode=disable"
$env:BCRYPT_COST="12"
$env:APP_ADDR=":8080"
```

**macOS/Linux:**
```bash
export DB_DSN='postgres://postgres:ваш_пароль@localhost:5432/pz9_auth?sslmode=disable'
export BCRYPT_COST='12'
export APP_ADDR=':8080'
```

## Запуск проекта

### 1. Клонирование и настройка
```bash
git clone https://github.com/MrFandore/Goland.git
cd Goland/Practica_9
```

### 2. Установка зависимостей
```bash
go mod download
```

### 3. Запуск сервера
```bash
go run cmd/api/main.go
```

## Проверка работы

После запуска сервер доступен по адресу: http://localhost:8080

### Доступные эндпоинты:

**Регистрация пользователя:**
```
POST /auth/register
```
Тело запроса: `{"email":"user@example.com", "password":"Secret123!"}`

**Вход пользователя:**
```
POST /auth/login
```
Тело запроса: `{"email":"user@example.com", "password":"Secret123!"}`

## Примеры тестирования

### Через командную строку (PowerShell):
```powershell
# Регистрация нового пользователя
Invoke-WebRequest -Uri "http://localhost:8080/auth/register" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"email":"user@example.com","password":"Secret123!"}'

# Попытка повторной регистрации с тем же email
Invoke-WebRequest -Uri "http://localhost:8080/auth/register" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"email":"user@example.com","password":"AnotherPassword"}'

# Успешный вход
Invoke-WebRequest -Uri "http://localhost:8080/auth/login" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"email":"user@example.com","password":"Secret123!"}'

# Неудачный вход (неверный пароль)
Invoke-WebRequest -Uri "http://localhost:8080/auth/login" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"email":"user@example.com","password":"WrongPassword"}'

# Неудачный вход (несуществующий email)
Invoke-WebRequest -Uri "http://localhost:8080/auth/login" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"email":"nonexistent@example.com","password":"SomePassword"}'
```

### Через командную строку (curl):
```bash
# Регистрация нового пользователя
curl -X POST http://localhost:8080/auth/register \
-H "Content-Type: application/json" \
-d '{"email":"user@example.com","password":"Secret123!"}'

# Попытка повторной регистрации с тем же email
curl -X POST http://localhost:8080/auth/register \
-H "Content-Type: application/json" \
-d '{"email":"user@example.com","password":"AnotherPassword"}'

# Успешный вход
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{"email":"user@example.com","password":"Secret123!"}'

# Неудачный вход (неверный пароль)
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{"email":"user@example.com","password":"WrongPassword"}'

# Неудачный вход (несуществующий email)
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{"email":"nonexistent@example.com","password":"SomePassword"}'
```

## Ожидаемые ответы

### Успешная регистрация (201 Created):
```json
{
  "status": "ok",
  "user": {
    "id": 1,
    "email": "user@example.com"
  }
}
```

### Ошибка регистрации - email занят (409 Conflict):
```json
{
  "error": "email_taken"
}
```

### Успешный вход (200 OK):
```json
{
  "status": "ok",
  "user": {
    "id": 1,
    "email": "user@example.com"
  }
}
```

### Ошибка входа (401 Unauthorized):
```json
{
  "error": "invalid_credentials"
}
```

### Ошибка валидации (400 Bad Request):
```json
{
  "error": "email_required_and_password_min_8"
}
```

## Модель данных

### User (Пользователь)
```go
type User struct {
    ID           int64     `gorm:"primaryKey" json:"id"`
    Email        string    `gorm:"uniqueIndex;size:255;not null" json:"email"`
    PasswordHash string    `gorm:"size:255;not null" json:"-"`
    CreatedAt    time.Time `json:"createdAt"`
    UpdatedAt    time.Time `json:"updatedAt"`
}
```

## Особенности реализации

### Безопасное хранение паролей
- **bcrypt** - алгоритм хэширования с автоматической солью
- **Cost factor 12** - баланс между безопасностью и производительностью
- **Сравнение хэшей** - использование `bcrypt.CompareHashAndPassword`

### Валидация данных
- **Email** - обязательное поле, автоматическое приведение к нижнему регистру
- **Пароль** - минимум 8 символов
- **Уникальность email** - проверка на уровне базы данных

### Обработка ошибок
- **Общие сообщения** - "invalid_credentials" вместо конкретных ошибок
- **HTTP статусы** - корректные коды ответов (200, 400, 401, 409, 500)
- **Логирование** - пароли не попадают в логи

## Требования

- Go версии 1.21 или выше
- PostgreSQL 14 или выше
- Установленные зависимости:
  - GORM: `gorm.io/gorm`
  - Драйвер PostgreSQL: `gorm.io/driver/postgres`
  - Chi роутер: `github.com/go-chi/chi/v5`
  - Bcrypt: `golang.org/x/crypto/bcrypt`

## Установка зависимостей
```bash
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/go-chi/chi/v5
go get golang.org/x/crypto/bcrypt
```

## Решение проблем

**Ошибка подключения к базе данных:**
```
db connect: dial error
```
Убедитесь, что PostgreSQL запущен и DSN указан правильно.

**Ошибка миграции:**
```
migrate: duplicate key value violates unique constraint
```
База данных уже содержит таблицу users - удалите её или используйте другую базу.

**Ошибка уникальности email:**
```
email_taken
```
Email уже зарегистрирован - используйте другой email.

**Слабый пароль:**
```
email_required_and_password_min_8
```
Пароль должен содержать минимум 8 символов.

## Преимущества bcrypt

### Безопасность
```go
// Хэширование пароля с автоматической солью
hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)

// Проверка пароля
err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
```

### Защита от атак
- **Rainbow tables** - автоматическая соль делает предварительно вычисленные таблицы бесполезными
- **Brute force** - регулируемая сложность замедляет перебор паролей
- **Timing attacks** - постоянное время сравнения хэшей

## Настройка сложности bcrypt

Параметр `cost` влияет на время хэширования:
- **cost 10** - быстрый, менее безопасный
- **cost 12** - рекомендуемый для большинства приложений
- **cost 14** - высокий уровень безопасности
- **cost 16+** - очень медленный, для критически важных данных

## Пример использования в продакшене

```go
// Рекомендуемые настройки для продакшена
config := Config{
    DB_DSN:     os.Getenv("DATABASE_URL"),
    BcryptCost: 12, // На продакшене можно увеличить до 14
    Addr:       ":" + os.Getenv("PORT"),
}
```

## Дополнительные меры безопасности

### Rate Limiting
```go
// Рекомендуется добавить ограничение попыток входа
// Например: максимум 5 попыток входа в минуту с одного IP
```

### Валидация email
```go
// Можно добавить проверку формата email
import "regexp"

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func isValidEmail(email string) bool {
    return emailRegex.MatchString(email)
}
```

## Остановка сервера
Нажмите `Ctrl+C` в командной строке где запущен сервер.

## Миграция базы данных

При первом запуске GORM автоматически создаст таблицу:
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);
```

## Проверка данных в базе

После регистрации пользователя можно проверить данные:
```sql
SELECT id, email, created_at FROM users;
```

**Важно:** Поле `password_hash` содержит bcrypt хэш, который нельзя расшифровать.
