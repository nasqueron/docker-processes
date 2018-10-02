package stringutilities

import (
	"devcentral.nasqueron.org/source/docker-processes/internal/consoleutilities"
	"strings"
)

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

func PadField(text string, length int) string {
	textLen := len(text)

	if textLen < length {
		return strings.Repeat(" ", length - textLen) + text
	}

	if textLen > length {
		if consoleutilities.IsUTF8() {
			return text[:length-1] + "…"
		}

		return text[:length-1] + "+"
	}

	return text
}
