package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: ./tcmd [tweetURL]")
	}
	tweetURL := os.Args[1]
	tweetID, err := getTweetID(tweetURL)
	if err != nil {
		log.Fatal(err)
	}
	imgURL, err := url.Parse(fmt.Sprintf("https://tweet-card.now.sh/%s.png", tweetID))
	if err != nil {
		log.Fatal(err)
	}
	q := imgURL.Query()
	q.Add("lang", "ja")
	q.Add("tz", "9")
	q.Add("scale", "1")
	imgURL.RawQuery = q.Encode()
	log.Printf("%s", imgURL)
	_, _ = fmt.Fprintf(os.Stdout, "[![%s](%s)](%s)", tweetURL, imgURL, tweetURL)
}

func getTweetID(tweetURL string) (string, error) {
	ms := regexp.MustCompile(`^https://twitter.com/(.+)/status/(\d+)`).FindStringSubmatch(tweetURL)
	if len(ms) < 3 {
		return "", fmt.Errorf("invalid tweet URL: %s", tweetURL)
	}
	return ms[2], nil
}