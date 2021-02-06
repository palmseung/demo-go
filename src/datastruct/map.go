package main

import "fmt"

func main() {
	m := CreateMap()
	m.Add("AAA", "11111")
	m.Add("BBB", "12342")
	m.Add("CCC", "23441")

	fmt.Println("AAA = ", m.Get("AAA"))
	fmt.Println("BBB = ", m.Get("BBB"))
	fmt.Println("CCC = ", m.Get("CCC"))
}

func Hash(s string) int {
	h := 0
	A := 256
	B := 3571
	for i := 0; i < len(s); i++ {
		h = (h*A + int(s[i])) % B
	}

	return h
}

type KeyValue struct {
	key   string
	value string
}

type Map struct {
	keyArray [3571][]KeyValue
}

func (m *Map) Add(key, value string) {
	h := Hash(key)
	m.keyArray[h] = append(m.keyArray[h], KeyValue{key, value})
}

func (m *Map) Get(key string) string {
	h := Hash(key)
	for i := 0; i < len(m.keyArray[h]); i++ {
		if m.keyArray[h][i].value == key {
			return m.keyArray[h][i].value
		}
	}
	return ""
}

func CreateMap() *Map {
	return &Map{}
}
