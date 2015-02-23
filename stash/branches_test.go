package stash

import (
	"testing"
)

func TestBranchesList(t *testing.T) {
	branches, err := client.Branches.List(testProject, testRepo)
	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	if len(branches) != 0 {
		t.Errorf("Expected 0 branches, got: %v", len(branches))
	}
}

func TestBranchesListInvalidRepo(t *testing.T) {
	branches, err := client.Branches.List(testProject, "wrong")
	if err == nil {
		t.Error("Expected an error")
	}
	if branches != nil {
		t.Errorf("Did not expect any branches, got: %v", branches)
	}
}
