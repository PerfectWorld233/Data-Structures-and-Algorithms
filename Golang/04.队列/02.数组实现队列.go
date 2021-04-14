package main

import (
	"errors"
	"fmt"
	"os"
)

/**
题目描述：
实现一个队列，使其具有入队列，出队列，查看队列元素等功能。
队列介绍：
队列是一个有序列表，可以用数组或是链表来实现。
遵循先入先出的原则，即先存入队列的数据，要先取出。后存入要后取出。
*/

type Queue struct {
	maxSize int
	array   [5]int // 数组模拟队列
	front   int    // 指向队列的首部
	rear    int    // 指向队列的尾部
}

// 添加数据到队列
func (this *Queue) AddQueue(value int) (err error) {
	// 判断队列是否已经满了
	if this.rear == this.maxSize-1 {
		return errors.New("queue full")
	}

	this.rear++
	this.array[this.rear] = value
	return
}

// 从队列中取出数据

func (this *Queue) GetQueue() (value int, err error) {
	// 先判断队列是否为空
	if this.rear == this.front {
		return -1, errors.New("队列为空")
	}

	this.front++
	value = this.array[this.front]
	return value, err
}

// 显示队列，找到队首， 遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("queue is ")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}

}

func main() {
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	var key string
	var value int

	for {
		fmt.Println("1.输入 add 表示添加数据到队列")
		fmt.Println("2.输入 get 表示从队列获取数据")
		fmt.Println("3.输入 show 表示显示队列")
		fmt.Println("4.输入 exit 表示结束队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要输入的队列数")
			fmt.Scanln(&value)
			err := queue.AddQueue(value)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列 OK ")
			}

		case "get":
			value, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了", value)
			}
			fmt.Println("get")

		case "show":
			queue.ShowQueue()

		case "exit":
			os.Exit(0)

		}
	}

}
