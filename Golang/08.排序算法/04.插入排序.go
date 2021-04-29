package main

import "fmt"

/**  插入排序：
基本思想：类似于打牌，将手里的牌从小到大插入排序。
		每一步将一个待排序的数据插入到前面已经排好序的有序序列中，直到插完所有元素为止。

时间复杂度：O(n²)
空间复杂度：O(1)，只需要一个额外空间用于交换
稳定性：插入排序是稳定的排序算法，满足条件nums[j] > nums[j + 1]才发生交换，这个条件可以保证值相等的元素的相对位置不变。

参考文章： https://studygolang.com/articles/13437
		https://blog.csdn.net/lengyuezuixue/article/details/81173737

*/

func InsertSort(data []int) []int {
	length := len(data)
	if length <= 1 {
		return nil
	}

	for i := 1; i < length-1; i++ {
		for j := i; j > 0 && data[j] < data[j-1]; j-- { // 如果 取出的牌小于左边的牌， 则交换位置
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
	return data
}

func main() {
	data := []int{3, 4, 1, 9, 7, 5, 8}
	result := InsertSort(data)
	fmt.Println(result)

}
