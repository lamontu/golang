package main

import "fmt"
import "time"

func main() {
  fmt.Println("main start")
  timer1 := time.NewTimer(time.Second * 2)

  //<-timer1.C
  fmt.Println(<-timer1.C)

  fmt.Println("Timer 1 expired")
  timer1.Stop()
  stop1 := timer1.Stop()
  if stop1 {
    fmt.Println("Timer 1 stopped")
  }
  fmt.Println("================================")

  timer2 := time.NewTimer(time.Second * 2)
  go func() {
    fmt.Println("go start")
    fmt.Println(<-timer2.C)
    fmt.Println("Timer 2 expired")
    fmt.Println("go end")
  }()

  time.Sleep(time.Second * 1)
  //time.Sleep(time.Second * 2)

  fmt.Println("before stop")
  stop2 := timer2.Stop()
  fmt.Println("after stop")

  if stop2 {
    fmt.Println("Timer2 stopped")
  }

  fmt.Println("main end")
}
