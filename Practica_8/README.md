# Практика 8 - Работа с MongoDB для документного хранения данных

## Структура проекта
```
Practica_8/
├── cmd/
│   └── main.go
├── internal/
│   ├── db/
│   │   └── mongo.go
│   └── notes/
│       ├── model.go
│       ├── repo.go
│       └── handler.go
├── docker-compose.yaml
├── go.mod
├── go.sum
└── README.md
```

## Описание файлов
- **cmd/main.go** - главный файл для запуска сервера и подключения к MongoDB
- **internal/db/mongo.go** - подключение к MongoDB с настройками таймаутов
- **internal/notes/model.go** - модель данных для заметок
- **internal/notes/repo.go** - репозиторий с CRUD операциями для MongoDB
- **internal/notes/handler.go** - HTTP обработчики для REST API
- **docker-compose.yaml** - конфигурация для запуска MongoDB в Docker
- **go.mod** - файл зависимостей Go
- **go.sum** - контрольные суммы зависимостей

## Запуск MongoDB

### 1. Запуск через Docker Compose (рекомендуется)
```bash
# Запуск MongoDB
docker-compose up -d

# Проверка статуса
docker-compose ps

# Остановка
docker-compose down
```

### 2. Подключение к MongoDB
```bash
# Подключение через MongoDB Shell
docker exec -it mongo-dev mongosh -u root -p secret --authenticationDatabase admin

# В консоли MongoDB выполните:
show dbs
use pz8
show collections
```

## Настройка переменных окружения

Создайте файл `.env` в корне проекта:
```
MONGO_URI=mongodb://root:secret@localhost:27017/?authSource=admin
MONGO_DB=pz8
HTTP_ADDR=:8080
```

**Windows PowerShell:**
```powershell
$env:MONGO_URI="mongodb://root:secret@localhost:27017/?authSource=admin"
$env:MONGO_DB="pz8"
$env:HTTP_ADDR=":8080"
```

**macOS/Linux:**
```bash
export MONGO_URI='mongodb://root:secret@localhost:27017/?authSource=admin'
export MONGO_DB='pz8'
export HTTP_ADDR=':8080'
```

## Запуск проекта

### 1. Клонирование и настройка
```bash
git clone https://github.com/MrFandore/Goland.git
cd Goland/Practica_8
```

### 2. Установка зависимостей
```bash
go mod download
```

### 3. Запуск сервера
```bash
go run cmd/main.go
```

## Проверка работы

После запуска сервер доступен по адресу: http://localhost:8080

### Доступные эндпоинты:

**Проверка здоровья сервера:**
```
GET /health
```
Ожидаемый результат: `{"status":"ok"}`

**Создание заметки:**
```
POST /api/v1/notes
```
Тело запроса: `{"title":"Заголовок заметки", "content":"Текст заметки"}`

**Получение списка заметок:**
```
GET /api/v1/notes?q=поиск&limit=10&skip=0
```
Параметры:
- `q` - поиск по заголовку (опционально)
- `limit` - количество записей (по умолчанию 20, максимум 200)
- `skip` - пропуск записей (по умолчанию 0)

**Получение заметки по ID:**
```
GET /api/v1/notes/{id}
```

**Частичное обновление заметки:**
```
PATCH /api/v1/notes/{id}
```
Тело запроса: `{"title":"Новый заголовок"}` или `{"content":"Новый текст"}`

**Удаление заметки:**
```
DELETE /api/v1/notes/{id}
```

## Примеры тестирования

### Через командную строку (PowerShell):
```powershell
# Создание заметки
Invoke-WebRequest -Uri "http://localhost:8080/api/v1/notes" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"title":"Первая заметка","content":"Содержание первой заметки"}'

# Получение списка заметок
Invoke-WebRequest -Uri "http://localhost:8080/api/v1/notes" -Method GET

# Поиск заметок
Invoke-WebRequest -Uri "http://localhost:8080/api/v1/notes?q=первая&limit=5" -Method GET

# Обновление заметки (замените {id} на реальный ID)
Invoke-WebRequest -Uri "http://localhost:8080/api/v1/notes/{id}" `
-Method PATCH `
-Headers @{"Content-Type"="application/json"} `
-Body '{"content":"Обновленное содержание"}'

# Удаление заметки
Invoke-WebRequest -Uri "http://localhost:8080/api/v1/notes/{id}" -Method DELETE
```

### Через командную строку (curl):
```bash
# Создание заметки
curl -X POST http://localhost:8080/api/v1/notes \
-H "Content-Type: application/json" \
-d '{"title":"Первая заметка","content":"Содержание первой заметки"}'

