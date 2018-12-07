package domain

import (
	"fmt"
	"time"
)

type ImageTweet struct {
	TextTweet
	Url string
}

func NewImageTweet(user string,text string,url string) *ImageTweet{
	time := time.Now()
	tweetText := TextTweet{-1,user,text,&time}
	tweet := ImageTweet{tweetText,url}
	return &tweet
}

func (tweet *ImageTweet) PrintableTweet()string{
	return  fmt.Sprintf("@%s: %s\n%s", tweet.User,tweet.Text,tweet.Url)
}

func (tweet *ImageTweet) String()string{
	return tweet.PrintableTweet()
}
