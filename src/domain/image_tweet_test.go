package domain_test

import (
	"github.com/tweet/src/domain"
	"testing"
)

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	tweet := domain.NewImageTweet("grupoesfera", "This is my image",
		"http://www.grupoesfera.com.ar/common/img/grupoesfera.png")
	// Operation
	text := tweet.PrintableTweet()
	// Validation
	expectedText := "@grupoesfera: This is my image\nhttp://www.grupoesfera.com.ar/common/img/grupoesfera.png"

	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}

func TestCanGetAStringFromImageTweet(t *testing.T) {
	// Initialization
	tweet := domain.NewImageTweet("grupoesfera", "This is my tweet","img")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet\nimg"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}





