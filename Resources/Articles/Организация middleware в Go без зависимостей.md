---
updated_at: 2026-02-21T01:43:55.384+10:00
tags:
  - middleware
  - go_stdlib
  - routing
  - handlers
  - function_composition
  - route_grouping
  - dependencies
  - article
---
Источник:
https://habr.com/ru/companies/otus/articles/988234/

Много лет я использовал сторонние пакеты, чтобы удобнее структурировать и управлять middleware в Go-веб-приложениях. В небольших проектах я часто брал [alice](https://github.com/justinas/alice), чтобы собирать «цепочки» middleware, которые можно переиспользовать на разных маршрутах. А в более крупных приложениях, где много middleware и маршрутов, я обычно использовал роутер вроде [chi](https://github.com/go-chi/chi) или [flow](https://github.com/alexedwards/flow), чтобы делать вложенные «группы» маршрутов со своим набором middleware для каждой группы.

Но после того как в [Go](https://otus.pw/salX/) 1.22 в http.ServeMux появилась новая функциональность сопоставления по шаблонам (pattern matching), я по возможности стал убирать сторонние зависимости из логики маршрутизации и переходить на одну лишь стандартную библиотеку.

Однако полный переход на стандартную библиотеку оставляет хороший вопрос: *как организовать и управлять middleware без использования сторонних пакетов?*

### Почему управление middleware — это проблема?

Если в приложении всего несколько маршрутов и middleware-функций, проще всего оборачивать ваши обработчики в нужные middleware для каждого маршрута отдельно. Примерно так:

```go
// На этом маршруте middleware нет.
mux.Handle("GET /static/", http.FileServerFS(ui.Files))

// Оба этих маршрута используют middleware requestID и logRequest.
mux.Handle("GET /", requestID(logRequest(http.HandlerFunc(home))))
mux.Handle("GET /article/{id}", requestID(logRequest(http.HandlerFunc(showArticle))))

// На этом маршруте дополнительно используются middleware authenticateUser и requireAdminUser.
mux.Handle("GET /admin", requestID(logRequest(authenticateUser(requireAdminUser(http.HandlerFunc(showAdminDashboard))))))
```

Это работает и не требует внешних зависимостей, но вы, вероятно, представляете, какие минусы появятся по мере роста числа маршрутов:

- Есть повторения в объявлениях маршрутов.
- Становится сложнее читать код и быстро понимать, какие маршруты используют один и тот же набор middleware.
- Это выглядит довольно «хрупко»: в большом приложении, если нужно добавить, убрать или поменять местами middleware на множестве маршрутов, легко пропустить один из них и не заметить ошибку.

### Альтернатива alice

Как я вкратце упоминал выше, пакет [alice](https://github.com/justinas/alice) позволяет объявлять и переиспользовать «цепочки» middleware. Мы могли бы переписать пример выше, используя alice, вот так:

```go
mux := http.NewServeMux()

// Создаём базовую цепочку middleware.
baseChain := alice.New(requestID, logRequest)

// Расширяем базовую цепочку middleware аутентификации для маршрутов только для админов.
adminChain := baseChain.Append(authenticateUser, requireAdminUser)

// На этом маршруте middleware нет.
mux.Handle("GET /static/", http.FileServerFS(ui.Files))

// Публичные маршруты используют базовые middleware.
mux.Handle("GET /", baseChain.ThenFunc(home))
mux.Handle("GET /article/{id}", baseChain.ThenFunc(showArticle))

// Админские маршруты с дополнительными middleware аутентификации.
mux.Handle("GET /admin", adminChain.ThenFunc(showAdminDashboard))
```

На мой взгляд, этот код заметно чище, и в значительной степени снимает три проблемы, о которых мы говорили выше.

Но если вам не хочется добавлять `alice` в зависимости, можно воспользоваться функцией `slices.Backward`, появившейся в Go 1.23, и буквально в несколько строк написать собственный тип chain:

```go
type chain []func(http.Handler) http.Handler

func (c chain) thenFunc(h http.HandlerFunc) http.Handler {
    return c.then(h)
}

func (c chain) then(h http.Handler) http.Handler {
    for _, mw := range slices.Backward(c) {
        h = mw(h)
    }
    return h
}
```

После этого тип `chain` можно использовать при объявлении маршрутов вот так:

```go
mux := http.NewServeMux()

// Создаём базовую цепочку middleware.
baseChain := chain{requestID, logRequest}

// Расширяем базовую цепочку middleware аутентификации для маршрутов только для админов.
adminChain := append(baseChain, authenticateUser, requireAdminUser)

mux.Handle("GET /static/", http.FileServerFS(ui.Files))

mux.Handle("GET /", baseChain.thenFunc(home))
mux.Handle("GET /article/{id}", baseChain.thenFunc(showArticle))
mux.Handle("GET /admin", adminChain.thenFunc(showAdminDashboard))
```

Синтаксис здесь не в точности такой же, как в `alice`, но очень близкий, а по поведению это, по сути, то же самое.

Если вам интересно применить этот подход в своём коде, я выложил тесты для типа chain [в этом gist](https://gist.github.com/alexedwards/219d88ebdb9c0c9e74715d243f5b2136).

### Альтернатива chi и похожим роутерам

В крупных приложениях, когда у меня есть много разных middleware, которые используются на множестве разных маршрутов, функциональность группировки маршрутов, которую дают роутеры вроде [chi](https://github.com/go-chi/chi) и [flow](https://github.com/alexedwards/flow), всегда очень выручала.

По сути, они позволяют создавать *группы* маршрутов с определённым набором middleware. Причём эти группы можно вкладывать друг в друга: дочерние группы «наследуют» middleware родительских групп и могут дополнять их своим набором.

Давайте посмотрим на пример с использованием `chi` — насколько я помню, это был первый роутер, который поддержал такой стиль группировки маршрутов.

```go
r := chi.NewRouter()
r.Use(recoverPanic) // «Глобальный» middleware, используется на всех маршрутах.

r.Method("GET", "/static/", http.FileServerFS(ui.Files))

// Создаём группу маршрутов.
r.Group(func(r chi.Router) {
    // Добавляем middleware для группы.
    r.Use(requestID)
    r.Use(logRequest)

    // Маршруты, объявленные внутри группы, будут использовать этот набор middleware.
    r.Get("/", home)
    r.Get("/article/{id}", showArticle)

    // Создаём вложенную группу маршрутов. Любые маршруты в этой группе будут
    // использовать middleware, объявленные в самой группе, и в родительских группах.
    r.Group(func(r chi.Router) {
        r.Use(authenticateUser)
        r.Use(requireAdminUser)

        r.Get("/admin", showAdminDashboard)
    })
})
```

Но если вы хотите остаться в рамках стандартной библиотеки, то сделать собственную реализацию роутера, которая оборачивает `http.ServeMux` и поддерживает группы middleware в похожем стиле, не так уж сложно:

```go
type Router struct {
    globalChain []func(http.Handler) http.Handler
    routeChain  []func(http.Handler) http.Handler
    isSubRouter bool
    *http.ServeMux
}

func NewRouter() *Router {
    return &Router{ServeMux: http.NewServeMux()}
}

func (r *Router) Use(mw ...func(http.Handler) http.Handler) {
    if r.isSubRouter {
        r.routeChain = append(r.routeChain, mw...)
    } else {
        r.globalChain = append(r.globalChain, mw...)
    }
}

func (r *Router) Group(fn func(r *Router)) {
    subRouter := &Router{
        routeChain:  slices.Clone(r.routeChain),
        isSubRouter: true,
        ServeMux:    r.ServeMux,
    }
    fn(subRouter)
}

func (r *Router) HandleFunc(pattern string, h http.HandlerFunc) {
    r.Handle(pattern, h)
}

func (r *Router) Handle(pattern string, h http.Handler) {
    for _, mw := range slices.Backward(r.routeChain) {
        h = mw(h)
    }
    r.ServeMux.Handle(pattern, h)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    var h http.Handler = r.ServeMux
    for _, mw := range slices.Backward(r.globalChain) {
        h = mw(h)
    }
    h.ServeHTTP(w, req)
}
```

А затем вы можете использовать тип `Router` в своём коде вот так:

```go
r := NewRouter()
r.Use(recoverPanic)

r.Handle("GET /static/", http.FileServerFS(ui.Files))

r.Group(func(r *Router) {
    r.Use(requestID)
    r.Use(logRequest)

    r.HandleFunc("GET /", home)
    r.HandleFunc("GET /article/{id}", showArticle)

    r.Group(func(r *Router) {
        r.Use(authenticateUser)
        r.Use(requireAdminUser)

        r.HandleFunc("GET /admin", showAdminDashboard)
    })
})
```

Теги:

- [middleware](https://habr.com/ru/search/?target_type=posts&order=relevance&q=[middleware])
- [стандартная библиотека Go](https://habr.com/ru/search/?target_type=posts&order=relevance&q=[%D1%81%D1%82%D0%B0%D0%BD%D0%B4%D0%B0%D1%80%D1%82%D0%BD%D0%B0%D1%8F+%D0%B1%D0%B8%D0%B1%D0%BB%D0%B8%D0%BE%D1%82%D0%B5%D0%BA%D0%B0+Go])
- [маршрутизация](https://habr.com/ru/search/?target_type=posts&order=relevance&q=[%D0%BC%D0%B0%D1%80%D1%88%D1%80%D1%83%D1%82%D0%B8%D0%B7%D0%B0%D1%86%D0%B8%D1%8F])
- [обработчики](https://habr.com/ru/search/?target_type=posts&order=relevance&q=[%D0%BE%D0%B1%D1%80%D0%B0%D0%B1%D0%BE%D1%82%D1%87%D0%B8%D0%BA%D0%B8])
- [композиция функций](https://habr.com/ru/search/?target_type=posts&order=relevance&q=[%D0%BA%D0%BE%D0%BC%D0%BF%D0%BE%D0%B7%D0%B8%D1%86%D0%B8%D1%8F+%D1%84%D1%83%D0%BD%D0%BA%D1%86%D0%B8%D0%B9])
- [группировка маршрутов](https://habr.com/ru/search/?target_type=posts&order=relevance&q=[%D0%B3%D1%80%D1%83%D0%BF%D0%BF%D0%B8%D1%80%D0%BE%D0%B2%D0%BA%D0%B0+%D0%BC%D0%B0%D1%80%D1%88%D1%80%D1%83%D1%82%D0%BE%D0%B2])
- [зависимости](https://habr.com/ru/search/?target_type=posts&order=relevance&q=[%D0%B7%D0%B0%D0%B2%D0%B8%D1%81%D0%B8%D0%BC%D0%BE%D1%81%D1%82%D0%B8])

Хабы:
- [Блог компании OTUS](https://habr.com/ru/companies/otus/articles/)
- [Go](https://habr.com/ru/hubs/go/)
- [Программирование](https://habr.com/ru/hubs/programming/)
- [Веб-разработка](https://habr.com/ru/hubs/webdev/)

Источник:
https://habr.com/ru/companies/otus/articles/988234/
