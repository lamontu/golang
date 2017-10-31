package main

import "time"
import "fmt"

func main() {
  fmt.Println("main start")
  c1 := make(chan string, 1)
  go func() {
    fmt.Println("go 1 start")
    time.Sleep(time.Second * 2)
    c1 <- "result 1"
    fmt.Println("go 1 end")
  }()

  select {
    case res := <-c1:
    fmt.Println(res)
    case <-time.After(time.Second * 1):
    fmt.Println("timeout 1")
  }
  //time.Sleep(time.Second * 1)  // for "go end" to output if no below

  c2 := make(chan string, 1)
  go func() {
    fmt.Println("go 2 start")
    time.Sleep(time.Second * 2)
    c2 <- "result 2"
    fmt.Println("go 2 end")
  }()

  select {
    case res := <-c2:
    fmt.Println(res)
    case <-time.After(time.Second * 3):
    fmt.Println("timeout 2")
  }

  fmt.Println("main finish")
}
