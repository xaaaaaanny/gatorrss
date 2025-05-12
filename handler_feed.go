package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/xaaaaaanny/gatorrss/internal/database"
	"time"
)

func handlerAggregate(s *state, cmd command) error {
	const feedURL = "https://www.wagslane.dev/index.xml"

	rssFeed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	fmt.Println(rssFeed)
	return nil
}

func handlerCreateFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("not enought arguments for addFeed command")
	}
	title := cmd.Args[0]
	url := cmd.Args[1]

	userName := s.Config.CurrentUserName
	currentUser, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}
	userID := currentUser.ID

	feedParam := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      title,
		Url:       url,
		UserID:    userID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParam)
	if err != nil {
		return err
	}
	fmt.Printf("Feed %v successfuly added\n", feed.Name)
	return nil
}
