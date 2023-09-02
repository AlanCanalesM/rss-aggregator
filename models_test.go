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
