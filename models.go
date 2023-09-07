package main

import (
	"time"

	"github.com/AlanCanalesM/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	APIKey     string    `json:"apikey"`
}

func databaseUserToUser(dbUser database.User) User {

	return User{
		ID:         dbUser.ID,
		Created_at: dbUser.CreatedAt,
		Updated_at: dbUser.UpdatedAt,
		Name:       dbUser.Name,
		APIKey:     dbUser.Apikey,
	}
}

type Feed struct {
	ID         uuid.UUID `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	URL        string    `json:"url"`
	UserID     uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {

	return Feed{
		ID:         dbFeed.ID,
		Created_at: dbFeed.CreatedAt,
		Updated_at: dbFeed.UpdatedAt,
		Name:       dbFeed.Name,
		URL:        dbFeed.Url,
		UserID:     dbFeed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {

	feeds := []Feed{}

	for _, feed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(feed))
	}

	return feeds

}

type FeedFollow struct {
	ID         uuid.UUID `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	UserID     uuid.UUID `json:"user_id"`
	FeedID     uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {

	return FeedFollow{
		ID:         dbFeedFollow.ID,
		Created_at: dbFeedFollow.CreatedAt,
		Updated_at: dbFeedFollow.UpdatedAt,
		UserID:     dbFeedFollow.UserID,
		FeedID:     dbFeedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {

	feedFollows := []FeedFollow{}

	for _, feedFollow := range dbFeedFollows {

		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(feedFollow))
	}

	return feedFollows

}
