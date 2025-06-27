package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zhengheetong/gator/internal/database"
);

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedbyURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't create feed_follow: %w", err)
	}

	ff, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	fmt.Println("Feeds Follow Created Successfully")
	printFeedFollow(s,ff)
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %w <url>", cmd.Name)
	}
	url := cmd.Args[0]
	feed, _ := s.db.GetFeedbyURL(context.Background(), url)

	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't Delete Feed Follow: %w", err)
	}
	fmt.Println("Feeds Unfollowed Successfully")

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	items, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't find following feeds: %w", err)
	}
	
	fmt.Println("Following Feed:")
	for i := 0; i < len(items); i++ {
		feed, _ := s.db.GetFeedbyID(context.Background(), items[i].FeedID)
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
	}
	return nil
}

func printFeedFollow(s *state,ff database.FeedFollow){
	user, _ := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	feed, _ := s.db.GetFeedbyID(context.Background(), ff.FeedID)
	fmt.Printf(" * ID:	%v\n", ff.ID)
	fmt.Printf(" * User:	%v\n", user.Name)
	fmt.Printf(" * Feed:	%v\n", feed.Name)
}

