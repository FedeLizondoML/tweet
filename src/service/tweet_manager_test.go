package service_test

import (
	"github.com/tweet/src/service"
	"testing"
)

func TestPublishedTweetIsSave(  t *testing.T ){
	 tweet := "This is my first tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}
}