package advancedsearch

import (
	"sort"
)

func CompareSingleWord(compare string, compared string) bool {
	if compare[0] == compared[0] {
		for j := 0; j < len(compare)-1; j++ {
			if compare[j+1] == compared[j+1] {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func CompareMultipleWords(compare string, compared []string) (bool, []string) {
	sort.Strings(compared)
	finalArray := make([]string, 0, len(compare))
	var res bool = false

	for i := 0; i < len(compared); i++ {
		word := compared[i]

		if compare[0] == word[0] {
			if CompareSingleWord(compare, compared[i]) {
				res = true
				finalArray = append(finalArray, compared[i])
			}
		}
	}

	return res, finalArray
}
