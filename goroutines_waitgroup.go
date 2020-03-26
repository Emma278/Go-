package main

/*
The go concurrency primitive is the goroutine. You can spin tens of thousands
 of go routines simultaneously within the same machine.
*/

import "fmt"
import "net/http"
import "sync"

// The function returnType makes a HTTP call, looks at the header, and returns
// the type.
func returnType(url string) {
  resp, err := http.Get(url)
  if err != nil {
    fmt.Printf("error: %s\n", err)
    return
  }

  defer resp.Body.Close()
  ctype := resp.Header.Get("content-type")
  fmt.Printf("%s -> %s\n", url, ctype)
}

func main() {
  urls := []string{
    "https://golang.org",
    "https://api.github.com",
    "https://httpbin.org/xml",
  }

  var wg sync.WaitGroup
// return the type of every url in urls
// and also use concurrence goroutine to speed up the script
  for _, url := range urls {
    wg.Add(1)
    go func(url string) {
      returnType(url)
      wg.Done()
    } (url)
  }
  wg.Wait()
}
