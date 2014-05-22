package stash

import (
	"errors"
	"os"
)

// Instance of the Stash client that we'll use for our unit tests
var client *Client

var (
	// Dummy user that we'll use to run integration tests
	testUser string

	// Dummy repo that we'll use to run integration tests
	testRepo string
)

var (
	// OAuth Consumer Key registered with Stash
	consumerKey string

	// OAuth Consumer Secret registered with Stash
	consumerSecret string

	// A valid access token issues for the `testUser` and `consumerKey`
	accessToken string
	tokenSecret string
	privateKey  string
)

func init() {
	consumerKey = os.Getenv("S_CONSUMER_KEY")
	accessToken = os.Getenv("S_ACCESS_TOKEN")
	tokenSecret = os.Getenv("S_TOKEN_SECRET")
	privateKey = os.Getenv("S_PRIVATE_KEY")
	testUser = os.Getenv("S_USER")
	testRepo = os.Getenv("S_REPO")

	switch {
	case len(consumerKey) == 0:
		panic(errors.New("must set the S_CONSUMER_KEY environment variable"))
	case len(consumerSecret) == 0:
		panic(errors.New("must set the S_CONSUMER_SECRET environment variable"))
	case len(accessToken) == 0:
		panic(errors.New("must set the S_ACCESS_TOKEN environment variable"))
	case len(tokenSecret) == 0:
		panic(errors.New("must set the S_TOKEN_SECRET environment variable"))
	case len(testUser) == 0:
		panic(errors.New("must set the S_USER environment variable"))
	case len(testRepo) == 0:
		panic(errors.New("must set the S_REPO environment variable"))
	}

	client = New("http://api.example.com", consumerSecret, accessToken,
		tokenSecret, privateKey)
}
