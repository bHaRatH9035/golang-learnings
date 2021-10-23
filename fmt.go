package main

import "fmt"

type Obj struct {
	value int
}

func main() {
	fmt.Println("Hello World!")
	fmt.Printf("verbose+: %+v\n", Obj{value: 10})
	fmt.Printf("verbose#: %#v\n", Obj{value: 20`})
}
