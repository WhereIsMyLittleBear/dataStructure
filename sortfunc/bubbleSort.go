package sortfunc

// BubbleSort 冒泡排序
func BubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] { //如果前面的比后一个数大 则交换相邻元素
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
