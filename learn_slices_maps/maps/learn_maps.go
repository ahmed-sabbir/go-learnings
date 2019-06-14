package main

import "fmt"

func main() {

	basicOperations()
	fmt.Println("----")
	mapLiteralOperations()
	fmt.Println("----")
	mapIsReferenceType()
}

func mapIsReferenceType() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	var m4 map[string]int
	fmt.Println("m4[\"goodbye\"]:", m4["goodbye"])
	m4 = m
	m4["goodbye"] = 400
	fmt.Println("m4[\"goodbye\"]:", m4["goodbye"])
	fmt.Println("m[\"goodbye\"]:", m["goodbye"])

	modMap(m)
	fmt.Println("m[\"cheese\"]:", m["cheese"])
}

func modMap(m map[string]int) {
	m["cheese"] = 20
}
func mapLiteralOperations() {
	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	fmt.Println("b:", m2["b"])
	delete(m2, "b")
	fmt.Println("b:", m2["b"])
}

func basicOperations() {
	m := make(map[string]int)
	m["hello"] = 300
	h := m["hello"]
	fmt.Println("hello in m:", h)
	fmt.Println("a in m", m["a"])

	if v, ok := m["hello"]; ok {
		fmt.Println("hello in m:", v)
	}

	printIfExistsInMap("hello", m)
	printIfExistsInMap("zello", m)

	m["hello"] = 20
	fmt.Println("hello in m:", m["hello"])
}

func printIfExistsInMap(k string, m map[string]int) {
	if v, ok := m[k]; ok {
		fmt.Println(k, "in m:", v)
	} else {
		fmt.Println(k, "not in map")
	}
}
