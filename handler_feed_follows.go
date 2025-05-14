package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/xaaaaaanny/gatorrss/internal/database"
	"time"
)

func handlerFeedFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enought arguments for follow command")
	}
	url := cmd.Args[0]

	userName := s.Config.CurrentUserName
	currentUser, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}
	userID := currentUser.ID

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}
	feedID := feed.ID

	feedFollowParam := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feedID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParam)
	if err != nil {
		return err
	}

	fmt.Println("Feed successfully followed")

	return nil
}

func handlerListFollowFeeds(s *state, cmd command) error {
	userName := s.Config.CurrentUserName
	currentUser, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}
	userID := currentUser.ID

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), userID)
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		return fmt.Errorf("No feeds followed yet")
	}

	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}
	return nil
}
