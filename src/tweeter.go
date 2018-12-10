package main

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/gin-gonic/gin"
	"github.com/tweet/src/domain"
	"github.com/tweet/src/service"
	"net/http"
	"strconv"
)

func main() {
	tweeterWriter:= service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager( tweeterWriter )
	go launchServer(tweetManager)
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your user: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			tweet := domain.NewTextTweet(user, text)

			_,_ = tweetManager.PublishTweet(tweet)

			c.Printf("Tweet \n%s\nwas sent\n",tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows last tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := tweetManager.GetLastTweet()

			if tweet == nil{
				c.Println("Tweet are Empty")

			} else {
				c.Println(tweet)
			}

			return
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := tweetManager.GetTweets()

			if tweets == nil{
				c.Println("Tweets are Empty")

			} else {
				for _,tweet:= range tweets{
					c.Println(tweet)
				}
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetById",
		Help: "Shows tweet by id",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			c.Print("Write tweet id: ")


			idString := c.ReadLine()
			id, err := strconv.Atoi(idString)

			if err != nil {
				c.Println("id invalid")
				return
			}

			tweet := tweetManager.GetTweetById(id)
			if tweet == nil{
				c.Println("Tweets not found")
			} else {
				c.Println(tweet)
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsByUser",
		Help: "Shows tweets by user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)
			c.Print("Write user name : ")

			user := c.ReadLine()

			tweets := tweetManager.GetTweetsByUser(user);

			if tweets == nil{
				c.Println("Tweets not found")
				return
			}

			for _,tweet := range tweets{
				c.Println(tweet)
			}

			return
		},
	})

	shell.Run()

}

func launchServer(manager *service.TweetManager){
	router := gin.Default()

	router.GET("tweets/:usuario", func(context *gin.Context) {
		parametro := context.Param("usuario")
		response := ""
		for _,tweet := range manager.GetTweetsByUser(parametro){
			response += tweet.String() + "\n"
		}
		//context.String(http.StatusOK, response )
		context.JSON(http.StatusOK,manager.GetTweetsByUser(parametro))
	})

	router.GET("tweets", showAllTweets(manager) )

	router.POST("tweet",addTextTweet(manager))

	err := router.Run()
	fmt.Print(err)
}

func showAllTweets( manager *service.TweetManager )gin.HandlerFunc{
	return gin.HandlerFunc(
		func( c *gin.Context ){
			tweets := manager.GetTweets()
			c.JSON(http.StatusOK,tweets)
			})
}

func addTextTweet( manager *service.TweetManager )gin.HandlerFunc{
	return gin.HandlerFunc(func(c *gin.Context) {
		tweet := &domain.TextTweet{}
		_,err := manager.PublishTweet(tweet)
		if err != nil{
			c.JSON(http.StatusNotFound,"Error")
		}else{
			c.JSON(http.StatusOK ,tweet)
		}

	})

}