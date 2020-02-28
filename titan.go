package main

import (
	t "github.com/hayate212/TitanFramework/.titan"
	"github.com/hayate212/TitanFramework/app"
)

func main() {
	w := t.NewWorker(t.WorkerConfig{Address: "localhost:9999"})
	app.Init(w)
	w.Run()
}
