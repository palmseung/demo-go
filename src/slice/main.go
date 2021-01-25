package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 0, 5) // 0-> length, 5 -> capacity

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))
	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("cap(b) = %d\n", cap(b))

	b = append(b, 1) // b에다가 1을 추가한 slice를 다시 b에 대입한다.

	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("cap(b) = %d\n", cap(b))
}
