package main

import "fmt"

func main() {
  messages := make(chan string)

  go func() {
    fmt.Println("goroutine start")
    messages <- "ping"
    fmt.Println("goroutine end")
  }()

  fmt.Println("main step 1")
  msg := <-messages
  fmt.Println(msg)
  fmt.Println("main step 2")
}
