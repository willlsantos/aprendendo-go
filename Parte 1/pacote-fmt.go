package main

import (
	"fmt"
)

func main() {
	x := "Oi Bom dia\n Como vai?\tespero \"que\" tudo bem." //Isso aqui é um interpretaded string litherals
	//São strings que são interpretadas, o /n significa uma nova linha, o /t um tab e etc
	//y := `"Oi Bom dia\n Como vai?\tespero \"que\" tudo bem."` //Nada do que foi adicionado foi interpretado, ficou de forma "crua"
	fmt.Println(x)
	sprint()
}

func sprint() {
	x := "Oi"
	y := "bom dia"
	z := fmt.Sprint(x, " ", y)

	fmt.Println(z)
}

/*
- Format printing: documentação.
    - Grupo #1: Print → standard out
        - func Print(a ...interface{}) (n int, err error)
        - func Println(a ...interface{}) (n int, err error)
        - func Printf(format string, a ...interface{}) (n int, err error)
            - Format verbs. (%v %T)
    - Grupo #2: Print → string, pode ser usado como variável
        - func Sprint(a ...interface{}) string
        - func Sprintf(format string, a ...interface{}) string
        - func Sprintln(a ...interface{}) string
    - Grupo #3: Print → file, writer interface, e.g. arquivo ou resposta de servidor
        - func Fprint(w io.Writer, a ...interface{}) (n int, err error)
        - func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
        - func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
*/
