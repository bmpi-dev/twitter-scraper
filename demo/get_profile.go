package main

import (
    "fmt"
    twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
    scraper := twitterscraper.New()
    profile, err := scraper.GetProfile("madawei2699")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", profile)
}