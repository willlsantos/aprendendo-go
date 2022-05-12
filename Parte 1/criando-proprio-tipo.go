/* Aprendendo a declarar nossos próprios tipos
Lembrando que tipos são fixo, uma vez declarada a variavel como de um certo tipo,
Ela será assim até final. Ou seja, ela é imutável. */
package main

import (
	"fmt"
)

type hotdog int //Tipo base hotdog é do tipo Int

var b hotdog //Aqui criamos uma variavel chamada b do tipo hotdog

func main() {
	x := 10
	fmt.Printf("%T\n", x) //Variavel do tipo int
	fmt.Printf("%T", b)   //Variavel do tipo hotdog
}
