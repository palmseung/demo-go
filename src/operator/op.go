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

	var c bool
	c = 3 > 4

	fmt.Println(" 3 > 4", c)

	var d bool
	d = 3 < 4

	var e bool
	e = 23 > 2

	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("result = ", d && e)

	if 3 > 4 {
		fmt.Println("거짓")
	}
	fmt.Println("참")

	h := 3
	if h == 3 {
		fmt.Println("참")
	}

}
