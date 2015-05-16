package main

import (
	"fmt"
)

func main() {
	classname := "15-440"

	fmt.Println("Hello,", classname)

	fmt.Println("Hello half-class", classname[0:3])

	fooname := classname[0:3]

	fmt.Println("Hello fooclass: ", fooname)

}
