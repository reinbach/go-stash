package stash

import (
	"encoding/json"
	"fmt"
)

type HookDetail struct {
	Key           string `"json:key"`
	Name          string `"json:name"`
	Type          string `"json:type"`
	Description   string `"json:description"`
	Version       string `"json:version"`
	ConfigFormKey string `"json:configFormKey"`
}

type Hook struct {
	Enabled bool        `"json:enabled"`
	Details *HookDetail `"json:details"`
}

type HookResource struct {
	client *Client
}

// Enable hook for named repository
func (r *RepoResource) CreateHook(project, slug, hook_key, link string) (*Hook, error) {
	hookConfig := map[string]string{"url": link}
	values, err := json.Marshal(hookConfig)
	if err != nil {
		return nil, err
	}

	hook := Hook{}
	enablePath := fmt.Sprintf("/projects/%s/repos/%s/settings/hooks/%s/enabled",
		project, slug, hook_key)

	// Enable hook
	if err := r.client.do("PUT", "core", enablePath, nil, nil, &hook); err != nil {
		return nil, err
	}

	// Set hook
	updatePath := fmt.Sprintf("/projects/%s/repos/%s/settings/hooks/%s/settings",
		project, slug, hook_key)
	if err := r.client.do("PUT", "core", updatePath, nil, values, &hook); err != nil {
		return nil, err
	}

	return &hook, nil
}
