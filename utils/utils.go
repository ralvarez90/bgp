package utils

import (
	"fmt"
	"strings"
)

// Imprime un separador de linea en consola.
func PrintSeparator(n int) {
	fmt.Println(strings.Repeat("-", n))
}

// Imprime mensaje y espera entrada del usuario
func SystemPause() {
	fmt.Print("\nPress any key to continue. . . ")
	fmt.Scanln()
}
