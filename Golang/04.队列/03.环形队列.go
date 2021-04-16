package main

import (
	"errors"
	"fmt"
)

/**
数组模拟循环队列
q.front == q.rear   可能是队空， 也可能是对满
所以判断队空与队满可以使用两种方法
1、 使用标记区别队空与队满
2、 少用一个元素  q.rear + 1 = q.front (本代码使用该方法)
*/

type CircleQueue struct {
	maxSize int    // 最大容量
	array   [5]int // 数组
	front   int    // 指向队列首
	rear    int    // 指向队尾
}

func (this *CircleQueue) Push(val int) (err error) {
	// 先判断队列是否满了
	if this.IsFull() {
		return errors.New("queue is full")
	}
	// this.rear 在队列尾部， 不包含最后的元素
	this.array[this.rear] = val
	this.rear = (this.rear + 1) % this.maxSize
	return
}

func (this *CircleQueue) Pop() (value int, err error) {
	if this.IsEmpty() {
		return -1, errors.New("queue is empty")
	}
	// front 是指向队首， 且包含队首元素
	value = this.array[this.front]
	this.front = (this.front + 1) % this.maxSize
	return
}

// 判断是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.rear == this.front
}

// 判断环形队列是否满了
func (this *CircleQueue) IsFull() bool {
	return (this.rear+1)%this.maxSize == this.front
}

// 判断环形队列中有多少个元素
func (this *CircleQueue) Size() int {
	return (this.rear + this.maxSize - this.front) % this.maxSize
}

// 查看队列中的所有元素
func (this *CircleQueue) Show() {
	fmt.Println("环形队列情况如下")
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	tempfront := this.front
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempfront, this.array[tempfront])
		tempfront = (tempfront + 1) % this.maxSize
	}

	fmt.Println()
}

func main() {

	circlequeue := &CircleQueue{
		maxSize: 5,
		front:   0,
		rear:    0,
	}

	circlequeue.Push(1)
	circlequeue.Push(2)
	circlequeue.Push(3)
	circlequeue.Push(4)
	//circlequeue.Pop()
	circlequeue.Show()
	size := circlequeue.Size()
	fmt.Println(size)

}
