package main

import (
	"fmt"
	"strconv"
)

func main() {

	//eh preciso converter explicitamente de um tipo para outro

	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	//cuidado

	fmt.Println(string(123)) //valor unicode

	//int para string

	fmt.Println(strconv.Itoa(notaFinal))

	//string para int

	num, _ := strconv.Atoi("123") //func que retorna dois valores (int e erro)

	fmt.Println(num)

	b, _ := strconv.ParseBool("1")

	if b {
		fmt.Println("Verdadeiro")
	}
}
