package stash

import (
	"fmt"
	"net/url"
)

type Project struct {
	Key string `"json:key"`
}

type Repo struct {
	Name    string   `"json:name"`
	Slug    string   `"json:slug"`
	Public  bool     `"json:public"`
	ScmId   string   `"json:scmId"`
	Project *Project `"json:project"`
}

type Response struct {
	Values        []Repo `"json:values"`
	Size          int    `"json:size"`
	NextPageStart int    `"json:nextPageStart"`
	IsLastPage    bool   `"json:isLastPage"`
}

type RepoResource struct {
	client *Client
}

// Get list of repositories
func (r *RepoResource) List() ([]Repo, error) {
	repos := []Repo{}
	response := Response{}
	path := "/repos"
	params := url.Values{}

	for {
		if response.NextPageStart != 0 {
			params.Set("start", fmt.Sprintf("%d", response.NextPageStart))
		}

		if err := r.client.do("GET", "core", path, params, nil, &response); err != nil {
			return nil, err
		}

		repos = append(repos, response.Values...)

		if response.IsLastPage {
			break
		}
	}

	return repos, nil
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
