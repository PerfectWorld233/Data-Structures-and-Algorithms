package main

import "fmt"

/**
1、获取链表的长度
2、遍历链表
3、删除指定位置的节点
4、在指定index位置插入数据
5、搜寻data数据
6、拿取LinkList中index位置的元素
7、给linklist增加数据
8、链表翻转
*/

type Node struct { // 定义节点元素
	Data interface{} // 定义数据域
	Next *Node       // 定义地址域 (指向下一个节点) 链域
}

type List struct { // 定义链表
	headNode *Node // 定义头结点
}

type Method interface {
	//IsEmpty()
	//Length()
	//ShowList()

}

// 判断链表是否为空
func (this *List) IsEmpty() bool {
	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

// 获取链表的长度
func (this *List) Length() int {
	cur := this.headNode // 获取指针头结点
	count := 0           // 定义一个计数器，初始化为0
	for cur != nil {
		cur = cur.Next // 地址逐个位移
		count++        // 如果头结点不为空， 则 count ++
	}
	return count
}

// 遍历单链表
func (this *List) ShowList() {
	if !this.IsEmpty() {
		cur := this.headNode
		for {
			fmt.Printf("%s,", cur.Data)
			if cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}
		}
	}
}

// 在头部插入元素
func (this *List) AddElem(data interface{}) *Node {
	node := &Node{Data: data} // 新建节点的值域
	node.Next = this.headNode // 给新建的节点的地址域赋值
	this.headNode = node
	return node
}

// 在尾部插入元素
func (this *List) AppendElem(data interface{}) {
	node := &Node{Data: data} // 创建新元素， 通过 data 参数进行数据域的赋值
	if this.IsEmpty() {       // 判断是否为空节点， 如果为空节点则直接把新元素作为头结点
		this.headNode = node
	} else {
		cur := this.headNode  // 定义变量存储头结点
		for cur.Next != nil { // 通过该节点的地址域是否为空 来判断是否为尾节点
			cur = cur.Next // 链表进行位移, 直到获取尾节点
		}
		cur.Next = node // 将尾节点地址域赋值
	}
}

// 查看链表是否包含某个元素
func (this *List) IsContain(data interface{}) bool {
	cur := this.headNode
	for cur.Next != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}

// 在链表中间某个位置插入元素
func (this *List) InsertElem(index int, data interface{}) {

	if index < 0 { // 如果小于 0,  则在头部添加节点
		this.AddElem(data)
	} else if index > this.Length() { // 如果大于 0, 则在尾部添加节点
		this.AppendElem(data)
	} else {
		pre := this.headNode
		count := 0

		for count < (index - 1) {
			pre = pre.Next
			count++
		}
		// 循环跳出后， pre 指向 index-1 的位置
		node := &Node{Data: data} // 创建新元素
		node.Next = pre.Next      // 新元素的地址域存放的是上一个链表的地址域的值
		pre.Next = node           // 上一个链表的地址域存储的地址指向新元素地址

	}

}

// 删除指定位置的元素
func (this *List) DeleteAtIndex(index int) {
	pre := this.headNode
	if index < 0 {
		this.headNode = pre.Next // 如果位置等于 0； 则删除头结点
	} else if index > this.Length() { // 如果输入的值 大于结点长度，则报错
		fmt.Println("超出链表长度")
		return
	} else {
		count := 0                                  //
		for count != (index-1) && pre.Next != nil { // 如果 index = 1 （首结点）， 则跳出循环，直接删除； 如果 index > 1 ; 则开始遍历，直到末尾结点
			count++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}

}

// 删除指定值的 元素
func (this *List) DeleteElem(data interface{}) {
	pre := this.headNode  // 定义变量存储头结点
	if pre.Data == data { // 如果头结点的值 就是要删除的值 , 则直接 删除， 结点地址域位移
		this.headNode = pre.Next
	}
	for pre.Next != nil {
		if pre.Next.Data == data {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next // 结点位移
		}
	}
}

// 链表翻转
func (this *List) ReverseList() {

	var pre *Node        // 声明 一个空节点
	cur := this.headNode // 将当前链表赋值给 cur
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
		this.headNode = pre
	}
}

// 测试
func main() {
	list := List{}
	list.AppendElem("1")
	list.AppendElem("1")
	list.AppendElem("2")
	list.AppendElem("3")
	list.AppendElem("5")
	list.AppendElem("8")
	list.AppendElem("13")
	list.AppendElem("21") // 尾部添加节点
	//list.ShowList()				// 查看链表
	//list.DeleteAtIndex(4)			// 删除指定位置的元素
	//list.DeleteElem("5")			// 删除指定值的元素
	//list.AddElem("50")			// 头部添加节点
	//Bool := list.IsContain("10")  // 判断链表是否包含 某个元素
	//Bool := list.IsEmpty()		// 判断链表是否为空
	//length := list.Length()		// 获取链表长度
	//fmt.Println(length)
	//fmt.Println(Bool)

	list.ReverseList() // 链表翻转

	list.ShowList() // 查看链表所有元素

}
