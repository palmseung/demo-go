package main

import "fmt"

func main() {
	a := []int{1, 2}
	b := append(a, 3)

	fmt.Printf("%p %p \n", a, b) // %p는 주소를 찍는다.

	c := make([]int, 2, 8)
	d := append(c, 4)

	fmt.Printf("%p %p \n", c, d)

}
