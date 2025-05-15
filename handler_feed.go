package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/xaaaaaanny/gatorrss/internal/database"
	"time"
)

func handlerCreateFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("not enought arguments for addFeed command")
	}
	title := cmd.Args[0]
	url := cmd.Args[1]

	feedParam := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      title,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParam)
	if err != nil {
		return err
	}
	fmt.Printf("Feed %v successfuly created\n", feed.Name)

	feedFollowParam := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParam)
	if err != nil {
		return err
	}
	fmt.Println("Feed successfully followed")
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		return fmt.Errorf("No feeds yet")
	}

	for _, feed := range feeds {
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		user, err := s.db.GetUsernameByUserId(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println(user.Name)
	}
	return nil
}
