package main

import "fmt"

func main() {

	var b byte = 3
	fmt.Println(b)

	i := 4
	fmt.Println(i)

	i += 4
	fmt.Println(i)

	i -= 4
	fmt.Println(i)

	i /= 2
	fmt.Println(i)

	i *= 4
	fmt.Println(i)

	i %= 2
	fmt.Println(i)

	x, y := 4, 6
	fmt.Println(x + y)

	x, y = y, x
	fmt.Println(x)
	fmt.Println(y)
}
