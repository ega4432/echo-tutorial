package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yoshimitsuEgashira/echo-tutorial/twitter-api/auth"

	"github.com/labstack/echo"
)

type Tweet struct {
	User string `json:"user"`
	Text string `json:"text"`
}

type TweetTemplate struct {
	User       string `json:"user"`
	Text       string `json:"text"`
	Id         string `json:"id"`
	ScreenName string `json:"screen_name"`
	Date       string `json:"date"`
	TweetId    string `json:"tweet_id"`
}

func SearchTweet(c echo.Context) error {
	keyword := c.FormValue("keyword")
	api, err := auth.ConnectTwitterAPI()
	if err != nil {
		fmt.Printf("Failed to connect : %v", api)
		return c.JSON(http.StatusInternalServerError, nil)
	}
	searchResult, err := api.GetSearch(`"`+keyword+`"`, nil)
	if err != nil {
		fmt.Printf("Failed to getsearch : %v", err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}
	tweets := make([]*Tweet, 0)

	for _, data := range searchResult.Statuses {
		tweet := new(Tweet)
		tweet.Text = data.FullText
		tweet.User = data.User.Name

		tweets = append(tweets, tweet)
	}
	return c.JSON(http.StatusOK, tweets)
}

func GetSearch(c echo.Context) error {
	value := c.QueryParam("value")
	api, err := auth.ConnectTwitterAPI()
	if err != nil {
		log.Printf("Failed to connect : %s", err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}
	searchResult, err := api.GetSearch(`"`+value+`"`, nil)
	if err != nil {
		log.Printf("Failed to GetSearch : %v", err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}
	tweets := make([]*TweetTemplate, 0)
	for _, data := range searchResult.Statuses {
		tweet := new(TweetTemplate)
		tweet.Text = data.FullText
		tweet.User = data.User.Name
		tweet.Id = data.User.IdStr
		tweet.ScreenName = data.CreatedAt
		tweet.Date = data.CreatedAt
		tweet.TweetId = data.IdStr
		tweets = append(tweets, tweet)
	}
	return c.Render(http.StatusOK, "tweet.html", tweets)
}
