package sortfunc

import "fmt"

// QuickSort 入口
func QuickSort(arr []int) []int {
	length := len(arr) - 1
	return Qsort(arr, 0, length)
}

// Qsort 从数列中挑出一个元素，称为 “基准”（pivot）；
//重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之/后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
//递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
func Qsort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		Qsort(arr, left, partitionIndex-1)
		Qsort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left                     //0
	index := pivot + 1                //1
	for i := index; i <= right; i++ { //i=1 privot=0 i=2 privot=0
		if arr[i] < arr[pivot] { //当右边比左边小时
			arr[i], arr[index] = arr[index], arr[i] //交换顺序
			index++
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	return index - 1
}

// 第一种写法
func quickSort(arr []int, left, right int) {
	temp := arr[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && arr[j] >= temp {
			j--
		}
		if j >= p {
			arr[p] = arr[j]
			p = j
		}
		for i <= p && arr[i] <= temp {
			i++
		}
		if i <= p {
			arr[p] = arr[i]
			p = i
		}
	}
	arr[p] = temp
	if p-left > 1 {
		quickSort(arr, left, p-1)
	}
	if right-p > 1 {
		quickSort(arr, p+1, right)
	}
}

// QuickSort1 1
func QuickSort1(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

// Quick2Sort 第二种写法
func Quick2Sort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mid, i := arr[0], 1
	head, tail := 0, len(arr)-1
	for head < tail {
		fmt.Println(arr)
		if arr[i] > mid {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		} else {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		}
	}
	arr[head] = mid
	Quick2Sort(arr[:head])
	Quick2Sort(arr[head+1:])
}

// Quick3Sort 第三种写法
func Quick3Sort(a []int, left int, right int) {
	if left >= right {
		return
	}
	explodeIndex := left
	for i := left + 1; i <= right; i++ {
		if a[left] >= a[i] {
			//分割位定位++
			explodeIndex++
			a[i], a[explodeIndex] = a[explodeIndex], a[i]
		}
	}
	//起始位和分割位
	a[left], a[explodeIndex] = a[explodeIndex], a[left]
	Quick3Sort(a, left, explodeIndex-1)
	Quick3Sort(a, explodeIndex+1, right)

}
