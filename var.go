package main

import (
	"fmt"
)

var y = 10 //variavel criada com scope do package inteito

func main() {
	z := 20 //variavel criada com scope apenas deste codeblock
	teste(z)
}

func teste(x int) {
	fmt.Println(y)
	fmt.Println(x)
}
