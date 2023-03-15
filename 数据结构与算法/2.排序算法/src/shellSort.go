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

func ShellSort(arr []int) {
  //使用Knuth序列确定间隔
	h := 1
	for h <= len(arr)/3 {
		h = 3*h + 1
	}
	var j int
  for gap:=h; gap >= 1;gap = (gap -1)/3 {
		for i := gap; i < len(arr); i++ {
			cur := arr[i]
			for j = i; j >= gap && cur < arr[j-gap]; j -= gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = cur
		}
	}
}

func main() {
	data := make([]int, 30)
	for i := 0; i < 30; i++ {
		data[i] = rand.Intn(100)
	}
	fmt.Println(data)
	ShellSort(data)
	fmt.Println(data)
  fmt.Println(Check(data))
}
