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

func BubbleSort(origin []int) {
	n := len(origin)
	//外层控制次数
	var i, j int
	for i = 0; i < n-1; i++ {
		for j = 0; j < n-1-i; j++ {
			if origin[j] > origin[j+1] {
				Swap(&origin[j], &origin[j+1])
			}
		}
	}
}

func main() {
	data := make([]int, 30)
	for i := 0; i < 30; i++ {
		data[i] = rand.Intn(100)
	}
	fmt.Println(data)
	BubbleSort(data)
	fmt.Println(data)
  fmt.Println(Check(data))
}
