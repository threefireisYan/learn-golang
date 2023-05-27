package main

import "fmt"

func main() {
	var arr = []int{92, 12, 44, 2, 3, 64, 22, 102, 52}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func QuickSort(arr []int, low int, hight int) {
	if low < hight {
		pos := Sort(arr, low, hight)
		QuickSort(arr, low, pos-1)
		QuickSort(arr, pos+1, hight)
	}
}

func Sort(arr []int, low int, hight int) (pos int) {

	point := arr[low]
	for low < hight {
		for low < hight && arr[hight] >= point {
			hight--
		}
		arr[low] = arr[hight]
		for low < hight && arr[low] <= point {
			low++
		}
		arr[hight] = arr[low]
	}

	pos = low
	arr[low] = point
	return pos
}
