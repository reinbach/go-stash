package stash

import (
	"fmt"
)

type Author struct {
	Name  string `"json:name"`
	Email string `"json:emailAddress"`
}

type Commit struct {
	Hash      string `"json:id"`
	Author    *Author
	Timestamp string `"json:authorTimestamp"`
	Message   string `"json:message"`
}

type CommitResource struct {
	client *Client
}

//Get commit data for commit hash
func (r *CommitResource) Get(project, slug, commitId string) (*Commit, error) {
	commit := Commit{}
	path := fmt.Sprintf("/projects/%s/repos/%s/commits/%s", project, slug,
		commitId)

	if err := r.client.do("GET", "core", path, nil, nil, &commit); err != nil {
		return nil, err
	}

	return &commit, nil
}
