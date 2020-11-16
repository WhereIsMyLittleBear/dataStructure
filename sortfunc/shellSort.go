package sortfunc

// ShellSort 希尔排序
func ShellSort(arr []int) []int {
	length := len(arr)
	gap := length / 2
	for gap > 0 {
		for i := gap; i < length; i++ {
			j := i - gap
			temp := arr[i]
			for j >= gap-1 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j = j - gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 2
	}
	return arr
}
