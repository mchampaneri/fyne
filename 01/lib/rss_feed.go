package lib

import (
	"github.com/mmcdole/gofeed"
)

type Feed struct {
	Title, Image, Link string
}

// test : "http://feeds.twit.tv/twit.xml"
func ReadFeed(link string) ([]Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(link)
	data := make([]Feed, 0)
	if err != nil {
		return data, err
	}
	for l, i := range feed.Items {
		feed := Feed{
			Title: i.Title,
			Link:  i.Link,
		}
		if l > 100 {
			break
		}
		data = append(data, feed)
	}
	return data, nil
}
