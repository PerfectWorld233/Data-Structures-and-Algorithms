package main


/**
	堆是具有以下性质的完全二叉树：
	每个结点的值都大于或等于其左右孩子结点的值，称为大顶堆;
	或者每个结点的值都小于或等于其左右孩子结点的值，称为小顶堆

	大顶堆：arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2]
	小顶堆：arr[i] <= arr[2i+1] && arr[i] <= arr[2i+2]

	堆排序时间复杂度： O(nlogn)
	参考文章： https://www.cnblogs.com/chengxiao/p/6129630.html
			https://blog.csdn.net/qq_41478279/article/details/84259510
*/
func heapSort(input []int)  {
	inputLen := len(input)
	if inputLen == 0 {
		return
	}
	for i := 0; i < inputLen; i ++ {

	}
}

func minAjust(input []int)  {
	inputLength := len(input)
	if inputLength <= 1 {
		return
	}
	for i := inputLength / 2 -1 ; i >= 0; i -- {
		if (2 * i + 1 <= inputLength -1 ) && (input[i] >= input[])
	}
}






func main()  {

}