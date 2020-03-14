package titan

import (
	"fmt"
	"log"
	"net"
	"reflect"

	"github.com/hayate212/seviper"
)

type Worker struct {
	Config  WorkerConfig
	Handles *EventHandles
}

type WorkerConfig struct {
	Address        string
	Port           int
	MaxRequestSize int
}

func NewWorker(args ...interface{}) *Worker {
	if args[0] == nil {
		return &Worker{
			//default setting
			Config: WorkerConfig{
				Address:        "localhost",
				Port:           9999,
				MaxRequestSize: 255,
			},
			Handles: NewEventHandles(),
		}
	}
	return &Worker{Config: args[0].(WorkerConfig), Handles: NewEventHandles()}
}

func (w *Worker) AddEventHandle(i interface{}) bool {
	return w.Handles.AddHandle(i)
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
					return
				}
				buff := rawbuff[:n]
				r := seviper.NewReader(buff)
				name := r.ToString()
				fmt.Printf("raw: %v\nlength: %v\nname: %v\n", buff, n, name)
				if result, ok := w.Handles.Proc(name, r.Backward()); ok {
					conn.Write(ToBytes(result))
				}
			}
		}()
	}
}

func ToBytes(r []reflect.Value) []byte {
	w := seviper.NewWriter()
	for _, v := range r {
		//fmt.Println(v.Kind())
		switch fmt.Sprintf("%v", v.Kind()) {
		case "string":
			w.Write(v.String())
		case "int":
			w.Write(v.Int())
		case "float32":
			w.Write(float32(v.Float()))
		case "float64":
			w.Write(v.Float())
		case "slice":
			if v.Type() == reflect.TypeOf([]byte{}) {
				w.Bytes = append(w.Bytes, v.Bytes()...)
			}
		}
	}
	return w.Bytes
}
