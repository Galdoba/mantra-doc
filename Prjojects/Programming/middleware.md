---
updated_at: 2026-02-17T16:49:39.207+10:00
tags:
  - tips
---
# Общая идея
Полученное от `frontend` реквест, перед тем как добратьля до бизнес-логики может (а чаще должен) пройти ряд манипуляций (логирование, аутентификация и т.д.). Так же полезно было бы преобразовать данные запроса во вменяемый ввод для функции сервиса.

Все эти манипуляции чтобы не переносить из хэндлера в хэндлер, можно вынести в отдельные мини-хендлеры и составлять из них цепочку мидлварей. после чего все передать в муксер. Выглядеть это может так:

Объявляем интерфейс мидлваря:
```go
type Middleware interface { 
	Wrap(next http.Handler) http.Handler  //Это что-то что оборачивает http.Handler
	                                  //и возвращает следуюший хендлер
}
```

Далее объявляем пустые типы и привязываем к ним функцию обертку, внутри которой прописываем/возвращаем наш следующий блок.
ВАЖНО: если блок по какой-то причине не удался он может принудительно завершить свое выполнение. 
ВОПРОС: Для передачи данных между блоками мы можем пользоваться полями структур мидлваря?

```go
type AuthMiddleware struct {}

func (*AuthMiddleware) Wrap(next http.Handler) http.Handler {
 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  authToken := r.Header.Get("Authorization")
  if len(authToken) > 0 {
   next.ServeHTTP(w, r)
  } else {
   http.Error(w, "Unauthorized", http.StatusUnauthorized)
  }
 })
}

type LogMiddleware struct {}

func (*LogMiddleware) Wrap(next http.Handler) http.Handler {
 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  startTime := time.Now()
  next.ServeHTTP(w, r)
  endTime := time.Since(startTime)
  log.Printf("[%v] %s %s took %v", startTime.Format(time.RFC3339), r.Method, r.URL.Path, endTime)
 })
}

```

Затем мы все это сшиваем в главном файле:
```go
func main() {
 initDB()

 router := http.NewServeMux()

 middlewares := []Middleware{
  new(AuthMiddleware),
  new(LogMiddleware),
 }

 finalHandler := http.HandlerFunc(trackVisitHandler)
 for _, m := range middlewares {
  finalHandler = m.Wrap(finalHandler)
 }

 router.Handle("/trackvisit", finalHandler)

 server := &http.Server{
  Addr:           ":8080",
  Handler:        router,
  ReadTimeout:    10 * time.Second,
  WriteTimeout:   10 * time.Second,
  MaxHeaderBytes: 1 << 20,
 }

 fmt.Println("Server listening on port :8080...")
 err := server.ListenAndServe()
 if err != nil {
  log.Fatal(err)
 }
}
```

ВОПРОС: В этом подходе в каком порядке будут вызываться обернутые хендлеры?

---

Middleware в Go: Общая логика, механизмы и реализация

# Определение

Middleware («промежуточное ПО») — это компонент программного обеспечения, расположенный между клиентом и сервером, выполняющий дополнительные обработчики запросов перед передачей их конечному обработчику. Это позволяет внедрять кросс-функциональные возможности, такие как аутентификация, журналирование, сжатие и кэширование.


Интерфейсы и методы в HTTP-сервисах

# Логика работы метода Wrap

Метод Wrap используется для обертывания существующего обработчика (http.Handler) новым функционалом. Обработчик принимает два аргумента: оригинальный обработчик и новая логика, которая должна выполняться до или после основного обработчика.
```go
type HandlerWrapper func(http.Handler) http.Handler
```

Типичный шаблон метода Wrap выглядит следующим образом:
```go
func MyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Предварительная обработка запроса (1)
        
        next.ServeHTTP(w, r.WithContext(ctx)) //Вызов следующего уровня (2)

        // Постобработка запроса (3)
    })
}
```

*коментарий:* Нижние слои впихиваются в середину. Они начинаются позже, но заканчиваются раньше. ([[LIFO]])

## Механизмы передачи данных через context.Context
Данные передаются между уровнями middleware через объект контекста (context.Context), позволяющий хранить значения и устанавливать дедлайны для обработки запросов. Ключом к передаче данных служит создание дочернего контекста с необходимыми значениями и передача его дальше по цепочке обработчиков.

Пример передачи данных:
```go
ctx := context.WithValue(r.Context(), keyUserID, userID)
next.ServeHTTP(w, r.WithContext(ctx))
```

