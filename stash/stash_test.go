package stash

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/reinbach/go-stash/oauth1"
)

var (
	testURL      string
	testProject  string
	testRepo     string
	testUser     string
	testPassword string
	consumerKey  string
	privateKey   string
	userAccess   string
	userSecret   string
	client       *Client
)

func init() {
	testURL = os.Getenv("STASH_URL")
	testProject = os.Getenv("STASH_PROJECT")
	testRepo = os.Getenv("STASH_REPO")
	testUser = os.Getenv("STASH_USER")
	testPassword = os.Getenv("STASH_PASSWORD")
	consumerKey = os.Getenv("STASH_CONSUMER_KEY")
	privateKey = os.Getenv("STASH_PRIVATE_KEY")

	switch {
	case len(testURL) == 0:
		panic(errors.New("must set the STASH_URL environment variable"))
	case len(testProject) == 0:
		panic(errors.New("must set the STASH_PROJECT environment variable"))
	case len(testRepo) == 0:
		panic(errors.New("must set the STASH_REPO environment variable"))
	case len(testUser) == 0:
		panic(errors.New("must set the STASH_USER environment variable"))
	case len(testPassword) == 0:
		panic(errors.New("must set the STASH_PASSWORD environment variable"))
	case len(consumerKey) == 0:
		panic(errors.New("must set the STASH_CONSUMER_KEY environment variable"))
	case len(privateKey) == 0:
		panic(errors.New("must set the STASH_PRIVATE_KEY environment variable"))
	}

	c, err := GetAuthorizedClient()
	if err != nil {
		fmt.Println("Failed to get client: ", err)
		os.Exit(1)
	}
	client = c
}

func TestGetFullApiUrl(t *testing.T) {
	url := client.GetFullApiUrl("core")
	if url != fmt.Sprintf("%s/rest/api/1.0", testURL) {
		t.Errorf("Core API URL is invalid, got: ", url)
	}
}

func GetAuthorizedClient() (*Client, error) {
	var consumer = oauth1.Consumer{
		RequestTokenURL:       testURL + "/plugins/servlet/oauth/request-token",
		AuthorizationURL:      testURL + "/plugins/servlet/oauth/authorize",
		AccessTokenURL:        testURL + "/plugins/servlet/oauth/access-token",
		CallbackURL:           oauth1.OOB,
		ConsumerKey:           consumerKey,
		ConsumerPrivateKeyPem: privateKey,
	}

	// Generate a Request Token
	requestToken, err := consumer.RequestToken()
	if err != nil {
		fmt.Println("error here: ", err)
		return nil, err
	}

	// TODO need to handle responses from service cleanly
	// need to read response
	// and make request based on response
	// handling user logged in/out

	url, _ := consumer.AuthorizeRedirect(requestToken)
	c := &http.Client{}
	resp, err := c.Get(url)
	if err != nil {
		fmt.Println("Something happened in authorize redirect: ", err)
	}

	// exchange for an access token
	verifier := resp.Header.Get("oauth_verifier")
	accessToken, err := consumer.AuthorizeToken(requestToken, verifier)
	if err != nil {
		return nil, err
	}

	// create the Stash client
	var client = New(
		testURL,
		consumerKey,
		accessToken.Token(),
		accessToken.Secret(),
		privateKey,
	)

	return client, nil
}
