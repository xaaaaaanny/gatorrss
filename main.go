package main

import (
	"fmt"
	"github.com/xaaaaaanny/gatorrss/internal/config"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("can`t read config file: %v", err)
	}
	fmt.Println(cfg)

	err = cfg.SetUser("xanny")
	if err != nil {
		log.Fatalf("can`t set user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("can`t read config file: %v", err)
	}

	fmt.Println(cfg)
}
