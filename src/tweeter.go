package main

import (
	"github.com/abiosoft/ishell"
	"github.com/tweet/src/domain"
	"github.com/tweet/src/service"
	"strconv"
)

func main() {
	tweetManager := service.NewTweetManager()
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