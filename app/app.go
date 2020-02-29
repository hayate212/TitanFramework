package app

import (
	"fmt"

	t "github.com/hayate212/TitanFramework/.titan"
)

func Init(w *t.Worker) {
	w.EventHandler.Set("test", TestFunc, []t.ArgType{t.INT, t.INT})
}

func TestFunc(args ...interface{}) {
	fmt.Println(args[0].(int), args[1].(int))
}
