package main

import "fmt"

func main() {
  // the slice s is a make function, consisting of an integer array of size 3.
  s := make([]int, 3)

  // slices are similar to arrays in that:
  // - you can index into a slice or an array:
  s[0] = 1
  s[1] = 2
  s[3] = 3
  fmt.Println(s)

  // - you access the content of a slice similarly to arrays:
  fmt.Println(s[0])

  // - print length works the same way too:
  fmt.Println(len(s))

  // slices DIFFER from arrays in that:
  // - the append function is unique to slices
  s = append(s, 4, 5, 6, 7)
  fmt.Println(s)

  // - you can slice out information from a slice, similar to how you would
  // slice a list in python:
  fmt.Println(s[1:3]) // prints out index 1 through 3.

  // - make a concise slice def:
  t := []int{100, 200, 300} // slice t is an integer array with values defined.
  fmt.Println(t)

  
}
