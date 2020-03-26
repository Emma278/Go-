package main

import (
  	"bufio"
  	"crypto/md5"
  	"fmt"
  	"io"
  	"log"
  	"os"
  	"strings"
)

/*
checker that checks signatures and filenames available, and displays them in a map.
For this program to work, you need the file md5sum.txt in the program directory, together with a number of log files.
*/

// parse the signature file, return a map of path->signature
func parseSignaturesFile(path string)(map[string]string, error) {
  // first step: try to open the file. If successful, close.
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  // create a map and use a scanner to go over the map.
  sigs := make(map[string]string)
  scanner := bufio.NewScanner(file)
  // every line num is 1. As long as the scan is true, increment the lnum with 1
  for lnum := 1; scanner.Scan(); lnum++ {
    // test to see if the lnum is the right format
    fields := strings.Fields(scanner.Text())
    if len(fields) != 2 {
      return nil, fmt.Errorf("%s:%d bad line", path, lnum)
    }
    // if it is, set the signature of this file in the map.
    sigs[fields[1]] = fields[0]
  }
  // return any scanning errors here
  if err := scanner.Err(); err != nil {
    return nil, err
  }
  // if everything is ok, return the map of file signatures.
  return sigs, nil
}

func fileMD5(path string) (string, error) {
  file, err := os.Open(path)
  if err != nil {
    return "", err
  }
  defer file.Close()

  // create new md5 hash signature
  hash := md5.New()
  if _, err := io.Copy(hash, file); err != nil {
    return "", err
  }

  return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// result type from worker
type result struct {
  path string
  match bool
  err error
}

// create the worker:
func md5Worker(path string, sig string, out chan *result) {
  r := &result{path: path}
  s, err := fileMD5(path)
  if err != nil {
    r.err = err
    out <- r
    return
  }

  r.match = (s == sig)
  out <- r
}


func main() {
  sigs, err := parseSignaturesFile("md5sum.txt")
  if err != nil {
    log.Fatalf("error: can't read signature file - %s", err)
  }

  out := make(chan *result)
  for path, sig := range sigs {
    go md5Worker(path, sig, out)
  }

  ok := true
  for range sigs {
    r := <-out
    switch {
    case r.err != nil:
      fmt.Printf("%s: error - %s\n", r.path, r.err)
      ok = false
    case !r.match:
      fmt.Printf("%s: signature mismatch\n", r.path)
      ok = false
    }

    if !ok {
      os.Exit(1)
    }
  }
}