## Очередь выполнения хэндлеров
Обработчики выполняются последовательно, начиная с самого внешнего слоя и заканчивая внутренним основным обработчиком. Каждый промежуточный слой вызывает следующий обработчик посредством вызова метода ServeHTTP. Таким образом формируется цепочка вызовов.

Порядок выполнения важен, поскольку первый добавленный middleware будет первым вызванным, но последним завершённым (LIFO).

### Принудительное завершение обработки всей цепочки
Промежуточный обработчик может прервать выполнение остальной части цепочки путём завершения обработки непосредственно в своем обработчике. Например, возврат статуса ошибки:
```go
if err != nil {
    http.Error(w, err, http.StatusUnauthorized)//отправляет ответ клиенту с ошибкой
    return //принудительно заканчивает работу хэнлдера
}

```

Это предотвратит дальнейшую обработку текущего middleware и не запустит следующий, провоцируя тем самым постобработку уже вызывавшихся обработчиков.

### Пропуск отдельных уровней
Промежуточный обработчик может пропустить уровень обработки, передав управление следующему обработчику без изменений состояния. Это полезно, когда определённые условия позволяют избежать дополнительной обработки.
```go
if skipCondition() {
    next.ServeHTTP(w, r)
    return //после обработки код обычного сценария не выполняется
}
//обычный сценарий
```

## Перехват и обработка паники

Механизм перехвата паники позволяет ловить исключительные ситуации, возникающие в процессе обработки запроса. Промежуточный обработчик может перехватить панику и вернуть соответствующий HTTP-статус клиенту.
```go
defer func() {
    if rec := recover(); rec != nil {
        log.Println("Panic recovered:", rec)
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    }
}()
```

## Различные подходы к реализации middleware
### Использование отдельных пакетов middleware (последовательное подключение)

Этот подход подразумевает последовательное применение middleware вручную, путём компоновки функционала.

Код примера:

```go
package main

import (
    "log"
    "net/http"
)

// Простое middleware для логгирования
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("[%s] %s\n", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    }
}

// Middleware для проверки токенов
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("X-Auth-Token")
        if token == "" || token != "secret-token" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    }
}

// Основная бизнес-логика
func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!\n"))
}

func main() {
    handler := loggerMiddleware(authMiddleware(helloHandler))
    http.Handle("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Как это работает:

Каждое middleware обертывает следующее звено цепочки, формируя составной обработчик.
Порядок подключения middleware определяется явно при создании общего обработчика.

### Создание контейнеров для middleware (Router или Chain)

Второй подход предполагает использование специальных структур, позволяющих удобно конфигурировать цепочку middleware.

Пакеты и инструменты:

Один из популярных подходов в экосистеме Go — использование сторонних библиотек вроде Gorilla Mux или Echo Framework, однако здесь мы реализуем простой собственный контейнер для демонстрации идеи.

Код примера:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

type MiddlewareChain struct {
    handlers []func(http.HandlerFunc) http.HandlerFunc
    finalHandler http.HandlerFunc
}

// Добавляет новое middleware в цепочку
func (c *MiddlewareChain) Use(mw func(http.HandlerFunc) http.HandlerFunc) {
    c.handlers = append(c.handlers, mw)
}

// Выполняет построение итогового обработчика
func (c *MiddlewareChain) Build() http.HandlerFunc {
    var finalHandler = c.finalHandler
    for i := len(c.handlers)-1; i >= 0; i-- { // Применяем middleware в обратном порядке
        finalHandler = c.handlers[i](finalHandler)
    }
    return finalHandler
}

// Простое middleware для логгирования
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("[%s] %s\n", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    }
}

// Middleware для проверки токенов
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("X-Auth-Token")
        if token == "" || token != "secret-token" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    }
}

// Основная бизнес-логика
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    chain := &MiddlewareChain{
        finalHandler: helloHandler,
    }

    chain.Use(loggerMiddleware)
    chain.Use(authMiddleware)

    http.Handle("/", chain.Build())
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Как это работает:

Мы создали специальный объект MiddlewareChain, который хранит список middleware и автоматически строит общий обработчик.
Порядок применения middleware контролируется контейнером, позволяя легко расширять цепочку новыми компонентами.

Заключение

Первый подход проще в понимании и подходит для небольших проектов, где чёткий порядок middleware заранее определён и редко меняется. Второй подход удобнее для больших систем, где нужно гибкое управление middleware и частые изменения в конфигурациях.


## Шаблонная реализация метода Wrap
Ниже представлен общий шаблон реализации middleware-обертки:
```go
package main

