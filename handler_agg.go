package main

import (
	"context"
	"fmt"
	"github.com/xaaaaaanny/gatorrss/internal/database"
	"time"
)

func handlerAggregate(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not enought arguments for aggregate command")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %v", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	
	return nil
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("Can`t get feed to fetch: %v\n", err)
		return
	}
	fmt.Printf("Found feed to fetch \n")
	scrapeFeed(s, feed)
}

func scrapeFeed(s *state, feed database.Feed) {
	err := s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Printf("Can`t mark feed: %v\n", err)
		return
	}

	fetchedFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("Can`t fetch feed: %v\n", err)
		return
	}

	for _, item := range fetchedFeed.Channel.Item {
		fmt.Printf("Post found: %v\n", item.Title)
	}

	fmt.Printf("Feed %s collected, %v posts found", feed.Name, len(fetchedFeed.Channel.Item))
}
