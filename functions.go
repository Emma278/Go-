package main

import "fmt"

// function that adds a and b together
func add(a int, b int) int {
  return a + b
}

// go can return more than one value:
// divmod that returns quotient (kvot) AND remainder (rest)
func divmod(a int, b int) (int, int) { // (variables and type)(number of returns and their type)
  return a / b, a % b
}

// Another way of writing the above:
// Numerator = täljare
// Denominator = nämnare
// func divmod(numerator, denominator int64) (quotient, remainder int64) {
//    quotient = numerator / denominator // integer division, decimals are truncated
//    remainder = numerator % denominator
//    return
//}

// print the result of the two functions above
func main() {
  val := add (1, 2)
  fmt.Println(val)

  div, mod := divmod(7, 2) // here, we set the var values div = 7, mod = 2
  fmt.Printf("div=%d, mod=%d\n", div, mod) // Printfunction prints the var values as %d = int base 10 (https://golang.org/pkg/fmt/)
}

// The above prints the sum of a and b (which is 3),
// and the quotient of 7 / 2 as an integer (3), with the remainder 1.
