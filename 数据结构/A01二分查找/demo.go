package main

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		splitdata := arr[0]
		// 比我小
		low := make([]int, 0, 0)
		// 比我大
		high := make([]int, 0, 0)
		// 与我一样大
		mid := make([]int, 0, 0)
		mid = append(mid, splitdata)
		for i := 1; i < len(arr); i++ {
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high)
		myarr := append(append(mid, low...), high...)
		return myarr
	}
}

func main() {
	arr := []int{4,19,111,237,6,5,10,11,223}
	fmt.Println(QuickSort(arr))
}
