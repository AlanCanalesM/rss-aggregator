package main

import (
	"time"

	"github.com/AlanCanalesM/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

// User represents the User model in the application.
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"apikey"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
}

// Feed represents the Feed model in the application.
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

// FeedFollow represents the FeedFollow model in the application.
type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

// Post represents the Post model in the application.
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

// databaseUserToUser converts a database User to an application User.
func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.Apikey,
		Username:  dbUser.Username,
		Password:  dbUser.Password,
	}
}

// databaseFeedToFeed converts a database Feed to an application Feed.
func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		URL:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

// databaseFeedsToFeeds converts a slice of database Feeds to application Feeds.
func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, len(dbFeeds))
	for i, feed := range dbFeeds {
		feeds[i] = databaseFeedToFeed(feed)
	}
	return feeds
}

// databaseFeedFollowToFeedFollow converts a database FeedFollow to an application FeedFollow.
func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

// databaseFeedFollowsToFeedFollows converts a slice of database FeedFollows to application FeedFollows.
func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := make([]FeedFollow, len(dbFeedFollows))
	for i, feedFollow := range dbFeedFollows {
		feedFollows[i] = databaseFeedFollowToFeedFollow(feedFollow)
	}
	return feedFollows
}

// databasePostToPost converts a database Post to an application Post.
func databasePostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		desc := dbPost.Description.String
		description = &desc
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

// databasePostsToPosts converts a slice of database Posts to application Posts.
func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := make([]Post, len(dbPosts))
	for i, post := range dbPosts {
		posts[i] = databasePostToPost(post)
	}
	return posts
}
