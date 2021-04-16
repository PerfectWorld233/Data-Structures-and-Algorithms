package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5}

	fmt.Println("倒序前：", data)

	reverse(data)
	fmt.Println("倒序后：", data)
}

func reverse(a []int) {
	length := len(a)
	for i := 0; i < length/2; i++ {
		a[i], a[length-i-1] = a[length-i-1], a[i]
	}
}
