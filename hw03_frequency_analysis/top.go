package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var clearWordRegex = regexp.MustCompile(`(\w+[\.\,\!\?]*\w+)|[а-яА-Яa-zA-Z-]{2,}|[а-яА-Яa-zA-Z]`)

func Top10(s string) []string {
	// Place your code here.
	StrFields := clearWordRegex.FindAllString(s, -1)
	StrMap := make(map[string]int)

	for _, word := range StrFields {
		word = strings.ToLower(word)
		StrMap[word]++
	}

	sortedKeys := make([]string, 0, len(StrMap))

	for key := range StrMap {
		sortedKeys = append(sortedKeys, key)
	}

	sort.Slice(sortedKeys, func(i, j int) bool {
		if StrMap[sortedKeys[i]] != StrMap[sortedKeys[j]] {
			return StrMap[sortedKeys[i]] > StrMap[sortedKeys[j]]
		}
		return sortedKeys[i] < sortedKeys[j]
	})

	if 10 < len(sortedKeys) {
		sortedKeys = sortedKeys[:10]
	}

	return sortedKeys
}
