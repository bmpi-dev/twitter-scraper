package main

import (
	"context"
	"fmt"

	twitterscraper "github.com/bmpi-dev/twitter-scraper"
	"github.com/pemistahl/lingua-go"
)

func main() {

	languages := []lingua.Language{
		lingua.English,
		lingua.Chinese,
	}

	detector := lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		Build()

	scraper := twitterscraper.New()

	for tweet := range scraper.GetTweets(context.Background(), "madawei2699", 3000) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		if language, exists := detector.DetectLanguageOf(tweet.Text); exists {
			tweet.LanguageType = language.String()
		}
		fmt.Println("============>")
		fmt.Println(tweet)
	}
}
