package main

import "fmt"

/**
Trie树，即字典树，又称单词查找树或键树，是一种树形结构，是一种哈希树的变种。典型应用是用于统计和排序大量的字符串（但不仅限于字符串），所以经常被搜索引擎系统用于文本词频统计。
	优点 ：最大限度地减少无谓的字符串比较，查询效率比较高，对于一些范围较小的或者内存不敏感的应用可以使用
		特别适用自动补全类应用

	缺点： trie树比较消耗内存：因为他没一层都保存一个字典表，就算这层的节点只有一个也要有一组表
		 使用的是指针类型，内存不连续对存储不友好，性能打折扣 优点：

	参考文章： https://www.cnblogs.com/gaopeng527/p/6699679.html
*/

// 字典树节点
type TrieNode struct {
	children map[interface{}]*TrieNode
	isEnd    bool
}

// 构造字典树节点
func newTrieNode() *TrieNode {
	return &TrieNode{children: make(map[interface{}]*TrieNode), isEnd: false}
}

// 字典树
type Trie struct {
	root *TrieNode
}

// 构造字典树
func NewTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

// 向字典树中插入一个单词
func (trie *Trie) Insert(word string) {
	node := trie.root
	for i := 0; i < len(word); i++ {
		_, ok := node.children[word[i]]
		if !ok {
			node.children[word[i]] = newTrieNode()
		}
		node = node.children[word[i]]
	}
	node.isEnd = true
}

// 搜索字典树中是否存在指定单词
func (trie *Trie) Search(word string) bool {
	node := trie.root
	for i := 0; i < len(word); i++ {
		_, ok := node.children[word[i]]
		if !ok {
			return false
		}
		node = node.children[word[i]]
	}
	return node.isEnd
}

// 判断字典树中是否存在指定前缀的单词
func (trie *Trie) StartsWith(prefix string) bool {
	node := trie.root
	for i := 0; i < len(prefix); i++ {
		_, ok := node.children[prefix[i]]
		if !ok {
			return false
		}
		node = node.children[prefix[i]]
	}
	return true
}

func main() {

	obj := NewTrie()
	obj.Insert("apple")
	obj.Insert("android")
	obj.Insert("and")
	fmt.Println(obj.Search("an"))
	fmt.Println(obj.StartsWith("an"))

}
