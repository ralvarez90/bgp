package ch02gobasics

import "fmt"

// SLICES
// Son arrays de tamaño dinámico. Para definir slices
// no es necesario especificar la longitud si se
// declara e inicializan.

// Se puedne de clarar de la forma
// var sliceName []type

// o con la función Make. Que recibe un tipo, una longitud
// inicial.

// Si un slice es declarado pero no inicializado este almacena
// un nil.
func LS09Slices(desc string) {
	fmt.Println(desc)

	// ejemplo 1,
	var colors = []string{"RED", "GREEN", "BLUE"}
	fmt.Printf("Colors: %v with type: %T\n", colors, colors)

	// ejemplo 2,
	var edades = make([]float64, 10, 20)
	printSliceDetails(edades)

	// ejemplo 3,
	var someElements []float64
	fmt.Println("someElements:", someElements)
	printSliceDetails(someElements)
}
