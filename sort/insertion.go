package mySort

// 插入排序
func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		current := arr[i] // 当前值
		preIndex := i - 1 // 前一个值的索引
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex] // 后移
			preIndex--
		}
		arr[preIndex+1] = current
	}
}
