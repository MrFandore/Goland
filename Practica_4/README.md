# Практика 4 - CRUD сервис задач с Chi роутером

## Структура проекта
```
Practica_4/
├── internal/
│   └── task/
│       ├── model.go
│       ├── repo.go
│       └── handler.go
├── pkg/
│   └── middleware/
│       ├── logger.go
│       └── cors.go
├── main.go
├── go.mod
├── requests.md
└── README.md
```

## Описание файлов
- **main.go** - главный файл для запуска сервера
- **internal/task/model.go** - структура задачи
- **internal/task/repo.go** - хранилище задач в памяти
- **internal/task/handler.go** - обработчики CRUD операций
- **pkg/middleware/logger.go** - логирование запросов
- **pkg/middleware/cors.go** - настройки CORS
- **go.mod** - файл зависимостей Go
- **requests.md** - примеры запросов

## Запуск проекта

### 1. Скачайте проект
```bash
git clone https://github.com/MrFandore/Goland.git
cd Goland/Practica_4
```

### 2. Установите зависимости
```bash
go mod download
```

### 3. Запустите сервер
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
Ожидаемый результат: `OK`

**Получить список всех задач:**
```
GET /api/tasks
```
Ожидаемый результат: список задач в формате JSON

**Создать новую задачу:**
```
POST /api/tasks
```
Тело запроса: `{"title":"Название задачи"}`

**Получить задачу по ID:**
```
GET /api/tasks/1
```
Ожидаемый результат: информация о задаче с ID=1

**Обновить задачу:**
```
PUT /api/tasks/1
```
Тело запроса: `{"title":"Новое название","done":true}`

**Удалить задачу:**
```
DELETE /api/tasks/1
```
Ожидаемый результат: статус 204 без тела

## Примеры тестирования

### Через командную строку (PowerShell):
```powershell
Invoke-WebRequest -Uri http://localhost:8080/api/tasks `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"title":"Выучить chi"}'

Invoke-WebRequest -Uri http://localhost:8080/api/tasks `
-Method GET

Invoke-WebRequest -Uri http://localhost:8080/api/tasks/1 `
-Method GET

Invoke-WebRequest -Uri http://localhost:8080/api/tasks/1 `
-Method PUT `
-Headers @{"Content-Type"="application/json"} `
-Body '{"title":"Выучить chi глубже","done":true}'

Invoke-WebRequest -Uri http://localhost:8080/api/tasks/1 `
-Method DELETE

Invoke-WebRequest -Uri http://localhost:8080/api/tasks `
-Method GET | Select-Object -ExpandProperty Content

```

### Через командную строку (curl):
```bash
# Создать задачу
curl -X POST http://localhost:8080/api/tasks -H "Content-Type: application/json" -d '{"title":"Выучить chi"}'

# Получить список задач
curl http://localhost:8080/api/tasks

# Получить задачу по ID
curl http://localhost:8080/api/tasks/1

# Обновить задачу
curl -X PUT http://localhost:8080/api/tasks/1 -H "Content-Type: application/json" -d '{"title":"Выучить chi глубже","done":true}'

# Удалить задачу
curl -X DELETE http://localhost:8080/api/tasks/1
```

## Особенности проекта
- Использование Chi роутера для маршрутизации
- Полный CRUD (Create, Read, Update, Delete) для задач
- Хранение данных в оперативной памяти
- Автоматическая генерация ID задач
- Логирование всех запросов
- Поддержка CORS для веб-приложений
- Валидация входных данных

## Требования
- Установленный Go версии 1.23 или выше
- Установленный Git

## Решение проблем

**Порт занят:**
Измените порт в файле main.go с ":8080" на другой

**Go не установлен:**
Скачайте и установите Go с https://golang.org/dl/

**Файлы не найдены:**
Убедитесь, что вы в папке Practica_4 и присутствует файл go.mod

## Остановка сервера
Нажмите Ctrl+C в командной строке где запущен сервер
