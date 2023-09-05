package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/AlanCanalesM/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func Test_databaseUserToUser(t *testing.T) {
	type args struct {
		dbUser database.User
	}

	testUuid := uuid.New()
	timeStamp := time.Now()
	tests := []struct {
		name string
		args args
		want User
	}{
		{
			name: "Valid user",
			args: args{
				database.User{
					ID:        testUuid,
					CreatedAt: timeStamp,
					UpdatedAt: timeStamp,
					Name:      "Alan",
				},
			},
			want: User{
				ID:         testUuid,
				Created_at: timeStamp,
				Updated_at: timeStamp,
				Name:       "Alan",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := databaseUserToUser(tt.args.dbUser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("databaseUserToUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_databaseFeedsToFeeds(t *testing.T) {
	testUuid := uuid.New()
	timeStamp := time.Now()
	type args struct {
		dbFeeds []database.Feed
	}
	tests := []struct {
		name string
		args args
		want []Feed
	}{
		{
			name: "Valid test",
			args: args{
				dbFeeds: []database.Feed{
					database.Feed{
						ID:        testUuid,
						CreatedAt: timeStamp,
						UpdatedAt: timeStamp,
						Name:      "someblog",
						Url:       "myurl",
						UserID:    testUuid,
					},
				},
			},
			want: []Feed{
				Feed{
					ID:         testUuid,
					Created_at: timeStamp,
					Updated_at: timeStamp,
					Name:       "someblog",
					URL:        "myurl",
					UserID:     testUuid,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := databaseFeedsToFeeds(tt.args.dbFeeds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("databaseFeedsToFeeds() = %v, want %v", got, tt.want)
			}
		})
	}
}
