package titan

import "reflect"

type EventHandles map[string]*EventHandle

func NewEventHandles() *EventHandles {
	return &EventHandles{}
}

func (hs *EventHandles) AddHandle(i interface{}, opt ...string) bool {
	name := ""
	if len(opt) > 0 {
		name = opt[0]
	} else {
		t := reflect.TypeOf(i)
		if t.Kind() != reflect.Ptr {
			return false
		}
		name = t.Elem().Name()
	}
	if name != "" {
		(*hs)[name] = NewEventHandle(i)
		return true
	}
	return false
}

func (hs *EventHandles) Proc(name string, args []byte) ([]reflect.Value, bool) {
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
	h := (*hs)[class]
	if args, ok := h.BytesToArgs(method, args); ok {
		return h.Proc(method, args)
	}
	return nil, false
}
