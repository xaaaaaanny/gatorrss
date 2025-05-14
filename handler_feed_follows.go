package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/xaaaaaanny/gatorrss/internal/database"
	"time"
)

func handlerFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enought arguments for follow command")
	}
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}
	feedID := feed.ID

	feedFollowParam := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParam)
	if err != nil {
		return err
	}

	fmt.Println("Feed successfully followed")

	return nil
}

func handlerListFollowFeeds(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds followed yet")
		return nil
	}

	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}
	return nil
}

func handlerDeleteFeedFromUser(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enought arguments for delete command")
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	args := database.DeleteFeedFromUserParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFeedFromUser(context.Background(), args)
	if err != nil {
		return err
	}
	fmt.Printf("%v unfollowed %v feed", user.Name, feed.Name)

	return nil
}
