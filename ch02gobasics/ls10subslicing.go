package ch02gobasics

import "fmt"

// SUB-SLICING
// Podemos obtener sub slices de slices mediante rangos
// como en python.
func LS10Subslicing(desc string) {
	fmt.Println(desc)

	marks := []float64{10, 12.6, 20.9, 37.54, 48.75, 50.0}
	printSliceDetails(marks)
	fmt.Println("--->")
	fmt.Println("original   ->", marks)
	fmt.Println("marks[1:]  ->", marks[1:])
	fmt.Println("marks[:4]  ->", marks[:4])
	fmt.Println("marks[1:3] ->", marks[1:3])
	fmt.Println("marks[:]   ->", marks[:])

}

func printSliceDetails(x []float64) {
	fmt.Printf("Lenght: %d, Capacity: %d, Slice: %v\n",
		len(x), cap(x), x)
	if x == nil {
		fmt.Println("El slice es nill...")
	}
}
