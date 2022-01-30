package advancedsearch

func CompareSingleWord(compare string, comapared string) bool {
	if compare[0] == comapared[0] {
		for j := 0; j < len(compare)-1; j++ {
			if compare[j+1] == comapared[j+1] {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	CompareSingleWord("Ma", "Macbook")
}
