package main

import (
	"fmt"
  "math/rand"
)

func Check(arr []int)bool{
  for i:=0;i<len(arr)-1;i++{
    if arr[i]>arr[i+1]{
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

func selectSort(arr []int) {
	var cur int
	n := len(arr)
	for i := 0; i < n-1; i++ {
		//选择出最小值
		cur = i
		for j := i; j < n; j++ {
			if arr[j] < arr[cur] {
				cur = j
			}
		}
		Swap(&arr[i], &arr[cur])
	}
}

func main() {
	data := make([]int, 30)
	for i := 0; i < 30; i++ {
		data[i] = rand.Intn(100)
	}
	fmt.Println(data)
	selectSort(data)
	fmt.Println(data)
  fmt.Println(Check(data))
}
