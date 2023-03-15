package main

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)>>1] {
		parentIndex := (index - 1) >> 1
		swap(arr, index, parentIndex)
		index = parentIndex
	}
}

func heapify(arr []int, heapsize int) {
	index := 0
  for sonIndex:=0; sonIndex < heapsize;index = sonIndex {
		sonIndex := index<<1 + 1
		if sonIndex +1< heapsize && arr[sonIndex] < arr[sonIndex+1] {
			sonIndex++
		}
		if arr[sonIndex] > arr[index] {
			swap(arr, index, sonIndex)
		} else {
			break
		}
	}
}
