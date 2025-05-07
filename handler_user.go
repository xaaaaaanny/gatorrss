package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no second argument for login command")
	}
	err := s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("can`t set user: %v", err)
	}

	fmt.Println("User has been set")
	return nil
}
