package arrays

import "testing"

func TestContains(t *testing.T) {
	str1 := "teyushen/star-go"
	str2 := "teyushen"


	strArray := make([]string, 0)
	strArray = append(strArray, "teyushen", "Golang")

	Contains(strArray, str2)

	if Contains(strArray, str1) {
		t.Errorf("it should be false")
		return
	}

	if !Contains(strArray, str2) {
		t.Errorf("it should be true")
		return
	}

}