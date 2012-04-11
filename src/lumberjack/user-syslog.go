package lumberjack

import (
  "fmt"
  "os"
  "net"
)

func UserLogReader(channel chan []byte) {
  addr, _ := net.ResolveUnixAddr("unixgram", "/dev/log")
  conn, _ := net.ListenUnixgram("unixgram", addr)

  // Ensure world writable access
  os.Chmod("/dev/log", 0666)

  data := make([]byte, 4096)

  for {
    length, _, err := conn.ReadFrom(data)
    channel <- data[0 : length]
  }
}
