package advancedsearch

import (
	"testing"
)

func TestCompareSingleWord(t *testing.T) {
	result := CompareSingleWord("mac", "Macbook")
	waited := true

	if result != waited {
		t.Errorf("Waited '%t', Result '%t'", waited, result)
	}
}

func TestCompareMultipleWords(t *testing.T) {
	arrayToPass := []string{"Macbook", "Mouse", "Macaco", "Balde", "Copo", "Macarofe"}
	result, _ := CompareMultipleWords("maca", arrayToPass)
	waited := true

	if result != waited {
		t.Errorf("Waited '%t', Result '%t'", waited, result)
	}
}
