package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/ChimeraCoder/anaconda"
	"github.com/labstack/echo"
)

type Tweet struct {
	User string `json:"user"`
	Text string `json:"text"`
}

type TweetTemplate struct {
	User       string `json:"user"`
	Text       string `json:"text"`
	ScreenName string `json:"screen_name"`
	Id         string `json:"id"`
	Date       string `json:"date"`
	TweetId    string `json:"tweet_id"`
}

//type Tweets *[]Tweet

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
	value := c.QueryParam("value")
	return c.Render(http.StatusOK, "hello", value)
}

func main() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("./twitter-api/public/views/*.html")),
	}
	e.Renderer = t
	e.POST("/search", searchTweet)
	e.GET("/hello", Hello)
	e.Logger.Fatal(e.Start(":1232"))
}

func searchTweet(c echo.Context) error {
	keyword := c.FormValue("keyword")
	api, err := connectTwitterAPI()
	if err != nil {
		fmt.Printf("Failed to connect : %v", api)
		return c.JSON(http.StatusInternalServerError, api)
	}
	fmt.Printf("auth info : %#v", api)
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

func connectTwitterAPI() (*anaconda.TwitterApi, error) {
	row, err := ioutil.ReadFile("./twitter-api/token.json")
	if err != nil {
		fmt.Printf("Failed to read json file: %s", err.Error())
		return nil, err
	}
	var twitterAccount TwitterAccount
	err = json.Unmarshal(row, &twitterAccount)
	if err != nil {
		fmt.Printf("Failed to parse to json: %s", err.Error())
		return nil, err
	}
	return anaconda.NewTwitterApiWithCredentials(twitterAccount.AccessToken, twitterAccount.AccessTokenSecret, twitterAccount.ConsumerKey, twitterAccount.ConsumerSecret), nil
}

type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}
