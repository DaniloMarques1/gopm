package main

import (
	"errors"
	"log"
	"os"

	"github.com/danilomarques1/gopm/cmd"
)

func main() {
	command, err := parseCmdArgs()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Run(command)
}

func parseCmdArgs() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("You should provide a command. Type help for instructions")
	}
	command := os.Args[1]
	return command, nil
}
