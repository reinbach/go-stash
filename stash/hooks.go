package stash

import (
	"fmt"
	"net/url"
)

type Hook struct {
	Enabled bool `"json:enabled"`
}

type HookResource struct {
	client *Client
}

// Enable hook for named repository
func (r *RepoResource) CreateHook(project, slug, hook_key, link string) (*Hook, error) {
	values := url.Values{}
	values.Add("url", link)

	hook := Hook{}
	path := fmt.Sprintf("/projects/%s/repos/%s/settings/hooks/%s/enabled",
		project, slug, hook_key)

	if err := r.client.do("PUT", path, nil, values, &hook); err != nil {
		return nil, err
	}

	return &hook, nil
}
