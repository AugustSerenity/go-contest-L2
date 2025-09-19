### Функция or (объединение done-каналов)
Реализовать функцию, которая будет объединять один или более каналов done (каналов сигнала завершения) в один. Возвращаемый канал должен закрываться, как только закроется любой из исходных каналов.

Сигнатура функции может быть такой:
```go
var or func(channels ...&lt;-chan interface{}) &lt;-chan interface{}
```
Пример использования функции:
```go
sig := func(after time.Duration) &lt;-chan interface{} {
   c := make(chan interface{})
   go func() {
      defer close(c)
      time.Sleep(after)
   }()
   return c
}

start := time.Now()
&lt;-or(
   sig(2*time.Hour),
   sig(5*time.Minute),
   sig(1*time.Second),
   sig(1*time.Hour),
   sig(1*time.Minute),
)
fmt.Printf("done after %v", time.Since(start))
```