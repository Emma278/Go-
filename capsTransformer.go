package main

import "fmt"
import "io"
import "os"

// Capper implements io.Writer and turns everything to uppercase.
type Capper struct {
  wtr io.Writer
}

// In go, characters are called runes. Convert the difference into a byte:
// "diff is a byte of lowercase minus uppercase"
func (c *Capper) Write(p []byte) (n int, err error) {
  diff := byte('a' - 'A')

// here we make a buffer that is a slice of bytes, of length p.
  out := make([]byte, len(p))
// the condition for i where c is a range of p is as follows:
  for i, c := range p {
// if c is a or z or any character in between, withdraw diff from c and
// set the "out" buffer of i to c.
  if c >='a' && c <= 'z' {
    c -= diff

    }
    out[i] = c
  }
// Then return the converted value.
  return c.wtr.Write(out)
}

// Print the result.
func main() {
  c := &Capper{os.Stdout}
  fmt.Fprintln(c, "Hello there")
}
