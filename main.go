package main

import (
	"fmt"
	"github.com/xaaaaaanny/gatorrss/internal/config"
	"log"
	"os"
)

type state struct {
	Config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("can`t read config file: %v", err)
	}

	appState := &state{
		Config: &cfg,
	}

	existCommands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	if len(os.Args) < 2 {
		log.Fatal("less than 2 arguments")
	}

	existCommands.register("login", handlerLogin)
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = existCommands.run(appState, command{Name: cmdName, Args: cmdArgs})
	fmt.Println(err)
}
