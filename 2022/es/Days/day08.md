## Configurar tu entorno DevOps para Go y Hello World

Antes de entrar en algunos de los fundamentos de Go debemos instalar Go en nuestra estación de trabajo y hacer lo que cada módulo de "aprendizaje de programación 101" nos enseña, que es crear la aplicación Hello World. 

Como esto va a ser caminar a través de los pasos para conseguir la instalación de Go vamos a tratar de documentar el proceso en imágenes para facilitar el seguimiento del proceso. _(Las capturas de pantalla pueden estar desactualizadas aun siendo tomadas en la última versión existente en el momento de escribir estas líneas)_.

En primer lugar, vamos a dirigirnos a [go.dev/dl](https://go.dev/dl/) y nos encontraremos con algunas opciones de descarga disponibles para distintos Sistemas Operativos.

![](Images/Day8_Go1.png)

Probablemente sepas qué sistema operativo tiene el equipo donde quiere efectuar la instalación, así que selecciona la descarga apropiada y luego podemos empezar a instalar. Estoy usando Windows para este tutorial, básicamente podemos dejar todos los valores predeterminados en el asistente de instalación. 

![](Images/Day8_Go2.png)

También tenga en cuenta que si tienes una versión anterior de Go instalada tendrás que eliminarla antes de instalar, Windows tiene esta opción en el asistente de instalación, para eliminar e instalar de una.

Una vez que haya terminado, debe abrir un prompt del sistema/terminal (CMD o powershell en Windows) y para comprobar que la instalación de Go se ha efectuado podemos comprobar la versión con:

```shell
go version
```

Si no obtienes la salida que vemos a continuación en la imagen, entonces Go no estará instalado y tendrás que volver sobre tus pasos para encontrar el error.

![](Images/Day8_Go3.png)

A continuación comprobaremos nuestro entorno para Go. Siempre es bueno comprobar que tus directorios de trabajo están configuradas correctamente, como puedes ver a continuación necesitamos asegurarnos que tenemos el siguiente directorio en el sistema.

![](Images/Day8_Go4.png)

¿Lo has comprobado? ¿Estás siguiendo la pista? Probablemente obtendrás algo como lo que se muestra a continuación si tratas de navegar allí. (Desde Windows)

![](Images/Day8_Go5.png)

En Linux queda creado el directorio con:

```shell
sudo apt install golang
```
Teniendo esta estructura:
```shell
#input
tree go -L 3

#output
go
└── pkg
    └── mod
        ├── cache
        ├── cloud.google.com
        ├── github.com
        ├── go.etcd.io
        ├── golang.org
        ├── google.golang.org
        ├── gopkg.in
        ├── go.starlark.net@v0.0.0-20200306205701-8dd3e2ee1dd5
        ├── helm.sh
        ├── k8s.io
        ├── oras.land
        └── sigs.k8s.io
```

Ok, en Windows debemos crear ese directorio. Se puede hacer fácilmente con el comando mkdir en PowerShell. También tenemos que crear 3 carpetas dentro de la carpeta Go como verás también a continuación.

```shell
mkdir go

cd go

mkdir pkg,bin,src
```

![](Images/Day8_Go6.png)

Ya tenemos Go instalado y nuestro directorio de trabajo listo para la acción. Ahora necesitamos un entorno de desarrollo integrado (IDE). Hay muchos disponibles que puedes usar pero el más común y el que el autor usa es Visual Studio Code o Code. (El traductor también utiliza vscode pero en equipos remotos es mejor con [Vim](https://gitea.vergaracarmona.es/man-linux/Guia-VIM)) Puedes aprender más sobre los IDEs [aquí](https://www.youtube.com/watch?v=vUn5akOlFXQ).

Si no has descargado e instalado VSCode en tu equipo, puedes hacerlo [aquí](https://code.visualstudio.com/download). Como podrás ver tendrás opciones para los distintos Sistemas Operativos.

![](Images/Day8_Go7.png)

Al igual que con la instalación de Go, vamos a descargar e instalar y mantener los valores predeterminados. Una vez completado puede abrir VSCode, seleccionar Abrir Archivo y navegar a nuestro directorio Go que hemos creado anteriormente.

![](Images/Day8_Go8.png)

Puede que te aparezca un popup sobre la seguridad, léela y luego si lo consideras dale a Sí para confiar en los autores. (En este tutorial no nos hacemos responsables más tarde si empiezas a abrir cosas en las que no deberías confiar).

Ahora deberías ver las tres carpetas que creamos anteriormente. Hacemos clic derecho en la carpeta src y creamos una nueva carpeta llamada `Hello`.

![](Images/Day8_Go9.png)

Hasta aquí todo muy fácil, diría que facilísimo. Ahora vamos a crear nuestro primer programa Go sin entender nada de lo que pongamos en esta siguiente fase.

Creamos un archivo llamado `main.go` dentro de la carpeta `Hello`. Tan pronto como pulses enter en el main.go se te preguntará si quieres instalar la extensión Go y también los paquetes, que puedes comprobar en el archivo pkg vacío que hicimos unos pasos atrás y notar que deberíamos tener algunos paquetes nuevos ahí dentro. ¿Correcto?

![](Images/Day8_Go10.png)

Ahora vamos a poner en marcha esta aplicación `Hola Mundo`, copia el siguiente código en tu nuevo archivo main.go y guárdalo.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello #90DaysOfDevOps")
}
```

Es de entender que el anterior código puede no tener ningún sentido, pero veremos algo más sobre funciones, paquetes, entre otras cosas, en próximos días. Por ahora, vamos a ejecutar nuestra aplicación. 

De vuelta en la terminal y en nuestra carpeta `Hello` podemos comprobar que todo está funcionando. Usando el siguiente comando podemos comprobar si nuestro programa de aprendizaje genérico está funcionando.

```shell
go run main.go
```
Nos dirá Hola.

![](Images/Day8_Go11.png)

Pero esto no termina aquí, ¿qué pasa si ahora queremos tomar nuestro programa y ejecutarlo en otras máquinas Windows? Podemos hacerlo construyendo nuestro binario con el siguiente comando

```shell
go build main.go
```

![](Images/Day8_Go12.png)

Si ejecutamos esto, veríamos la misma salida:

```shell
#input
./main.exe

#output
Hello #90DaysOfDevOps
```

## Recursos

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

Nos vemos en el [Día 9](day09.md).

![](Images/Day8_Go13.png)
