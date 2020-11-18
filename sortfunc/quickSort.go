package sortfunc

import "fmt"

// QuickSort 从数列中挑出一个元素，称为 “基准”（pivot）；
//重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之/后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
//递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mid, i := arr[0], 1 //mid:基准 ;i=1
	left, right := 0, len(arr)-1
	for left < right {
		fmt.Println(arr)
		if arr[i] > mid { //左边的数比基准大时 将左边的数和右边的数交换 因为应该在基准右边 否则
			arr[i], arr[right] = arr[right], arr[i]
			right--
		} else {
			arr[i], arr[left] = arr[left], arr[i]
			left++
			i++
		}
	}
	//arr[left] = mid
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}
