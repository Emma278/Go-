package main

import "fmt"



func main() {
  nums := []int{16, 8, 42, 4, 23, 15}
  // max is index 0 of nums (which is 16)
  max := nums [0]
  // for any index, value is a number in nums starting from index 1 (8):
  for _, value := range nums[1:] {
    // if that value is larger than max, then the new max is value.
    // The loop continues until no values are larger than max.
    if value > max {
      max = value
    }
  }
  fmt.Println(max)
}
