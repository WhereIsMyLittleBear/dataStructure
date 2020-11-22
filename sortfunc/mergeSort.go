package sortfunc

// MergeSort 归并排序
func MergeSort(data []int) []int {
	length := len(data)
	if length <= 1 {
		return data
	}
	num := length / 2
	left := MergeSort(data[:num])
	right := MergeSort(data[num:])
	return merge(left, right)
}
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
