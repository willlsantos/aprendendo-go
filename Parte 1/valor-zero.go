/* Usar := sempre que possivel (Declarador curto de operação)
   Usar o var para criar variaveis package-level scope
   Declaração vs Inicialização vs Atribuição de valor das variaveis. Ex: Caixas postais

   Os zeros:
	- ints: 0
	- floats: 0.0
	- booleans: false
	- strings: ""
	- pointes, functions, interfaces, slices, channels, maps: nil */

package main

import (
	"fmt"
)

var x int
var a int
var b float64
var c string
var d bool

func main() {
	//x = 10 //Inicialização e atribuição de valor da variavel
	//x = 20 //Apenas atribuição de valor da variavel
	//x := 10
	//Abaixo print de variaveis declaradas porém não inicializadas.
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	fmt.Printf("%v, %T\n", x, x)
}
