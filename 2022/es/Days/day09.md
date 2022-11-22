## Vamos a explicar el código de Hello World

### ¿Cómo funciona Go?

En el [Día 8](day08.md) realizamos la instalación de Go en el equipo y luego creamos nuestra primera aplicación Go.

En esta sección, vamos a echar un vistazo más profundo al código y entender algunas cosillas más sobre el lenguaje Go.

### ¿Qué es la compilación?

Antes de ver las [5 líneas del código de Hello](Go/hello.go) necesitamos entender un poco la compilación.

Los lenguajes de programación que utilizamos habitualmente, como Python, Java, Go y C++, son lenguajes de alto nivel. Esto significa que son legibles para el ser humano. Cuando una máquina intenta ejecutar un programa, éste debe tener una forma que la máquina pueda entender. Con lo cuál, tenemos que traducir nuestro código legible para el ser humano a código legible para la máquina (Código máquina). Esta acción es a lo que se le llama **compilar**.

![](Images/Day9_Go1.png)

Si no sabes de lo que estamos hablando mejor vuelve a ver lo que hicimos en el [Día 8](day08.md), creamos un simple Hello World en el fichero `main.go` y luego compilamos nuestro ejecutable con el comando `go build main.go`.

### ¿Qué son los paquetes?

Un paquete es una colección de archivos fuente en el mismo directorio que se compilan juntos. En el caso que nos ocupa con nuestro paquete, un grupo de archivos .go en el mismo directorio. ¿Recuerdas nuestra carpeta Hello del [día 8](day08.md)? Cuando realizas programas Go más complejos puedes encontrar que tienes la carpeta1, la carpeta2 y la carpeta3 que contienen diferentes ficheros .go que componen tu programa con múltiples paquetes.

Utilizamos paquetes para poder reutilizar el código de otras personas, no tenemos que escribir todo desde cero. Tal vez queremos una calculadora como parte de nuestro programa, probablemente podrías encontrar un paquete Go existente que contenga las funciones matemáticas que podrías importar a tu código ahorrando mucho tiempo. La belleza de compartir el código y el Open Source.

Go te anima a organizar tu código en paquetes para que sea fácil reutilizar y mantener el código fuente.

### Hello #90DaysOfDevOps línea a línea

Ahora echemos un vistazo a nuestro archivo main.go de Hello #90DaysOfDevOps recorriendo cada línea.

![](Images/Day9_Go2.png)

En la primera línea tienes `package main`. Significa que este fichero pertenece a un paquete llamado main. Todos los ficheros .go deben pertenecer a un paquete, también deben tener `package <lo-que-sea>` en la línea inicial.

Un paquete puede llamarse como quieras. En este caso tenemos que llamarlo `main` ya que es el punto de partida del programa que va a estar en este paquete, esto es una regla. 

![](Images/Day9_Go3.png)

Siempre que queramos compilar y ejecutar nuestro código tenemos que decirle a la máquina dónde tiene que empezar la ejecución. Lo hacemos escribiendo una función llamada main. La máquina buscará una función llamada main para encontrar el punto de entrada del programa.

Una función es un bloque de código que puede hacer alguna tarea específica y puede ser utilizada en todo el programa.

Puedes declarar una función con cualquier nombre usando `func` pero en este caso, necesitamos nombrarla `main` porque, como hemos comentado, es donde comienza el código.

![](Images/Day9_Go4.png)

A continuación, vamos a ver la línea 3 de nuestro código, la importación, esto significa que quieres traer otro paquete a tu programa principal. fmt es un paquete estándar que se está proporcionado por Go, este paquete contiene la función `Println()` y porque importamos esto podemos usarla en la línea 6. Hay varios paquetes estándar que puedes incluir en tu programa y aprovecharlos o reutilizarlos en tu código ahorrando la molestia de tener que escribir desde cero. [Go Standard Library](https://pkg.go.dev/std)

![](Images/Day9_Go5.png)

El `Println()` que tenemos aquí es una forma de escribir la salida estándar en la terminal donde se ha ejecutado el ejecutable con éxito. Siéntete libre de cambiar el mensaje entre los paréntesis.

![](Images/Day9_Go6.png)

### TLDR

- **Línea 1** = Este archivo estará en el paquete llamado `main` y es necesario que se llame `main` porque incluye el punto de entrada del programa.
- **Línea 3** = Para que podamos utilizar el `Println()` tenemos que importar el paquete fmt para utilizarlo en la línea 6.
- **Línea 5** = El punto de entrada real, es la función `main`.
- **Línea 6** = Esto nos permitirá imprimir "Hello #90DaysOfDevOps" en nuestro terminal.

## Recursos

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

Nos vemos en el [Día 10](day10.md).
