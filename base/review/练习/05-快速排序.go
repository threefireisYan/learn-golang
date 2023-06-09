package main

import "fmt"

func main() {
	var arr1 = []int{23, 78, 3, 12, 6, 72, 8, 2, 3, 46}

	SortQuick(arr1, 0, len(arr1)-1)

	fmt.Println(arr1)

}

// 普通快速排序
func SortQuick(array []int, begin, end int) {
	if begin < end {
		// 进行切分
		loc := partition(array, begin, end)
		// 对左部分进行快排
		SortQuick(array, begin, loc-1)
		// 对右部分进行快排
		SortQuick(array, loc+1, end)
	}
}

// 切分函数，并返回切分元素的下标
func partition(array []int, begin, end int) int {
	low := begin + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
	hight := end     // array[end]是数组的最后一位

	// 没重合之前
	for low < hight {
		if array[low] > array[begin] {
			array[low], array[hight] = array[hight], array[low] // 交换
			hight--
		} else {
			low++
		}
	}

	/* 跳出while循环后，i = j。
	 * 此时数组被分割成两个部分  -->  array[begin+1] ~ array[i-1] < array[begin]
	 *                        -->  array[i+1] ~ array[end] > array[begin]
	 * 这个时候将数组array分成两个部分，再将array[i]与array[begin]进行比较，决定array[i]的位置。
	 * 最后将array[i]与array[begin]交换，进行两个分割部分的排序！以此类推，直到最后i = j不满足条件就退出！
	 */
	if array[low] >= array[begin] { // 这里必须要取等“>=”，否则数组元素由相同的值组成时，会出现错误！
		low--
	}

	array[begin], array[low] = array[low], array[begin]
	return low
}
