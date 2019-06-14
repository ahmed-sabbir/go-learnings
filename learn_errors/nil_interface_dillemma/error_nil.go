package main

import "fmt"

func reallyNil() error {
	var e error
	fmt.Println("e is nil:", e == nil)
	return e
}

//MyError is a custom error
type MyError struct {
	A       int
	B       int
	Message string
}

func (me *MyError) Error() string {
	return fmt.Sprintf("this is my custom error %s", me.Message)
}

func notReallyNil() error {
	var me *MyError
	fmt.Println("my error is nil:", me == nil)
	return me
}

func main() {
	e := reallyNil()
	me := notReallyNil()
	fmt.Println("in main, e is nill:", e == nil)
	fmt.Println("in main, me is nil:", me == nil)
}
