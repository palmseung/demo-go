package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintSungjuk() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) InputSungjuk(class string, grade string) {
	s.class = class
	s.grade = grade
}

func main() {
	var s Student = Student{name: "tucker", age: 23, class: "수학", grade: "A+"}
	s.InputSungjuk("과학", "C")
	s.PrintSungjuk()

	// var a int
	// a = 1

	// Increase(a)
	// fmt.Println(a)
	// IncreaseWithPointer(&a)
	// fmt.Println(a)
}

func Increase(x int) {
	x++
}

func IncreaseWithPointer(x *int) {
	*x = *x + 1
}

// func main() {
// 	var a int
// 	var p *int

// 	a = 1
// 	p = &a

// 	fmt.Println(a, p, *p)
// }
