package domain

import (
	"fmt"
	"strings"
	"time"
)

type TextTweet struct {
	Id int
	User string
	Text string
	Date *time.Time
}

const MAX_CHARACTERS_PER_TWEET int = 140

func NewTextTweet(user string,text string) *TextTweet{
	time  := time.Now()
	tweet := TextTweet{-1,user,text,&time}
	return &tweet
}

func (tweet *TextTweet) UserIsEmpty() bool {
	return strings.TrimSpace( tweet.User ) == ""
}

func (tweet *TextTweet) TextIsEmpty() bool {
	return strings.TrimSpace( tweet.Text ) == ""
}

func (tweet *TextTweet) TextHasMoreCharactersThanMaxCharactersPerTweet() bool{
	return len(tweet.Text) > MAX_CHARACTERS_PER_TWEET
}

func (tweet *TextTweet) IsEquals(anotherTweet *TextTweet) bool {
	if tweet == nil || anotherTweet == nil{
		return tweet == nil && anotherTweet == nil
	}
	return tweet.User == anotherTweet.User && tweet.Text == anotherTweet.Text
}

func (tweet *TextTweet) PrintableTweet()string{
	return  fmt.Sprintf("@%s: %s", tweet.User,tweet.Text)
}

func (tweet *TextTweet) String()string{
	return tweet.PrintableTweet();
}

func (tweet *TextTweet) GetUser()string{
	return tweet.User
}

func (tweet *TextTweet) GetText()string{
	return tweet.Text
}

func (tweet *TextTweet) GetId() int {
	return tweet.Id
}

func (tweet *TextTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *TextTweet) GetDate() *time.Time{
	return tweet.Date
}