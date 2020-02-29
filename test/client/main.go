package main

import (
	"net"
	"time"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:9999")
	buff := []byte{00, 00, 00, 00}
	for {
		conn.Write(buff)
		buff[0]++
		time.Sleep(time.Second * 1)
	}
}
