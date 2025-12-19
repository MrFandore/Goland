# Практическое задание №14: Оптимизация запросов к БД. Использование connection pool

**ФИО**: Пряшников Дмитрий Максимович
**Группа**: ПИМО-01-25

![cat-silly.gif](Imagine%2Fcat-silly.gif)

**Цель**: Оптимизация запросов к БД. Использование connection pool

## Задание
1. Научиться находить «узкие места» в SQL-запросах и устранять их (индексы, переписывание запросов, пагинация, батчинг)
2. Освоить настройку пула подключений (connection pool) в Go и параметры его тюнинга
3. Научиться использовать EXPLAIN/ANALYZE, базовые метрики (pg_stat_statements), подготовленные запросы и транзакции
4. Применить техники уменьшения N+1 запросов и сокращения аллокаций на горячем пути

## Описание проекта и требования
### Структура проекта:
```
├── README.md
├── cmd
│   └── api
│       └── main.go
├── docker-compose.yaml
├── go.mod
├── go.sum
└── internal
    ├── config
    │   └── config.go
    ├── model
    │   └── note.go
    ├── pagination
    │   └── cursor.go
    ├── storage
    │   ├── postgres
    │   │   ├── queries.go
    │   │   └── repo.go
    │   └── redis
    │       └── cache.go
    └── transport
        └── http
            ├── handlers.go
            ├── respond.go
            └── server.go
```

## Запуск проекта

### 1. Клонирование репозитория
```bash
git clone https://github.com/MrFandore/Goland.git
cd Goland/Practica_14
```

### 2. Пример файла .env
```env
# Remote Postgres
DB_DSN=postgres://root:root@http://address:5432/pz9_bcrypt?sslmode=disable

# HTTP
HTTP_ADDR=:8087

# Redis (локально через docker compose)
REDIS_ADDR=127.0.0.1:6379
REDIS_PASSWORD=kek
REDIS_DB=0
CACHE_TTL_SECONDS=45
```

### 6. Проверка того что оно работает
![img.png](Imagine%2Fimg.png)    
![img_1.png](Imagine%2Fimg_1.png)
## Примеры запросов к API

### Создание заметки
```bash
curl -s -X POST http://178.72.139.210:8087/notes \
  -H 'Content-Type: application/json' \
  -d '{"title":"redis tips","content":"..."}'
```
![img_2.png](Imagine%2Fimg_2.png)

### Получение заметки по ID (с кэшированием второго запроса)
```bash
curl -s http://178.72.139.210:8087/notes/1
curl -s http://178.72.139.210:8087/notes/1
```
![img_3.png](Imagine%2Fimg_3.png)

### Список с keyset-пагинацией + поиск по title (FTS)
```bash
curl -s "http://178.72.139.210:8087/notes?limit=20&q=redis"
curl -s "http://178.72.139.210:8087/notes?limit=20&q=redis&cursor=eyJjcmVhdGVkX2F0IjoiMjAyNS0xMi0xMlQxMDoyNzozNy43NjU3WiIsImlkIjoxfQ"
```
![img_6.png](Imagine%2Fimg_6.png)

### Обновление заметки
```bash
curl -s -X PATCH http://178.72.139.210:8087/notes/1 \
  -H 'Content-Type: application/json' \
  -d '{"title":"redis tips v2","content":"updated"}'
```
![img_4.png](Imagine%2Fimg_4.png)

### Пакетное получение заметок (батч вместо N+1)
```bash
curl -s "http://178.72.139.210:8087/notes/batch?ids=1,2,3,4"
```
![img_5.png](Imagine%2Fimg_5.png)

## Оптимизации, примененные в проекте

### Проблемы до оптимизации:
1. **Пагинация через OFFSET**: время выполнения запросов росло почти линейно с номером страницы, так как PostgreSQL приходилось сканировать все предыдущие строки
2. **Проблема N+1 запросов**: множество отдельных запросов по первичному ключу увеличивало время ответа из-за network round-trips
3. **Неэффективный поиск по title**: запросы не использовали GIN-индекс, что приводило к полному сканированию таблицы (Seq Scan)

### Примененные оптимизации:
1. **Переписывание запросов и добавление индексов**:
    - Замена OFFSET на keyset-пагинацию по паре (created_at, id)
    - Замена N+1 запросов на батч-запрос с использованием `ANY`
    - Приведение поисковых запросов к форме, соответствующей индексам

2. **Индексы**:
    - Композитный B-tree индекс на (created_at, id) для keyset-пагинации
    - GIN индекс на tsvector для полнотекстового поиска (FTS)

3. **Настройка пула подключений**:
    - MaxConns: 20
    - MinConns: 5
    - MaxConnLifetime: 1 час


### 7. **Заключение**
![img_6.png](..%2FPractica_16%2FImagine%2Fimg_6.png)
