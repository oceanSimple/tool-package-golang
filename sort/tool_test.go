package mySort

import (
	"log"
	"sort"
	"testing"
)

func TestGetArray(t *testing.T) {
	log.Println(GetArray())
}

func TestShuffle(t *testing.T) {
	arr := GetArray()
	log.Println(Shuffle(arr))
}

func TestCompareArray(t *testing.T) {
	arr1 := GetArray()
	arr2 := Shuffle(GetArray())
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})
	log.Println(CompareArray(arr1, arr2))
}
