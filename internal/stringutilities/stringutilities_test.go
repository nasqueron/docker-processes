package stringutilities

import "testing"

func TestContains (t *testing.T) {
	fruits := []string {"apple", "banana", "cherry"}

	if !Contains(fruits, "apple") {
		t.Errorf("Search was incorrect for apple, actual: false, expected: true")
	}

	if Contains(fruits, "dragonfruit") {
		t.Errorf("Search was incorrect for dragonfruit, actual: true, expected: false")
	}

	if Contains(fruits, "") {
		t.Errorf("Search was incorrect for empty string, actual: true, expected: false")
	}

	if Contains([]string{}, "") {
		t.Errorf("Search was incorrect for empty slice, actual: true, expected: false")
	}
}

func TestGetLongestWord (t *testing.T) {
	fruits := []string {"apple", "banana", "cherry", "dragonfruit"}

	longestWord := GetLongestWord(fruits)
	if longestWord != "dragonfruit" {
		t.Errorf("Search was incorrect, actual: %s, expected: dragonfruit", longestWord)
	}

	longestWord = GetLongestWord([]string{})
	if longestWord != "" {
		t.Errorf("Search was incorrect, actual: %s, expected: empty string", longestWord)
	}
}