package main

import "fmt"

// just a string drill:

func main() {
  book := "The colour of Magic"
  fmt.Println(book)

  fmt.Println(len(book))

// prints value (T) and type (string) of character 0
  fmt.Printf("book[0] = %v (type %T)\n", book, book)

// strings in go are immutable. I can't change a
// character like below.
//  book[0] = 116

// Slice(start:end):
  fmt.Println(book[4:11])

// Slice(no end):
  fmt.Println(book[4:])

// Slice(no start):
  fmt.Println(book[:4])

// concatenate a string with + :
  fmt.Println("t" + book[1:])

// strings are unicode so special signs are no probs:
  fmt.Println("It's just $2!")

// multi-line:
  poem := `
  The road goes on
  and on
  and oooooon
  ...
  `
  fmt.Println(poem)
}
