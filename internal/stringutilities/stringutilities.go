package stringutilities

func Contains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}

func GetLongestWord(haystack []string) string {
	maxLen := 0
	longestWord := ""

	for _, word := range haystack {
		wordLen := len(word)
		if wordLen > maxLen {
			longestWord = word
			maxLen = wordLen
		}
	}

	return longestWord
}

