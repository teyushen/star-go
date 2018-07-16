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
		return
	}

}

func TestSaveRepos(t *testing.T) {
	repo1 := RepoConfig{"me", []string{"star-go", "star-go1"}}
	repo2 := RepoConfig{"you", []string{"no-star", "no-star1"}}
	SaveRepos([]RepoConfig{repo1, repo2}, 1)
	repos := GetRepos(1)

	if len(repos) != 2 {
		t.Errorf("Expected repo size is '2', but got '%v'", len(repos))
		return
	}

	if o := repos[0].Owner; o != "me" {
		t.Errorf("Expected Owner is 'me', but got '%s'", o)
		return
	}
	if o := repos[1].Owner; o != "you" {
		t.Errorf("Expected Owner is 'you', but got '%s'", o)
		return
	}
	if size := len(repos[0].RepoNames); size != 2 {
		t.Errorf("Expected me repo size is '2', but got '%d'", size)
		return
	}
	if size := len(repos[1].RepoNames); size != 2 {
		t.Errorf("Expected you repo size is '2', but got '%d'", size)
		return
	}

}

func TestAppendRepos(t *testing.T) {
	TestSaveRepos(t)
	repo1 := RepoConfig{"me", []string{"star-go2", "star-go3"}}
	repo2 := RepoConfig{"they", []string{"star1", "star3"}}
	AppendRepos(1, repo1, repo2)
	repos := GetRepos(1)

	if len(repos) != 3 {
		t.Errorf("Expected repo size is '3', but got '%v'", len(repos))
		return
	}

	if size := len(repos[0].RepoNames); size != 4 {
		t.Errorf("Expected RepoName size is 4, but got '%d'", size)
		return
	}

	if o := repos[2].Owner; o != "they" {
		t.Errorf("Expected Owner is 'they', but got '%s'", o)
		return
	}
	if size := len(repos[2].RepoNames); size != 2 {
		t.Errorf("Expected RepoName size is 2, but got '%d'", size)
		return
	}

}

func TestMergeRepo(t *testing.T) {

	existRepos := make([]RepoConfig, 0)
	existRepos = append(existRepos, RepoConfig{"Dennis", []string{"star-go", "oauth2"}})

	repo1 := RepoConfig{"Dennis", []string{"23"}}
	repo2 := RepoConfig{"golang", []string{"go"}}


	if !mergeRepo(existRepos, repo1) {
		t.Error("Expected true, but got false")
		return
	}
	if mergeRepo(existRepos, repo2) {
		t.Error("Expected false, but got true")
		return
	}

}


