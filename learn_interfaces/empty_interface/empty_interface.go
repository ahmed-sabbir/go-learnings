package main

import "fmt"

func main() {
	emptyInterfaceAssertionOps()
	emptyInterfaceSwitchOps()
}

func emptyInterfaceAssertionOps() {
	var i interface{}
	i = "Hello"
	j := i.(string)
	k, ok := i.(int)
	fmt.Println(i, j, k, ok)

	//m:=i.(int)
	//fmt.Println(m)
}

func emptyInterfaceSwitchOps() {
	doStuff(10)
	doStuff("Hello")
	doStuff(true)
}

func doStuff(i interface{}) {
	switch i := i.(type) {
	case int:
		fmt.Println("Double of i is:", i+i)
	case string:
		fmt.Println("i is", len(i))
	default:
		fmt.Println("don't want to live anymore")
	}
}
