package main

import "fmt"

func main() {
	var arr1 = []int{92, 12, 44, 2, 3, 64, 22, 102, 52}
	QuickSorts(arr1, 0, len(arr1)-1)
	fmt.Println(arr1)
}

func QuickSorts(arr []int, low int, hight int) {
	//递归循环的时候如果再次划分的时候高位值或低位值仅一位，则不继续进行排序换位
	if low < hight {
		//	获取基准坐标值
		pos := Sorts(arr, low, hight)
		QuickSorts(arr, low, pos-1)
		QuickSorts(arr, pos+1, hight)
	}
}

func Sorts(arr []int, low int, hight int) (pos int) {
	//先获取基准值,以low下标为基准
	point := arr[low]
	//判断low 是否小于 hight
	for low < hight {
		//判断高位值是否大于基准值
		for low < hight && arr[hight] >= point {
			//如果值大于等于基准值，那么就是高位值处于基准值右侧，无须移动高位值换位。直接判断次高位值
			hight--
		}
		//如果小于基准值，那么就是高位值应当处于基准值左侧，需要移动高位值到左侧。因为low下标位置的已经被赋予给point
		//所以仅需要将hight的值赋值给low即可
		arr[low] = arr[hight]
		for low < hight && arr[low] <= point {
			low++
		}
		arr[hight] = arr[low]
	}
	//如果low等于hight就会跳出循环
	//此时将取出的基准值放入基准下标
	arr[low] = point
	pos = low
	return pos
}
