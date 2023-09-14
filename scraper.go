package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/AlanCanalesM/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

// This function will start the scraping process
// It will fetch the feeds from the database
// and then it will scrape each feed
// It will run on a goroutine and it will run every timeBetweenRequest duration
func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {

	log.Printf("Scraping on %v goroutines every %s durarion", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {

		feeds, err := db.GetNextFeedToFetch(context.Background(), int32(concurrency))

		if err != nil {
			log.Println("Error fetching feeds: ", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {

			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}

		wg.Wait()
	}
}

// This function will scrape a feed
// It will fetch the feed from the url
// and then it will save the posts in the database
func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marked feed: ", err)
	}

	rssFeed, err := urlToFeed(feed.Url)

	if err != nil {
		log.Println("Error fetching feed: ", err)
	}

	for _, item := range rssFeed.Channel.Item {

		description := sql.NullString{}

		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)

		if err != nil {
			fmt.Printf("We could not parse the date %v with err %v", item.PubDate, err)
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
			fmt.Println("We couldn't create the post", err)
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))

}
