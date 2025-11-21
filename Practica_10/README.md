# Практика 10 - JWT аутентификация и авторизация

## Структура проекта
```
Practica_10/
├── internal/
│   ├── core/
│   │   ├── service.go
│   │   └── user.go
│   ├── http/
│   │   ├── middleware/
│   │   │   ├── authn.go
│   │   │   └── authz.go
│   │   └── router.go
│   ├── platform/
│   │   ├── config/
│   │   │   └── config.go
│   │   └── jwt/
│   │       └── jwt.go
│   └── repo/
│       └── user_mem.go
├── main.go
├── go.mod
├── go.sum
└── README.md
```

## Описание файлов
- **main.go** - главный файл для запуска сервера
- **internal/core/service.go** - бизнес-логика и обработчики эндпоинтов
- **internal/core/user.go** - модель пользователя
- **internal/http/router.go** - настройка маршрутов и middleware
- **internal/http/middleware/authn.go** - аутентификация JWT токенов
- **internal/http/middleware/authz.go** - авторизация по ролям (RBAC)
- **internal/platform/config/config.go** - конфигурация приложения
- **internal/platform/jwt/jwt.go** - работа с JWT токенами (RS256)
- **internal/repo/user_mem.go** - хранилище пользователей в памяти
- **go.mod** - файл зависимостей Go

## Запуск проекта

### 1. Перейдите в папку проекта
```bash
cd Goland/Practica_10
```

### 2. Установите зависимости
```bash
go mod download
```

### 3. Запустите сервер
```bash
go run main.go
```

Или с настройкой переменных окружения:
```bash
export APP_PORT=8080
export JWT_SECRET=your-secret-key
export JWT_TTL=15m
go run main.go
```

## Проверка работы

После запуска сервер доступен по адресу: http://localhost:8080

### Доступные эндпоинты:

**Аутентификация:**
```
POST /api/v1/login
```
Тело запроса: `{"Email":"admin@example.com","Password":"secret123"}`
Ожидаемый результат: access и refresh токены

**Обновление токенов:**
```
POST /api/v1/refresh
```
Тело запроса: `{"refresh_token":"your-refresh-token"}`
Ожидаемый результат: новая пара токенов

**Получить информацию о текущем пользователе:**
```
GET /api/v1/me
```
Заголовок: `Authorization: Bearer <access-token>`

**Получить пользователя по ID (ABAC):**
```
GET /api/v1/users/{id}
```
Заголовок: `Authorization: Bearer <access-token>`
- Пользователи могут получать только свой профиль
- Админы могут получать любой профиль

**Статистика (только для админов):**
```
GET /api/v1/admin/stats
```
Заголовок: `Authorization: Bearer <access-token>`

## Примеры тестирования

### Тестовые пользователи:
- **Админ:** `admin@example.com` / `secret123`
- **Пользователь:** `user@example.com` / `secret123`

### Через командную строку (curl):

```bash
# Логин как админ
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"Email":"admin@example.com","Password":"secret123"}'

# Логин как пользователь
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"Email":"user@example.com","Password":"secret123"}'

# Получить информацию о себе (замените ACCESS_TOKEN)
curl http://localhost:8080/api/v1/me \
  -H "Authorization: Bearer ACCESS_TOKEN"

# Обновить токены (замените REFRESH_TOKEN)
curl -X POST http://localhost:8080/api/v1/refresh \
  -H "Content-Type: application/json" \
  -d '{"refresh_token":"REFRESH_TOKEN"}'

# Получить статистику (только для админа)
curl http://localhost:8080/api/v1/admin/stats \
  -H "Authorization: Bearer ACCESS_TOKEN"

# Получить пользователя по ID
curl http://localhost:8080/api/v1/users/1 \
  -H "Authorization: Bearer ACCESS_TOKEN"
```

### Через PowerShell:

```powershell
# Логин
$login = Invoke-WebRequest -Uri http://localhost:8080/api/v1/login `
  -Method POST `
  -Headers @{"Content-Type"="application/json"} `
  -Body '{"Email":"admin@example.com","Password":"secret123"}'

$tokens = $login.Content | ConvertFrom-Json
$accessToken = $tokens.access_token

# Получить информацию о себе
Invoke-WebRequest -Uri http://localhost:8080/api/v1/me `
  -Headers @{"Authorization"="Bearer $accessToken"}

# Получить статистику
Invoke-WebRequest -Uri http://localhost:8080/api/v1/admin/stats `
  -Headers @{"Authorization"="Bearer $accessToken"}
```

## Особенности проекта

- **JWT аутентификация** с access (15 мин) и refresh (7 дней) токенами
- **RS256 алгоритм** с поддержкой ротации ключей (kid)
- **RBAC авторизация** с ролями "admin" и "user"
- **ABAC правила** - пользователи могут читать только свой профиль
- **Rate limiting** - 5 попыток входа за 5 минут с одного IP
- **Blacklist refresh токенов** - предотвращение повторного использования
- **Stateless архитектура** - не требует хранения сессий на сервере
- **Middleware цепочка** - AuthN → AuthZ → обработчик

## Требования
- Установленный Go версии 1.24 или выше
- Установленный Git

## Решение проблем

**Порт занят:**
Измените порт через переменную окружения `APP_PORT`

**Ошибки JWT:**
Убедитесь, что установлены все зависимости: `go mod download`

**403 Forbidden:**
Проверьте роль пользователя и права доступа к эндпоинту

**401 Unauthorized:**
Проверьте корректность токена и срок его действия
