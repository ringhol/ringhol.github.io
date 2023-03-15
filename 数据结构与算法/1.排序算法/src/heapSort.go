package main

import (
	"fmt"
	"math/rand"
)

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func Check(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func heapInsert(arr []int, index int) {
	for index>0&& arr[index] > arr[(index-1)>>1] {
		parentIndex := (index - 1) >> 1
		swap(arr, index, parentIndex)
		index = parentIndex
	}
}

func heapify(arr []int, heapsize int) {
	index := 0
	for sonIndex := index<<1+1; sonIndex < heapsize; index,sonIndex=sonIndex,sonIndex<<1+1 {
		if sonIndex+1 < heapsize && arr[sonIndex] < arr[sonIndex+1] {
			sonIndex++
		}
		if arr[sonIndex] > arr[index] {
			swap(arr, index, sonIndex)
		} else {
			break
		}
	}
}

func HeapSort(arr []int) {
	//首先建立初始大根堆
	size := len(arr)
	heapSize := 0
	for i := 0; i < size; i++ {
		heapInsert(arr, i)
		heapSize++
	}
	//取出最大的数，放在数组最后，原来堆中最后一个数放到根节点，同时重新构建大根堆
	for i := 0; i < size-1; i++ {
		max := arr[0]
		arr[0] = arr[heapSize-1]
		arr[heapSize-1] = max
		heapSize--
		heapify(arr, heapSize)
	}
}

func main() {
	data := make([]int, 30)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(100)
	}
	fmt.Println(data)
	HeapSort(data)
	fmt.Println(data)
	fmt.Println("check: ", Check(data))

}
