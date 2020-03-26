package main

import "fmt"
import "time"

func main() {
  ch1, ch2 := make(chan int), make(chan int)

  go func() {
    ch1 <- 42
  }()

  // if val is from ch1, print "got (val) from ch1"
  select {
  case val := <-ch1:
    fmt.Printf("got %d from ch1\n", val)
  case val := <-ch2:
    fmt.Printf("got %d from ch2\n", val)
  }

  fmt.Println("-----")
  // out makes a channel of floats.
  // after 0.1 s, 3.14 gets passed to out.
  out := make(chan float64)
  go func() {
    time.Sleep(100 * time.Millisecond)
    out <- 3.14
  }()

  // if val is the same as the item in out, print what we got.
  // if nothing returns after 0.2 s, print "timeout".
  select {
  case val := <-out:
    fmt.Printf("got %f\n", val)
  case <-time.After(200 * time.Millisecond):
    fmt.Println("timeout")
  }
}
