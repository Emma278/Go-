package main

import "fmt"

type Trade struct {
  Symbol  string  // Stock symbol
  Volume  int     // Number of shares
  Price   float64 // Trade price
  Buy     bool    // True if buy trade, false if sell trade
}

// Value returns the trade value
func (t *Trade) Value() float64 { // This func has a POINTER RECEIVER
                                  // to the Trade object. Value() is the method
                                  // name, and last is the return type float.
  value := float64(t.volume) * t.Price // convert volume to a float so that
                                       // value contains only one data type.
  if t.Buy {
    value = -value
  }

  return value
}

func main() {
  t := Trade{
    Symbol: "MSFT",
    Volume: 10,
    Price:  99.98,
    Buy:    true,
  }
}
