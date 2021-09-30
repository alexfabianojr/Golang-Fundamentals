package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Banna" == "Banna")
	fmt.Println(3 != 2)
	fmt.Println(3 < 2)
	fmt.Println(3 > 2)
	fmt.Println(3 <= 2)
	fmt.Println(3 >= 2)

	data1 := time.Unix(0, 0)
	data2 := time.Unix(0, 0)

	fmt.Println(data1 == data2)
	fmt.Println(data1.Equal(data2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Joao"}
	p2 := Pessoa{"Joao"}

	fmt.Println(p1 == p2)
}
