package domain

import (
	"fmt"
	"time"
)

type QuoteTweet struct {
	TextTweet
	quote Tweet
}

func NewQuoteTweet(user string,text string,quote Tweet) *QuoteTweet{
	time := time.Now()
	tweetText := TextTweet{-1,user,text,&time}
	tweet := QuoteTweet{tweetText,quote}
	return &tweet
}

func (tweet *QuoteTweet) PrintableTweet()string{
	return  fmt.Sprintf("@%s: %s \"%s\"", tweet.User,tweet.Text,tweet.quote)
}

func (tweet *QuoteTweet) String()string{
	return tweet.PrintableTweet()
}
