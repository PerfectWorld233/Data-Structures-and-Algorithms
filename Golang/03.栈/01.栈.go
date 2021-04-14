package main

import "fmt"

// 1、 push 向栈中压入一个元素
// 2、 pop  从栈顶取出一个元素
// 3、 isEmpty 判断栈是否为空
// 4、 length 获取栈中的元素个数
// 5、 top 		查询栈顶元素

type stack struct {
	cache []int // 切片
}

// 入栈
func (sk *stack) push(n int) {
	sk.cache = append(sk.cache, n)
}

// 获取栈中的元素个数
func (sk *stack) length() int {
	return len(sk.cache)
}

func (sk *stack) print() {

	fmt.Printf("%v", *sk)
}

// 出栈
func (sk *stack) pop() int {
	if sk.length() == 0 {
		return 0
	}
	item := sk.cache[sk.length()-1]

	sk.cache = sk.cache[:len(sk.cache)-1]

	return item

}

func main() {

	stack := new(stack)
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)

	stack.print()
}
