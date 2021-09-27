package main

import (
	"fmt"
	"math"
)

func main() {

	a, b := 3, 2

	fmt.Println("Soma: ", a+b)
	fmt.Println("Subtracao: ", b-a)
	fmt.Println("Divisao: ", a/b)
	fmt.Println("Multiplicacao: ", a*b)
	fmt.Println("Modulo: ", a%b)

	//bitwise

	fmt.Println("E: ", a&b)   //11 & 10 = 10
	fmt.Println("Ou: ", a|b)  // 11 | 10 = 11
	fmt.Println("Xor: ", a^b) // 11 ^ 10 = 01

	//outras operacoes

	c, d := 2.0, 3.0

	fmt.Println("Maior:", math.Max(float64(a), float64(b)))
	fmt.Println("Menor:", math.Min(c, d))
	fmt.Println("Exponencial: ", math.Pow(c, d))
}
