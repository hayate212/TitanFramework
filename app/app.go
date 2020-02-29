package app

import (
	"fmt"

	"github.com/hayate212/seviper"

	t "github.com/hayate212/TitanFramework/.titan"
)

func Init(w *t.Worker) {
	w.EventHandler.Set("test", TestFunc, []t.ArgType{t.INT, t.INT})
}

func TestFunc(args ...interface{}) []byte {
	fmt.Println(args[0].(int), args[1].(int))
	w := seviper.NewWriter()
	w.Write(args[0].(int))
	w.Write(args[1].(int))
	return w.Bytes
}
