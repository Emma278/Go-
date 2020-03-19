package main

import "fmt"

func doubleAt(values []int, i int) { // doubleAt takes a slice of values, and their indexes,
  values[i] *= 2 // and then doubles the value at that index.
}

func double(n int) { // And this func should return a double value for a var, right?
  n *= 2
}

func doublePtr(n *int) { // **
  *n *= 2
}

func main() {
  values := []int{1, 2, 3, 4} // this is a slice of int values
  doubleAt(values, 2) // so, this should double the value at index 2 (which is 3))
  fmt.Println(values) // and print the new values [1, 2, 6, 4].

  val := 10
  double(val)
  fmt.Println(val)
  doublePtr(&val) // ***
  fmt.Println(val)

// Go passes integers to functions, by value.
// So Go creates a copy of the integer and passes it to the function.
// Changes to the integer inside the function, won't affect the original value.
// But when Go passes a slice or map to a function, it passes it by reference.
// This means that inside the function body, you're working with the exact same object
// that was passed and any changes to the slice you make inside the function,
// is over.

// Ergo: double should work if val is a slice or map. But if it's not:

// Make double work by using pointers (they will be passed to an object but you
// will not use risky pointer arithmetic as in C or C++):

// ** This new function doublePtr() gets a pointer * to an integer. Inside the
// function body we are DE-REFERENCING the pointer, and assign a value to it.
// *** when we have a pointer we can call doublePtr with AMPERSAND value, and
// then print it again.
// More on pointers in go: https://tour.golang.org/moretypes/1
// https://gist.github.com/josephspurrier/7686b139f29601c3b370



}
