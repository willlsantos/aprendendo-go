package main

import (
	"fmt"
)

type hotdog int

var b hotdog = 10

func main() {
	x := 10
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", b)

	x = int(b)
	conversao()
}

func conversao() {
	x := 10
	//fmt.Printf("%v\n", x)
	x = int(b) //Convertendo a variavel do tipo hotdog para tipo int
	fmt.Printf("%v\n", x)
}
