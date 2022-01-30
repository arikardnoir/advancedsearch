package advancedsearch

import "testing"

func TestCompareWords(t *testing.T) {
	resultado := CompareSingleWord("Mac", "Macbook")
	esperado := true

	if resultado != esperado {
		t.Errorf("esperado '%t', resultado '%t'", esperado, resultado)
	}
}
