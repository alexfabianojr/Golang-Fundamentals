package main

import "fmt"

func main() {

	i := 1

	//Go nao tem aritmetica de ponteiro
	//ponteiro eh uma referencia de memoria

	var p *int = nil // nil = null

	p = &i //"pegue esta variavel e me devolva o endereco de memoria dela"

	fmt.Println(i)

	fmt.Println(p)

	*p++ //desreferenciando (pegando o valor pela referencia e incrementando ela)

	fmt.Println(i)

	fmt.Println(p)

	fmt.Println(p, *p, i, &i)
}
