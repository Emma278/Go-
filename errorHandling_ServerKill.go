package main

/*
Script that identifies a pid file and kills the server with that pid.
It also counts for errors such as: cannot find/open pid file,
file contains bad pid value.
*/

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
  data, err := ioutil.ReadFile(pidFile)
  if err != nil {
    return errors.Wrap(err, "can't open pid file (is server running?)")
  }

  if err := os.Remove(pidFile); err != nil {
    // ok to continue if this fails
    log.Printf("warning: can't remove pid file - %s", err)
  }

  strPID := strings.TrimSpace(string(data))
  pid, err := strconv.Atoi(strPID)
  if err != nil {
    return errors.Wrap(err, "bad process ID")
  }

  // "kill it"
  fmt.Printf("killing server with pid=%d\n", pid)
  return nil
}

func main() {
  if err := killServer("server.pid"); err != nil {
    fmt.Fprintf(os.Stderr, "error: %s\n", err)
    os.Exit(1)
  }
}
