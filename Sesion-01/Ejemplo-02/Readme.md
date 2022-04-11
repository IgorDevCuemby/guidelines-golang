## Tipos de dato

### OBJETIVO

- Usar los distintos tipos de datos disponibles en Go
- Utiilizar funciones de cast para hacer conversiones de tipo

#### REQUISITOS

1. 

#### DESARROLLO
Los tipos de datos básicos de Go, los bloques de construcción fundamentales a partir de los cuales se construyen los tipos más complejos, se pueden dividir en tres subcategorías:

* Booleanos que contienen solo un bit de información, verdadero o falso, que representa alguna conclusión o estado lógico.

* Numéricos que representan simples (coma flotante de varios tamaños y enteros con y sin signo) o números complejos

* Texto que representan una secuencia inmutable de puntos de código Unicode.

- Booleanos:

Los datos booleanos, representa los dos valores de verdad lógicos, existe de alguna forma en todos los lenguajes de programación que se hayan inventado. está representado por el tipo bool, un tipo entero especial de 1-bit que tiene dos valores posibles:

        - True (bool) 

        - False (bool) 
```
#Datos booleanos
and := true && false
fmt.Println(and) // false

or := true || false
fmt.Println(or) // true

not := !true
fmt.Println(not) // false
```
Curiosamente, Go no incluye un operador lógico XOR. hay un operador, pero está reservado para la operación XOR bit a bit

- Numéricos:
Go tiene una pequeña colección de números enteros con nombre sistemático, de punto flotante y con y sin signo:

        enteros con signo
        - int8, int16, int32, int64

        enteros sin signo
        - uint8, uint16, uint32, uint64

        punto flotante
        - float32, float64
```
#Datos numéricos
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

```

Para la mayoría de los usos, generalmente tiene sentido usar solo int y float64.

- Numéricos complejos:
Go ofrece dos tamaños de números complejos, si te sientes un poco imaginativo complex64 y complex128. Estos se pueden expresar como un literal imaginario mediante un punto flotante seguido inmediatamente por un i:

        - complex64, complex128
```
#Datos complejos
var x complex64 = 3.1415i
fmt.Println(x) // (0+3.1415i)

var y complex128 = 3.1415i
fmt.Println(y) // (0+3.1415i)

```        


- Texto:

        - string (str): Cadenas de texto (cadena de caracteres)



