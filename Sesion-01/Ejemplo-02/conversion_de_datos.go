package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Conversion de datos
	//Convierte un valor de tipo string a un valor de tipo int
	var numero int
	numero, _ = strconv.Atoi("10")
	fmt.Println(numero)

	//Convierte un valor de tipo int a un valor de tipo string
	var numero2 string
	numero2 = strconv.Itoa(10)
	fmt.Println(numero2)

	//Convierte un valor de tipo string a un valor de tipo float64
	var numero3 float64
	numero3, _ = strconv.ParseFloat("10.5", 64)
	fmt.Println(numero3)

	//Convierte un valor de tipo float64 a un valor de tipo string
	var numero4 string
	numero4 = strconv.FormatFloat(10.5, 'f', 2, 64)
	fmt.Println(numero4)

	//Convierte un valor de tipo string a un valor de tipo bool
	var numero5 bool
	numero5, _ = strconv.ParseBool("true")
	fmt.Println(numero5)

	//Convierte un valor de tipo bool a un valor de tipo string
	var numero6 string
	numero6 = strconv.FormatBool(true)
	fmt.Println(numero6)
}
