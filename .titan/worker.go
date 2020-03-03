package titan

import (
	"fmt"
	"log"
	"net"
	"reflect"

	"github.com/hayate212/seviper"
)

type Worker struct {
	Config WorkerConfig
	Handle *EventHandle
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
			Handle: NewEventHandle(nil),
		}
	}
	return &Worker{Config: args[0].(WorkerConfig), Handle: NewEventHandle(nil)}
}

func (w *Worker) SetEventHandle(i interface{}) {
	w.Handle = NewEventHandle(i)
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
				name := r.ToString()
				if args, ok := w.Handle.BytesToArgs(name, r.Backward()); ok {
					if result, ok := w.Handle.Proc(name, args); ok {
						conn.Write(ToBytes(result))
					}
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
