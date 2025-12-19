# Практика 13 — Профилирование Go-приложения (pprof)

**ФИО**: Пряшников Дмитрий Максимович  
**Группа**: ПИМО-01-25  

![f_960531e1f54b16da.gif](Imagine%2Ff_960531e1f54b16da.gif)

## Цель работы
Научиться подключать и использовать профилировщик `pprof` для анализа CPU, памяти и горутин, измерять время выполнения функций и находить узкие места в коде.

## Структура проекта

```
pprof-lab/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   └── work/
│       ├── slow.go
│       ├── timer.go
│       └── slow_test.go
├── go.mod
└── README.md
```

## Клонирование и запуск проекта

### 1. Клонирование репозитория
```bash
git clone https://github.com/MrFandore/Goland.git
cd Goland/Practica_13
```

### 2. Запуск сервера с pprof
```bash
go run ./cmd/api/main.go
```

Сервер запустится на `178.72.139.210:8086`.  
Профилировщик будет доступен по адресу: http://178.72.139.210:8086/debug/pprof/
![img.png](Imagine%2Fimg.png)

### Эндпоинты
```http://178.72.139.210:8086/work```
![img.png](Imagine%2Fimg.png)

```http://178.72.139.210:8086/debug/pprof/```
![img_1.png](Imagine%2Fimg_1.png)
![img_2.png](Imagine%2Fimg_2.png)

```http://178.72.139.210:8086/debug/pprof/profile?seconds=30```
[profile](Imagine%2Fprofile)

### 3. Заключение
![img_3.png](Imagine%2Fimg_3.png)
