package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/reinbach/go-stash/oauth1"
	"github.com/reinbach/go-stash/stash"
)

var (
	testURL     string
	consumerKey string
	privateKey  string
)

func init() {
	testURL = os.Getenv("STASH_URL")
	consumerKey = os.Getenv("STASH_CONSUMER_KEY")
	privateKey = os.Getenv("STASH_PRIVATE_KEY")

	switch {
	case len(testURL) == 0:
		panic(errors.New("must set the STASH_URL environment variable"))
	case len(consumerKey) == 0:
		panic(errors.New("must set the STASH_CONSUMER_KEY environment variable"))
	case len(privateKey) == 0:
		panic(errors.New("must set the STASH_PRIVATE_KEY environment variable"))
	}
}

func GetAuthorized() (*stash.Client, error) {
	var consumer = oauth1.Consumer{
		RequestTokenURL:       testURL + "/plugins/servlet/oauth/request-token",
		AuthorizationURL:      testURL + "/plugins/servlet/oauth/authorize",
		AccessTokenURL:        testURL + "/plugins/servlet/oauth/access-token",
		CallbackURL:           oauth1.OOB,
		ConsumerKey:           consumerKey,
		ConsumerPrivateKeyPem: privateKey,
	}

	// Step 1: Generate a Request Token. This is a temporary token that is
	// used for having the user authorize an access token and to sign the
	// request to obtain said access token.
	requestToken, err := consumer.RequestToken()
	if err != nil {
		return nil, err
	}

	fmt.Println("\nRequest Token")
	fmt.Println(" - oauth_token: ", requestToken.Token())
	fmt.Println(" - oauth_token_secret: ", requestToken.Secret())

	// Step 2: Redirect to the provider. Since this is a CLI script we do not
	// redirect. In a web application you would redirect the user to the URL
	// below.
	uri, _ := consumer.AuthorizeRedirect(requestToken)

	fmt.Println("\nGo to the following link in your browser:")
	fmt.Println(uri)

	// scan for user input of response
	var verifier string
	fmt.Println("Are you done (Enter verification code)?")
	fmt.Scan(&verifier)

	// Step 3: Once the consumer has redirected the user back to the
	// oauth_callback URL you can request the access token the user has
	// approved. You use the request token to sign this request. After this
	// is done you throw away the request token and use the access token
	// returned. You should store this access token somewhere safe, like a
	// database, for future use.
	accessToken, err := consumer.AuthorizeToken(requestToken, verifier)
	if err != nil {
		return nil, err
	}

	fmt.Println("\nAccess Token")
	fmt.Println(" - oauth_token: ", accessToken.Token())
	fmt.Println(" - oauth_token_secret: ", accessToken.Secret())
	fmt.Println("You may now access protected resources using the access tokens above\n")

	// create the Stash client
	var client = stash.New(
		testURL,
		consumerKey,
		accessToken.Token(),
		accessToken.Secret(),
		privateKey,
	)

	return client, nil
}

func main() {
	if _, err := GetAuthorized(); err != nil {
		fmt.Println("Error: ", err)
	}
}
