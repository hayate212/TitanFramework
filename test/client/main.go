package main

import (
	"net"
	"time"

	"github.com/hayate212/seviper"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		return
	}
	c := 0
	for {
		w := seviper.NewWriter()
		w.Write("test")
		w.Write(c)
		w.Write(c * c)
		_, err := conn.Write(w.Bytes)
		if err != nil {
			return
		}
		c += 16
		time.Sleep(time.Second * 1)
	}
}
