package main

import "fmt"

func insertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		for i := 1; i < len(arr); i++ {
			backup := arr[i]
			j := i - 1
			for j >= 0 && arr[j] > backup {
				arr[j + 1] = arr[j]
				j--
			}
			arr[j+1] = backup
		}
		return arr
	}
}

func main() {
	arr := []int{10,19,8,3,7,4,6,5,10}
	fmt.Println(insertSort(arr))
}
