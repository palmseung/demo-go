package main

import "fmt"

func main() {
	var A [10]int

	for i := 0; i < len(A); i++ {
		A[i] = i * i
	}
	fmt.Println(A)

	s := "hello 월드"

	fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]), ", ")
	}
	fmt.Println()

	B := "hello 월드"
	b := []rune(B)
	fmt.Println(len(b))

	for i := 0; i < len(b); i++ {
		fmt.Print(string(b[i]), ", ")
	}

}