import (
    "context"
    "log"
    "net/http"
)

// KeyUserID is a type-safe way to store the User ID in Context
type KeyUserID string

const userIDKey = KeyUserID("user_id")

// AuthMiddleware checks for authentication before proceeding with request handling
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userID := extractUserIDFromToken(r.Header.Get("Authorization"))
        if userID == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        ctx := context.WithValue(r.Context(), userIDKey, userID)
        r = r.WithContext(ctx)

        defer func() {
            if rec := recover(); rec != nil {
                log.Printf("Panic occurred during processing: %v\n", rec)
                http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            }
        }()

        next.ServeHTTP(w, r)
    })
}

// LoggingMiddleware logs incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s\n", r.Method, r.URL.Path, r.RemoteAddr)
        next.ServeHTTP(w, r)
    })
}

// Main entry point
func main() {
    mux := http.NewServeMux()
    mux.Handle("/", LoggingMiddleware(AuthMiddleware(http.HandlerFunc(handle))))

    http.ListenAndServe(":8080", mux)
}

// handle simulates an actual endpoint
func handle(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value(userIDKey).(string)
    w.Write([]byte("Hello, authenticated user " + userID))
}

// extractUserIDFromToken extracts user ID from token
func extractUserIDFromToken(token string) string {
    // Simulated extraction logic
    return "user123"
}
```

Этот пример демонстрирует простой цикл обработки запроса с использованием двух middleware: авторизации и логирования. В реальной практике middleware применяется для гораздо большего спектра задач, включая сжатие, кэширование, контроль производительности и многое другое.


---


# Типовая цепочка middleware (order of execution):

1. Logger: Создаёт логгер и привязывает его к запросу, регистрирует базовые метрики (методы, пути, IP, заголовки и др.).
	Цель: фиксировать происходящее на уровне инфраструктуры.
2. Recovery/Panic Handler: Перехватывает паники и обрабатывает исключения. Обычно логирует ошибку и формирует подходящий ответ клиенту (например, HTTP 500 Internal Server Error).
	Цель: обеспечить отказоустойчивость и предотвращение каскадирования ошибок.
3. Metrics Collector: Подключает сборщик метрик (Prometheus, DataDog и т.п.) для мониторинга производительности и здоровья сервисов.
	Цель: мониторинг нагрузки и эффективности работы приложений.
4. CORS Middleware: Обрабатывает CORS-заголовки, проверяя происхождение запросов и устанавливая необходимые заголовки Access-Control-Allow-Origin и т.д.
	Цель: защита от атак типа CSRF/CORS.
5. Rate Limiter: Ограничивает частоту запросов, защищаясь от DDoS-атак и злоупотребления API.
	Цель: предотвратить перегрузку сервера чрезмерным количеством запросов.
6. Request Validation: Базовая проверка валидности HTTP-запроса (правильность JSON, наличие обязательных полей и т.д.). Этот шаг часто осуществляется сразу после базовой аутентификации.
	Цель: избегать лишней нагрузки на последующие слои обработки.
7. Authentication: Определяет личность клиента (токены JWT, сессии, OAuth и т.д.).
	Цель: убедиться, что запрос поступил от доверенного субъекта.
8. Authorization: Устанавливает права доступа для прошедшего проверку пользователя. Может включать RBAC, ABAC и другие модели контроля доступа.
	Цель: ограничить доступ к ресурсам согласно ролям и полномочиям пользователя.
9. Context Enrichment: Заполняет контекст дополнительными полезными данными (IP пользователя, геолокация, временные зоны и т.д.).
	Цель: передать дополнительную информацию нижним уровням для принятия решений.
10. Business Logic Preprocessing: Преобразование входящего запроса в понятную форму для последующей бизнес-логики (например, десериализация JSON в структуры данных).
	Цель: подготовить данные для передачи в рабочие компоненты приложения.
11. Cache Layer: Кэширование результатов, если это применимо (например, запросы к редким данным или медленно изменяемым сущностям).
	Цель: ускорить выполнение последующих аналогичных запросов.
12. Main Business Logic: Основной обработчик, выполняющий специфическую работу (например, создание заказа, обновление профиля пользователя и т.д.).
	Цель: реализовать основную бизнес-функцию.
13. Response Postprocessing: Обратное преобразование результата бизнес-логики в пригодный для отправки клиенту вид (например, сериализация обратно в JSON/XML).
	Цель: подготовка финального ответа.
14. Error Handling: Если какая-то обработка завершилась неудачей, возвращаем осмысленное сообщение об ошибке пользователю.
	Цель: информировать пользователя о причинах сбоя и предложить возможные решения.

Сначала идут инфраструктурные middleware (logger, recovery, metrics collector).
Затем следуют защитные механизмы (rate limiter, authentication, authorization).
Далее располагаются уровни предварительной подготовки данных (request validation, context enrichment, preprocessing).
Потом идёт логика основного обработчика.
Завершают процесс постобработка и формирование конечного ответа.

Эта структура позволяет разделить ответственность каждого компонента.


---

# подходы к реализации middleware
Давай рассмотрим оба подхода на примере реализации в Go с использованием стандартной библиотеки net/http.

1. Использование отдельных пакетов middleware (последовательное подключение)

Этот подход подразумевает последовательное применение middleware вручную, путём компоновки функционала.

Код примера:

```go
package main

