Что выведет программа?

Объяснить работу конвейера с использованием select.

```go
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func asChan(vs ...int) <-chan int {
  c := make(chan int)
  go func() {
    for _, v := range vs {
      c <- v
      time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    }
    close(c)
  }()
  return c
}

func merge(a, b <-chan int) <-chan int {
  c := make(chan int)
  go func() {
    for {
      select {
      case v, ok := <-a:
        if ok {
          c <- v
        } else {
          a = nil
        }
      case v, ok := <-b:
        if ok {
          c <- v
        } else {
          b = nil
        }
      }
      if a == nil && b == nil {
        close(c)
        return
      }
    }
  }()
  return c
}

func main() {
  rand.Seed(time.Now().Unix())
  a := asChan(1, 3, 5, 7)
  b := asChan(2, 4, 6, 8)
  c := merge(a, b)
  for v := range c {
    fmt.Print(v)
  }
}
```
Ответ:
Будут выведены числа от 1 до 8 включительно в случайном порядке,
т.к. функция `asChan` использует `rand.Intn(...)` для случайной задержки между отправками значений и `select` выбирает случайный готовый канал, если готовы сразу оба канала.
После того как данные из обоих каналов были получены и переменным 
`a` и `b` были присвоены значения `nil`, результирующий канал `c` будет закрыт
и произойдет выход из функции `merge`.
В функции main происходит чтение данных из результирующего канала `c` и завершение работы после закрытия обоих входных каналов. 