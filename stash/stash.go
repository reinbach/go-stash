package stash

import (
	"errors"
)

var (
	ErrNilClient = errors.New("client is nil")
)

// New creates an instance of the Stash Client
func New(apiUrl, consumerKey, accessToken, tokenSecret string) *Client {
	c := &Client{}
	c.ApiUrl = apiUrl
	c.ConsumerKey = consumerKey
	c.ConsumerSecret = "dont't care"
	c.AccessToken = accessToken
	c.TokenSecret = tokenSecret

	c.Repos = &RepoResource{c}
	c.Users = &UserResource{c}
	c.Branches = &BranchResource{c}
	c.Commits = &CommitResource{c}
	c.Contents = &ContentResource{c}
	return c
}

type Client struct {
	ApiUrl         string
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	TokenSecret    string

	Repos    *RepoResource
	Users    *UserResource
	Branches *BranchResource
	Commits  *CommitResource
	Contents *ContentResource
}

// Guest Client that can be used to access
// public APIs that do not require authentication.
var Guest = New("", "", "", "")
