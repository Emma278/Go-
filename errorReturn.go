package main

import "fmt"
import "math"

// this function uses the fact that go can return several values (here, float
// and error), to return an error message if the square root of n is negative.
// error is a built-in type in go (just as float64, int, string).
// nil is how go expresses "null" or "NaN".
func sqrt(n float64) (float64, error) {
  if n < 0 {
    return 0.0, fmt.Errorf("sqrt of negative value (%f)", n)
  }
  return math.Sqrt(n), nil
}

func main() {
  s1, err := sqrt(2.0)
  if err != nil { // "if error is not nil / if an error has happened"
    fmt.Printf("ERROR: %s\n", err) // Print the error.
  } else {
      fmt.Println(s1) // Otherwise, print the sqrt output.
  }

  s2, err := sqrt(-2.0)
  if err != nil {
    fmt.Printf("ERROR: %s\n", err)
  } else {
    fmt.Println(s2)
  }
}

// Go also contains a Panic function, somewhat similar to exceptions in other
// languages. Use of panic is discouraged in Go and considered an anti-pattern.
// The way to signal an error is to return it.
