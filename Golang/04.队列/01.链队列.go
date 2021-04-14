package main

import "fmt"

type QueueNode struct {
	Data interface{}
	Next *QueueNode
}

// 创建链列
func (queue *QueueNode) Create(Data ...interface{}) {
	if queue == nil {
		return
	}
	if len(Data) == 0 {
		return
	}

	for _, v := range Data {
		newNode := new(QueueNode)
		newNode.Data = v

		queue.Next = newNode
		queue = queue.Next
	}

}

// 打印链列
func (queue *QueueNode) Print() {
	if queue == nil {
		return
	}
	for queue != nil {
		if queue.Data != nil {
			fmt.Println(queue.Data, " ")
		}
		queue = queue.Next
	}

	fmt.Println()

}

// 链列个数
func (queue *QueueNode) length() int {
	if queue == nil {
		return -1
	}

	i := 0
	for queue.Next != nil {
		i++
		queue = queue.Next
	}
	return i
}

// 入列

func (queue *QueueNode) Push(Data interface{}) {

	if queue == nil {
		return
	}
	if Data == nil {
		return
	}

	// 找到队列末尾
	for queue.Next != nil {
		queue = queue.Next
	}

	// 创建新节点  将新节点加入队列末尾
	newNode := new(QueueNode)
	newNode.Data = Data

	queue.Next = newNode

}

// 出队
func (queue *QueueNode) Pop() {
	// 队头出列
	if queue == nil {
		return
	}

	// 记录列队第一个节点
	queue.Next = queue.Next.Next

}

func main() {

	queue := new(QueueNode)
	queue.Create(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Print()

	queue.Pop()
	queue.Print()
	fmt.Println(queue.length())

}
