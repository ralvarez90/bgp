package ch02gobasics

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

// PRINTLN
// Empleamos la función fmt.Println para imprimir
// mensajes por la consola.
func LS01HelloWorld(desc string) {
	fmt.Println(desc)
	fmt.Println("Hello World desde Golang!")
}

// VARIABLES
// Empleamos la palabra reservada var para declarar
// variables.
func LS02Variables(desc string) {
	fmt.Println(desc)

	// empleamos la palabra reservada var
	// para declarar variables. Existen
	// - boolean
	// - string
	// - uint8|16|32|64
	// - int8|16|32|64
	// - float32
	// - float64
	// - complex

	// Con alias, tenemos otros dipos de datos que
	// son derivados pero se les asigna otro alias.
	// byte -> uint8
	// uint -> uint32
	// int  -> int32
	// rune -> int32
	// uintptr

	// Colecciones
	// Arrays
	// Slices
	// Maps
	// Structs
	// Enums

	// Tipos en los que el lenguaje se organiza
	// Funciones
	// Interfaces
	// Channels
	// Punteros

	// Constantes
	// Se declarar unsando la palabra reservada const. No
	// permite cambiarle el valor durante la ejecución del
	// programa.

	// Declaración explícita
	var strVariable string = "i am a string"
	fmt.Printf("strVariable: %s\n", strVariable)

	// asigna valor default ""
	var anotherStrVar string
	fmt.Printf("anotherStrVar: %s\n", anotherStrVar)

	// entero
	var defaultInt int
	fmt.Printf("defaultInt: %d\n", defaultInt)

	// Declaración implícita
	myName := "Rodrigo Álvarez"
	fmt.Printf("myName: %s\n", myName)

	// Declaración de constante
	const PI_VALUE float64 = 3.14159
	fmt.Printf("Constante PI_VALUE: %f\n", PI_VALUE)
}

// INPUT DATA
// Se ingresan algunos valores por teclado usando las
// diferentes alternativas.
// De preferencia use bufio para entrada de datos.
func LS03EntradaDeDatos(desc string) {
	fmt.Println(desc)

	// Forma 1, uso de fmt.Scan, fmt.Scanf y fmt.Scanln para
	// leer info por teclado a partir del os.Stdin
	fmt.Println("1. Usando fmt.Scanf")
	var name string
	var age int
	fmt.Print("Input your name: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Input yout age: ")
	fmt.Scanf("%d", &age)
	fmt.Printf("Data{name: %s, age: %d}\n", name, age)
	fmt.Println("--->")

	fmt.Print("Input yout name: ")
	fmt.Scanln(&name)
	fmt.Print("Input your age: ")
	fmt.Scanln(&age)
	fmt.Printf("Data{name: %s, age: %d}\n", name, age)
	fmt.Println("--->")

	// Forma 2, uso de Fscan, Fscanf, Fscanln para leer entradas
	// de datos de un io.Reader específico.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter some text: ")
	input, _ := reader.ReadString('\n')
	fmt.Printf("You have entered: %s", input)
	fmt.Println("--->")
}

// MATH PACKAGE
func LS04MathPackage(desc string) {
	fmt.Println(desc)

	// suma de enteros
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Printf("Integer sum: %d\nType: %T\n", intSum, intSum)

	// suma de flotantes
	f1, f2, f3 := 123.3, 23.2, 2.00009
	floatSum := f1 + f2 + f3
	fmt.Printf("Float sum: %f\nType: %T\n", floatSum, floatSum)

	// uso de redonde
	floatSum = math.Round(floatSum)
	fmt.Printf("Rounded sum: %f\n", floatSum)

	circleRadius := 15.6
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference)
}

// DATES AND TIMES
func LS05DatesAndTimes(desc string) {
	fmt.Println(desc)

	// objeto Time
	now := time.Now()
	fmt.Println("Current time:", now)

	// format date
	formatDate := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go was lanched at:", formatDate)

	// my birthday
	myBirthday := time.Date(1990, time.March, 10, 12, 0, 0, 0, time.Local)
	fmt.Println("My birthday is:", myBirthday)
	fmt.Println("My birthday is:", myBirthday.Format(time.ANSIC))
	fmt.Println("My birthday is:", myBirthday.Format(time.DateOnly))

}
