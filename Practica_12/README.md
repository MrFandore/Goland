# Практика 12 — Подключение Swagger/OpenAPI

**ФИО**: Пряшников Дмитрий Максимович  
**Группа**: ПИМО-01-25

![cat-silly.gif](Imagine%2Fcat-silly.gif)
## Цель работы
Подключить автоматическую генерацию документации OpenAPI/Swagger к проекту notes-api из практики 11, используя подход code-first с пакетом swaggo/swag.

## Структура проекта после подключения Swagger

```
notes-api/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── core/
│   │   └── note.go
│   ├── http/
│   │   ├── handlers/
│   │   │   └── notes.go
│   │   └── router.go
│   └── repo/
│       └── note_mem.go
├── docs/            
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── api/
│   └── openapi.yaml        
├── go.mod
├── go.sum
└── README.md
```

## Установка зависимостей

```bash
go get github.com/swaggo/http-swagger
go install github.com/swaggo/swag/cmd/swag@latest
```

## Добавление аннотаций в код

### Верхнеуровневые аннотации (cmd/api/main.go)
```go
// Package main Notes API server.
//
// @title Notes API
// @version 1.0
// @description Учебный REST API для заметок (CRUD).
// @contact.name Backend Course
// @contact.email example@university.ru
// @BasePath /api/v1
package main
```

### Аннотации над обработчиками (internal/http/handlers/notes.go)
```go
// ListNotes godoc
// @Summary Список заметок
// @Description Возвращает список заметок с пагинацией и фильтром по заголовку
// @Tags notes
// @Param page query int false "Номер страницы"
// @Param limit query int false "Размер страницы"
// @Param q query string false "Поиск по title"
// @Success 200 {array} core.Note
// @Header 200 {integer} X-Total-Count "Общее количество"
// @Failure 500 {object} map[string]string
// @Router /notes [get]
func (h *Handlers) ListNotes(w http.ResponseWriter, r *http.Request) { /* ... */ }
```

## Генерация документации

```bash
swag init -g cmd/api/main.go -o docs
```

## Подключение Swagger UI к серверу

В router.go добавить:
```go
import httpSwagger "github.com/swaggo/http-swagger"

r.Get("/docs/*", httpSwagger.WrapHandler)
```

## Запуск проекта

```bash
go run ./cmd/api/main.go
```

Документация будет доступна по адресу: http://localhost:8080/docs/index.html

## Отчёт по работе

### 1. Фрагмент аннотаций над методами
```go
// CreateNote godoc
// @Summary Создать заметку
// @Tags notes
// @Accept json
// @Produce json
// @Param input body NoteCreate true "Данные новой заметки"
// @Success 201 {object} core.Note
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /notes [post]
func (h *Handlers) CreateNote(w http.ResponseWriter, r *http.Request) { /* ... */ }
```

### 2. Скриншот работающей страницы Swagger UI

![img.png](..%2FPractica_13%2FImagine%2Fimg.png)
![img_1.png](Imagine%2Fimg_1.png)
![img_2.png](Imagine%2Fimg_2.png)

### 3. Заключение
![img_3.png](Imagine%2Fimg_3.png)

