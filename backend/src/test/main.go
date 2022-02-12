package main

import (
	"fmt"
	"os"
	"log"
	"encoding/json"
)

type configuration struct {
	ServerAddress string `json:"server"`
}

func checkError(err error) {
	if err != nil {
		log.Fetal(err)
		os.Exit(1)
	}
}

func main() {
	file, err := os.Open("config.json")
	checkError(err)

	config := new(configuration)
	err = json.NewDecoder(file).Decode(config)
	checkError(err)

	test.RunServer()
}
