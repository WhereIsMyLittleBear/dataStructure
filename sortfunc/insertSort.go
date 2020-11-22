package sortfunc

// InsertSort 插入排序
func InsertSort(arr []int) []int {
	length := len(arr)
	for i := 1; i < length; i++ {
		j := i - 1
		temp := arr[i]
		for j >= 0 && arr[j] > temp { //当前一个数比后一个数大时，将后一个数依次向前移
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
	return arr
}

// InsertSort2 第二种
func InsertSort2(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- { //当前一个数比后一个数大时，将后一个数依次向前移
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}
