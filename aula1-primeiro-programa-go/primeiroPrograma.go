// Programas executaveis iniciam pelo pacote main
package main

/*
Os codigos em Go sao organizados em pacotes
e para usa-los eh necessario declarar um ou varios imports
*/
import (
	"fmt"
)

// A funcao de entrada de um programa Go eh a funcao main
// A main deve ser unica na hierarquia mais alta do projeto
// Por isso dividimos o curso em pastas
func main() {
	fmt.Print("Primeiro!")
	fmt.Print("Programa!")
}
