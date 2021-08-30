package main

import "errors"

/**
平衡二叉树（Balanced Binary Tree）又被称为AVL树
它是一 棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树
*/

type AVLNode struct {
	Left, Right *AVLNode    // 指向左孩子和右孩子
	Data        interface{} // 该节点存储数据
	height      int         // 该节点的高度
}

// 比较两个节点中 Data 的大小
type comparator func(a, b interface{}) int

// compare 指针
var compare comparator

func NewAVLNode(data interface{}) *AVLNode {
	return &AVLNode{
		Left:   nil,
		Right:  nil,
		Data:   data,
		height: 1,
	}
}

// 新建 AVL 树
func CreatAVLTree(data interface{}, myfunc comparator) (*AVLNode, error) {
	if data == nil && myfunc == nil {
		return nil, errors.New("is empty")
	}
	compare = myfunc
	return NewAVLNode(data), nil
}

// 获取节点数据
func (avlNode *AVLNode) GetData() interface{} {
	return avlNode.Data
}

// 设置节点数据
func (avlNode *AVLNode) SetData(data interface{}) {
	if avlNode == nil {
		return
	}
	avlNode.Data = data

}

// 获取节点的右孩子节点
func (avlNode *AVLNode) GetRight() *AVLNode {
	if avlNode == nil {
		return nil
	}
	return avlNode.Right
}

// 获取节点的左孩子节点

func (avlNode *AVLNode) GetLeft() *AVLNode {
	if avlNode == nil {
		return nil
	}
	return avlNode.Left
}

// 获取节点的高度

func (avlNode *AVLNode) GetHeight() int {
	if avlNode == nil {
		return 0
	}
	return avlNode.height
}

// 比较两个子树的高度
func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// 逆时针旋转 左旋
func (avlNode *AVLNode) LeftRotate() *AVLNode {
	headNode := avlNode.Right
	avlNode.Right = headNode.Left
	headNode.Left = avlNode

	avlNode.height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
	headNode.height = Max(headNode.Left.GetHeight(), headNode.Right.GetHeight()) + 1
	return headNode
}

// 先顺时针旋转再逆时针旋转  先右旋，再左旋

// 查找指定节点
func (avlNode *AVLNode) Find(data interface{}) *AVLNode {
	var finded *AVLNode

	switch compare(data, avlNode.Data) { // compare(a,b)   a > b ,返回 大于0的数; a < b 返回小于 0 的数
	case -1:
		finded = avlNode.Left.Find(data)
	case 1:
		finded = avlNode.Right.Find(data)
	case 0:
		return avlNode
	}
	return finded
}

// 查找最小节点
func (avlNode *AVLNode) FindMin() *AVLNode {
	var finded *AVLNode

	if avlNode.Left != nil {
		finded = avlNode.Left.FindMin()
	} else {
		finded = avlNode
	}
	return finded
}

// 查找最大节点
func (avlNode *AVLNode) FindMax() *AVLNode {
	var finded *AVLNode
	if avlNode.Right != nil {
		finded = avlNode.Right.FindMin()
	} else {
		finded = avlNode
	}
	return finded
}

// 调整 AVL 树的高度
func (avlNode *AVLNode) adjust() *AVLNode {
	if avlNode.Right.GetHeight()-avlNode.Left.GetHeight() == 2 {
		if avlNode.Right.Right.GetHeight() > avlNode.Right.Left.GetHeight() {
			avlNode = avlNode.LeftR
		}
	}
}

// 删除数据
func (avlNode *AVLNode) Delete(value interface{}) *AVLNode {
	if avlNode == nil {
		return nil
	}

	switch compare(value, avlNode.Data) {
	case 1:
		avlNode.Right = avlNode.Right.Delete(value)
	case -1:
		avlNode.Left = avlNode.Left.Delete(value)
	case 0:
		// 找到数据， 删除节点
		if avlNode.Left != nil && avlNode.Right != nil {
			avlNode.Data = avlNode.Right.FindMin().Data
			avlNode.Right = avlNode.Right.Delete(avlNode.Data)
		} else if avlNode.Left != nil {
			avlNode = avlNode.Left
		} else {
			avlNode = avlNode.Right
		}

	}

	if avlNode != nil {
		avlNode.height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
		avlNode = avlNode.adjust()
	}

}

/**
参考文章： https://studygolang.com/articles/21882
*/

/**
参考文章： https://blog.csdn.net/qq_36183935/article/details/80315808
*/
