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
	Port    int
}

func NewWorker(args ...interface{}) *Worker {
	var config WorkerConfig
	if args[0] != nil {
		config = args[0].(WorkerConfig)
	} else {
		config = WorkerConfig{
			Address: "localhost",
			Port:    9999,
		}
	}
	return &Worker{config}
}

func (w *Worker) Run() {
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%v", w.Config.Address, w.Config.Port))
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
