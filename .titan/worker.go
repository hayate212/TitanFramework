package titan

import (
	"fmt"
	"log"
	"net"

	"github.com/hayate212/seviper"
)

type Worker struct {
	Config       WorkerConfig
	EventHandler WorkerEventHandler
}

type WorkerConfig struct {
	Address        string
	Port           int
	MaxRequestSize int
}

func NewWorker(args ...interface{}) *Worker {
	var config WorkerConfig
	var we *WorkerEventHandler
	if args[0] != nil {
		config = args[0].(WorkerConfig)
	} else {
		//default setting
		config = WorkerConfig{
			Address:        "localhost",
			Port:           9999,
			MaxRequestSize: 255,
		}
	}
	if args[1] != nil {
		we = args[1].(*WorkerEventHandler)
	} else {
		we = NewWorkerEventHandler()
	}
	return &Worker{config, *we}
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
		go func() {
			fmt.Println(conn.RemoteAddr())
			for {
				rawbuff := make([]byte, w.Config.MaxRequestSize)
				n, err := conn.Read(rawbuff)
				if err != nil {
					continue
				}
				buff := rawbuff[:n]
				fmt.Printf("%v\nlength:%v\n", buff, n)
				r := seviper.NewReader(buff)
				key := r.ToString()
				if e, ok := w.EventHandler[key]; ok {
					e.Run(r.Backward())
				}
			}
		}()
	}
}
