package main

import "fmt"
import "time"

func main() {
  ch := make(chan int) // let's make a channel of integers.

  /*
  below, ch blocks everything because it creates a "deadlock", blocking
  anything to get pushed to the channel.
  <-ch
  fmt.Println("Here")
  */

  // try this instead.
  go func() {
    // send number of the channel...
    ch <- 353
  }()

  // ...then receive from the channel...
  val := <-ch
  fmt.Printf("got %d\n", val)
  fmt.Println("-----")

  // ...now let's see if we can send multiple values to the channel with a
  // for loop that will send values 0, 1, 2 per second:
  go func() {
    for i := 0; i < 3; i++ {
      fmt.Printf("sending %d\n", i)
      ch <- i
      time.Sleep(time.Second)
    }
  }()

  // and send those values as they come, to the channel, and print what ch
  // received:
  for i := 0; i < 3; i++ {
    val := <-ch
    fmt.Printf("received %d\n", val)
    fmt.Println("-----")

  // The outcome will be an assurance that we both sended and received each
  // value successfully.

  // Now let's assume that we don't know how many items we have in the channel.
  // We'll signal that the work is done by closing the channel:

  go func() {
    for i := 0; i < 3; i++ {
      fmt.Printf("sending %d\n", i)
      ch <- i
      time.Sleep(time.Second)
    }
    close(ch)
  }()

  // we use the range keyword to iterate over the values, and then print what
  // we received.
  for i := range ch {
    fmt.Printf("received %d\n", i)
  }
}
}
