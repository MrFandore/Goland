# Практическое задание №15: Unit-тестирование функций (testing, testify)

**ФИО**: Пряшников Дмитрий Максимович  
**Группа**: ПИМО-01-25

![f_960531e1f54b16da.gif](Imagine%2Ff_960531e1f54b16da.gif)
---
## Задание:
* Освоить базовые приёмы unit-тестирования в Go с помощью стандартного пакета testing.
* Научиться писать табличные тесты, подзадачи t.Run, тестировать ошибки и паники.
* Освоить библиотеку утверждений testify (assert, require) для лаконичных проверок.
* Научиться измерять покрытие кода (go test -cover) и формировать html-отчёт покрытия.
* Подготовить минимальную структуру проектных тестов и общий чек-лист качества тестов.
---

### Описание проекта и требования:

#### Структура проекта (практическая часть):
```
pz15-tests/
├── internal/
│   ├── mathx/
│   │   ├── mathx.go
│   │   └── mathx_test.go
│   ├── stringsx/
│   │   ├── stringsx.go
│   │   └── stringsx_test.go
│   └── service/
│       ├── repo.go
│       ├── service.go
│       └── service_test.go
├── go.mod
└── go.sum
```

#### Запуск проекта:
### 1. **Клонируем репозиторий:**
```bash
git clone https://github.com/MrFandore/Goland.git
cd Goland/Practica_15
```

### 2. **Тестим работу:**

 * ```go test ./...```
![img.png](Imagine%2Fimg.png)
 * ```go test -v ./internal/...```
![img_1.png](Imagine%2Fimg_1.png)
 * ```go test -cover ./...```
![img_2.png](Imagine%2Fimg_2.png)
 * ```go test -coverprofile=coverage.out ./... ```
 * ```go tool cover -html=coverage.out```
![img_3.png](Imagine%2Fimg_3.png)

### 3. **Заключение**
![img_5.png](..%2FPractica_16%2FImagine%2Fimg_5.png)