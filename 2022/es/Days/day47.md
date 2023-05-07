## Docker Networking & Security

Hasta ahora hemos visto como desplegar contenedores, pero no sabemos aún cómo funcionan las redes y cuál es la seguridad mínima que necesitan. Cómo bien anuncia el título, este es el plan de hoy.

### Docker Networking Basics

Vamos al grano, abre un terminal y escribe el comando `docker network`. Este es el comando principal para configurar y gestionar redes de los contenedores.

Puedes ver como podemos usar el comando y todos los sub-comandos disponibles. Podemos crear nuevas redes, listar las existentes, inspeccionarlas y eliminarlas.

![](Images/Day47_Containers1.png)

Vamos a echar un vistazo a las redes existentes que tenemos desde nuestra instalación usando el comando `docker network list`.

Cada red tiene un ID y un NOMBRE únicos. Cada red también está asociada a un único controlador. Observa que la red "bridge" y la red "host" tienen el mismo nombre que sus respectivos controladores.

![](Images/Day47_Containers2.png)

A continuación, podemos echar un vistazo más profundo a nuestras redes con el comando `docker network inspect`.

Si ejecuto `docker network inspect bridge` puedo obtener todos los detalles de configuración de ese nombre de red específico. Esto incluye nombre, ID, controladores, contenedores conectados y mucho más como puedes ver.

![](Images/Day47_Containers3.png)

### Docker: Bridge Networking

Como has visto arriba una instalación estándar de Docker Desktop nos da una red pre-construida llamada `bridge` Si vuelves a mirar el comando `docker network list`, verás que la red llamada bridge está asociada con el driver `bridge`. Que tengan el mismo nombre no quiere decir que sean la misma cosa. Están conectadas pero no son lo mismo.

La salida anterior también muestra que la red bridge tiene un ámbito local. Esto significa que la red sólo existe en este host Docker. Esto es cierto para todas las redes que utilizan el controlador bridge que proporciona la red de un solo host.

Todas las redes creadas con el controlador bridge se basan en un bridge Linux (también conocido como conmutador virtual).

### Conectar un Contenedor

Por defecto la red bridge es asignada a los nuevos contenedores, lo que significa que a menos que especifiques una red todos los contenedores serán conectados a la red bridge.

Vamos a crear un nuevo contenedor con el comando `docker run -dt ubuntu sleep infinity`.

El comando sleep de arriba sólo va a mantener el contenedor funcionando en segundo plano para que podamos trastear con él.

![](Images/Day47_Containers4.png)

Si a continuación comprobamos nuestra red bridge con `docker network inspect bridge` veremos que tenemos un contenedor que coincide con lo que acabamos de desplegar porque no hemos especificado ninguna red.

![](Images/Day47_Containers5.png)

También podemos bucear en el contenedor usando `docker exec -it 3a99af449ca2 bash`, tendrás que usar `docker ps` para obtener el ID del contenedor que has creado.

En principio, nuestra imagen no tiene nada con lo que hacer ping, así que tenemos que ejecutar el siguiente comando: `apt-get update && apt-get install -y iputils-ping`. Luego vamos a hacer ping a una dirección de interfaz externa, a ver que pasa: `ping -c5 www.90daysofdevops.com`

![](Images/Day47_Containers6.png)

Ahora podemos parar el contenedor con `docker stop 3a99af449ca2` y utilizar `docker ps` para buscar su ID de contenedor, veremops que esto eliminará nuestro contenedor.

### Configurar NAT para conectividad externa

En este paso, iniciaremos un nuevo contenedor NGINX y asignaremos el puerto 8080 en el host Docker al puerto 80 dentro del contenedor. Esto significa que el tráfico que llega al host Docker en el puerto 8080 pasará al puerto 80 dentro del contenedor.

Inicie un nuevo contenedor basado en la imagen oficial de NGINX ejecutando `docker run --name web1 -d -p 8080:80 nginx`.

![](Images/Day47_Containers7.png)

Revise el estado del contenedor y las asignaciones de puertos ejecutando `docker ps`.

![](Images/Day47_Containers8.png)

La línea superior muestra el nuevo contenedor web1 ejecutando NGINX. Fíjate en el output del contenedor en la asignación de puertos - `0.0.0.0:8080->80/tcp`. Asigna el puerto 8080 en todas las interfaces del host al puerto 80 dentro del contenedor web1. Esta asignación de puertos es lo que hace que el servicio web del contenedor sea accesible desde fuentes externas (a través de la dirección IP de los hosts Docker en el puerto 8080).

Ahora necesitamos nuestra dirección IP para nuestro host anfitrión, podemos hacer esto yendo a nuestro terminal WSL y utilizando el comando `ip addr`.

![](Images/Day47_Containers9.png)

Entonces podemos tomar esta IP y abrir un navegador y dirigirnos a `http://172.25.218.154:8080/` (Ten en cuenta que tiene que se tu IP, seguramente sea diferente). Esto confirma que NGINX es accesible.

![](Images/Day47_Containers10.png)