import (
    "log"
    "net/http"
)

// Простое middleware для логгирования
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("[%s] %s\n", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    }
}

// Middleware для проверки токенов
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("X-Auth-Token")
        if token == "" || token != "secret-token" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    }
}

// Основная бизнес-логика
func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!\n"))
}

func main() {
    handler := loggerMiddleware(authMiddleware(helloHandler))
    http.Handle("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Как это работает:

Каждое middleware обертывает следующее звено цепочки, формируя составной обработчик.
Порядок подключения middleware определяется явно при создании общего обработчика.

2. Создание контейнеров для middleware (Router или Chain)

Второй подход предполагает использование специальных структур, позволяющих удобно конфигурировать цепочку middleware.

Пакеты и инструменты:

Один из популярных подходов в экосистеме Go — использование сторонних библиотек вроде Gorilla Mux или Echo Framework, однако здесь мы реализуем простой собственный контейнер для демонстрации идеи.

Код примера:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

type MiddlewareChain struct {
    handlers []func(http.HandlerFunc) http.HandlerFunc
    finalHandler http.HandlerFunc
}

// Добавляет новое middleware в цепочку
func (c *MiddlewareChain) Use(mw func(http.HandlerFunc) http.HandlerFunc) {
    c.handlers = append(c.handlers, mw)
}

// Выполняет построение итогового обработчика
func (c *MiddlewareChain) Build() http.HandlerFunc {
    var finalHandler = c.finalHandler
    for i := len(c.handlers)-1; i >= 0; i-- { // Применяем middleware в обратном порядке
        finalHandler = c.handlers[i](finalHandler)
    }
    return finalHandler
}

// Простое middleware для логгирования
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("[%s] %s\n", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    }
}

// Middleware для проверки токенов
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("X-Auth-Token")
        if token == "" || token != "secret-token" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    }
}

// Основная бизнес-логика
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    chain := &MiddlewareChain{
        finalHandler: helloHandler,
    }

    chain.Use(loggerMiddleware)
    chain.Use(authMiddleware)

    http.Handle("/", chain.Build())
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Как это работает:

Мы создали специальный объект MiddlewareChain, который хранит список middleware и автоматически строит общий обработчик.
Порядок применения middleware контролируется контейнером, позволяя легко расширять цепочку новыми компонентами.

#### ВАЖНО
Если мы хотим создать кастомное Middleware то нужно делать дополнительную обертку:
```go
func NewLoggerMiddlewareHandler(global *log.Logger) func(http.HandlerFunc) http.HandlerFunc {//обертка с дополнительными параметрями возвращающая
	return func(next http.HandlerFunc) http.HandlerFunc {//HandleFunc, который
		return func(w http.ResponseWriter, r *http.Request) {//возвращает тело функции
			start := time.Now()
			next(w, r)
			if local, ok := r.Context().Value("request_logger").(*log.Logger); ok {
				local.Printf("request specific: %v", http.StatusAccepted)
			}
			duration := time.Since(start)
			global.Printf("global: %s %s completed for %v", r.Method, r.Pattern, duration)

		}
	}
}

```

Заключение

Первый подход проще в понимании и подходит для небольших проектов, где чёткий порядок middleware заранее определён и редко меняется. Второй подход удобнее для больших систем, где нужно гибкое управление middleware и частые изменения в конфигурациях.