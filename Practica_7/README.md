# Практика 7 - Работа с Redis для кэширования данных

## Структура проекта
```
Practica_7/
├── internal/
│   └── cache/
│       └── cache.go
├── main.go
├── go.mod
├── go.sum
└── README.md
```

## Описание файлов
- **main.go** - главный файл для запуска HTTP сервера с эндпоинтами для работы с Redis
- **internal/cache/cache.go** - реализация кэша с методами для работы с Redis
- **go.mod** - файл зависимостей Go
- **go.sum** - контрольные суммы зависимостей

## Установка и запуск Redis

### 1. Установка Redis
**Windows:**
```bash
# Скачать с официального сайта Redis для Windows
# Или использовать WSL с Redis
```

**macOS:**
```bash
brew install redis
brew services start redis
```

**Linux (Ubuntu/Debian):**
```bash
sudo apt update
sudo apt install redis-server
sudo systemctl start redis-server
```

**Docker (все платформы):**
```bash
docker run --name redis -p 6379:6379 -d redis
```

### 2. Проверка работы Redis
```bash
# Подключиться к Redis CLI
redis-cli

# Проверить подключение
ping
# Ожидаемый ответ: PONG

# Тестовые команды
SET test "Hello Redis"
GET test
```

## Запуск проекта

### 1. Клонирование и настройка
```bash
git clone https://github.com/MrFandore/Goland.git
cd Goland/Practica_7
```

### 2. Установка зависимостей
```bash
go mod download
```

### 3. Запуск приложения
```bash
go run main.go
```

## Проверка работы

После запуска сервер доступен по адресу: http://localhost:8080

### Доступные эндпоинты:

**Установка значения с TTL:**
```
GET /set?key=имя_ключа&value=значение
```
Ожидаемый результат: `OK: ключ=значение (TTL 10s)`

**Получение значения:**
```
GET /get?key=имя_ключа
```
Ожидаемый результат: `VALUE: ключ=значение`

**Проверка времени жизни ключа:**
```
GET /ttl?key=имя_ключа
```
Ожидаемый результат: `TTL for ключ: оставшееся_время`

## Примеры тестирования

### Через командную строку (PowerShell):
```powershell
# Установка значения
Invoke-WebRequest -Uri "http://localhost:8080/set?key=username&value=JohnDoe" -Method GET

# Получение значения
Invoke-WebRequest -Uri "http://localhost:8080/get?key=username" -Method GET

# Проверка TTL
Invoke-WebRequest -Uri "http://localhost:8080/ttl?key=username" -Method GET

# Установка другого значения
Invoke-WebRequest -Uri "http://localhost:8080/set?key=email&value=test@example.com" -Method GET
```

### Через командную строку (curl):
```bash
# Установка значения
curl "http://localhost:8080/set?key=username&value=JohnDoe"

# Получение значения
curl "http://localhost:8080/get?key=username"

# Проверка TTL
curl "http://localhost:8080/ttl?key=username"

# Установка другого значения
curl "http://localhost:8080/set?key=email&value=test@example.com"
```

### Через браузер:
Просто перейдите по ссылкам в адресной строке:
```
http://localhost:8080/set?key=test&value=hello
http://localhost:8080/get?key=test
http://localhost:8080/ttl?key=test
```

## Реализация кэша

### Структура кэша
```go
type Cache struct {
    rdb *redis.Client
}
```

### Доступные методы

**Set - сохранение значения:**
```go
func (c *Cache) Set(key string, value string, ttl time.Duration) error
```
Сохраняет значение в Redis с указанным временем жизни (TTL).

**Get - получение значения:**
```go
func (c *Cache) Get(key string) (string, error)
```
Возвращает значение по ключу. Если ключ не существует или истек TTL, возвращает ошибку.

**TTL - проверка времени жизни:**
```go
func (c *Cache) TTL(key string) (time.Duration, error)
```
Возвращает оставшееся время жизни ключа.

## Особенности реализации

- **Автоматическое истечение** - ключи автоматически удаляются через 10 секунд
- **In-memory хранилище** - все данные хранятся в оперативной памяти Redis
- **Высокая производительность** - время ответа измеряется миллисекундами
- **Простота использования** - минимальный API для базовых операций

## Пример использования в коде

```go
// Создание клиента Redis
c := cache.New("localhost:6379")

// Сохранение значения на 10 секунд
err := c.Set("user:1", "Alice", 10*time.Second)

// Получение значения
value, err := c.Get("user:1")

// Проверка TTL
ttl, err := c.TTL("user:1")
```

## Требования

- Go версии 1.21 или выше
- Redis сервер версии 6 или выше
- Установленная библиотека: `github.com/redis/go-redis/v9`

## Установка зависимостей
```bash
go get github.com/redis/go-redis/v9
```

## Решение проблем

**Ошибка подключения к Redis:**
```
connect: connection refused
```
Убедитесь, что Redis сервер запущен на localhost:6379

**Ключ не найден:**
```
key not found
```
Ключ либо не существует, либо истек его TTL

**Redis не установлен:**
Используйте Docker для быстрого запуска:
```bash
docker run --name redis -p 6379:6379 -d redis
```

**Порт занят:**
Измените порт в файле main.go с ":8080" на другой

## Преимущества использования Redis

1. **Высокая скорость** - операции в памяти выполняются за микросекунды
2. **Автоматическое истечение** - TTL для автоматической очистки данных
3. **Простота использования** - интуитивно понятный API
4. **Надежность** - отказоустойчивая архитектура
5. **Масштабируемость** - поддержка кластеризации

## Сценарии использования

### Кэширование базы данных
```go
// Псевдокод для кэширования результатов БД
func GetUserProfile(userID int) (*User, error) {
    // Сначала проверяем кэш
    if cached, err := cache.Get(fmt.Sprintf("user:%d", userID)); err == nil {
        return parseUser(cached), nil
    }
    
    // Если нет в кэше, идем в базу
    user, err := db.GetUser(userID)
    if err != nil {
        return nil, err
    }
    
    // Сохраняем в кэш на 5 минут
    cache.Set(fmt.Sprintf("user:%d", userID), user.ToJSON(), 5*time.Minute)
    return user, nil
}
```

### Хранение сессий
```go
// Создание сессии
sessionID := generateSessionID()
cache.Set(sessionID, userData, 24*time.Hour)

// Проверка сессии
userData, err := cache.Get(sessionID)
```

## Остановка сервера
Нажмите `Ctrl+C` в командной строке где запущен сервер.

## Остановка Redis
**Если Redis запущен через Docker:**
```bash
docker stop redis
```

**Если Redis запущен как служба:**
```bash
# macOS
brew services stop redis

# Linux
sudo systemctl stop redis-server
```

## Дополнительные команды Redis

Для более сложных операций можно использовать Redis CLI:
```bash
redis-cli

# Просмотр всех ключей
KEYS *

# Удаление ключа
DEL key_name

# Установка TTL для существующего ключа
EXPIRE key_name 60

# Просмотр информации о Redis
INFO
```
