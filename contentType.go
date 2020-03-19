package main

import "fmt"
import "http"

// function that gets a URL and returns the value of Content-Type response HTTP
// header. Function returns an error if it can't perform the GET request.

func contentType(url string) (string, error) {
  resp, err := http.Get(url)
  if err != nil {
    return "", err
  }

  defer resp.Body.Close() // close the body of resp

  ctype := resp.Header.Get("Content-Type")
  if ctype == "" { // return an error if Content-Type header not found
    return "", fmt.Errorf("can't find Content-Type header")
  }

  return ctype, nil
}

func main() {
  ctype, err := contentType("https://linkedin.com")
  if err != nil {
    fmt.Printf("ERROR: %s\n", err)
  } else {
    fmt.Println(ctype)
  }
}
