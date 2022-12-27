## Automatizar tareas con scripts bash

El shell que vamos a ver hoy es el bash pero cubriremos otro shell mañana: ZSH.

BASH - **B**ourne **A**gain **Sh**ell

Casi podríamos dedicar una sección entera de 7 días a los scripts de shell al igual que los lenguajes de programación, bash nos da la capacidad de trabajar junto a otras herramientas de automatización para hacer las cosas.

Mucha gente ha llegado a configurar scripts de shell realmente complejos para hacer que suceda algo. Algunas cosas muy importantes en los negocios dependen de scripts. No necesitaremos entender al completo el shell/bash scripting para este propósito, que nunca está de más, pero este no es el camino que cogeremos. Lo que sí que haremos es aprender lo básico de shell/bash scripting para trabajar junto a nuestras herramientas de automatización y para tareas ad-hoc.

Un ejemplo de esto podría ser el VAGRANTFILE que usamos para crear nuestra MV. Podríamos envolver esto en un simple script bash que borrara y renovara esto cada lunes por la mañana para que tengamos una copia fresca de nuestra MV Linux cada semana. También podríamos añadir toda la pila de software que necesitamos en dicha máquina Linux y así hacerlo todo a través de un único script bash, una única orden.

Otra cosa es que las preguntas sobre scripts prácticos son cada vez más frecuentes en las entrevistas para los puestos relacionados con la Administración de sistemas, entornos cloud y filosofía DevOps.

Vamos a darle duro que tenemos mucho por delante en el día de hoy.
### Empecemos

Como con otros conceptos que estamos cubriendo en estos 90 días, la única manera real de aprender es prácticando. No os quedéis tan solo leyendo estas lineas y dando el capítulo como hecho. La experiencia práctica va a ayudar a empapar todo esto en tu memoria muscular, aunque tardes algunos días más merece la pena que practiques todo lo que puedas.

Para empezar vamos a necesitar un editor de texto. En el [Día 17](day17.md) cubrimos los dos editores de texto más comunes y un poco sobre cómo usarlos. (Espero que hayas visto más por tu cuenta ;)

Vamos a ponernos manos a la obra y a crear nuestro primer script de shell. Así, en frío, creamos un archivo.

```shell
touch 90DaysOfDevOps.sh
```

Y lo abrimos con el editor nano (El preferido de los valencianos).

```shell
nano 90DaysOfDevOps.sh
``` 

Esto abrirá nuestro nuevo shell script en blanco en nano. De nuevo, puedes elegir el editor de texto que prefieras.

La primera línea de todos los scripts bash deberá ser algo así 

```bash
#!/usr/bin/bash
``` 
Esta es la ruta a tu binario bash. Esta línea conocida como `SHEBANG`, `HASHBANG` o `SHARPBANG`, conseguiréis llamar o indicar que intérprete de comandos o consola de sistema, debe ejecutar o interpretar este script.

Sin embargo, deberías comprobarlo en el terminal ejecutando 

```shell
which bash
``` 

Este comando nos permitirá ver la ubicación del binario.

Si no estás usando Ubuntu y sigues perdiendo el tiempo con Sistemas Operativos privativos, entonces también podrías intentar 

```shell
whereis bash
``` 

Sin embargo, puede ver otras rutas listadas en scripts de shell ya creados que podrían incluir:

- `#!/bin/bash`
- `#!/usr/bin/env bash`

En la siguiente línea de nuestro script, lo ideal es añadir un comentario explicando el propósito del script o al menos alguna información para quien venga detrás. Que puedes ser tú mismo en un par de semanas, así que explicatelo para tí, pero como si tuvieses alzheimer. Los comentarios en los scripts son con almohadilla: `#`. Esto nos permite comentar líneas particulares en nuestro fichero para que no afecten al código. Se utiliza para proporcionar descripciones de los comandos, del flujo del código, información del autor, mención a fuentes, etc. Como comenté antes, es indispensable si se comparte esto.

figlet, el programa que instalamos en días anteriores, sirve para crear arte asci, el cual queda muy vistoso para iniciar scripts.

![](Images/Day19_Linux1.png)

Todos los comandos que repasamos en el [Día 15](día15.md) podrían usarse en los scripts como un simple comando para probar.

Añadamos un simple bloque de código a nuestro script.

