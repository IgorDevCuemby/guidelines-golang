## Tipo de contenedores: Arrays, Slices y Maps

### OBJETIVO

- Conocer los distintos tipos de contenedores en Go
- Utilizar arrays, slices y maps para almacenar datos
#### REQUISITOS

```
go --version
```
#### DESARROLLO

Go tiene tres tipos de contenedores principales que se pueden usar para almacenar valores de elementos. [contenedores.go](./contenedores.go).

- ArrayArray: son secuencias de longitud fija de cero o más elementos de un tipo en particular.
- Slice: una abstracción alrededor de una matriz que se puede cambiar de tamaño en tiempo de ejecución.
- Map: una colección de pares clave-valor, donde la clave es un tipo de dato de referencia, y el valor es un tipo de dato de valor. Estructura de datos asociativa que permite que distintas claves se emparejen arbitrariamente o se "asignen a" valores.