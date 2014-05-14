package stash

import (
	"testing"
)

func Test_Users(t *testing.T) {
	// FIND the currently authenticated user
	curr, err := client.Users.Current()
	if err != nil {
		t.Error(err)
	}

	// verify we get back data
	if curr.Username == "" {
		t.Errorf("No value for Username")
	}
}
