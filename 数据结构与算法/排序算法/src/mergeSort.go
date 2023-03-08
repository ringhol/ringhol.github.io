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

func MergeSort(arr []int) []int {
	//如果长度为1则不需要排序
	if len(arr) == 1 {
		return arr
	}
	res := make([]int, len(arr))
	//归并的两个数组必须已经有序
	//先对左侧数组进行归并排序确保有序,在对右侧数组进行归并确保有序，然后对这两个数组进行归并
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	//对这两部分进行归并
	i := 0
	j := 0
  k:=0
	for  i<len(left)&&j<len(right)&&k<len(res) {
		if  left[i] < right[j] {
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
    k++
	}
  //处理剩余的数据
  for i<len(left){
    res[k] = left[i];i++;k++
  }
  for j<len(right){
    res[k] = right[j];j++;k++
  }
	return res
}

func main() {
	data := make([]int, 30)
	for i := 0; i < 30; i++ {
		data[i] = rand.Intn(100)
	}
	fmt.Println(data)
	res := MergeSort(data)
	fmt.Println(res)
	fmt.Println(Check(res))
}
