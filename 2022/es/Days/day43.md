## Qué es Docker y cómo instalarlo

En el post anterior, he mencionado Docker al menos una vez y eso es porque Docker es el culpable en gran parte de la popularización de los contenedores, a pesar de que han existido desde hace mucho tiempo.

Vamos a utilizar y explicar un poco Docker, pero también debemos mencionar la [Open Container Initiative (OCI)](https://www.opencontainers.org/), que es una organización de estándares de la industria que fomenta la innovación, evitando el peligro de la dependencia del proveedor. Gracias a la OCI, tenemos donde elegir entre las herramientas de contenedores, incluyendo Docker, [CRI-O](https://cri-o.io/), [Podman](http://podman.io/), [LXC](https://linuxcontainers.org/), entre otros.

Docker es un framework de software para construir, ejecutar y gestionar contenedores. El término "docker" puede referirse tanto a las herramientas (los comandos y un daemon) como al formato de archivo Dockerfile.

Utilizaremos Docker Personal, que es gratuito (para la educación y el aprendizaje). Esto incluye todo lo esencial que necesitamos para obtener una buena base de conocimientos de los contenedores y herramientas.

Probablemente vale la pena desglosar algunas de las herramientas "docker" para saber por qué se utilizan. El término docker puede referirse al proyecto docker en general, que es una plataforma para que los desarrolladores y administradores desarrollen, envíen y ejecuten aplicaciones. También puede ser una referencia al daemon Docker que se ejecuta en el host que gestiona las imágenes y contenedores también llamado Docker Engine.

### Docker Engine

Docker Engine es una tecnología de contenerización de código abierto para construir y contenerizar tus aplicaciones. Docker Engine actúa como una aplicación cliente-servidor con:

- Un servidor con un proceso demonio de larga ejecución dockerd.
- Las API especifican interfaces que los programas pueden utilizar para hablar con el demonio Docker y darle instrucciones.
- Una interfaz de línea de comandos (CLI) docker cliente.

Lo anterior fue tomado de la documentación oficial de Docker y el específico [Docker Engine Overview](https://docs.docker.com/engine/)

### Docker Desktop

Disponemos de un Docker Desktop para sistemas Windows y macOS. Un entorno de desarrollo docker ligero y fácil de instalar. Una aplicación nativa del sistema operativo que aprovecha las capacidades de virtualización en el sistema operativo anfitrión. Nadie lo utiliza, pero es una buena opción para comenzar.

Puedes construir, depurar, probar, empaquetar y enviar aplicaciones Dockerizadas en Windows o macOS.

En Windows, también podemos aprovechar WSL2 y Microsoft Hyper-V. Veremos algunas de las ventajas de WSL2 a medida que avancemos, ya que como Windows no llega lo mejor es utilizar las opciones Linux.

Debido a la difícil integración con las capacidades del hipervisor en el sistema operativo anfitrión, docker proporciona la capacidad de ejecutar sus contenedores con sistemas operativos Linux que no dan tantos problemas.

### Docker Compose

Docker compose es una herramienta que te permite ejecutar aplicaciones más complejas sobre múltiples contenedores. Con la ventaja de ser capaz de utilizar un único archivo y comando para hacer girar su aplicación. Es el principio de la orquestación de contenedores, un concepto que tendremos que interiorizar para dejar de tratar a nuestras máquinas como mascotas y comenzar a verlas como ganado.

### Docker Hub

Un recurso centralizado para trabajar con Docker y sus componentes, un registro de imágenes que nos facilitará mucho la existencia. Hay un montón de servicios adicionales que pueden ser utilizados en parte con la automatización o integrados en GitHub, así como el escaneo de seguridad.

### Dockerfile

Un dockerfile es un archivo de texto que contiene comandos que también se pueden ejecutar manualmente para construir una imagen docker. Docker puede construir imágenes automáticamente leyendo las instrucciones que tenemos en nuestro dockerfile gracias al lenguaje declarativo.

## Instalando Docker Desktop

La [documenación de Docker](https://docs.docker.com/engine/install/) es asombrosa y si quieres aprender de verdad deberías echarle un vistazo y tenerla como referencia. Vamos a utilizar Docker Desktop en Windows con WSL2. Ya he ejecutado la instalación en la máquina que estamos utilizando.

![](Images/Day43_Containers1.png)

Toma nota antes de seguir adelante en los requisitos del sistema antes de [Instalar Docker Desktop en Windows](https://docs.docker.com/desktop/windows/install/). Si estás usando macOS, inclusive para la arquitectura de CPU basada en M1, puedes echar un vistazo a [Instalar Docker Desktop en macOS](https://docs.docker.com/desktop/mac/install/). Linux también puede ejecutar Docker Desktop virtualizado: https://docs.docker.com/desktop/install/linux-install/

Voy a ejecutar a través de la instalación de Docker Desktop para Windows en otra máquina Windows y registrar el proceso de abajo.

## Recursos

- [TechWorld with Nana - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
- [Programming with Mosh - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=pTFZFxd4hOI)
- [Docker Tutorial for Beginners - What is Docker? Introduction to Containers](https://www.youtube.com/watch?v=17Bl31rlnRM&list=WL&index=128&t=61s)
- [WSL 2 with Docker getting started](https://www.youtube.com/watch?v=5RQbdMn04Oc)
- [En español] En los [apuntes](https://vergaracarmona.es/apuntes/) del traductor:
  - [Preparación de entorno de pruebas local para docker](https://vergaracarmona.es/preparacion-de-entorno-de-pruebas-local-para-docker/)
  - [Uso básico de docker](https://vergaracarmona.es/uso-basico-de-docker/)
  - [Una breve historia sobre contenedores](https://vergaracarmona.es/breve-historia-de-contenedores/)
  - [Desplegar con docker-compose los servicios Traefik y Portainer](https://vergaracarmona.es/desplegar-con-docker-compose-los-servicios-traefik-y-portainer/)

Nos vemos en el [Día 44](day44.md)
