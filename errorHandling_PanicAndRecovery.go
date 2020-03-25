package main

// this file shows examples of panic code: But, panic is discouraged in go,
// so follow the Hitch-Hikers' Guide to the Galaxy, and DON'T PANIC plz.

import "fmt"
// import "os"

func safeValue(vals []int, index int) int {
  defer func() {
    if err := recover(); err != nil {
      fmt.Printf("ERROR: %s\n", err)
    }
  }()

  return vals[index]
}

func main() {

  /*
  vals := []int{1, 2, 3}
  v := vals[10] // Panic: you can't call vals index 10 if it's not declared
  fmt.Println(v)
  */

  /*
  file, err := os.Open("no-such-file")
  if err != nil {
    panic(err)
  }
  defer file.Close() // if no errors, close file
  fmt.Println("file opened")
  */

  // call saveValues slice of integers 1, 2 and 3, at index 10:
  v := safeValue([]int{1, 2, 3}, 10)
  fmt.Println(v)
}
