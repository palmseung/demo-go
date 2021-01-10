package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {

		for j := 1; j <= 9; j++ {
			if i == 3 && j == 2 {
				continue
			}

			fmt.Printf("%v X %v = %v \n", i, j, i*j)
		}
		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
	fmt.Println()

	for i := 4; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	for i := 1; i <= 5; i++ {
		if i < 3 {
			for j := 0; j < i; j++ {
				fmt.Printf("*")
			}
		}

		if i >= 3 {
			for j := i; j <= 5; j++ {
				fmt.Printf("*")
			}
		}

		fmt.Println()
	}

	for i := 1; i <= 7; i++ {
		if i <= 4 {
			for j := 4; j > i; j-- {
				fmt.Print(" ")
			}

			for j := 1; j <= 2*i-1; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}

		if i >= 5 {
			for j := 5; j <= i; j++ {
				fmt.Print(" ")
			}

			for j := 1; j <= -2*i+15; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}

	}
}
