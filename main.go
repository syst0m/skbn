package main

import (
	"log"
	"os"

	"github.com/syst0m/skbn/cmd"
)

func main() {
	cmd := cmd.NewRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		log.Fatal("Failed to execute command")
	}
}