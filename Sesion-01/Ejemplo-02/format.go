package main

import "fmt"

func main() {
	//Fromato de entrada y salida E/S
	//%v es el formato por defecto
	fmt.Printf("Formato por defecto: %v\n", 10)

	//%T muestra el tipo de dato
	fmt.Printf("Tipo de dato: %T\n", 10)

	// %% muestra el %
	fmt.Printf("signo porcentual: %%\n")

	//%t muestra un valor booleano
	fmt.Printf("valor booleano: %t\n", true)

	//%d muestra un valor entero
	fmt.Printf("valor entero: %d\n", 10)

	//%b muestra un valor en binario
	fmt.Printf("valor en binario: %b\n", 10)

	//%f muestra un valor flotante
	fmt.Printf("valor flotante: %f\n", 10.5)

	//%s muestra un valor string
	fmt.Printf("valor string: %s\n", "Hola Mundo")

	//%q muestra un valor string
	fmt.Printf("valor string: %q\n", "Hola Mundo")

}
