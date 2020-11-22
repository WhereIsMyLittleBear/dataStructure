package sortfunc

// SelectSort 选择排序
func SelectSort(arr []int) []int {
	/* 选择排序：首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。 */
	length := len(arr)
	var minIndex int
	for i := 0; i < length-1; i++ {
		minIndex = i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] { // 寻找最小的数
				minIndex = j // 将最小数的索引保存
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i] //把最小的数放在最前面
	}
	return arr
}