```bash
mkdir 90DaysOfDevOps
cd 90DaysOfDevOps
touch Day19
ls
```

Luego puedes guardar esto y salir de tu editor de texto, si ejecutamos el script con 

```shell
./90DaysOfDevOps.sh
``` 

deberías obtener un mensaje de permiso denegado. Puedes comprobar los permisos de este fichero usando el comando

```shell
ls -al
```
y podrás ver resaltado que no tenemos derechos ejecutables sobre este fichero.

![](Images/Day19_Linux2.png)

We can change this using `chmod +x 90DaysOfDevOps.sh` and then you will see the `x` meaning we can now execute our script.

![](Images/Day19_Linux3.png)

Podemos cambiar esto usando 

```shell
chmod +x 90DaysOfDevOps.sh
``` 
y entonces verás la `x` que significa que ahora podemos ejecutar nuestro script.

![](Images/Day19_Linux4.png)

Bastante básico, pero puedes empezar a ver cómo esto lo puedes utilizar para llamar a otras herramientas para hacer tu vida más fácil y automatizar las cosas.

### Variables, Condicionales

Gran parte de esta sección es una repetición de lo que vimos aprendiendo [Golang](day11.md), pero vale la pena que lo repasemos de nuevo para asentar conocimientos.

#### Variables

Las variables nos permiten definir un término particular que se utiliza varias veces a lo largo de un script. Sirve para simplificar el código que puede ser potencialmente complejo.

Es muy simple declarar una variable, añádela con una línea limpia en el script.

```bash
challenge="90DaysOfDevOps"
```

De esta forma, donde usemos `$challenge` estaremos dando el valor de la variable y si cambiamos este valor de la variable se reflejará en todos los lugares donde la hayamos usado.

![](Images/Day19_Linux5.png)

Si ahora ejecutamos nuestro script `sh` veremos la impresión que se ha añadido a nuestro script.

![](Images/Day19_Linux5.png)

También podemos pedir un input del usuario para que establezca el valor de la variables utilizando lo siguiente:

```bash
echo "Enter your name"
read name
```

Esto definiría la variable `$name`. Con lo que podemos utilizar este valor más adelante.

#### Condicionales

Para saber a quién tenemos en nuestro reto y cuántos días han completado, podemos ayudarnos con los condicionales 
- `if`
- `if-else`
- `else-if`, 

Atentos al siguiente código.

```bash
#!/bin/bash
#   ___   ___  ____                   ___   __ ____              ___
#  / _ \ / _ \|  _ \  __ _ _   _ ___ / _ \ / _|  _ \  _____   __/ _ \ _ __  ___
# | (_) | | | | | | |/ _` | | | / __| | | | |_| | | |/ _ \ \ / / | | | '_ \/ __|
#  \__, | |_| | |_| | (_| | |_| \__ \ |_| |  _| |_| |  __/\ V /| |_| | |_) \__ \
#    /_/ \___/|____/ \__,_|\__, |___/\___/|_| |____/ \___| \_/  \___/| .__/|___/
#                           |___/                                    |_|
#
# This script is to demonstrate bash scripting!
#
# Variables to be defined

ChallengeName=#90DaysOfDevOps
TotalDays=90

# User Input

echo "Enter Your Name"
read name
echo "Welcome $name to $ChallengeName"
echo "How Many Days of the $ChallengeName challenge have you completed?"
read DaysCompleted

if [ $DaysCompleted -eq 90 ]
then
  echo "You have finished, well done"
elif [ $DaysCompleted -lt 90 ]
then
  echo "Keep going you are doing great"
else
  echo "You have entered the wrong amount of days"
fi
```

Puedes ver que estamos realizando algunas comparaciones o comprobando valores entre sí. Aquí tenemos diferentes opciones que vale la pena destacar.

- `eq` - si los dos valores son iguales devolverá TRUE
- `ne` - si los dos valores no son iguales devolverá TRUE
- `gt` - si el primer valor es mayor que el segundo valor devolverá TRUE
- `ge` - si el primer valor es mayor o igual que el segundo valor devolverá TRUE
- `lt` - si el primer valor es menor que el segundo valor devolverá TRUE
- `le` - si el primer valor es menor o igual que el segundo valor devolverá TRUE

También podemos usar scripts bash para determinar información sobre archivos y carpetas, esto se conoce como condiciones de archivo.

- `-d file` TRUE si el fichero es un directorio
- `-e file` TRUE si el fichero existe
- `-f file` TRUE si la cadena proporcionada es un fichero
- `-g file` TRUE si el id de grupo está establecido en un fichero
- `-r file` TRUE si el fichero es legible
- `-s file` TRUE si el fichero tiene un tamaño distinto de cero

```bash
FILE="90DaysOfDevOps.txt"
if [ -f "$FILE" ]
then
  echo "$FILE is a file"
