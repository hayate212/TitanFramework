package titan

import (
	"fmt"
	"reflect"

	"github.com/hayate212/seviper"
)

type EventHandle struct {
	i interface{}
	v reflect.Value
}

func NewEventHandle(i interface{}) *EventHandle {
	return &EventHandle{i: i, v: reflect.ValueOf(i)}
}

func (h *EventHandle) Proc(name string, args []reflect.Value) ([]reflect.Value, bool) {
	m := h.v.MethodByName(name)
	if m.Kind() != reflect.Func {
		return nil, false
	}
	return m.Call(args), true
}

func (h *EventHandle) BytesToArgs(name string, bytes []byte) ([]reflect.Value, bool) {
	m := h.v.MethodByName(name)
	if m.Kind() != reflect.Func {
		return nil, false
	}
	r := seviper.NewReader(bytes)
	t := m.Type()
	args := make([]reflect.Value, t.NumIn())
	for i := 0; i < len(args); i++ {
		k := t.In(i).Kind()
		switch fmt.Sprintf("%v", k) {
		case "string":
			args[i] = reflect.ValueOf(r.ToString())
		case "int":
			args[i] = reflect.ValueOf(r.ToInt())
		case "float32":
			args[i] = reflect.ValueOf(r.ToFloat32())
		case "float64":
			args[i] = reflect.ValueOf(r.ToFloat64())
		}
	}
	return args, true
}
