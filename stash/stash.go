package stash

import (
	"errors"
	"fmt"
)

var (
	ErrNilClient = errors.New("client is nil")
)

// New creates an instance of the Stash Client
func New(apiUrl, consumerKey, accessToken, tokenSecret, privateKey string) *Client {
	c := &Client{}
	c.ApiUrl = fmt.Sprintf("%s/rest/api/1.0", apiUrl)
	c.ConsumerKey = consumerKey
	c.ConsumerSecret = "dont't care"
	c.ConsumerPrivateKeyPem = privateKey
	c.AccessToken = accessToken
	c.TokenSecret = tokenSecret

	c.Repos = &RepoResource{c}
	c.Branches = &BranchResource{c}
	c.Commits = &CommitResource{c}
	c.Contents = &ContentResource{c}
	c.Hooks = &HookResource{c}
	return c
}

type Client struct {
	ApiUrl                string
	ConsumerKey           string
	ConsumerSecret        string
	ConsumerPrivateKeyPem string
	AccessToken           string
	TokenSecret           string

	Repos    *RepoResource
	Branches *BranchResource
	Commits  *CommitResource
	Contents *ContentResource
	Hooks    *HookResource
}

// Guest Client that can be used to access
// public APIs that do not require authentication.
var Guest = New("", "", "", "", "")
