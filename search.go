package advancedsearch

import (
	"sort"
	"strings"
)

func CompareSingleWord(compare string, compared string) bool {
	query := strings.ToUpper(compare)
	word := strings.ToUpper(compared)

	if query[0] == word[0] {
		for j := 0; j < len(query)-1; j++ {
			if query[j+1] == word[j+1] {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func CompareMultipleWords(compare string, compared []string) (bool, []string) {
	compare = strings.ToUpper(compare)
	sort.Strings(compared)
	finalArray := make([]string, 0, len(compare))
	var res bool = false

	for i := 0; i < len(compared); i++ {
		word := strings.ToUpper(compared[i])

		if compare[0] == word[0] {
			if CompareSingleWord(compare, compared[i]) {
				res = true
				finalArray = append(finalArray, compared[i])
			}
		}
	}

	return res, finalArray
}
