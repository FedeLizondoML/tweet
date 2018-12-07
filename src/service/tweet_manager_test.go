package service_test

import (
	"github.com/tweet/src/domain"
	"github.com/tweet/src/service"
	"testing"
)

func TestPublishedTweetIsSave(  t *testing.T ){
	var tweet *domain.TextTweet
	user := "fede"
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	tweetManager := service.NewTweetManager()

	_,_ = tweetManager.PublishTweet(tweet)

	publishTweet := tweetManager.GetLastTweet()

	if !isValidTweet(t,tweet,user,text){
		t.Errorf("Expected tweet is %s \nbut is %s: %s",user,publishTweet.GetUser(),publishTweet.GetText())
	}

	if publishTweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T){
	var tweet domain.Tweet

	var user string
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)
	tweetManager := service.NewTweetManager()

	var err error
	_,err = tweetManager.PublishTweet(tweet)

	if err != nil && err.Error() != "user is required"{
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished( t *testing.T ){
	var tweet domain.Tweet
	tweetManager := service.NewTweetManager()
	user := "fede"
	var text string
	tweet = domain.NewTextTweet(user, text)


	var err error
	_,err = tweetManager.PublishTweet(tweet)

	if err != nil && err.Error() != "text is required"{
		t.Error("Expected error is text is required")
	}
}

func TestTweetWitchExceeding140CharactersIsNotPublished(t *testing.T) {
	var tweet domain.Tweet

	user := "fede"
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ante erat, commodo at dignissim ac, vehicula ac eros. In viverra sed nibh ac cras amet."
	tweet = domain.NewTextTweet(user, text)

	tweetManager := service.NewTweetManager()
	var err error
	_,err = tweetManager.PublishTweet(tweet)

	if err != nil && err.Error() != "text has more characters than permited"{
		t.Error("Expected error is text is required")
	}
}

func TestCanPublishAndRetriveMoreThanOneTweet(t *testing.T) {
	tweetManager := service.NewTweetManager()
	var tweet,secondTweet domain.Tweet

	tweet = domain.NewTextTweet("fede","primerTweet")
	secondTweet = domain.NewTextTweet("fede","segundoTweet")

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)

	publishTweets := tweetManager.GetTweets()
	if len(publishTweets) != 2{
		t.Errorf("Expected size is 2 but was %d",len(publishTweets))
		return
	}

	firstPublishTweet := publishTweets[0]
	secondPublishTweet := publishTweets[1]


	if !areValidTweet(t,firstPublishTweet,tweet){
		t.Errorf("Expected tweet is user %s and text %s but was user %s and text %s",
			tweet.GetUser(),tweet.GetText(),firstPublishTweet.GetUser(),firstPublishTweet.GetText())
		return
	}

	if !areValidTweet(t,secondPublishTweet,secondTweet){
		t.Errorf("Expected tweet is user %s and text %s but was user %s and text %s",
			secondTweet.GetUser(),secondTweet.GetText(),secondPublishTweet.GetUser(),secondPublishTweet.GetText())
		return
	}

}


func TestCanRetriveTweetById(t *testing.T){
	tweetManager := service.NewTweetManager()

	user := "Marco Polo"
	text := "tweet"
	tweet := domain.NewTextTweet(user,text)

	id,_ := tweetManager.PublishTweet(tweet)

	publishTweet := tweetManager.GetTweetById(id)

	if !isValidTweet(t,publishTweet,user,text){
		t.Errorf("Expected tweet is user %s and text %s but was user %s and text %s",
			publishTweet.GetUser(),publishTweet.GetText(),user,text)
		return
	}
}


func TestCanCountThewTweetsSentByAnUser(t *testing.T){
	tweetManager := service.NewTweetManager()
	user := "Marco"
	anotherUser := "Pedro"
	text1 := "text1"
	text2 := "text2"
	text3 := "text3"
	text4 := "text4"

	tweetManager.PublishTweet(domain.NewTextTweet(user,text1))
	tweetManager.PublishTweet(domain.NewTextTweet(anotherUser,text2))
	tweetManager.PublishTweet(domain.NewTextTweet(user,text3))
	tweetManager.PublishTweet(domain.NewTextTweet(anotherUser,text1))
	tweetManager.PublishTweet(domain.NewTextTweet(anotherUser,text4))

	countOfTweets := tweetManager.CountTweetsByUser(user)

	if countOfTweets != 2{
		t.Errorf("Expected tweets is 2 but was %d",countOfTweets)
		return
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager()

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet := domain.NewTextTweet(user, text)
	secondTweet := domain.NewTextTweet(user, secondText)
	thirdTweet := domain.NewTextTweet(anotherUser, text)
	// publish the 3 tweets

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	if tweets == nil{
		t.Error("Tweets sin usar")
		return
	}

	if len(tweets) != 2 {
		t.Errorf("Expected 2 tweets but was %d ",len(tweets))
		return
	}
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet
	if !areValidTweet(t,tweet,firstPublishedTweet){
		t.Errorf("Expected 1º tweet to be user %s and text %s but was user %s and text %s",
			tweet.User,tweet.Text,firstPublishedTweet.GetUser(),firstPublishedTweet.GetText())
		return
	}

	if !areValidTweet(t,secondTweet,secondPublishedTweet){
		t.Errorf("Expected 1º tweet to be user %s and text %s but was user %s and text %s",
			secondTweet.User,secondTweet.Text,secondPublishedTweet.GetUser(),secondPublishedTweet.GetText())
		return
	}

}

func isValidTweet(t *testing.T,tweet domain.Tweet,user string,text string)bool{
	return tweet.GetUser() == user && tweet.GetText() == text
}

func areValidTweet(t *testing.T,tweet domain.Tweet,tweetToCompare domain.Tweet)bool{
	return isValidTweet(t,tweet,tweetToCompare.GetUser(),tweetToCompare.GetText())
}