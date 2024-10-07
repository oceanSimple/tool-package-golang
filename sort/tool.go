package mySort

import (
	"fmt"
	"math/rand"
	"time"
)

func TestSortFunc(name string, sort func(arr []int)) {
	fmt.Println("--" + name + "--")
	origin := GetArray()
	disorder := Shuffle(GetArray())

	// 计时
	startTime := time.Now()
	sort(disorder)
	endTime := time.Now()

	if CompareArray(origin, disorder) {
		fmt.Println("Success!")
	} else {
		fmt.Println("Error!")
	}

	fmt.Println("Time: ", endTime.Sub(startTime))
}

// 获得一个1-100的数组
func GetArray() []int {
	var arr []int
	for i := 1; i <= 100000; i++ {
		arr = append(arr, i)
	}
	return arr
}

// 将数组随机打乱
func Shuffle(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// 比较数组是否相等
func CompareArray(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
