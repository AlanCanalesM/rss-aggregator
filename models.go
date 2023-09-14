package main

import (
	"time"

	"github.com/AlanCanalesM/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

// This file will contain the models for the application
// We will have a model for each table in the database
// We will also have a function to convert from the database model to the application model

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

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(dbPost database.Post) Post {

	var description *string

	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {

	posts := []Post{}

	for _, dbPost := range dbPosts {
		posts = append(posts, databasePostToPost(dbPost))
	}

	return posts
}
