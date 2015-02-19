package stash

import (
	"testing"
)

func Test_Repos(t *testing.T) {
	repo, err := client.Repos.Find(testProject, testRepo)
	if err != nil {
		t.Error(err)
	}

	if repo.Slug != testRepo {
		t.Errorf("repo slug [%v]; want [%v]", repo.Slug, testRepo)
	}
}
