package main

import "fmt"

func ping(pings chan<- string, msg string) {
  fmt.Println("ping routine start")
  pings <- msg
  fmt.Println("ping routine end")
}

func pong(pings <-chan string, pongs chan<- string) {
  fmt.Println("pong routine start")
  msg := <-pings
  pongs <- msg
  fmt.Println("pong routine end")
}

func main() {
  fmt.Println("main start")
  pings := make(chan string, 1)
  pongs := make(chan string, 1)
  ping(pings, "passed message")
  pong(pings, pongs)
  fmt.Println(<-pongs)
  fmt.Println("main end")
}
