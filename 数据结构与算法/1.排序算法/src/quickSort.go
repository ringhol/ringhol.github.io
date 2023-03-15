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

func QuickSort(arr []int, i, j int) {
	if i >= j {
		return
	}
	tmp := arr[rand.Int()%(j-i)+i] //随机选择基准点
	l, r := i-1, j
	for cur := i; cur < r; {
		if arr[cur] < tmp { //小的放在左边
			swap(arr, cur, l+1)
			l++
			cur++
		} else if arr[cur] > tmp { //大的放在右边
			swap(arr, cur, r-1)
			r--
		} else { //相等的在中间不用管
			cur++
		}
	}
	QuickSort(arr, i, l+1)
	QuickSort(arr, r, j)

}

func Check(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func main() {
	data := make([]int, 60)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(100)
	}
	fmt.Println(data)
	QuickSort(data, 0, len(data))
	fmt.Println(data)
	fmt.Println("check: ", Check(data))
}
