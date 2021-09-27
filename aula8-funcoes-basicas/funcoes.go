package main

import "fmt"

//nome - parametros - retorno
func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	a := 2
	b := 3
	imprimir(somar(a, b))
}
