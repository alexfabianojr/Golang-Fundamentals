package main

import "fmt"

func comprar(trab1, trab2 bool) (bool, bool, bool) {

	comprarTv50 := trab1 && trab2
	comprarTv30 := trab1 != trab2 // ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv30, comprarSorvete
}

func main() {

	tv50, tv30, sorvete := comprar(true, true)

	fmt.Printf("Tv50: %t, Tv30: %t, Sorvete: %t, Saudavel: %t",
		tv50, tv30, sorvete, !sorvete)

}