Estas instrucciones se han sacado de "Docker Networking - DockerCon 2017". Aun siendo de la DockerCon 2017 sigue estando vigente a día de hoy. En ese tutorial siguen explicando otros conceptos de redes incluyyendo Docker Swarm, pero no lo veremos aquí. No obstante, aquí tenéis el enlace por si queréis echarle un vistazo: [Docker Networking - DockerCon 2017](https://github.com/docker/labs/tree/master/dockercon-us-2017/docker-networking) 

### Asegurando tus contenedores

Los contenedores proporcionan un entorno seguro para las cargas de trabajo frente a una configuración de servidor completa. Ofrecen la capacidad de dividir aplicaciones en componentes mucho más pequeños y poco acoplados, cada uno aislado del otro, lo que ayuda a reducir la superficie de ataque en general.

Pero no son inmunes a los hackers que buscan explotar los sistemas. Todavía tenemos que entender las trampas de seguridad de la tecnología y mantener las mejores prácticas.

### Alejarse del permiso de root

Todos los contenedores que hemos desplegado han estado utilizando el permiso root al proceso dentro de sus contenedores. Esto significa que tienen pleno acceso administrativo a su contenedor y entornos de acogida. Ahora a recorrer sabíamos que estos sistemas no iban a estar en funcionamiento por mucho tiempo. Pero ya has visto lo fácil que es ponerlo en marcha.

Podemos añadir algunos pasos a nuestro proceso para permitir que los usuarios no root sean quienes administren el contenedor, siguiendo así una buena práctica de seguridad. Al crear nuestro dockerfile podemos crear cuentas de usuario. Puedes encontrar este ejemplo también en la carpeta [containers](containers) del repositorio.

```
# Use the official Ubuntu 18.04 as base
FROM ubuntu:18.04
RUN apt-get update && apt-get upgrade -y
RUN groupadd -g 1000 basicuser && useradd -r -u 1000 -g basicuser basicuser
USER basicuser
```

También podemos usar `docker run --user 1009 ubuntu` el comando Docker run anula cualquier usuario especificado en tu Dockerfile. Por lo tanto, en el siguiente ejemplo, su contenedor siempre se ejecutará con el identificador de usuario 1009 con menos privilegios, con un nivel de permisos más bajo.

Sin embargo, este método no aborda el fallo de seguridad subyacente de la propia imagen. Por lo tanto, es mejor especificar un usuario no root en tu Dockerfile para que tus contenedores siempre se ejecuten de forma segura.

### Registro Privado

Hasta ahora hemos utilizado registros públicos en DockerHub. Con un registro privado de imágenes de contenedores establecido por la organización podremos evitar problemas de curiosos. Se puede alojar la imágen donde desees, existen registros OpenSources y también administrados por terceros. Esto da un control completo de las imágenes disponibles para el equipo de trabajo.

DockerHub está muy bien para darte una base, pero sólo te va a proporcionar un servicio básico en el que tienes que confiar mucho en el editor de imágenes.

### Liviano y limpio

He mencionado esto en todo momento sin relacionarlo con la seguridad. Pero el tamaño del contenedor puede afectar a la seguridad en términos de superficie de ataque. Si tienes recursos que no usas en tu aplicación entonces no los necesitas en tu contenedor, los puedes eliminar.

Esta es mi mayor preocupación con la extracción de las imágenes `latest` porque eso puede traer un montón de bloat (hinchazón) a sus imágenes. DockerHub muestra el tamaño comprimido para cada una de las imágenes en un repositorio.

En definitiva, busca siempre de que tus imágenes sean lo más pequeñas posibles. `docker image` es el comando para comprobar el tamaño de las imágenes en local.

![](Images/Day47_Containers11.png)

## Recursos

- [TechWorld with Nana - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
- [Programming with Mosh - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=pTFZFxd4hOI)
- [Docker Tutorial for Beginners - What is Docker? Introduction to Containers](https://www.youtube.com/watch?v=17Bl31rlnRM&list=WL&index=128&t=61s)
- [WSL 2 with Docker getting started](https://www.youtube.com/watch?v=5RQbdMn04Oc)
- [Blog on getting started building a docker image](https://stackify.com/docker-build-a-beginners-guide-to-building-docker-images/)
- [Docker documentation for building an image](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
- [YAML Tutorial: Everything You Need to Get Started in Minute](https://www.cloudbees.com/blog/yaml-tutorial-everything-you-need-get-started)
- [En español] En los [apuntes](https://vergaracarmona.es/apuntes/) del traductor:
  - [Preparación de entorno de pruebas local para docker](https://vergaracarmona.es/preparacion-de-entorno-de-pruebas-local-para-docker/)
  - [Uso básico de docker](https://vergaracarmona.es/uso-basico-de-docker/)
  - [Una breve historia sobre contenedores](https://vergaracarmona.es/breve-historia-de-contenedores/)
  - [Desplegar con docker-compose los servicios Traefik y Portainer](https://vergaracarmona.es/desplegar-con-docker-compose-los-servicios-traefik-y-portainer/)

Nos vemos en el [Día 48](day48.md)
