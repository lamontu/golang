package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
  fmt.Println("go start", id)
  for j := range jobs {
    fmt.Println("worker", id, "processing job", j)
    time.Sleep(time.Second)
    results <- j * 2
  }
  fmt.Println("go finish", id)
}

func main() {
  jobs := make(chan int, 100)
  results := make(chan int, 100)

  for w := 1; w <= 3; w++ {
    go worker(w, jobs, results)
  }

  for j := 1; j <= 9; j++ {
    jobs <- j
  }
  close(jobs)

  //<-results
  for a := 1; a <= 9; a++ {
    <-results
  }
}
