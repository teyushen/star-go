package app

import (
	"testing"
)

func TestReadRepos(t *testing.T) {
	writeToConfig(Repo{"me", "repoName"})
	repos := ReadRepos()

	if len(repos) != 1 {
		t.Errorf("Expected repo size is '1', but got '%v'", len(repos))
		return
	}

	if o := repos[0].Owner; o != "me" {
		t.Errorf("Expected Owner is 'me', but got '%s'", o)
	}
	//if r := repos[0].RepoName; r != "me" {
	//	t.Errorf("Expected RepoName is 'star-go', but got '%s'", r)
	//}


}

