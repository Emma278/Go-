package main

import "fmt"

// Defer is used for garbage handling and closing services once they are done.

func cleanup(name string) {
  fmt.Printf("Cleaning up %s\n", name) // printing: cleaning up {name}
}

func worker() {
  defer cleanup("A")    // defer will run even if there is an error or panic in
  fmt.Println("worker") //the service.
}

func main() {
  worker()
}
