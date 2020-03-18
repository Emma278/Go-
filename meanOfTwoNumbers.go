package main

import "fmt"

// Calculates the mean of two whole numbers (integers).

func main() {

//  var x int
//  var y int

//  x = 1
//  y = 2

// The above is really tedious so let's do this instead, same result:
  x, y := 1.0, 2.0
// 1.0 defaults to float64. 1 defaults to int.

// Printf for Printfunction. Needs a "%verb" declared to know what format the
// value will be printed in (%v is default base 10). %T is type (int, float etc)
// Printf cheatsheet here:
// https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
  fmt.Printf("x=%v, type of %T\n", x, x)
  fmt.Printf("y=%v, type of %T\n", y, y)

  mean := (x + y) / 2.0
  fmt.Printf("Result: %v, type of %T\n", mean, mean)
}
