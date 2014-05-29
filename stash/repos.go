package stash

import (
	"fmt"
)

type Project struct {
	Key string `"json:key"`
}

type Repo struct {
	Name     string `"json:name"`
	Slug     string `"json:slug"`
	Public   bool   `"json:public"`
	Project  *Project
	CloneUrl string `"json:cloneUrl"`
}

type RepoResource struct {
	client *Client
}

// Get the named repository
func (r *RepoResource) Find(project, slug string) (*Repo, error) {
	repo := Repo{}
	path := fmt.Sprintf("/projects/%s/repos/%s", project, slug)

	if err := r.client.do("GET", "core", path, nil, nil, &repo); err != nil {
		return nil, err
	}

	return &repo, nil
}
