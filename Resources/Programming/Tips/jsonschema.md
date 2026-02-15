---
updated_at: 2026-02-15T17:05:01.962+10:00
tags:
  - package
---
Библиотека для [[Validation|валидации]] [[JSON]] по стандарту [[JSON Schema]]. 

Удобный инструмент, чтобы проверять структуру данных без ручной проверки полей.

Чем полезна:
- Валидация входящих API-запросов  
- Проверка конфигурационных файлов  
- Контроль формата данных между микросервисами  
- Защита от некорректных или неожиданных данных  

Раньше для этого использовали сторонние решения или писали проверки вручную. Теперь есть официальная библиотека от Google с упором на производительность и соответствие стандарту.

Что умеет библиотека:

- Поддержка современных версий JSON Schema  
- Компиляция схемы один раз и быстрая валидация  
- Понятные сообщения об ошибках  
- Подходит для high-load сервисов  

Пример использования:
```go
package main

import (
 "fmt"
 "log"

 "github.com/google/jsonschema-go/jsonschema"
)

func main() {
 schemaJSON := []byte(`{
  "type": "object",
  "properties": {
   "name": { "type": "string" },
   "age":  { "type": "integer", "minimum": 0 }
  },
  "required": ["name"]
 }`)

 schema, err := jsonschema.Compile(schemaJSON)
 if err != nil {
  log.Fatal(err)
 }

 data := []byte(`{"name": "John", "age": 30}`)

 if err := schema.Validate(data); err != nil {
  fmt.Println("Invalid:", err)
 } else {
  fmt.Println("Valid JSON")
 }
}
```

▪️ Blog: opensource.googleblog.com/2026/01/a-json-schema-package-for-go.html
▪️Github: github.com/google/jsonschema-go
