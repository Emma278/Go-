package main

import "fmt"

// Bear with me, I'm from python so pointers are still foreign...

func main() {
  x := 15
  a := &x // The ampersand POINTS to the var where a will get it's value.
          // This is a "memory address".
  fmt.Println(*a) // the asterisk marks that a will get its' value from a
                  // memory address.
  *a = 5
  fmt.Println(x) // now we've changed the value of x to 5.

  // fun experiment:
  *a = *a**a // "whatever's through that a equals whatever's through that a
             // times whatever's through that a" is here equal to *a = 5*5
             // which means we just changed the value of x to 25.
  fmt.Println(x)
  fmt.Println(*a)
}

// Wrapup: &x means that it's a memory address to the variable x that you pass
// and when you want to read what's at that memory address you use the *.
