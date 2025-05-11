package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/xaaaaaanny/gatorrss/internal/database"
	"os"
	"time"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no second argument for login command")
	}
	user, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		os.Exit(1)
		return fmt.Errorf("cant login to not existing user: %v", err)
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("can`t set user: %v", err)
	}

	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("no second argument for register command")
	}
	name := cmd.Args[0]

	userParam := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}
	user, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		os.Exit(1)
		return fmt.Errorf("user already exist: %v", err)
	}

	user, err = s.db.CreateUser(context.Background(), userParam)
	if err != nil {
		return fmt.Errorf("cant create user: %v", err)
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("can`t set user: %v", err)
	}

	fmt.Printf("User - %v, %v, %v, %v was created\n", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user == s.Config.CurrentUserName {
			fmt.Printf("* %v (current)\n", user)
			continue
		}
		fmt.Printf("* %v\n", user)
	}

	return nil
}
