package main

import "fmt"
import "time"

func worker(done chan bool) {
  fmt.Print("goroutine working...")
  time.Sleep(time.Second)
  fmt.Println("goroutine done")

  done <- true
  fmt.Println("goroutine exit")
}

func main() {
  fmt.Println("main step 1")

  done := make(chan bool, 1)
  go worker(done)

  fmt.Println("main step 2")

  <-done

  fmt.Println("main step 3")
}
