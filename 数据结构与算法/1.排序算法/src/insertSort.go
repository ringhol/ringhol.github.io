package main

import (
	"fmt"
	"math/rand"
)

func Check(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func Swap(a, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}

func InsertSort(arr []int) {
	n := len(arr)
	var j int
	for i := 1; i < n; i++ {
		cur := arr[i]
		for j = i; j > 0 && cur < arr[j-1]; j-- {
			//向后挪
			arr[j] = arr[j-1]
		}
		//此时j为要插入的位置
		arr[j] = cur
	}
}

func main() {
	data := make([]int, 30)
	for i := 0; i < 30; i++ {
		data[i] = rand.Intn(100)
	}
	fmt.Println(data)
	InsertSort(data)
	fmt.Println(data)
  fmt.Println(Check(data))
}
