package main

import (
    "context"
    "fmt"
    twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
    scraper := twitterscraper.New()

    for tweet := range scraper.GetTweets(context.Background(), "madawei2699", 3000) {
        if tweet.Error != nil {
            panic(tweet.Error)
        }
        fmt.Println("============>")
        fmt.Println(tweet)
    }
}
