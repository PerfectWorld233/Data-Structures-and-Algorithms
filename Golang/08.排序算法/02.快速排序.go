package main

import "fmt"

/**
快速排序

1、选定一个基准值(任意选,以第一个为例)
2、定义左右指针指向左右两端
3、左指针往右移动，
	如果遇到大于基准值的数就把它和右指针的值调换位置,
  然后左指针不动,右指针开始向左移动,
	如果遇到小于基准值的数就把他和左指针的值调换位置,然后开始移动左指针,以此类推,知道左右指针相遇
4、递归上述过程直到排序结束
*/

func QuickSort(data []int) []int {
	// 长度小于等于1就直接结束
	if len(data) <= 1 {
		return nil
	}
	// 将第一个值选定为基准值
	flag := data[0]
	// 定义左右指针
	left, right := 0, len(data)-1

	for i := 1; i <= right; i++ {
		if data[i] > flag {
			data[i], data[right] = data[right], data[i]
			right--
		} else if data[i] < flag {
			data[i], data[left] = data[left], data[i]
			left++
		}
	}

	QuickSort(data[:left])   // 对左边快排
	QuickSort(data[left+1:]) // 对右边快排

	return data

}

func main() {
	data := []int{4, 2, 1, 6, 6, 9, 5, 8}
	result := QuickSort(data)
	fmt.Println(result)
}
