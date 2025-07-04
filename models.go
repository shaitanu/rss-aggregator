package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/rss-aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey		string `json:"api_key"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey: dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	Name      string	`json:"name"`
	Url       string	`json:"url"`
	UserID    uuid.UUID	`json:"user_id"`
}

type FeedFollow struct{
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	FeedID uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID	`json:"user_id"`
	

}

func datbaaseFeedtoFeed(dbFeed database.Feed)Feed{
return Feed{
	ID:dbFeed.ID,
	CreatedAt :dbFeed.CreatedAt,
	UpdatedAt: dbFeed.UpdatedAt,
	Name:dbFeed.Name,
	Url: dbFeed.Url,
	UserID: dbFeed.UserID,

}
}

func datbaaseFeedstoFeeds(dbFeeds []database.Feed)[]Feed{
feeds:=[]Feed{}
for _,dbFeed:=range dbFeeds{
feeds=append(feeds,datbaaseFeedtoFeed(dbFeed))	
}
return feeds
}

func databaseFeedFollowtoFeedFollow(dbFeedFollow database.FeedFollow)FeedFollow{
	return FeedFollow{
		ID: dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID: dbFeedFollow.UserID,
		FeedID: dbFeedFollow.FeedID,
	}
}
func datbaseFeedFollowstoFollows(dbFeedFollows []database.FeedFollow)[]FeedFollow{
feedFollows:=[]FeedFollow{}
for _,dbFeedFollow:=range dbFeedFollows{
feedFollows=append(feedFollows,databaseFeedFollowtoFeedFollow(dbFeedFollow))	
}
return feedFollows
}