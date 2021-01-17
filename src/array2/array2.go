package main

import "fmt"

func main() {
	// arr := [5]int{1, 2, 3, 4, 5}

	// //임시 변수 선언 없이 배열 순서 뒤집기
	// for i := 0; i < len(arr)/2; i++ {
	// 	arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	// }

	// fmt.Println(arr)

	//Radix Sort
	arr := [11]int{0, 5, 4, 9, 1, 2, 8, 3, 6, 4, 5}
	temp := [10]int{}

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		temp[idx]++
	}

	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}

	fmt.Println(arr)
}
