# Практика 5 - Работа с PostgreSQL в Go

## Структура проекта
```
Practica_5/
├── main.go
├── db.go
├── repository.go
├── go.mod
├── go.sum
├── .env.example
└── README.md
```

## Описание файлов
- **main.go** - главный файл для запуска приложения с демонстрацией всех операций
- **db.go** - подключение к PostgreSQL и настройка пула соединений
- **repository.go** - репозиторий с методами для работы с задачами в БД
- **go.mod** - файл зависимостей Go
- **go.sum** - контрольные суммы зависимостей
- **.env.example** - пример файла с переменными окружения

## Настройка базы данных

### 1. Установка и запуск PostgreSQL
**Windows:**
```bash
# Скачать с официального сайта и установить
# Запустить SQL Shell (psql)
```

**macOS:**
```bash
brew install postgresql
brew services start postgresql
```

**Linux (Ubuntu/Debian):**
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
```

### 2. Создание базы данных и таблицы
```sql
-- Подключиться к PostgreSQL
psql -U postgres

-- Создать базу данных
CREATE DATABASE todo;

-- Подключиться к базе
\c todo

-- Создать таблицу задач
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    done BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Проверить создание
SELECT * FROM tasks;
```

## Запуск проекта

### 1. Клонирование и настройка
```bash
git clone https://github.com/MrFandore/Goland.git
cd Goland/Practica_5
```

### 2. Установка зависимостей
```bash
go mod download
```

### 3. Настройка подключения к БД

**Вариант 1: Через переменные окружения**
```bash
# Создайте файл .env
echo "DATABASE_URL=postgres://postgres:ваш_пароль@localhost:5432/todo?sslmode=disable" > .env
```

**Вариант 2: Прямое указание в коде (только для разработки)**
Измените DSN в файле `main.go`:
```go
dsn = "postgres://postgres:ваш_пароль@localhost:5432/todo?sslmode=disable"
```

### 4. Запуск приложения
```bash
go run .
```

## Демонстрация работы

После запуска приложение автоматически выполнит:

### 1. Создание тестовых задач
- "Сделать ПЗ №5"
- "Купить кофе" 
- "Проверить отчёты"

### 2. Вывод всех задач
```
=== Tasks ===
#1 | Сделать ПЗ №5             | done=false | 2025-01-15T10:30:00Z
#2 | Купить кофе               | done=false | 2025-01-15T10:30:01Z
#3 | Проверить отчёты          | done=false | 2025-01-15T10:30:02Z
```

### 3. Фильтрация невыполненных задач
```
=== Undone Tasks ===
#1 | Сделать ПЗ №5             | done=false | 2025-01-15T10:30:00Z
#2 | Купить кофе               | done=false | 2025-01-15T10:30:01Z
#3 | Проверить отчёты          | done=false | 2025-01-15T10:30:02Z
```

### 4. Поиск задачи по ID
```
=== Task with ID=1 ===
#1 | Сделать ПЗ №5             | done=false | 2025-01-15T10:30:00Z
```

### 5. Массовое добавление задач
```
Массовое добавление задач завершено
```

## Доступные методы репозитория

### CreateTask
```go
id, err := repo.CreateTask(ctx, "Новая задача")
```
Создает одну задачу и возвращает её ID.

### ListTasks
```go
tasks, err := repo.ListTasks(ctx)
```
Возвращает все задачи из базы данных.

### ListDone
```go
// Невыполненные задачи
undoneTasks, err := repo.ListDone(ctx, false)

// Выполненные задачи  
doneTasks, err := repo.ListDone(ctx, true)
```
Возвращает задачи по статусу выполнения.

### FindByID
```go
task, err := repo.FindByID(ctx, 1)
```
Находит задачу по указанному ID.

### CreateMany
```go
titles := []string{"Задача 1", "Задача 2", "Задача 3"}
err := repo.CreateMany(ctx, titles)
```
Массовое добавление задач через транзакцию.

## Проверка через psql

После выполнения программы проверьте данные в БД:
```sql
\c todo
SELECT id, title, done, created_at FROM tasks ORDER BY id;
```

## Настройки пула соединений

В файле `db.go` настроены оптимальные параметры пула:

```go
db.SetMaxOpenConns(10)     // Максимум активных соединений
db.SetMaxIdleConns(5)      // Соединений в простое  
db.SetConnMaxLifetime(30 * time.Minute) // Время жизни соединения
```

**Обоснование настроек:**
- `MaxOpenConns=10` - оптимально для 4-ядерного процессора
- `MaxIdleConns=5` - баланс между скоростью и использованием памяти
- `ConnMaxLifetime=30m` - защита от "протухания" соединений

## Требования

- Go версии 1.21 или выше
- PostgreSQL 14 или выше
- Установленный драйвер pgx

## Установка драйвера PostgreSQL
```bash
go get github.com/jackc/pgx/v5/stdlib
go get github.com/joho/godotenv
```

## Решение проблем

**Ошибка подключения:**
```
openDB error: dial error
```
Проверьте:
- Запущен ли PostgreSQL
- Правильность пароля в DSN
- Наличие базы данных `todo`

**Ошибка аутентификации:**
```
pq: password authentication failed
```
Убедитесь, что используете правильный пароль для пользователя postgres.

**Таблица не найдена:**
```
relation "tasks" does not exist
```
Выполните SQL-скрипт создания таблицы из раздела "Настройка базы данных".

## Особенности реализации

- **Параметризованные запросы** - защита от SQL-инъекций через плейсхолдеры $1, $2
- **Транзакции** - массовое добавление через BeginTx/Commit
- **Context** - все методы принимают context для таймаутов и отмены
- **Connection Pool** - оптимальные настройки пула соединений
- **Сканирование результатов** - корректное преобразование типов PostgreSQL → Go

## Остановка приложения
Нажмите `Ctrl+C` в командной строке где запущено приложение.
