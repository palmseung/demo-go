package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	name  string
	class int

	sungjuk Sungjuk
}

type Sungjuk struct {
	name  string
	grade string
}

func (p Person) PrintName() {
	fmt.Print(p.name)
}

func (s Student) ViewSungjuk() {
	fmt.Println(s.sungjuk)
}

func main() {
	var p Person
	p1 := Person{"개똥이", 15}
	p2 := Person{name: "소똥이", age: 21}
	p3 := Person{name: "Carson"}
	p4 := Person{}

	fmt.Println(p, p1, p2, p3, p4)

	p.name = "Smith"
	p.age = 24

	// fmt.Println(p)

	p.PrintName()

	var s Student
	s.name = "철수"
	s.class = 1

	s.sungjuk.name = "수학"
	s.sungjuk.grade = "A"

	s.ViewSungjuk()
}
