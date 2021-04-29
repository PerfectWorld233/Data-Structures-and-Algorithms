package main

import "fmt"

/**
归并排序：
 首先将数组一份为二，分别为左数组和右数组
若左数组的长度大于1，那么对左数组实施归并排序
若右数组的长度大于1， 那么对右数组实施归并排序
将左右数组进行合并

时间复杂度:  O(n log n)
参考文章： https://www.cnblogs.com/chengxiao/p/6194356.html
	     https://blog.csdn.net/terrygmx/article/details/83381052
*/

func MergeSort(data []int) []int {
	length := len(data)
	if length < 2 {
		return nil
	}
	//  递归拆分
	m := length / 2
	LeftData := data[:m]
	LeftSize := len(LeftData)
	if LeftSize > 1 {
		LeftData = MergeSort(LeftData)
	}

	RightData := data[m:]
	RightSize := len(RightData)
	if RightSize > 1 {
		RightData = MergeSort(RightData)
	}

	fmt.Println(LeftData)
	fmt.Println(RightData)

	//  合并
	tmp := make([]int, 0) // 申请一个新的空间
	i, j := 0, 0

	//  将元素从小到大一次放入 新的切片
	for i < LeftSize && j < RightSize {
		if LeftData[i] < RightData[j] {
			tmp = append(tmp, LeftData[i])
			i++
		} else {
			tmp = append(tmp, RightData[j])
			j++
		}
	}
	if i < LeftSize {
		tmp = append(tmp, LeftData[i:]...)
	}
	if j < RightSize {
		tmp = append(tmp, RightData[j:]...)
	}
	return tmp
}

func main() {
	data := []int{1, 4, 8, 2, 6, 9, 5, 7}

	result := MergeSort(data)
	fmt.Println(result)
}
