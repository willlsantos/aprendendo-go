package main

import (
	"fmt"
)

//Variaveis declaradas aqui, funcionam em todo o código
//Para atribuir valores a variaveis existentes usamos o sinal =
//O valor de atribuição, só é atribuido a variaveis existentes

func main() { //Isso aqui é o codeblock.

	// a marmota := só funciona dentro do codeblock e é utilizada para criar novas variaveis

	x := 10
	y := "Olá bom dia"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("x: %v, %T\n", z, z)
}
