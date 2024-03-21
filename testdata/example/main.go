package main

import "fmt"

type Query interface { // want "gostyle.ifacenames"
	Do() error
}

type Getter interface {
	GetSomething() //  want "gostyle.getters"
}

func main() {
	fmt.Println("hello")
}
