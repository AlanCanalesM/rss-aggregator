package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/AlanCanalesM/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

// ScrapingConfig contains the configuration for the scraping process.
type ScrapingConfig struct {
	DB                  *database.Queries
	Concurrency         int
	TimeBetweenRequests time.Duration
}

// StartScraping initiates the scraping process.
func StartScraping(config ScrapingConfig) {
	log.Printf("Scraping on %v goroutines every %s duration", config.Concurrency, config.TimeBetweenRequests)

	ticker := time.NewTicker(config.TimeBetweenRequests)

	for range ticker.C {
		feeds, err := config.DB.GetNextFeedToFetch(context.Background(), int32(config.Concurrency))
		if err != nil {
			log.Println("Error fetching feeds: ", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(config.DB, wg, feed)
		}

		wg.Wait()
	}
}

// ScrapeFeed scrapes a feed and saves the posts to the database.
func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched: ", err)
		return
	}

	rssFeed, err := FetchRSSFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed: ", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}

		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing date %v with err %v", item.PubDate, err)
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: pubAt,
			Url:         item.Link,
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("Error creating the post: ", err)
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