# Получение списка заметок
curl "http://localhost:8080/api/v1/notes"

# Поиск заметок
curl "http://localhost:8080/api/v1/notes?q=первая&limit=5"

# Обновление заметки
curl -X PATCH http://localhost:8080/api/v1/notes/{id} \
-H "Content-Type: application/json" \
-d '{"content":"Обновленное содержание"}'

# Удаление заметки
curl -X DELETE http://localhost:8080/api/v1/notes/{id}
```

## Модель данных

### Note (Заметка)
```go
type Note struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Title     string             `bson:"title"         json:"title"`
    Content   string             `bson:"content"       json:"content"`
    CreatedAt time.Time          `bson:"createdAt"     json:"createdAt"`
    UpdatedAt time.Time          `bson:"updatedAt"     json:"updatedAt"`
}
```

## Особенности реализации

### Автоматические индексы
- **Уникальный индекс** на поле `title` для предотвращения дубликатов
- **Автоматическое создание** индексов при запуске приложения

### Обработка ошибок
- **404 Not Found** - когда заметка не найдена
- **409 Conflict** - при нарушении уникальности заголовка
- **400 Bad Request** - при неверных входных данных
- **500 Internal Server Error** - при ошибках сервера

### Таймауты
- **10 секунд** - таймаут подключения к MongoDB
- **3 секунды** - таймаут проверки подключения (ping)
- **5 секунд** - таймаут для каждого HTTP запроса

## Требования

- Go версии 1.21 или выше
- Docker и Docker Compose
- MongoDB драйвер для Go

## Установка зависимостей
```bash
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/bson
go get github.com/go-chi/chi/v5
```

## Преимущества MongoDB

### Гибкость данных
```json
{
  "_id": "651fba2f9a...",
  "title": "Гибкая заметка",
  "content": "Можно легко добавлять новые поля",
  "tags": ["mongodb", "nosql", "гибкость"],
  "metadata": {
    "priority": "high",
    "category": "technical"
  },
  "createdAt": "2025-01-15T10:30:00Z",
  "updatedAt": "2025-01-15T10:35:00Z"
}
```

### Производительность
- **In-memory операции** - быстрая работа с данными
- **Горизонтальное масштабирование** - поддержка шардирования
- **Гибкие запросы** - мощная система агрегации

## Решение проблем

**Ошибка подключения к MongoDB:**
```
mongo connect: connection refused
```
Убедитесь, что Docker контейнер с MongoDB запущен:
```bash
docker-compose ps
```

**Ошибка аутентификации:**
```
authentication failed
```
Проверьте правильность логина и пароля в MONGO_URI

**Дубликат заголовка:**
```
duplicate key error
```
Заголовки заметок должны быть уникальными - используйте другой заголовок

**Неверный ID:**
```
invalid ObjectID
```
Убедитесь, что передаете корректный ObjectID в формате HEX

## Пример документа в MongoDB

После создания заметки через API, в MongoDB будет документ:
```json
{
  "_id": ObjectId("651fba2f9a2b5c4e8f9a1b2c"),
  "title": "Пример заметки",
  "content": "Текст заметки",
  "createdAt": ISODate("2025-01-15T10:30:00Z"),
  "updatedAt": ISODate("2025-01-15T10:30:00Z")
}
```

## Дополнительные команды MongoDB

### Просмотр данных
```javascript
// Подключение к базе
use pz8

// Просмотр всех коллекций
show collections

// Просмотр всех заметок
db.notes.find().pretty()

// Поиск по заголовку
db.notes.find({title: {$regex: "пример", $options: "i"}})

// Статистика коллекции
db.notes.stats()
```

### Управление индексами
```javascript
// Просмотр индексов
db.notes.getIndexes()

// Создание текстового индекса
db.notes.createIndex({title: "text", content: "text"})

// Удаление индекса
db.notes.dropIndex("title_1")
```

## Остановка сервера
Нажмите `Ctrl+C` в командной строке где запущен сервер.

## Остановка MongoDB
```bash
docker-compose down
```

## Особенности архитектуры

### Документная модель vs Реляционная
| MongoDB (Документная) | PostgreSQL (Реляционная) |
|----------------------|--------------------------|
| Гибкая схема | Строгая схема |
| JSON-подобные документы | Таблицы и строки |
| Вложенные объекты | Связи через JOIN |
| Горизонтальное масштабирование | Вертикальное масштабирование |
