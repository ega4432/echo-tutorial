package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ChimeraCoder/anaconda"
)

type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

func ConnectTwitterAPI() (*anaconda.TwitterApi, error) {
	row, err := ioutil.ReadFile("./token.json")
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
