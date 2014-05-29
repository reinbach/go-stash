package stash

import (
	"fmt"
)

type Branch struct {
	ID         string `"json:id"`
	LatestHash string `"json:latestChangeset"`
}

type Branches struct {
	Branches []*Branch `"json:values"`
}

type BranchResource struct {
	client *Client
}

// Get list of branches for repo
func (r *BranchResource) List(project, slug string) ([]*Branch, error) {
	branches := Branches{}
	path := fmt.Sprintf("/projects/%s/repos/%s/branches", project, slug)

	if err := r.client.do("GET", "core", path, nil, nil, &branches); err != nil {
		return nil, err
	}

	return branches.Branches, nil
}
