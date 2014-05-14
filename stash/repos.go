package stash

import (
	"fmt"
)

type Project struct {
	Key string `"json:key"`
}

type Repo struct {
	Name    string `"json:name"`
	Slug    string `"json:slug"`
	Project *Project
}

type RepoResource struct {
	client *Client
}

// Get the named repository
func (r *RepoResource) Find(project, slug string) (*Repo, error) {
	repo := Repo{}
	path := fmt.Sprintf("/projects/%s/repos/%s", project, slug)

	if err := r.client.do("GET", path, nil, nil, &repo); err != nil {
		return nil, err
	}

	return &repo, nil
}

// Enable hook for named repository
func (r *RepoResource) CreateHook(project, slug, hook string) (*Repo, error) {
	repo := Repo{}
	path := fmt.Sprintf("/projects/%s/repos/%s/settings/hooks/%s/enabled",
		project, slug, hook)

	if err := r.clent.do("GET", path, nil, nil&repo); err != nil {
		return nil, err
	}

	return &repo, nil
}
