package stash

import (
	"fmt"
	"testing"
)

var (
	link = "link"
)

func GetHook() string {
	return fmt.Sprintf("%s?owner=%s&name=%s&branch=${refChange.name}&hash=${refChange.toHash}&message=${refChange.type}&author=", testURL, testProject, testRepo)
}

func TestCreateHook(t *testing.T) {
	hook, err := client.Hooks.CreateHook(testProject, testRepo, hookKey, GetHook())
	if err != nil {
		t.Errorf("Unexpected error on `client.Hooks.CreateHook()`, got %v", err)
	}
	if hook.Enabled != true {
		t.Error("Expected hook to be enabled")
	}
}

func TestDeleteHook(t *testing.T) {
	err := client.Hooks.DeleteHook(testProject, testRepo, hookKey, GetHook())
	if err != nil {
		t.Errorf("Unexpected error on `client.Hooks.DeleteHook()`, got %v", err)
	}
}
