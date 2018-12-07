package domain_test

import (
	"github.com/tweet/src/domain"
	"testing"
)



func TestTweet_CreateATweet(t *testing.T) {
	tweet := domain.NewTextTweet("User", "This is my tweet")

	if tweet.UserIsEmpty(){
		t.Errorf("The expected user text is %s but was '' ", tweet.User)
		return
	}

	if tweet.TextIsEmpty(){
		t.Errorf("The expected text is %s but was '' ", tweet.Text)
		return
	}

	if tweet.TextHasMoreCharactersThanMaxCharactersPerTweet(){
		t.Errorf("The expected legth is less than %d but was %d ", domain.MAX_CHARACTERS_PER_TWEET, len(tweet.Text))
		return
	}
}


func TestTweet_UserIsEmpty(t *testing.T) {
	tweet := domain.NewTextTweet("", "This is my tweet")
	if !tweet.UserIsEmpty(){
		t.Errorf("The expected user is '' but was %s", tweet.User)
	}
}

func TestTweet_TextIsEmpty(t *testing.T) {
	tweet := domain.NewTextTweet("User", "")
	if !tweet.TextIsEmpty(){
		t.Errorf("The expected text is '' but was %s", tweet.User)
	}
}

func TestTweet_UserIsEmptyAndTextIsEmpty(t *testing.T) {
	tweet := domain.NewTextTweet("", "")
	if !tweet.UserIsEmpty(){
		t.Errorf("The expected text is '' but was %s", tweet.User)
	}
}

func TestTextTweetPrintsUserAndText(t *testing.T) {
	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation

	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}

func TestCanGetAStringFromTweet(t *testing.T) {
	// Initialization
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}