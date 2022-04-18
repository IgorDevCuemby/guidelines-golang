## Tipos de dato

### OBJETIVO

- Usar los distintos tipos de datos disponibles en Go
- Utiilizar funciones de cast para hacer conversiones de tipo

#### REQUISITOS

```
go --version
```
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


- Texto - string (str):
Una cadena representa una secuencia de puntos de código Unicode. las cadenas en Go son inmutables:
una vez creado, no es posible cambiar el contenido de una cadena.

Go admite dos estilos de literales de cadena, el estilo de comillas dobles (o literales interpretados) y el estilo de comillas inversas (o literales de cadena sin formato). por ejemplo, los siguientes dos literales de cadena son equivalentes:
``` 
//Cadenas de texto

// La forma interpretada
"Hola\nworld!\n"

// la forma de cadena sin formato
`Hola
world!`

``` 
** go run Sesion-01/Ejemplo-02/tipos_de_datos.go **
#Variables 
las variables se pueden declarar utilizando la palabra clave ~~var~~ para emparejar un identificador con algún valor escrito, y se pueden actualizar en cualquier momento, con la forma general: [Tipos de Datos](./tipos_de_datos.go).

``` 

var name type = expression

``` 

sin embargo, existe una flexibilidad considerable en la declaración de variables:

```
//Con Inicialización
var foo int = 5

//Multiple declaraciones
var (
    bar int
    baz string
)
//Or
var bar, baz int = 1, 2

//Con inferencia de tipo
var qux = 42

//Multiples tipos mixtos
var b, f, s = true, 2.3, "four"

//Sin inicialización
var s string
```

:pushpin: go es muy obstinado sobre el desorden: lo odia. si declara una variable en una función pero no la usa, su programa se negará a compilar.

- Declaración de variables cortas:
go proporciona un poco de ~~azúcar sintáctico~~ que permite que las variables dentro de las funciones se declaren y asignen simultáneamente usando el operador := en lugar de una declaración ~~var~~ son un tipo implícito.
declaraciones de variables cortas tienen la forma general
```
//Declaración de variables cortas
name := expression

```

estos se pueden usar para declarar asignaciones únicas y múltiples:
```
//Con inicialización
percent := rand.Float64() * 100.0

//Multiple declaraciones a la vez
x, y := 1, 2
```

:warning: recuerda que := es una declaración, y = es una asignación. Un operador := que intente declarar variables existentes fallará en tiempo de compilación.
curiosamente (y a veces confuso), si una declaración de variable corta tiene una combinación de variables nuevas y existentes en su lado izquierdo, la declaración de variable corta actúa como una asignación a las variables existentes.

#### Formato E/S en Go

El paquete fmt de go implementa varias funciones para formatear la entrada y la salida. Los más utilizados de estos son (probablemente) fmt.Printf y fmt.Scanf, que se pueden usar para escribir en la salida estándar y leer desde la entrada estándar, respectivamente.
        
        ```
        //Escribir en la salida estándar
        func Printf(format string, a ...interface{}) (n int, err error)

        //Leer desde la entrada estándar
        func Scanf(format string, a ...interface{}) (n int, err error)
        ```
algunas de las banderas comunes que se usan en las cadenas de formato incluyen: [Formato.go](./format.go).

        - %v: el del formato por defecto.
        - %T: una representación de tipo del valor.
        - %% un signo de porcentaje literal. no consume ningún argumento.
        - %t un booleano. se imprime como true o false.
        - %b Integer en base 2.
        - %d Integer en base 10.
        - %f Decimal punto flotante decimal, pero, no exponente (e) 123.456 
        - %s String en formato de cadena.
        - %q String en formato de cadena con comillas.

#### El identificador Blank ( _ )

el identificador en blanco, representado por el operador _ (guión bajo), actúa como un marcador de posición anónimo. Puede usarse como cualquier otro identificador en una declaración, excepto que no introduce un enlace.

        ```
        str := "World!"

        _, err := fmt.Printf("Hello, %s\n", str)
        if err != nil {
            // Do something
        }
        ```

