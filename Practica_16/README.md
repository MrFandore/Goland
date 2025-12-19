# Практическое задание №16: Интеграционное тестирование API. Использование Docker для тестовой БД

**ФИО:** Пряшников Дмитрий Максимович
**Группа:** ПИМО-01-25

![cat-silly.gif](Imagine%2Fcat-silly.gif)
---

## Цель задания
- Освоить интеграционное тестирование REST API: проверка цепочки «маршрут → хендлер → сервис → репозиторий → реальная БД».
- Научиться поднимать изолированную тестовую среду БД в Docker.
- Освоить два подхода к инфраструктуре тестов:
  - **A.** Локальная среда через `docker-compose` (просто и наглядно).
  - **B.** Программный подъём контейнеров через `testcontainers-go` (изолированно и удобно для CI).
- Научиться инициализировать схему БД (миграции/auto-migrate), сидировать тестовые данные, очищать окружение.
- Внедрить интеграционные проверки CRUD-эндпоинтов (статусы, заголовки, JSON-ответы, эффекты в БД).

---

## Структура проекта

```
Practica_16/
├── README.md
├── docker-compose.yml
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── db/
│   │   └── migrate.go
│   ├── httpapi/
│   │   └── handlers.go
│   ├── models/
│   │   └── note.go
│   ├── repo/
│   │   └── postgres.go
│   ├── service/
│   │   └── service.go
│   └── integration/                      
│       ├── notes_integration_test.go     
│       └── notes_tc_integration_test.go                         
├── go.mod
└── go.sum
```

#### Запуск проекта:
### 1. Клонируем репозиторий:
```bash
git clone https://github.com/MrFandore/Goland.git
cd Goland/Practica_16
```

### 2. Идем и проверяем как оно работает ручками
```
curl -i -X POST http://178.72.139.210:8089/notes \
-H 'Content-Type: application/json' \
-d '{"title":"Hello","content":"World"}' 
```
![img.png](Imagine%2Fimg.png)

```
curl -i http://178.72.139.210:8089/notes/1
```
![img_1.png](Imagine%2Fimg_1.png)

### 3. Тестики тестики тестики...
![img_2.png](Imagine%2Fimg_2.png)
![img_3.png](Imagine%2Fimg_3.png)

### 4. Заключение
![img_4.png](Imagine%2Fimg_4.png)