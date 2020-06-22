package main

import "fmt"

func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		for i := 0; i < len(arr) - 1; i++ {
			isNeedExchange := false
			for j := 0; j < len(arr) - i - 1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					isNeedExchange = true
				}
			}
			if !isNeedExchange {
				break
			}
			fmt.Println(arr)
		}
		return arr
	}
}

func main() {
	arr := []int {11,9,2,8,3,7,4,6,5,10}
	fmt.Println(arr)
	fmt.Println(BubbleSort(arr))
}
