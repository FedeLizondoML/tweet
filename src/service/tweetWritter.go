package service

import "github.com/tweet/src/domain"

type TweetWriter interface {
	Write(tweet domain.Tweet)
}


