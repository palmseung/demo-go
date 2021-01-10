package main

import "fmt"

// func main() {
// 	a, b := func1(2, 3)

// 	fmt.Println(a, b)
// }

// func func1(x, y int) (int, int) {

// 	func2(x, y)
// 	return y, x
// }

// func func2(x, y int) {
// 	fmt.Printf("x = %d, y = %d", x, y)
// 	fmt.Println()

// }

// func main() {
// 	f1(10)
// }

// func f1(x int) {
// 	if x == 0 {
// 		return
// 	}

// 	fmt.Println(x)
// 	f1(x - 1)
// }

func main() {
	x := sum(10, 0)

	fmt.Println(x)
}

func sum(x int, s int) int {
	if x == 0 {
		return s
	}

	s += x
	return sum(x-1, s)
}
