package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zhengheetong/gator/internal/database"
)

func handlerFeedAdd(s *state, cmd command, user database.User) error {
	if len(cmd .Args) != 2{
		return fmt.Errorf("usage: %v <name> <url>",cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]
	id := uuid.New();
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: id,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: name,
		Url: url,
		UserID: user.ID,
	})
	if err != nil{
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	fmt.Println("Feeds Created Successfully")
	printFeed(feed)

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: id,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	fmt.Println("Feeds Follow Created Successfully")

	return nil
	
}

func handlerFeeds(s *state, cmd command) error {
	items, err := s.db.GetFeeds(context.Background())
	if err != nil{
		return fmt.Errorf("couldn't get feeds: %w", err)
	}
	fmt.Println("Name	|URL		|Created by")
	for i :=0; i < len(items); i++ {
		user,_ := s.db.GetUserByID(context.Background(), items[i].UserID)
		fmt.Printf("%v		|%v		|%v\n", items[i].Name, items[i].Url, user.Name)
	}
	
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
