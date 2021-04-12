package main

import (
	"fmt"
	"sync"
)

/**
1、添加头部节点操作AddHead
2、添加尾部节点操作AppendTail
4、插入任意节点操作Insert
5、删除任意节点操作Remove
8、获取指定位置的节点Get
9、搜索任意节点Search
10、获取链表大小GetSize
11、打印所有节点操作ShowList
12、反转所有节点操作Reverse
*/

// 定义节点
type DoubleNode struct {
	data interface{} // 值
	Prev *DoubleNode // 上一个节点指针
	Next *DoubleNode // 下一个节点指针
}

// 定义链表
type DoubleList struct {
	lock     *sync.RWMutex // 锁
	Capacity uint          // 最大容量
	Size     uint          // 当前容量
	Head     *DoubleNode   // 头结点
	Tail     *DoubleNode   // 尾部节点
}

// 链表初始化
func (list *DoubleList) Init() {
	list.lock = new(sync.RWMutex)
	list.Size = 0
	list.Head = nil
	list.Tail = nil
}

// 1、添加头部节点
//	 判断链表是否为空
func (list *DoubleList) AddHead(node *DoubleNode) bool {
	if list.Capacity == 0 {
		return false // 空链表
	}
	list.lock.Lock() //
	defer list.lock.Unlock()

	if list.Head == nil { // 只有一个节点
		list.Head = node
		list.Tail = node
	} else {
		list.Head.Prev = node // 将旧头部节点
		node.Next = list.Head
	}
	return true

}

// 2、添加尾部节点
func (list *DoubleList) AppendTail(node *DoubleNode) bool {
	// 先判断是否有容量
	if list.Capacity == 0 {
		return false
	}
	list.lock.Lock()
	defer list.lock.Unlock()

	// 判断尾部是否为空
	if list.Tail == nil { // 即节点为空
		list.Tail = node
		list.Head = node
	} else {
		list.Tail.Next = node //旧的尾部下个节点指向新节点
		node.Prev = list.Tail // 追加新节点时,先将node的上节点指向旧的尾部节点
		list.Tail = node      // 设置新的尾部节点
		list.Tail.Next = nil  // 新的尾部节点为空
	}

	list.Size++
	return true

}

// 4、添加任意位置元素
func (list *DoubleList) Insert(index uint, node *DoubleNode) bool {
	if list.Size == list.Capacity { // 容量满了
		return false
	}
	if list.Size == 0 {
		return list.AddHead(node)
	}

	list.lock.Lock()
	defer list.lock.Unlock()

	if index == 0 {
		node.Next = list.Head
		list.Head = node
		list.Head.Prev = nil
		list.Size++
		return true
	}

	nextNode := list.Get(index)
	node.Prev = nextNode.Prev
	node.Next = nextNode
	nextNode.Prev.Next = node
	nextNode.Prev = node
	list.Size++
	return true
}

// 5、删除指定位置节点
func (list *DoubleList) Delete(index uint) bool {
	if index > list.Size-1 {
		return false
	}

	list.lock.Lock()
	defer list.lock.Unlock()

	if index == 0 {
		if list.Size == 1 {
			list.Head = nil
			list.Tail = nil
		} else {
			list.Head.Next.Prev = nil
			list.Head = list.Head.Next
		}
		list.Size--
		return true
	}

	if index == list.Size-1 {
		list.Tail.Prev.Next = nil
		list.Tail = list.Tail.Prev
		list.Size--
		return true
	}

	node := list.Get(index)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	list.Size--
	return true
}

// 打印双向链表信息
func (list *DoubleList) ShowList() {
	if list == nil || list.Size == 0 {
		fmt.Println("this doubleList is empty")
		return
	}

	list.lock.RLock()
	defer list.lock.RUnlock()

	fmt.Printf("this double list size is %d\n", list.Size)

	ptr := list.Head
	for ptr.Next != nil {
		fmt.Printf("data is %v\n", ptr.data)
		ptr = ptr.Next
	}

}

// 打印翻转链表
func (list *DoubleList) Reverse() {
	if list == nil || list.Size == 0 {
		fmt.Println("this doubleList is empty")
		return
	}

	list.lock.RLock()
	defer list.lock.RUnlock()
	fmt.Printf("this double list size is %d\n", list.Size)

	ptr := list.Tail
	for ptr.Prev != nil {
		fmt.Printf("data is %v\n", ptr.data)
		ptr = ptr.Prev
	}
}

// 获取链表长度
func (list *DoubleList) Length() int {
	return list.Length()
}

// 获取指定位置的节点
func (list *DoubleList) Get(index uint) *DoubleNode {
	if list.Size == 0 || index > list.Size-1 {
		return nil
	}

	if index == 0 {
		return list.Head
	}

	node := list.Head
	var i uint
	for i = 1; i <= index; i++ {
		node = node.Next
	}
	return node
}

// 查找指定数据所在的节点
func (list *DoubleList) Search(data interface{}) *DoubleNode {
	if list.Size == 0 {
		return nil
	}

	node := list.Head
	for ; node.Next != nil; node = node.Next {
		if node.data == data {
			break
		}

	}

	return node
}
