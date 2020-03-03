package app

import (
	"fmt"

	t "github.com/hayate212/TitanFramework/.titan"
)

type Handle struct{}

func Init(w *t.Worker) {
	w.SetEventHandle(&Handle{})
}

func (h *Handle) TestFunc(x, y int) string {
	fmt.Println(x, y)
	return "success"
}
