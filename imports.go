package main

import (
  "fmt"
  "log"
  "os"

  toml "github.com/pelletier/go-toml"
)

/*
To use go-toml, we need to go get github.com/pelletier/go-toml.
It's that easy to get a package outside of the standard library.
Give the package a local name by naming it before the import quote.
Once the program is built, all external packages will be built into the
executable so no one else need to go get external packages before running
the program.

The rest of this exercise is essentially a config type, opening a toml file,
decoding it, and displaying the username and password.
And standard error management.

To run the program, you need a config.toml file in the program dir.
*/

type Config struct {
  Login struct{
    User      string
    Password  string
  }
}

func main() {
  file, err := os.Open("config.toml")
  if err != nil {
    log.Fatalf("error: can't open config file - %s", err)
  }
  defer file.Close()

  cfg := &Config{}
  dec := toml.NewDecoder(file)
  if err := dec.Decode(cfg); err != nil {
    log.Fatalf("error: can't decode configuration file - %s", err)
  }

  fmt.Printf("%+y\n", cfg)

}
