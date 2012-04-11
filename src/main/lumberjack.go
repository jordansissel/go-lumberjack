package main

import (
  "lumberjack"
  "time"
  "fmt"
)

func main() {
  /* Limit channel buffer size to 1 */
  c := make(chan []byte, 1)
  go lumberjack.UserLogReader(c)

  for {
    data := <- c
    fmt.Printf("%.*s\n", len(data), data)
    time.Sleep(10000 * time.Millisecond)
  }
}
