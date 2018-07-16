package app

import "testing"

func TestParseToRepo(t *testing.T) {
	str1 := "teyushen/star-go"
	str2 := "teyushen"
	str3 := "teyushen/star-go/star-go"

	repoConfig1, error1 := ParseToRepo(str1)

	if error1 != nil || repoConfig1.Owner != "teyushen" || repoConfig1.RepoNames[0] != "star-go" {
		t.Errorf("Expected owner: 'teyushen', repoName: 'star-go', but got owner:'%s', repoNAme: %s", repoConfig1.Owner, repoConfig1.RepoNames[0])
		return
	}

	_, error2 := ParseToRepo(str2)

	if error2 == nil {
		t.Errorf("something wrong")
		return
	}

	_, error3 := ParseToRepo(str3)

	if error3 == nil {
		t.Errorf("something wrong")
		return
	}

}
