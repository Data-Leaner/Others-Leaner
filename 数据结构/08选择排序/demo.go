package main

import "fmt"

// 依次挑选最大的值
func SelectSortMax(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length - 1; i++ {
			// 标记最小值索引
			min := i
			for j := i + 1; j < length; j++ {
				if arr[min] < arr[j] {
					min = j
				}
			}
			// 数据交换
			if i != min {
				arr[i], arr[min] = arr[min], arr[i]
			}
			fmt.Println(arr)
		}
		return arr
	}
}

func main() {
	arr := []int{1,9,2,8,3,7,4,6,5,10}
	SelectSortMax(arr)
}
