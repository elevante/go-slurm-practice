package main

import "fmt"

func foo() {
	panic("Паника в foo()")
}

func main() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	fmt.Println("Start")
	foo()
	fmt.Println("End")
}
