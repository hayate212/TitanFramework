package main

import (
	"log"

	t "github.com/hayate212/TitanFramework/.titan"
	"github.com/hayate212/TitanFramework/app"
)

func main() {
	var config t.TitanConfig
	err := t.ConfigLoad(&config)
	if err != nil {
		log.Fatal(err)
	}
	we := t.NewWorkerEventHandler()
	w := t.NewWorker(
		t.WorkerConfig{
			Address:        config.Address,
			Port:           config.Port,
			MaxRequestSize: config.MaxRequestSize,
		},
		we,
	)
	app.Init(w)
	w.Run()
}
