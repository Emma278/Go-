package main

import "fmt"

// fizzBuzz prints the number series 1-20, BUT:
// If a number is divisible by 3, it prints fizz
// If a number is divisible by 5, it prints buzz
// If a number is divisible by both 3 and 5, it prints fizzbuzz

func main() {
  for i := 1; i <= 20; i++ {
    if i%3 == 0 && i%5 == 0 {
      fmt.Println("fizzbuzz")
    } else if i%3 == 0 {
      fmt.Println("fizz")
    } else if i%5 == 0 {
      fmt.Println("buzz")
    } else {
      fmt.Println(i)
    }
  }
}
