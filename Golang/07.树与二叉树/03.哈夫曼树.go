package main

// 基于最小堆实现Huffman树

// 最小堆
type MinHeap struct {
	Size int
	Heap []*HuffmanTree
}

type HuffmanTree struct {
	Left, Right *HuffmanTree
	Wight       int
}
