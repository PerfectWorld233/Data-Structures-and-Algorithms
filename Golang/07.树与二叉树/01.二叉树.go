package main

import (
	"fmt"
	"math/rand"
)

/**
二叉树的定义：
1、每个节点最多有两棵子树, 没有子树或者只有一颗子树也是可以的
2、左子树和右子树是有顺序的, 次序不能任意颠倒
即使树中某结点只有一颗树, 也要区分是右树还是左树


二叉树的遍历方式：
1、 前序遍历		访问根节点，遍历左子树， 遍历右子树
2、 中序遍历		遍历左子树， 访问根节点， 遍历右子树
3、 后续遍历		遍历左子树， 遍历右子树， 访问根节点
*/

// 定义二叉树的节点
type Node struct {
	value int
	left  *Node
	right *Node
}

// 打印节点的值
func (node *Node) Show() {
	fmt.Printf("%d", node.value)
}

// 设置节点的值
func (node *Node) SetValue(value int) {
	node.value = value
}

// 创建节点
func CreateNode(value int) *Node {
	return &Node{value, nil, nil}
}

/**
 	二叉树节点的插入:
	把数据插入到指定的位置，这里的数据唯一不重复。
	1.先判断当前结点和插入的值的比较，如果等于，就跳过：
	2.如果数据大于本节点，就插入到右孩子，如果右孩子为nil，直接插入，有值继续比较插入
	3.如果数据小于本节点，就插入到左孩子，如果左孩子为nil，直接插入，有值继续比较插入
		对于有序的二叉树，左子节点小于当前节点，右子节点大于当前节点

*/
func (node *Node) Insert(n *Node) {
	if n.value == node.value {
		return
	}

	if n.value > node.value {
		if node.right == nil {
			node.right = n
		} else {
			node.right.Insert(n)
		}
	} else {
		if node.left == nil {
			node.left = n
		} else {
			node.left.Insert(n)
		}
	}

}

/**
查找节点
1.先判断当前结点是否是需要查找的结点
2.如果不是，则分别往该结点的左孩子和右孩子递归，直到当前结点为 nil 为止。
*/
func (node *Node) FindNode(n *Node, x int) *Node {
	// 目标节点为空
	if n == nil {
		return nil
	}
	// 单前节点就是目标节点
	if n.value == x {
		return n
	} else {
		p := node.FindNode(n.left, x)
		if p != nil {
			return p
		}
		return node.FindNode(n.right, x)
	}

}

// 求树的高度
// 参数： 根节点
// 树的高度= max(左子树高度， 右子树高度) + 1
func (node *Node) GetTreeHigh(n *Node) int {
	if n == nil {
		return 0
	} else {
		lHigh := node.GetTreeHigh(n.left)
		rHigh := node.GetTreeHigh(n.right)
		if lHigh > rHigh {
			return lHigh + 1
		} else {
			return rHigh + 1
		}
	}

}

// 递归前序遍历二叉树
// 参数： 根节点
// 返回值： nil
func (node *Node) PreOrder(n *Node) {
	if n != nil {
		fmt.Printf("%d\n", n.value)
		node.PreOrder(n.left)
		node.PreOrder(n.right)
	}
}

// 递归中序遍历二叉树
// 参数： 根节点
// 返回值： nil
func (node *Node) InOrder(n *Node) {
	if n != nil {
		node.InOrder(n.left)
		fmt.Printf("%d", n.value)
		node.InOrder(n.right)
	}
}

// 递归后序遍历二叉树
func (node *Node) PostOrder(n *Node) {
	if n != nil {
		node.PostOrder(n.left)
		node.PostOrder(n.right)
		fmt.Printf("%d", n.value)
	}
}

// 打印所有的叶子节点
func (node *Node) GetLeafNode(n *Node) {
	if n != nil {
		if n.left == nil && n.right == nil {
			fmt.Printf("%d", n.value)
		}

		node.GetLeafNode(n.left)
		node.GetLeafNode(n.right)
	}
}

// 测试效果啊
func main() {
	root := CreateNode(5)
	//root.left = CreateNode(2)
	//root.right = CreateNode(4)
	//
	//high := root.GetTreeHigh(root)
	//fmt.Println(high)
	//root.PostOrder(root)

	for i := 0; i < 5; i++ {
		n := rand.Intn(50)
		fmt.Println(n)
		fmt.Println("+++++++++++++")
		root.Insert(CreateNode(n))
	}
	//root.Show()
	fmt.Println("================")
	root.PreOrder(root)
	data := root.FindNode(root, 5)
	fmt.Println(data)
}
