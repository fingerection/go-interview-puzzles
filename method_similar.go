package main

import "fmt"

type OB struct {}

func (ob *OB) foo() {
	fmt.Println("*ob foo")
}

func (ob OB) foo() {
	fmt.Println("ob foo")
}

func main() {
	ob := OB{}
	pob := &ob
	ob.foo()
	pob.foo()
}

// Will the code pass compiling?
