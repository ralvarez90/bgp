package ch02gobasics

import (
	"fmt"
	"sort"
)

// MAPS
// Son colecciones desordenadas de pares key-value. Se
// puede considerar como una tabla hash. Los keys de cada
// elemento son los que se comparar para efectuar ordenamientos.

// Empleamos la función make para crear estas estructuras de
// datos.

// Cuando iteramos un mapa usando range, se itera key por key

// Empleamos los keys como índices y así agregarle elementos a un
// mapa
func LS13Maps(desc string) {
	fmt.Println(desc)
	fmt.Println("Sintaxis:")
	fmt.Println("map_variable := make(map[key_data_type]value_data_type)")

	// ejemplo
	states := make(map[string]string)
	states["NY"] = "New York"
	states["WA"] = "Washinton"
	states["CA"] = "California"
	states["MX"] = "México"

	// delete and element
	delete(states, "MX")
	states["MEX"] = "México"

	fmt.Println("States:", states)
	fmt.Println("--->")

	// slices content of keys
	keys := make([]string, len(states))
	i := 0
	for key := range states {
		keys[i] = key
		i++
	}

	// order keys
	sort.Strings(keys)
	fmt.Println("Order of keys after sort.Strings:", keys)
	fmt.Println("--->")

	for key, value := range states {
		fmt.Printf("states[%s] -> %s\n", key, value)
	}
}
