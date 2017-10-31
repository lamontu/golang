package main

import "fmt"
import "time"

func main() {
  fmt.Println("main start")
  jobs := make(chan int, 5)
  done := make(chan bool)

  go func() {
    fmt.Println("go start")
    for {
      j, more := <-jobs
      if more {
        fmt.Println("received job", j)
      } else {
        fmt.Println("received all jobs")
        done <- true
        return
      }
    }
    fmt.Println("go end")  // Not output
  }()

  for j := 1; j <= 3; j++ {
    jobs <- j
    fmt.Println("sent job", j)
  }
  close(jobs)
  fmt.Println("sent all jobs")

  <-done
  time.Sleep(time.Second * 2)

  fmt.Println("main end")
}
