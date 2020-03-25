package main

import "fmt"
import "os"
import "github.com/pkg/errors" // non-standard package: in a terminal, type
                               // go get github.com/pkg/errors
import "log"

type Config struct {
  // config fields go here
}

func readConfig(path string) (*Config, error) {
  // try to open the file
  file, err := os.Open(path)
  // if file doesn't open, return error
  if err != nil {
    return nil, errors.Wrap(err, "can't open configuration file")
  }
  // readConfig is done.
  defer file.Close()

  // cfg is Config.
  cfg := &Config{}
  // Parse file here
  return cfg, nil
}

// **Send stack trace to app.log (if you display it everyone will believe that
// there are bugs in the program)
func setupLogging() {
  out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    return
  }
  log.SetOutput(out)
}

func main() {
  // **call setupLogging
  setupLogging()
  // call readConfig
  cfg, err := readConfig("/path/to/config.toml")
  // if readConfig isn't responding, return error and exit program.
  if err != nil {
    fmt.Fprintf(os.Stderr, "error: %s\n", err)
    log.Printf("error: %+v", err) // **print err to app.log
    os.Exit(1)
  }
  // No errors: then print the config.
  fmt.Println(cfg)
}
