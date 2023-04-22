package ch02gobasics

import "fmt"

// FUNCIONES APPEND Y COPY
// Se emplean para copiar el contenido de un slices en otro
// y para agregar elementos en un slices.

// Podemos emplear la funcuón append para eliminar elementos
// de un mismo slice.
func LS11AppendAndCopy(desc string) {
	fmt.Println(desc)

	// declaramos slices
	var nums []float64
	printSliceDetails(nums)
	for i := 1; i <= 10; i++ {
		nums = append(nums, (float64(i)))
		printSliceDetails(nums)
	}

	fmt.Println("--->")

	// creamos copia
	copyOfNums := make([]float64, len(nums), cap(nums)*2)
	copy(copyOfNums, nums)
	fmt.Println("copyOfNums:", copyOfNums)
	printSliceDetails(copyOfNums)

	fmt.Println("--->")

	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println("Before:", colors)

	// eliminamos primer elemento
	colors = append(colors[1:])
	fmt.Println("Items after removing 1st element:", colors)
}
