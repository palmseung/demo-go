package main

import "fmt"

func main() {
	var m map[string]string
	m = make(map[string]string)

	m["abc"] = "def"
	fmt.Println(m["abc"])

	m2 := make(map[int]int)
	m2[1] = 1
	m2[2] = 0

	v, ok1 := m2[1]
	v2, ok2 := m2[2]
	v3, ok3 := m2[3]

	fmt.Println(v, ok1, v2, ok2, v3, ok3)

	delete(m2, 2)
	v2, ok2 = m2[2]
	fmt.Println(v, ok1, v2, ok2, v3, ok3)

	m2[2] = 12
	m2[3] = 14

	for key, value := range m2 {
		fmt.Println(key, value)
	}
}
