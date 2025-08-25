Что выведет программа?

Объяснить вывод программы.
```go
package main

type customError struct {
  msg string
}

func (e *customError) Error() string {
  return e.msg
}

func test() *customError {
  // ... do something
  return nil
}

func main() {
  var err error
  err = test()
  if err != nil {
    println("error")
    return
  }
  println("ok")
}
```
Ответ:
Вывод в консоль `error` т.к. (*customError, nil) != (nil, nil) 
Интерфейс error содержит информацию о конкретном типе (*customError)
Интерфейс с известным типом, но nil-значением не равен чистому nil