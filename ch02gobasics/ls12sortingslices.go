package ch02gobasics

import (
	"fmt"
	"math/rand"
	"sort"
)

// ORDENANDO SLICES
// Go implementa la función sort para ordenar distintos elementos
// del lenguaje. Esta se importa del paquete sort
func LS12SortingSlices(desc string) {
	fmt.Println(desc)

	// create and slice
	intSlices := make([]int, 5)
	for i := 0; i < 5; i++ {
		intSlices[i] = 1 + rand.Intn(100)
	}

	fmt.Println("original:", intSlices)

	// ordenamos
	sort.Ints(intSlices)
	fmt.Println("sorted:", intSlices)
}
