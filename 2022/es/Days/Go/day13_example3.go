package main

import (
	// other imports
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Credentials stores all of our access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// getClient is a helper function that will return a twitter client
// that we can subsequently use to send tweets, or to stream new tweets
// this will take in a pointer to a Credential struct which will contain
// everything needed to authenticate and return a pointer to a twitter Client
// or an error
func getClient(creds *Credentials) (*twitter.Client, error) {
	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Pass in your Access Token and your Access Token Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}
func main() {
	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}
	{
		const DaysTotal int = 90
		var remainingDays uint = 90
		challenge := "#90DaysOfDevOps"

		fmt.Printf("Welcome to the %v challenge.\nThis challenge consists of %v days\n", challenge, DaysTotal)

		var TwitterName string
		var DaysCompleted uint

		// asking for user input
		fmt.Println("Enter Your Twitter Handle: ")
		fmt.Scanln(&TwitterName)

		fmt.Println("How many days have you completed?: ")
		fmt.Scanln(&DaysCompleted)

		// calculate remaining days
		remainingDays = remainingDays - DaysCompleted

		//fmt.Printf("Thank you %v for taking part and completing %v days.\n", TwitterName, DaysCompleted)
		//fmt.Printf("You have %v days remaining for the %v challenge\n", remainingDays, challenge)
		// fmt.Println("Good luck")

		client, err := getClient(&creds)
		if err != nil {
			log.Println("Error getting Twitter Client, this is expected if you did not supply your Twitter API tokens")
			log.Println(err)
		}

		message := fmt.Sprintf("Hey I am %v I have been doing the %v for %v days and I have %v Days left", TwitterName, challenge, DaysCompleted, remainingDays)
		tweet, resp, err := client.Statuses.Update(message, nil)
		if err != nil {
			log.Println(err)
		}
		log.Printf("%+v\n", resp)
		log.Printf("%+v\n", tweet)
	}

}
