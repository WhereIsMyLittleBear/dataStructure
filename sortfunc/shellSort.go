package sortfunc

// ShellSort 希尔排序
func ShellSort(arr []int) []int {
	length := len(arr)
	//gap是间隔数，通过不断缩小间隔，为了避免特殊序列无法被排序，所以必须保证最后一次 间隔必须为1，间隔为1的时候只是检查是否排序成功，如果成功了只会比较，不会交换，但是部分特殊序列会在大间隔的时候是有序的，此时只有比较没有交换，而小间隔时交换次数多，时间复杂度高，导致没有真正得到优化
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
