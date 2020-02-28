package titan

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type TitanConfig struct {
	Address string `json:"Address"`
	Port    int    `json:"Port"`
}

func ConfigLoad(config *TitanConfig) error {
	buff, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(buff, &config); err != nil {
		return err
	}
	return nil
}
