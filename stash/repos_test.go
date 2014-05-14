package stash

import (
	"testing"
)

func Test_Repos(t *testing.T) {
	// FIND the named repo
	repo, err := client.Repos.Find(testUser, testRepo)
	if err != nil {
		t.Error(err)
	}

	if repo.Slug != testRepo {
		t.Errorf("repo slug [%v]; want [%v]", repo.Slug, testRepo)
	}
}
