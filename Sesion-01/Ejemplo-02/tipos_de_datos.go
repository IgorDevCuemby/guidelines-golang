package main

import "fmt"

func main() {
	// Datos Booleanos
	and := true && false
	fmt.Println(and) // false

	or := true || false
	fmt.Println(or) // true

	not := !true
	fmt.Println(not) // false

	//Datos numéricos
	int8 := int8(127)
	fmt.Println(int8) // 127

	int16 := int16(32767)
	fmt.Println(int16) // 32767

	int32 := int32(2147483647)
	fmt.Println(int32) // 2147483647

	int64 := int64(9223372036854775807)
	fmt.Println(int64) // 9223372036854775807

	float32 := float32(3.14)
	fmt.Println(float32) // 3.14

	float64 := float64(3.141592653589793)
	fmt.Println(float64) // 3.141592653589793

	//Datos complejos
	var x complex64 = 3.1415i
	fmt.Println(x) // (0+3.1415i)

	var y complex128 = 3.1415i
	fmt.Println(y) // (0+3.1415i)

}
