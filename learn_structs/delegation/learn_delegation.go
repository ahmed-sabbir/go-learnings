package main

import "fmt"

type Foo struct {
	A int
	b string
}

type Bar struct {
	C string
	F Foo
}

// Bazz has an embedded Foo
type Bazz struct {
	D string
	Foo
}

func main() {
	f := Foo{A: 10, b: "hello"}
	b1 := Bar{C: "Fred", F: f}
	fmt.Println(b1.F.b)

	b2 := Bazz{D: "Nancy", Foo: f}
	fmt.Println(b2.A)

	var f2 Foo = b2.Foo
	fmt.Println(f2)
}
