package lumberjack

import (
  "syscall"
  "fmt"
)

/* Read logs via klogctl and send each event to the chan 'channel' */
func KernelLogReader(channel chan []byte) {
  data := make([]byte, 16384)

  for {
    length, error := syscall.Klogctl(2, data)
    if error != nil {
      fmt.Println("Error:", error)
      continue;
    }
    fmt.Printf("Received(%d): %.*s\n", length, length, data)

    /* Send the message to the channel */
    channel <- data[0 : length]
  }
}

/* Example:
 * func main() {
 *   c := make(chan []byte)
 *   go KernelLogReader(c)
 *   
 *   for {
 *     data := <- c
 *     fmt.Printf("%.*s\n", len(data), data)
 *   }
 * }
 */