else
  echo "$FILE is not a file"
fi
```

![](Images/Day19_Linux7.png)

Siempre que tengamos ese archivo todavía en nuestro directorio deberíamos obtener el primer comando echo. Pero si eliminamos ese archivo entonces deberíamos obtener el segundo comando echo del script.

![](Images/Day19_Linux8.png)

Esto puede ser utilizado para ahorrar mucho tiempo al buscar en un sistema elementos específicos.

Este repositorio tiene lo que parece ser una interminable cantidad de scripts que puedes probar: [DevOps Bash Tools](https://github.com/HariSekhon/DevOps-Bash-tools/blob/master/README.md)

#### Ejemplo

*Escenario*: 

Tenemos nuestra empresa llamada "90DaysOfDevOps" funcionando desde hace un tiempo y ahora queremos ampliar el equipo de 1 persona a muchas más. Yo soy el único hasta ahora que conoce el proceso de incorporación por lo que queremos reducir ese cuello de botella mediante la automatización de algunas tareas.

*Requisitos*:

- Un usuario puede ser pasado como argumento en la línea de comandos.
- Se crea un usuario con el nombre del argumento.
- Se puede pasar una contraseña como argumento en la línea de comandos.
- Se establece la contraseña para el usuario.
- Se muestra un mensaje de éxito en la creación de cuenta.

Comencemos creando nuestro script de shell con 

```shell
touch create_user.sh
```

Antes de continuar hagámoslo ejecutable usando

```shell
chmod +x create_user.sh
```

Para empezar a editar nuestro script para el escenario planteado podemos usar

```shell
nano create_user.sh
```

O también vim, según cuál sea nuestro editor preferido.

```shell
vim create_user.sh
```

Podemos echar un vistazo al primer requisito "Un usuario puede ser pasado como un argumento en la línea de comandos". Podemos usar lo siguiente

```bash
#! /usr/bin/bash

# Un usuario puede ser pasado como un argumento en la línea de comandos
echo "$1"
```

![](Images/Day19_Linux9.png)

Ejecuta este script utilizando

```shell
./create_user.sh Michael
```

Veremos como sustituye el valor de la variable impresa con `echo` por el argumento que le hemos dado.

![](Images/Day19_Linux10.png)

A continuación podemos tomar ese segundo requisito "Se crea un usuario con el nombre del argumento" esto se puede hacer con el comando `useradd`. La opción `-m` es para crear el directorio home del usuario como /home/nombredeusuario

```bash
#! /usr/bin/bash

# Un usuario puede ser pasado como un argumento en la línea de comandos
echo "$1 user account being created."

# Se crea un usuario con el nombre del argumento.
sudo useradd -m "$1"
```

⚠️ *Cuidado: Si no proporcionas un nombre de cuenta de usuario entonces dará error ya que la variable `$1` no tiene valor.*

Podemos comprobar que esta cuenta ha sido creada con el comando

```shell
awk -F: '{ print $1}' /etc/passwd
```

![](Images/Day19_Linux11.png)

Nuestros siguientes requisitos son "Se puede pasar una contraseña como argumento en la línea de comandos" y "Se establece la contraseña para el usuario". 

⚠️ *Cuidado: Nunca vamos a hacer esto en producción, las contraseñas nunca deben pasar por la línea de comando ni por un script ya que quedan expuestas y puede ser un vector de ataque. Siempre pasarán por un proceso previo de codificación y encriptación. Esto solo lo hacemos para la practica de añadir argumentos que gestione el script con las variables.* 

```bash
#! /usr/bin/bash

# Un usuario puede ser pasado como un argumento en la línea de comandos
echo "$1 user account being created."

