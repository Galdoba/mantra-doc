---
updated_at: 2026-04-06T17:43:23.462+10:00
tags:
  - go_1_26
---
В Go появился быстрый immutable map - альтернатива стандартному map



Это map, который нельзя менять после создания,  зато он сильно быстрее и экономнее по памяти.

Зачем это нужно:
- много данных (миллионы ключей)  
- данные не меняются  
- важна скорость чтения  

Обычный map в Go:
- ~56 байт на ключ  

Immutable map:
- ~9 байт на ключ  
- почти в **3 раза быстрее** 

Как использовать

```go

package main

import (
    "fmt"
    "log"

    "github.com/lemire/constmap"
)

func main() {
    keys := []string{"apple", "banana", "cherry"}
    values := []uint64{100, 200, 300}

    cm, err := constmap.New(keys, values)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(cm.Map("banana")) // 200
}
```

Статья: https://lemire.me/blog/2026/03/29/a-fast-immutable-map-in-go/  
Код: https://github.com/lemire/constmap