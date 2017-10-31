package main

import "fmt"
import "time"

func main() {
  c1 := make(chan string)
  c2 := make(chan string)

  go func() {
    fmt.Println("go 1 start")
    time.Sleep(time.Second * 1)
    c1 <- "one"
    fmt.Println("go 1 end")
  }()

  go func() {
    fmt.Println("go 2 start")
    time.Sleep(time.Second * 2)
    c2 <- "two"
    fmt.Println("go 2 end")
  }()

  fmt.Println("before for loop")
  for i := 0; i < 2; i++ {
    select {
      case msg1 := <-c1:
      fmt.Println("received", msg1)
      case msg2 := <-c2:
      fmt.Println("received", msg2)
    }
  }
  fmt.Println("after for loop")
}
