package main

import "fmt"

func main() {
	a, b := func1(2, 3)

	fmt.Println(a, b)
}

func func1(x, y int) (int, int) {

	func2(x, y)
	return y, x
}

func func2(x, y int) {
	fmt.Printf("x = %d, y = %d", x, y)
	fmt.Println()

}
