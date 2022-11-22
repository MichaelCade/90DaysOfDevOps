Antes de empezar con los temas de hoy mandamos un gran saludo a [Techworld con Nana](https://www.youtube.com/watch?v=yyUHQIec83I) y a ese fantástico y conciso viaje a través de los fundamentos básicos de Go.

En el [Día 8](day08.md) configuramos nuestro entorno, en el [Día 9](day09.md) recorrimos el código de Hello #90DaysOfDevOps y en el [Día 10](day10.md)) vimos nuestro workspace Go y profundizamos un poco sobre la compilación y ejecución del código.

Hoy vamos a echar un vistazo a las Variables, Constantes y Tipos de Datos escribiendo un nuevo programa.

## Variables y Constantes en Go

Empecemos por planificar nuestra aplicación, parece una buena idea trabajar en un programa que nos diga cuántos días hemos permanecido en nuestro reto #90DaysOfDevOps.

Lo primero que hay que tener en cuenta aquí es que a medida que construyamos nuestra app, demos la bienvenida a nuestros asistentes y vayamos dando el feedback al usuario sobre el número de días que ha completado, a lo largo del programa, podríamos utilizar el término #90DaysOfDevOps muchas veces. Así que aquí tenemos una oportunidad de hacer una variables con #90DaysOfDevOps para nuestro programa.

- Las variables se utilizan para almacenar valores.
- Son como una pequeña caja con nuestra información o valores guardados (Así lo explican todos los docentes a quienes empiezan 😊).
- Podremos utilizar esta variable en todo el programa, lo que también beneficia a que si este reto o la variable cambia, sólo tendremos que cambiar el valor en un solo lugar. Esto significa que podemos trasladar esto a otros retos que tenemos en la comunidad con sólo cambiar el valor de esa variable. 

Para declarar esto en nuestro programa Go definimos un valor utilizando una **keywords** para las variables. Esto permanecerá dentro de nuestro bloque de código `func main` que verás más adelante. Puedes encontrar más información sobre Keywords [aquí](https://go.dev/ref/spec#Keywords).

Recuerda asegurarte de que los nombres de tus variables sean descriptivos. Si declaras una variable debes usarla o recibirás un error, esto es para evitar posible dead code, código que nunca se usa. Lo mismo ocurre con los paquetes que no se utilizan.

```go
var challenge = "#90DaysOfDevOps"
```

Con lo anterior establecido vamos a utilizarlo en el siguiente fragmento de código. Se puede ver en la salida de abajo que hemos utilizado la variable.

```go
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    fmt.Println("Welcome to", challenge, "")
}
```

Puedes encontrar el código anterior en el fichero [day11_example1.go](Go/day11_example1.go).

A continuación verás que hemos construido nuestro código con el ejemplo anterior y hemos obtenido la salida que se muestra a continuación.

![](Images/Day11_Go1.png)

Sabemos que este reto en concreto es de 90 días, pero el próximo, quizás sea de 100 por lo que definiremos una variable que nos ayude aquí también. Sin embargo, para nuestro programa, queremos definirla como una constante. Las constantes son como las variables, excepto que su valor no puede ser cambiado dentro del código (aunque podemos crear una nueva aplicación más adelante con el mismo código y cambiar la constante. El 90 no cambiará mientras estemos ejecutando esta aplicación)

Añadiendo `const` a nuestro código y añadiendo otra línea de código para imprimir esto.

```go
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    const daystotal = 90

    fmt.Println("Welcome to", challenge, "")
    fmt.Println("This is a", daystotal, "challenge")
}
```

Puedes encontrar el código anterior en[day11_example2.go](Go/day11_example2.go).

Si a continuación volvemos a pasar por ese proceso de `go build` y lo ejecutamos veremos a continuación el resultado.

![](Images/Day11_Go2.png)

Por último añadiremos otra variable para el número de días que hemos completado. Pero esto no será el final de nuestro programa, seguiremos con nuestro programa el [Día 12](day12.md) para añadir más funcionalidad. 

A continuación verás añadida la variable `dayscomplete` con el número de días completados.

```go
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    const daystotal = 90
    var dayscomplete = 11

    fmt.Println("Welcome to", challenge, "")
    fmt.Println("This is a", daystotal, "challenge and you have completed", dayscomplete, "days")
    fmt.Println("Great work")
}
```

Puedes encontrar el código anterior en [day11_example3.go](Go/day11_example3.go)

Volvamos a ejecutar compilando con `go build` y ejecutando el fichero compilado o directamente usando `go run`.

![](Images/Day11_Go3.png)

Aquí hay algunos ejemplos para hacer el código más fácil de leer y editar. Hasta ahora hemos estado usando `Println` pero podemos simplificar cambiando el `Printf` por `%v` lo que definirá nuestras variables en orden al final de la línea de código. También vamos a usar `\n` para un salto de línea.

Usando `%v` tendremos un valor por defecto, pero hay otras opciones que se pueden encontrar en la [documentación del paquete fmt](https://pkg.go.dev/fmt). 

Se pone interesante la cosa, puedes ver el código de ejemplo [day11_example4.go](Go/day11_example4.go)

Las variables también pueden ser definidas en un formato más simple aun. En vez de definir `var` y el "tipo" puedes codificarla de la siguiente manera para obtener la misma funcionalidad pero con un aspecto más limpio y sencillo. Esto sólo funcionará para las variables y no para las constantes.

```go
func main() {
    challenge := "#90DaysOfDevOps"
    const daystotal = 90
```

## Tipos de datos

En los ejemplos anteriores, no hemos definido el tipo de las variables, esto es porque podemos darle un valor y Go es lo suficientemente inteligente como para saber de qué tipo se trata o al menos puede inferir cuál es en base al valor que ha almacenado. Sin embargo, si queremos que un usuario ingrese esto requerirá un tipo específico.

Hasta ahora hemos utilizado Strings y Integers en nuestro código. Los integer para el número de días y los strings para el nombre del reto.

Es importante tener en cuenta que cada tipo de datos puede hacer cosas diferentes y se comporta de forma distinta. Por ejemplo, los integer pueden multiplicarse mientras que las cadenas no.

Hay cuatro categorías:

- **Tipo Basic**:  Los números, las cadenas y los booleanos entran en esta categoría.
- **Tipo Aggregate**: Los arrays y los structs entran en esta categoría.
- **Tipo Reference**: Pointers, slices, maps, functions, y channels se incluyen en esta categoría.
- **Interface type**

El tipo de datos es un concepto importante en la programación, pues especifica el tamaño y el tipo de los valores de las variables.

Go está tipado estáticamente, lo que significa que una vez que se define un tipo de variable, sólo puede almacenar datos de ese tipo.

Go tiene tres tipos de datos básicos:

- **bool**:  representa un valor booleano, es verdadero o falso
- **Numeric**: representa tipos integer, valores de punto float y tipos complejos.
- **string**: representa un valor de cadena

Este recurso está súper detallado sobre los tipos de datos: [Golang by example](https://golangbyexample.com/all-data-types-in-golang-with-examples/).

Si necesitamos definir un tipo en nuestra variable podemos hacerlo así:

```go
var TwitterHandle string
var DaysCompleted uint
```

Como Go implica variables donde se da un valor podemos imprimir esos valores con lo siguiente:

```go
fmt.Printf("challenge is %T, daystotal is %T, dayscomplete is %T\n", conference, daystotal, dayscomplete)
```

Hay muchos tipos diferentes de integer y floats.

- **int** =  números enteros
- **unint** = números enteros positivos
- **floating point types** = números que contienen un componente decimal

## Recursos

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

En el próximo día vamos a empezar a añadir alguna funcionalidad de entrada, para que el usuario pueda añadir datos en nuestro programa, por ejemplo, los días que se han completado.

Nos vemos en el [Día 12](day12.md).
