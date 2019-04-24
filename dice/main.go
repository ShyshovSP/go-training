package main

import (
	"fmt"
	"math/rand"
)

const n = 10000 // number or rollings

var x [11]int
var y = [11]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

func main() {

	for i := 1; i < n; i++ {

		var a, b, S int
		a = rand.Intn(6) + 1
		b = rand.Intn(6) + 1

		S = a + b
		var g int
		g = S - 2

		x[g]++

	}

	fmt.Println(y)
	fmt.Printf("%v", x)
}
