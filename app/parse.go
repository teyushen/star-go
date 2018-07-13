package app

import (
	"strings"
	"errors"
)

func ParseToRepo(str string) (RepoConfig, error) {
	if strings.ContainsAny(str, "/") && len(strings.Split(str, "/")) == 2 {
		return RepoConfig{strings.Split(str, "/")[0], []string{strings.Split(str, "/")[1]}}, nil
	}
	return RepoConfig{}, errors.New("Can not parse this string: " + str)
}