# Se crea un usuario con el nombre del argumento.
sudo useradd -m "$1"

# Se puede pasar una contraseña como argumento en la línea de comandos.
# Se establece la contraseña para el usuario.
sudo chpasswd <<< "$1":"$2"
```

Si luego ejecutamos este script con los dos parámetros 

```shell
./create_user.sh 90DaysOfDevOps password
```

Podéis ver en la siguiente imagen que creó nuestro usuario y contraseña. Luego entramos a ese usuario y confirmamos quien era el usuario activo con `whoami`.

![](Images/Day19_Linux12.png)

El requisito final es "Se muestra un mensaje de éxito en la creación de cuenta." ya lo tenemos en la línea superior del código y podemos ver en la captura de pantalla anterior que muestra `90DaysOfDevOps user account being created`. Esto se quedó de las pruebas con el parámetro `$1`.

Ahora, este script puede ser utilizado para integrar y configurar rápidamente nuevos usuarios en nuestros sistemas Linux. Pero para que quede más vistoso y lo pueda utilizar cualquiera, podemos añadir la entrada de usuario y password mediante lo que vimos anteriormente. Mira como queda el código de bonico.

```bash
#! /usr/bin/bash

echo "What is your intended username?"
read  username
echo "What is your password"
read  password

# Un usuario puede ser pasado como un argumento en la línea de comandos
echo "$1 user account being created."

# Se crea un usuario con el nombre del argumento.
sudo useradd -m "$1"

# Se puede pasar una contraseña como argumento en la línea de comandos.
# Se establece la contraseña para el usuario.
sudo chpasswd <<< "$1":"$2"

# Se muestra un mensaje de éxito en la creación de cuenta.
echo "The Account for $username has successfully been created"

```

Así los pasos son más interactivos,

![](Images/Day19_Linux14.png)

Para terminar, podemos dar un output de éxito para confirmar que la nueva cuenta de usuario se ha creado.

![](Images/Day19_Linux15.png)

Un detalle más sería ocultar la entrada de la contraseña con el flag `-s` en la línea de código 

```bash
read -s password
```

![](Images/Day19_Linux16.png)

Además, podríamos borrar el usuario creado con

```bash
sudo userdel test_user
```
Podéis ver el [Script completo](Linux/create-user.sh) con los comentarios y las impresiones en pantalla traducidas al castellano.

Una vez más, esto no es algo que vayas a utilizar en el día a día, pero forma parte de la base necesaria para cuando lo necesites. Además, pone de manifiesto la flexibilidad que proporciona el shell scripting. 

Piensa en las tareas repetitivas que haces todos los días y cómo podrías automatizarlas, la primera opción probablemente sea usar un script bash,pero existen territorios complejos, así que a practicar, practicar y practicar.

Para último ejemplo, echad un vistazo a este simple bash que ayuda a levantar un clúster Kubernetes usando minikube en una máquina local, junto con los servicios de datos y Kasten K10 para ayudar a demostrar los requisitos y necesidades en torno a la gestión de datos: [Project Pace](https://github.com/MichaelCade/project_pace/blob/main/singlecluster_demo.sh) Quizá sea muy avanazado ya que no hemos visto Kubernetes todavía, pero puede ayudar a ver posibilidades del scripting.

## Recursos

- [Bash in 100 seconds](https://www.youtube.com/watch?v=I4EWvMFj37g)
- [Bash script with practical examples - Full Course](https://www.youtube.com/watch?v=TPRSJbtfK4M)
- [Client SSH GUI - Remmina](https://remmina.org/)
- [The Beginner's guide to SSH](https://www.youtube.com/watch?v=2QXkrLVsRmk)
- [Vim in 100 Seconds](https://www.youtube.com/watch?v=-txKSRn0qeA)
- [Vim tutorial](https://www.youtube.com/watch?v=IiwGbcd8S7I)
- [Learn the Linux Fundamentals - Part 1](https://www.youtube.com/watch?v=kPylihJRG70)
- [Linux for hackers (don't worry you don't need to be a hacker!)](https://www.youtube.com/watch?v=VbEx7B_PTOE)
- [Webminal](https://www.webminal.org/) 
- [Guía básica de Vim](https://gitea.vergaracarmona.es/man-linux/Guia-VIM)

Nos vemos en el [Día 20](day20.md).
