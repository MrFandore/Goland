# Практика 11 - Проектирование REST API (CRUD для заметок)
**ФИО**: Пряшников Дмитрий Максимович  
**Группа**: ПИМО-01-25

![f_960531e1f54b16da.gif](Imagine%2Ff_960531e1f54b16da.gif)

## Структура проекта
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
├── api/
│   └── openapi.yaml
├── go.mod
├── go.sum
└── README.md
```

## Описание файлов
- **cmd/api/main.go** - точка входа приложения, инициализация сервера
- **internal/core/note.go** - модель данных "Заметка"
- **internal/http/handlers/notes.go** - HTTP-обработчики для операций с заметками
- **internal/http/router.go** - настройка маршрутизации API
- **internal/repo/note_mem.go** - in-memory репозиторий для хранения заметок
- **api/openapi.yaml** - документация OpenAPI/Swagger
- **go.mod** - файл зависимостей Go

## Запуск проекта

### 1. Клонирование репозитория
```bash
git clone https://github.com/MrFandore/Goland.git
cd Goland/Practica_13
```

### 2. Проверка установленных инструментов
```bash
# Проверка версии Go
go version

# Проверка версии Git
git --version
```

### 3. Установка зависимостей
```bash
# Инициализация модуля Go
go mod init notes-api

# Установка зависимостей
go mod tidy

# Или установите chi router напрямую
go get github.com/go-chi/chi/v5
```

### 4. Запуск сервера
```bash
# Запуск приложения
go run ./cmd/api/main.go
```

### 5. Проверка работы
После запуска сервер будет доступен по адресу: `http://localhost:8085`



Для проверки работы API выполните:
```bash
# Создание тестовой заметки
curl -X POST http://localhost:8085/api/v1/notes \
  -H "Content-Type: application/json" \
  -d '{"title":"Тестовая заметка","content":"Это проверка API"}'

# Получение всех заметок
curl http://localhost:8085/api/v1/notes
```


### 6. Остановка сервера
Для остановки сервера нажмите `Ctrl+C` в терминале.

### Примечания
- Убедитесь, что порт 8080 свободен
- Для macOS/Linux могут потребоваться права доступа
- При возникновении ошибок проверьте путь к проекту и структуру каталогов

### Заключение
![img_2.png](Imagine%2Fimg_2.png)
