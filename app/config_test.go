package app

import (
	"testing"
	"os"
)

func TestMain(m *testing.M) {
	retCode := m.Run()
	os.Exit(retCode)
}

func TestSaveUser(t *testing.T) {
	SaveUser(User{"123456"})
	user := GetUser()

	if user.Token != "123456" {
		t.Errorf("Expected User Token is '123456', but got '%s'", user.Token)
	}

}

func TestSaveRepos(t *testing.T) {
	SaveRepos([]Repo{Repo{"me", "star-go"}})
	repos := GetRepos()

	if len(repos) != 1 {
		t.Errorf("Expected repo size is '1', but got '%v'", len(repos))
		return
	}

	if o := repos[0].Owner; o != "me" {
		t.Errorf("Expected Owner is 'me', but got '%s'", o)
	}
	if r := repos[0].RepoName; r != "star-go" {
		t.Errorf("Expected RepoName is 'star-go', but got '%s'", r)
	}

}

func TestSaveRepo(t *testing.T) {
	SaveRepo(Repo{"me", "star-go"})
	repos := GetRepos()

	if len(repos) != 1 {
		t.Errorf("Expected repo size is '1', but got '%v'", len(repos))
		return
	}

	if o := repos[0].Owner; o != "me" {
		t.Errorf("Expected Owner is 'me', but got '%s'", o)
	}
	if r := repos[0].RepoName; r != "star-go" {
		t.Errorf("Expected RepoName is 'star-go', but got '%s'", r)
	}

}

func TestAppendRepo(t *testing.T) {
	SaveRepo(Repo{"me", "star-go"})
	AppendRepo(Repo{"you", "go-star"})
	repos := GetRepos()

	if len(repos) != 2 {
		t.Errorf("Expected repo size is '2', but got '%v'", len(repos))
		return
	}

	if o := repos[1].Owner; o != "you" {
		t.Errorf("Expected Owner is 'you', but got '%s'", o)
	}
	if r := repos[1].RepoName; r != "go-star" {
		t.Errorf("Expected RepoName is 'go-star', but got '%s'", r)
	}

}

func TestAppendRepos(t *testing.T) {
	SaveRepo(Repo{"me", "star-go"})
	AppendRepos([]Repo{Repo{"you", "go-star"}, Repo{"they", "no-star"}})
	repos := GetRepos()

	if len(repos) != 3 {
		t.Errorf("Expected repo size is '3', but got '%v'", len(repos))
		return
	}

	if o := repos[2].Owner; o != "they" {
		t.Errorf("Expected Owner is 'they', but got '%s'", o)
	}
	if r := repos[2].RepoName; r != "no-star" {
		t.Errorf("Expected RepoName is 'no-star', but got '%s'", r)
	}

}

