package mySort

func ShellSort(arr []int) {
	n := len(arr)

	step := n / 2
	for ; step >= 1; step /= 2 {
		for i := step; i < n; i++ {
			for j := i - step; j >= 0; j -= step {
				if arr[j] > arr[j+step] {
					arr[j], arr[j+step] = arr[j+step], arr[j]
					continue
				}
				break
			}
		}
	}
}
