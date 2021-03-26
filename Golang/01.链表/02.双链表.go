package main

/**
1、添加头部节点操作AppendHead
2、添加尾部节点操作AppendTail
3、追加尾部节点操作Append
4、插入任意节点操作Insert
5、删除任意节点操作Remove
6、删除头部节点操作RemoveHead
7、删除尾部节点操作RemoveTail
8、获取指定位置的节点Get
9、搜索任意节点Search
10、获取链表大小GetSize
11、打印所有节点操作Print
12、反转所有节点操作Reverse
*/

type DoubleNode struct {
	Key   int
	Value interface{}
	Prev  *DoubleNode
	Next  *DoubleNode
}
