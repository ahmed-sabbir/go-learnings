package main

import "fmt"

// Foo is foo
type Foo struct {
	A int
	b string
}

func main() {
	basicOperations()
	fmt.Println("------")
	structsAreValueTypes()
}

func structsAreValueTypes() {
	f := Foo{
		A: 20,
	}

	var f2 Foo
	f2 = f
	f2.A = 100
	fmt.Println("f2.A:", f2.A)
	fmt.Println("f.A", f.A)

	var f3 *Foo = &f
	f3.A = 200
	fmt.Println("f3.A:", f3.A)
	fmt.Println("f.A", f.A)
}

func basicOperations() {
	f := Foo{}
	fmt.Println(f)

	g := Foo{10, "Hello"}
	fmt.Println(g)

	h := Foo{
		b: "goodbye",
	}
	fmt.Println(h)

	h.A = 1000
	fmt.Println(h.A)
}
