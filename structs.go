package main

import "fmt"

// Structs is used to combine several fields into a data type.
// Below, Trade is a trade in stocks. Each field is a different data type within
// the Trade type. Now that the type struct is defined, we can use the type.
type Trade struct {
  Symbol  string  // Stock symbol
  Volume  int     // Number of shares
  Price   float64 // Trade price
  Buy     bool    // True if buy trade, false if sell trade
}

func main() { // we can define Trade values by position:
  t1 := Trade{"MSFT", 10, 99.98, true}
  fmt.Println(t1)

  fmt.Printf("%+v\n", t1) // with +v, we print the object type structs.

  fmt.Println(t1.Symbol) // Only print the value for Symbol in t1

  t2 := Trade{ // A structured way of defining values for var t2:
    Symbol: "MSFT",
    Volume: 10,
    Price:  99.98,
    Buy:    true,
  }
  fmt.Printf("%+v\n", t2) // Prints values AND their data type

  t3 := Trade{} // Assigns 0 value to the objects
  fmt.Printf("0 value of %+v\n", t3)
}

// Anything starting with an Uppercase letter, will be accessible for other
// packages. Otherwise it is only accessible from within the current package.
