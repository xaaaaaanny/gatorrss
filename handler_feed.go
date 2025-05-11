package main

import (
	"context"
	"fmt"
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
