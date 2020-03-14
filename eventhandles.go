package titan

import "reflect"

type EventHandles map[string]*EventHandle

func NewEventHandles() *EventHandles {
	return &EventHandles{}
}

func (hs *EventHandles) AddHandle(i interface{}) bool {
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Ptr {
		return false
	}
	name := t.Elem().Name()
	(*hs)[name] = NewEventHandle(i)
	return true
}

func (hs *EventHandles) Proc(name string, args []reflect.Value) ([]reflect.Value, bool) {
	d := -1
	for i, c := range name {
		if c == 0x2e {
			d = i
			break
		}
	}
	if 0 > d {
		return nil, false
	}

	class, method := name[0:d], name[d+1:]
	return (*hs)[class].Proc(method, args)
}
