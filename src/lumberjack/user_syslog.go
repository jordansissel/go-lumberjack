package lumberjack

import (
  "os"
  "net"
)

/*
 * This uses the '/dev/log' unix socket to read logs.
 * On linux, this should read events written using the syslog(3) libc api or
 * the logger(1) tool.
 *
 * DATA LOSS WARNING: /dev/log is a unix datagram socket and has a finite,
 * nonblocking packet buffer. Therefore, use of the syslog(3) libc api is
 * extremely discouraged due to data loss. If you care about your logs,
 * you must not use syslog(3).
 *
 * TODO(sissel): Allow providing the path to the log file as an argument.
 */
func UserLogReader(channel chan []byte) {
  /* TODO(sissel): Handle errors */
  addr, _ := net.ResolveUnixAddr("unixgram", "/dev/log")
  /* TODO(sissel): Handle errors */
  conn, _ := net.ListenUnixgram("unixgram", addr)

  /* Ensure world writable access.
   *
   * Fun fact, the logger(1) tool on Linux will silently exit successfully if
   * it cannot write to /dev/log. */
  os.Chmod("/dev/log", 0666)

  data := make([]byte, 4096)

  for {
    length, _, _ := conn.ReadFrom(data)
    channel <- data[0 : length]
  }
}
