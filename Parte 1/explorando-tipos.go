/*Em GO os tipos de variaveis são muito importantes, é a base de tudo
Tipos em GO são estáticos.
O Tipo da variavel pode ser deduzido pelo compilador ou declarado especificamente.
Se a variavel for declarada sem o valor, esse valor só poderá ser atribuido a
essa variavel, dentro de um Code Block.

Tipos de Dados Primitivos: int, string, bool
Tipos de Dados Compostos: slice, array, struct, map

O ato de definir, criar e estruturar tipos compostos chama-se composição.
*/
package main

import (
	"fmt"
)

var x int = 10 //ou posso declara var x = 10
var q = 10.2   // ou posso declarar assim.

func main() {

	fmt.Printf("O valor de x é: %v, o tipo da varíavel é: %T\n", x, x)
	fmt.Printf("O valor de q é: %v, o tipo da varíavel é: %T\n", q, q)
}
