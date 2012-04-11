package lumberjack

import (
  "syscall"
  "fmt"
)

/* Read logs via klogctl and send each event to the chan 'channel'
 *
 * This uses the 'klogctl' syscall on linux.
 * The first argument is the 'command' and '2' means "blocking read" 
 * which consumes the buffer but blocks until there is data.
 *
 * In general, the only data coming over this interface is usually kernel
 * logs via printk.
 *
 * Example code:
 *
 *
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

func main() {
  data := make([]byte, 16)
  syscall.Klogctl(10, data)
  fmt.Println(data)
}
