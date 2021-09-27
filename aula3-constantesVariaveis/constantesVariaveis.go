package main

import (
	"fmt"
	"math"
	//m "math" - podemos criar alias dessa forma ao import
)

func main() {
	//sintaxe basica: const/var + nome + tipo = valor
	const PI float64 = 3.1415

	//Nao sou obrigado a declarar o tipo em go (tipo inferido na declaracao do valor)
	var raio = 3.2

	// forma reduzida de criar um var e mais utilizada
	// := declaracao de novo var
	// = atribuindo valor novo
	area := PI * math.Pow(raio, 2)

	//em Go uma variavel nao usada da um erro de compilacao
	//todas variaveis declaradas devem ser usadas

	fmt.Println("A area da circunferencia eh: ", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	//declarando mais de uma variavel na mesma linha

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)
}
