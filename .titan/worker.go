package titan

import (
	"fmt"
	"log"
	"net"
)

type Worker struct {
	Config WorkerConfig
}

type WorkerConfig struct {
	Address string
}

func NewWorker(config WorkerConfig) *Worker {
	return &Worker{config}
}

func (w *Worker) Run() {
	server, err := net.Listen("tcp", w.Config.Address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		defer conn.Close()
		fmt.Println(conn.RemoteAddr())
	}
}
