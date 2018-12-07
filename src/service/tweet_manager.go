package service

import (
	"fmt"
	"github.com/tweet/src/domain"
)

type TweetManager struct {
	tweets []domain.Tweet
	tweetsByUser map[string][]domain.Tweet
}

func NewTweetManager() *TweetManager {
	tweetManager := &TweetManager{}
	tweetManager.InitializeService()
	return tweetManager
}


func (tweetManager *TweetManager) PublishTweet(tweetToPublish domain.Tweet) (int,error){

	if tweetToPublish.UserIsEmpty(){
		return -1,fmt.Errorf("user is required")
	}

	if tweetToPublish.TextIsEmpty(){
		return -1,fmt.Errorf("text is required")
	}

	if tweetToPublish.TextHasMoreCharactersThanMaxCharactersPerTweet() {
		return -1,fmt.Errorf("text has more characters than permited")
	}

	id := len(tweetManager.tweets)
	tweetToPublish.SetId(id)
	tweetManager.tweets = append(tweetManager.tweets,tweetToPublish)
	tweetManager.addTweetToUser(tweetToPublish.GetUser(),tweetToPublish)

	return id,nil
}

func (tweetManager *TweetManager) addTweetToUser(user string, tweet domain.Tweet)  {
	value,exist := tweetManager.tweetsByUser[user]

	if ! exist {
		value = make([]domain.Tweet,0)
	}

	tweetManager.tweetsByUser[user] = append( value, tweet )
}

func (tweetManager *TweetManager) GetTweetById(id int) domain.Tweet{
	for _,x := range tweetManager.tweets{
		if x.GetId() == id{
			return x
		}
	}
	return nil
}

func (tweetManager *TweetManager) CountTweetsByUser( user string ) int {

	countTweetsOfUser := 0

	for _,tweet := range tweetManager.tweets{
		if tweet.GetUser() == user{
			countTweetsOfUser++
		}
	}
	return  countTweetsOfUser
}


func (tweetManager *TweetManager) GetLastTweet() domain.Tweet{
	positionLastTweet := len(tweetManager.tweets) -1
	if positionLastTweet < 0{
		return nil
	}
	return tweetManager.tweets[positionLastTweet]

}

func (tweetManager *TweetManager) GetTweets() []domain.Tweet{
	return tweetManager.tweets
}

func (tweetManager *TweetManager) InitializeService() {
	tweetManager.tweets = make([]domain.Tweet,0)
	tweetManager.tweetsByUser =  make(map[string][]domain.Tweet)
}

func (tweetManager *TweetManager) GetTweetsByUser(user string)[]domain.Tweet{
	return tweetManager.tweetsByUser[user]
}

