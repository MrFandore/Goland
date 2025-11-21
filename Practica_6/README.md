# Практика 6 - Использование ORM GORM для работы с PostgreSQL

## Структура проекта
```
Practica_6/
├── internal/
│   ├── db/
│   │   └── postgres.go
│   ├── models/
│   │   └── models.go
│   └── httpapi/
│       ├── router.go
│       └── handlers.go
├── main.go
├── go.mod
├── go.sum
└── README.md
```

## Описание файлов
- **main.go** - главный файл для запуска сервера и автоматических миграций
- **internal/db/postgres.go** - подключение к PostgreSQL через GORM
- **internal/models/models.go** - модели данных (User, Note, Tag) с связями
- **internal/httpapi/router.go** - настройка маршрутов Chi роутера
- **internal/httpapi/handlers.go** - обработчики HTTP запросов
- **go.mod** - файл зависимостей Go
- **go.sum** - контрольные суммы зависимостей

## Настройка базы данных

### 1. Создание базы данных
```sql
-- Подключиться к PostgreSQL
psql -U postgres

-- Создать базу данных
CREATE DATABASE pz6_gorm;

-- Подключиться к базе
\c pz6_gorm
```

### 2. Настройка переменной окружения
**Windows PowerShell:**
```powershell
$env:DB_DSN="host=127.0.0.1 user=postgres password=ваш_пароль dbname=pz6_gorm port=5432 sslmode=disable"
```

**macOS/Linux:**
```bash
export DB_DSN='host=127.0.0.1 user=postgres password=ваш_пароль dbname=pz6_gorm port=5432 sslmode=disable'
```

## Запуск проекта

### 1. Клонирование и настройка
```bash
git clone https://github.com/MrFandore/Goland.git
cd Goland/Practica_6
```

### 2. Установка зависимостей
```bash
go mod download
```

### 3. Запуск сервера
```bash
go run main.go
```

## Проверка работы

После запуска сервер доступен по адресу: http://localhost:8080

### Доступные эндпоинты:

**Проверка здоровья сервера:**
```
GET /health
```
Ожидаемый результат: `{"status":"ok"}`

**Создание пользователя:**
```
POST /users
```
Тело запроса: `{"name":"Имя пользователя", "email":"email@example.com"}`

**Создание заметки:**
```
POST /notes
```
Тело запроса: `{"title":"Заголовок", "content":"Текст заметки", "userId":1, "tags":["go", "gorm"]}`

**Получение заметки по ID:**
```
GET /notes/1
```
Ожидаемый результат: информация о заметке с автором и тегами

## Примеры тестирования

### Через командную строку (PowerShell):
```powershell
# Проверка здоровья
Invoke-WebRequest -Uri http://localhost:8080/health -Method GET

# Создание пользователя
Invoke-WebRequest -Uri http://localhost:8080/users `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"name":"Alice","email":"alice@example.com"}'

# Создание заметки с тегами
Invoke-WebRequest -Uri http://localhost:8080/notes `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"title":"Первая заметка","content":"Текст заметки...","userId":1,"tags":["go","gorm"]}'

# Получение заметки
Invoke-WebRequest -Uri http://localhost:8080/notes/1 -Method GET
```

### Через командную строку (curl):
```bash
# Проверка здоровья
curl http://localhost:8080/health

# Создание пользователя
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{"name":"Alice","email":"alice@example.com"}'

# Создание заметки с тегами
curl -X POST http://localhost:8080/notes \
-H "Content-Type: application/json" \
-d '{"title":"Первая заметка","content":"Текст заметки...","userId":1,"tags":["go","gorm"]}'

# Получение заметки
curl http://localhost:8080/notes/1
```

## Модели данных

### User (Пользователь)
```go
type User struct {
    ID        uint   `gorm:"primaryKey"`
    Name      string `gorm:"size:100;not null"`
    Email     string `gorm:"size:200;uniqueIndex;not null"`
    Notes     []Note // Связь 1:N с заметками
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### Note (Заметка)
```go
type Note struct {
    ID        uint   `gorm:"primaryKey"`
    Title     string `gorm:"size:200;not null"`
    Content   string `gorm:"type:text"`
    UserID    uint   `gorm:"not null"`
    User      User
    Tags      []Tag `gorm:"many2many:note_tags;"` // Связь M:N с тегами
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### Tag (Тег)
```go
type Tag struct {
    ID        uint   `gorm:"primaryKey"`
    Name      string `gorm:"size:50;uniqueIndex;not null"`
    Notes     []Note `gorm:"many2many:note_tags;"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

## Особенности проекта

- **Автоматические миграции** - таблицы создаются автоматически через `AutoMigrate`
- **Связи между таблицами**:
  - **1:N** - один пользователь может иметь много заметок
  - **M:N** - заметки могут иметь много тегов, теги могут принадлежать многим заметкам
- **Preload загрузка** - автоматическая подгрузка связанных данных (User и Tags)
- **Валидация данных** - проверка обязательных полей и уникальности email/тегов
- **Пул соединений** - оптимизированные настройки пула GORM

## Настройки пула соединений GORM

```go
sqlDB.SetMaxIdleConns(10)     // Соединений в простое
sqlDB.SetMaxOpenConns(20)     // Максимум активных соединений
sqlDB.SetConnMaxLifetime(time.Hour) // Время жизни соединения
```

## Требования

- Go версии 1.21 или выше
- PostgreSQL 14 или выше
- Установленные зависимости:
  - GORM: `gorm.io/gorm`
  - Драйвер PostgreSQL: `gorm.io/driver/postgres`
  - Chi роутер: `github.com/go-chi/chi/v5`

## Установка зависимостей
```bash
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/go-chi/chi/v5
```

## Решение проблем

**Ошибка подключения к БД:**
```
DB_DSN ПУСТО, ТУТА НИКОГО НЕТУ))
```
Установите переменную окружения DB_DSN как описано выше.

**Ошибка дублирования email:**
```
duplicate key value violates unique constraint
```
Email должен быть уникальным - используйте другой email.

**Таблицы не создаются:**
```
relation "users" does not exist
```
Убедитесь, что AutoMigrate выполняется в main.go.

**Неправильный пароль/хост:**
```
connection failed
```
Проверьте параметры подключения в DB_DSN.

## Преимущества использования GORM

1. **Автоматизация** - автоматическое создание таблиц и связей
2. **Безопасность** - защита от SQL-инъекций через параметризацию
3. **Производительность** - умная загрузка связанных данных через Preload
4. **Удобство** - минимальный код для стандартных CRUD операций
5. **Гибкость** - поддержка сложных связей между таблицами

## Остановка сервера
Нажмите `Ctrl+C` в командной строке где запущен сервер.
