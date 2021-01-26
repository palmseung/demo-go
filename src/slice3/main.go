package main

import "fmt"

func RemoveLast(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

func RemoveFirst(a []int) ([]int, int) {
	return a[1:], a[0]
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b := a[4:8]
	b[0] = 1
	b[1] = 2
	fmt.Println(a) // slice 는 a를 잘라서 오는 것이 아니라, 일부분만 가리키기 때문에, b의 값을 바꿔도 a값도 바뀜
	fmt.Println(b)

	c := a[4:] //from idx 4 to the end
	fmt.Println(c)

	d := a[:4] // from the start to idx 4
	fmt.Println(d)

	e := a
	for i := 0; i < 5; i++ {
		var back int
		e, back = RemoveLast(e)
		fmt.Println(back)
	}
	fmt.Println(e)

	f := a
	for i := 0; i < 5; i++ {
		var first int
		f, first = RemoveFirst(f)
		fmt.Println(first)
	}
	fmt.Println(f)
}
