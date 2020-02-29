package titan

import "github.com/hayate212/seviper"

type EventHandle struct {
	f    func(args ...interface{}) []byte
	args []ArgType
}

type ArgType int

const (
	INT ArgType = iota
	STRING
	FLOAT
)

type WorkerEventHandler map[string]EventHandle

func NewWorkerEventHandler() *WorkerEventHandler {
	return &WorkerEventHandler{}
}

func (we *WorkerEventHandler) Set(key string, f func(args ...interface{}) []byte, args []ArgType) {
	(*we)[key] = EventHandle{f: f, args: args}
}

func (e *EventHandle) Run(argsbuff []byte) []byte {
	br := seviper.NewReader(argsbuff)
	switch len(e.args) {
	case 0:
		return e.f()
	case 1:
		x := getArgs(br, e.args[0])
		return e.f(x)
	case 2:
		x := getArgs(br, e.args[0])
		y := getArgs(br, e.args[1])
		return e.f(x, y)
	case 3:
		x := getArgs(br, e.args[0])
		y := getArgs(br, e.args[1])
		z := getArgs(br, e.args[2])
		return e.f(x, y, z)
	case 4:
		x := getArgs(br, e.args[0])
		y := getArgs(br, e.args[1])
		z := getArgs(br, e.args[2])
		a := getArgs(br, e.args[3])
		return e.f(x, y, z, a)
	case 5:
		x := getArgs(br, e.args[0])
		y := getArgs(br, e.args[1])
		z := getArgs(br, e.args[2])
		a := getArgs(br, e.args[3])
		b := getArgs(br, e.args[4])
		return e.f(x, y, z, a, b)
	case 6:
		x := getArgs(br, e.args[0])
		y := getArgs(br, e.args[1])
		z := getArgs(br, e.args[2])
		a := getArgs(br, e.args[3])
		b := getArgs(br, e.args[4])
		c := getArgs(br, e.args[5])
		return e.f(x, y, z, a, b, c)
	}
	return nil
}

func getArgs(r *seviper.Reader, argtype ArgType) interface{} {
	switch argtype {
	case INT:
		return r.ToInt()
	case STRING:
		return r.ToString()
	case FLOAT:
		return r.ToFloat32()
	}
	return nil
}
