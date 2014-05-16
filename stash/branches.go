package stash

import (
	"fmt"
)

type Branch struct {
	ID         string `"json:id"`
	LatestHash string `"json:latestChangeset"`
}

type Branches struct {
	Branch []*Branch `"json:values"`
}

type BranchResource struct {
	client *Client
}

// Get list of branches for repo
func (r *BranchResource) List(apiUrl, project, slug string) (*Branches, error) {
	branches := Branches{}
	path := fmt.Sprintf("/projects/%s/repos/%s/branches", project, slug)

	if err := r.client.do("GET", path, nil, nil, &branches); err != nil {
		return nil, err
	}

	return &branches, nil
}
