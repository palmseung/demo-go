package main

import "fmt"

func main() {
	a := 4
	b := 2

	fmt.Printf("a&b = %v\n", a&b)
	fmt.Printf("a|b = %v\n", a|b)
	fmt.Println("result = ", a^b) // Println 에서는 %v 사용x
	fmt.Println("^4 = ", ^a)
	fmt.Println("5%2 = ", 5%2)

	fmt.Println("a << 1 = ", a<<1)
	fmt.Println("a >> 1 = ", a>>1)
}
