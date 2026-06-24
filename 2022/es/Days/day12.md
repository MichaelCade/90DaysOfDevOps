## Obtención del input del usuario con Punteros y un programa terminado

Ayer, en el ([Día 11](day11.md)), creamos nuestro primer programa Go self-contained, creando el código para obtener el input del usuario como variables dentro de nuestro código y se les dio valores. Ahora pediremos al usuario su input para dar el valor a la variable del mensaje final.

¡Qué emocionante!

## Obteniendo el input del usuario

Antes de hacer esto vamos a echar un vistazo a nuestra aplicación para recorrer las variables que queremos como prueba antes de obtener la entrada del usuario.

Ayer terminamos con nuestro código con el aspecto de este fichero: [day11_example4.go](Go/day11_example4.go).

Hemos definido manualmente nuestras variables y constantes `challenge, daystotal, dayscomplete`.

Ahora vamos a añadir una nueva variable llamada `TwitterName`. Puedes ver como queda en el código en el fichero [day12_example1.go](Go/day12_example1.go) y si ejecutamos este código esta es nuestra salida.

![](Images/Day12_Go1.png)

Estamos en el día 12 y tendríamos que cambiar ese `dayscomplete` cada día y compilar nuestro código cada día si esto fuera hardcodeado, lo cual no suena muy atractivo.

Con lo cual, queremos que el usuario escriba los valores de su nombre en Twitter y del número de días completados para obtenerlos en nuestro programa. Para ello podemos utilizar otra función input del paquete `fmt`.

Recapitulando sobre el paquete `fmt`, diferentes funciones para la entrada y salida formateada (I/O)

- Imprimir mensajes
- Recoger el input del usuario
- Escribir en un archivo

En lugar de asignar en el programa el valor de una variable, queremos pedírsela al usuario en un input.

```go
fmt.Scan(&TwitterName)
```

Fíjate que también usamos `&` antes de la variable. Esto se conoce como puntero y lo veremos en la siguiente sección.

En el código del fichero [day12_example2.go](Go/day12_example2.go) puedes ver que estamos pidiendo al usuario que introduzca la información mencionada en las variables `TwitterName` y `DaysCompleted`.

Ahora vamos a ejecutar nuestro programa y verás que tenemos la entrada de ambas variables.

![](Images/Day12_Go2.png)

Vale, está muy bien, hemos conseguido un input del usuario y hemos imprimido un mensaje con los valores obtenidos. pero ¿qué pasa si conseguimos que nuestro programa nos diga cuántos días nos quedan en nuestro reto?

Para ello, vamos a crear una variable llamada `remainingDays` y le daremos el valor integer "90". A continuación, tenemos que cambiar este valor para imprimir los días restantes cuando recibamos la entrada del usuario de "días completados", es decir, hacer una resta.

```go
remainingDays = remainingDays - DaysCompleted
```

Puedes ver el aspecto de nuestro programa terminado en el fichero [day12_example2.go](Go/day12_example3.go).

Si ahora ejecutamos este programa puedes ver que se hace un simple cálculo basado en la entrada del usuario y el valor de la variable `remainingDays`.

![](Images/Day12_Go3.png)

## ¿Qué es un puntero? (Variables especiales)

Un puntero es una variable (especial) que apunta a la dirección de memoria de otra variable.

Puedes ver una explicación más detallada en [geeksforgeeks](https://www.geeksforgeeks.org/pointers-in-golang/).

Ahora, simplifiquemos nuestro código y mostremos con y sin el `&` delante de uno de nuestros comandos de impresión, esto nos da la dirección de memoria del puntero. El ejemplo del código lo puedes encontrar en el fichero [day12_example4.go](Go/day12_example4.go).

A continuación ejecuta este código y observa.

![](Images/Day12_Go4.png)

## Recursos

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

Nos vemos en el [Día 13](day13.md).
