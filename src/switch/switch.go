package main

import "fmt"

func main() {
	x := 33

	switch x {
	case 31:
		fmt.Println("x = 31")
	case 32:
		fmt.Println("x = 32")
	case 33:
		fmt.Println("x = 33")
	}

	// golang에는 while문이 없고..., for문으로 해결...!
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}

		if i == 5 {
			break
		}
		fmt.Println(i)

	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
}
