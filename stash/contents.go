package stash

import (
	"fmt"
)

type Lines struct {
	Text string `"json:text"`
}

type Content struct {
	Data []*Lines
}

type ContentResource struct {
	client *Client
}

// Get content data for file
func (r *ContentResource) Find(project, slug, path, commitId string) (*Content, error) {
	content := Content{}
	url_path := fmt.Sprintf("/projects/%s/repos/%s/browse/%s?at=%s", project,
		slug, path, commitId)

	if err := r.client.do("GET", url_path, nil, nil, &content); err != nil {
		return nil, err
	}

	return &content, nil
}
