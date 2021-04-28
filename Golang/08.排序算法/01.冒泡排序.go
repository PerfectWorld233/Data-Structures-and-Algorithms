package main

import "fmt"

/**
冒泡排序
从小到大排序
*/

func BubbleSort(data []int) []int {
	length := len(data)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data

}

func main() {
	data := []int{1, 3, 4, 9, 6, 8}
	result := BubbleSort(data)
	fmt.Println(result)
}
