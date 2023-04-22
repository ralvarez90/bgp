package ch02gobasics

import "fmt"

// Inicialización de "Objetos"
// Go permite inicializar elementos compuestos de dos formas
// diferentes. Una es usando new y otra make.

// New únicamente reserva la memoria del objeto complejo pero
// no lo inicializa en memoria, por ejemplo cuando un Map es
// es alojado usando new, unicamente regresa el puntero a dicho
// map.

// New retorna unicamente un puntero.

// Make function, aloja e inicializa los objetos en memoria.
func LS06MakeAndNew(desc string) {
	fmt.Println(desc)

	// forma correcta
	string_Map := make(map[string]int)
	string_Map["Marks"] = 46
	fmt.Println(string_Map)
}

// Punteros
// Son en escencia variables que almacenan direcciones de
// memoria.
func LS07PointersDataType(desc string) {
	fmt.Println(desc)

	// vars
	value1 := 42
	var ptrToValue1 = &value1

	var ptrToValue2 *int
	value2 := 100
	ptrToValue2 = &value2

	fmt.Println("address of value1:", &value1)
	fmt.Println("address of value1:", ptrToValue1)

	fmt.Println("address of value2:", &value2)
	fmt.Println("address of value2:", ptrToValue2)

	fmt.Println("content of value1:", value1)
	fmt.Println("content of value1:", *ptrToValue1)

	fmt.Println("content of value2:", value2)
	fmt.Println("content of value2:", *ptrToValue2)
	fmt.Println("--->")

	// Modificamos variable desde puntero
	value3 := 32.14
	ptrFloatValue3 := &value3
	fmt.Println("value3:", value3)

	*ptrFloatValue3 = *ptrFloatValue3 / 31
	fmt.Println("new value3:", value3)
	fmt.Println("new value3:", *ptrFloatValue3)
}

// ARRAYS
// Son colecciones de un mismo tipo de tamaño fijo. Es decir
// su capacidad y su longitud son iguales siempre.
// Se emplea la función len para obtener la longitud de un
// array.
func LS08Arrays(desc string) {
	fmt.Println(desc)

	// Sintaxis general
	// var arrayName[size] datatype
	var someIntegers [10]int
	for i := 0; i < len(someIntegers); i++ {
		someIntegers[i] = -1
	}
	fmt.Println("someIntegers:", someIntegers)
	fmt.Println()

	// arreglo de flotatnes
	var someFloats = [5]float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("someFloats:", someFloats)

	// otro
	var somePares = []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	fmt.Println("somePares:", somePares)
	fmt.Println()

	// declarando e inicializando
	var array2 [10]int
	var i, j int
	j = 1
	for i = 0; i < 10; i++ {
		array2[i] = i*j + 50
		j += 6
	}

	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("array2=", array2)
}
