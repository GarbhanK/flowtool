package main

import (
	"github.com/garbhank/flowtool/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("Error executing command entrypoint: %s", err.Error())
	}
}

