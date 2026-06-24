## Alternativas a Docker

Dije al principio de esta sección que íbamos a usar Docker, simplemente porque tiene muchos recursos ya que la comunidad es muy grande, pero también porque fue de aquí de donde vinieron los impulsos para popularizar los contenedores. Te animo a que veas algo de la [historia de los contenedores](https://vergaracarmona.es/breve-historia-de-contenedores/), me pareció muy útil para entender todos los conceptos que trae.

Pero como he aludido existen alternativas a Docker si pensamos en lo que es Docker y de qué se ocupa: una plataforma para desarrollar, probar, desplegar y gestionar aplicaciones.

Quiero destacar algunas alternativas a Docker que podrías ver o verás en el futuro.

### Podman

¿Qué es Podman? [Podman](https://podman.io/) es un motor de contenedores sin demonio para desarrollar, gestionar y ejecutar contenedores OCI en tu sistema Linux. Los contenedores se pueden ejecutar como root o en modo rootless.

Al igual que Docker no hay ningún requisito para la virtualización, pero utilizará el sistema operativo subyacente y por eso no se puede ejecutar en Windows a pelo. Se puede ejecutar bajo WSL2 pero no es tan visual como la experiencia con Docker Desktop, que está hecho para quienes usan el ratón. 

También hay un cliente remoto de Windows donde se puede conectar a una máquina virtual Linux donde se ejecutarán los contenedores. Con el Ubuntu en WSL2 de la versión 20.04, siguiendo los siguientes pasos podrás instalar Podman en tu instancia WSL.

```Shell
echo "deb https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/xUbuntu_20.04/ /" |
sudo tee /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list
```

Para añadir la clave GPG:

```Shell
curl -L "https://download.opensuse.org/repositories/devel:/kubic:\
/libcontainers:/stable/xUbuntu_20.04/Release.key" | sudo apt-key add -
```

Ejecute una actualización del sistema con `sudo apt-get update && sudo apt-get upgrade`. Por último, podemos instalar podman usando `sudo apt install podman`. Espero que os déis cuenta de la incomodidad que es usar un sistema operativo que no sea Linux para trabajar con contenedores. Suele pasar con la mayoría de aplicaciones de Sistemas o desarrollo. Así como hemos visto la optimización de los contenedores, deberéis tener en cuenta que el sistema operativo anfitrión también es importante.

Ahora podemos utilizar muchos de los mismos comandos que hemos estado utilizando para docker. Como puedes ver a continuación, he utilizado `podman images` y no tengo nada después de la instalación. Luego, he utilizado `podman pull ubuntu` para bajar la imagen del contenedor ubuntu.

![](Images/Day48_Containers1.png)

Podemos ejecutar nuestra imagen de Ubuntu usando `podman run -dit ubuntu` y `podman ps` para ver nuestra imagen en ejecución.

![](Images/Day48_Containers2.png)

Para entrar en ese contenedor podemos ejecutar `podman attach dazzling_darwin` el nombre de tu contenedor seguramente será diferente.

![](Images/Day48_Containers3.png)

Si te estás moviendo de docker a podman también es común cambiar el archivo de configuración para tener `alias docker=podman` de esa manera cualquier comando que ejecute con docker utilizará podman.

### LXC

LXC es un motor de contenedorización que permite a los usuarios crear múltiples entornos de contenedores Linux aislados. A diferencia de Docker, LXC actúa como un hipervisor para crear múltiples máquinas Linux con archivos de sistema separados, y características de red. Existió antes que Docker y volvió a aparecer debido a las deficiencias de Docker.

Sin embargo, LXC es tan ligero como Docker y fácil de desplegar.

### Containerd

Un runtime de contenedor independiente. Containerd aporta simplicidad y robustez, así como, por supuesto, portabilidad. Containerd era antes una herramienta que se ejecutaba como parte de los servicios de contenedores Docker hasta que Docker decidió graduar sus componentes en componentes independientes.

Es un proyecto de la Cloud Native Computing Foundation, lo que lo sitúa en la misma clase que herramientas de contenedores populares como Kubernetes, Prometheus y CoreDNS. De hecho, a partir de la versión 1.11, Kubernetes utiliza containerd como su runtime de contenedor predeterminado mientras que antes era Docker.

### Otras herramientas Docker

También podríamos mencionar herramientas y opciones alrededor de Rancher, y VirtualBox pero podemos cubrirlas en más detalle en otra ocasión.

[**Gradle**](https://gradle.org/)

- Las exploraciones de construcción permiten a los equipos depurar sus scripts de forma colaborativa y realizar un seguimiento del historial de todas las construcciones.
- Las opciones de ejecución dan a los equipos la capacidad de construir continuamente de modo que cada vez que se introducen cambios, la tarea se ejecuta automáticamente.
- El diseño personalizado del repositorio ofrece a los equipos la posibilidad de tratar cualquier estructura de directorios de archivos como un repositorio de artefactos.

[**Packer**](https://packer.io/)

- Posibilidad de crear varias imágenes de máquina en paralelo para ahorrar tiempo a los desarrolladores y aumentar la eficacia.
- Los equipos pueden depurar fácilmente las compilaciones mediante el depurador de Packer, que inspecciona los fallos y permite a los equipos probar soluciones antes de reiniciar las compilaciones.
- Compatibilidad con muchas plataformas mediante plugins para que los equipos puedan personalizar sus compilaciones.

[**Logspout**](https://github.com/gliderlabs/logspout)

- Herramienta de registro: la capacidad de personalización de la herramienta permite a los equipos enviar los mismos registros a varios destinos.
- Los equipos pueden gestionar fácilmente sus archivos porque la herramienta sólo requiere acceso al socket Docker.
- Completamente de código abierto y fácil de desplegar.

[**Logstash**](https://www.elastic.co/products/logstash)

- Personaliza tu pipeline utilizando el marco de trabajo pluggable de Logstash.
- Analice y transforme fácilmente sus datos para el análisis y para ofrecer valor empresarial.
- La variedad de salidas de Logstash le permite dirigir sus datos hacia donde desee.

[**Portainer**](https://www.portainer.io/)

- Utilice plantillas prediseñadas o cree las suyas propias para desplegar aplicaciones.
- Cree equipos y asigne funciones y permisos a sus miembros.
- Sepa qué se está ejecutando en cada entorno mediante el panel de control de la herramienta.

## Recursos

- [TechWorld with Nana - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
- [Programming with Mosh - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=pTFZFxd4hOI)
- [Docker Tutorial for Beginners - What is Docker? Introduction to Containers](https://www.youtube.com/watch?v=17Bl31rlnRM&list=WL&index=128&t=61s)
- [WSL 2 with Docker getting started](https://www.youtube.com/watch?v=5RQbdMn04Oc)
- [Blog on getting started building a docker image](https://stackify.com/docker-build-a-beginners-guide-to-building-docker-images/)
- [Docker documentation for building an image](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
- [YAML Tutorial: Everything You Need to Get Started in Minute](https://www.cloudbees.com/blog/yaml-tutorial-everything-you-need-get-started)
- [Podman | Daemonless Docker | Getting Started with Podman](https://www.youtube.com/watch?v=Za2BqzeZjBk)
- [LXC - Guide to building an LXC Lab](https://www.youtube.com/watch?v=cqOtksmsxfg)
- [En español] En los [apuntes](https://vergaracarmona.es/apuntes/) del traductor:
  - [Preparación de entorno de pruebas local para docker](https://vergaracarmona.es/preparacion-de-entorno-de-pruebas-local-para-docker/)
  - [Uso básico de docker](https://vergaracarmona.es/uso-basico-de-docker/)
  - [Una breve historia sobre contenedores](https://vergaracarmona.es/breve-historia-de-contenedores/)
  - [Desplegar con docker-compose los servicios Traefik y Portainer](https://vergaracarmona.es/desplegar-con-docker-compose-los-servicios-traefik-y-portainer/)

Nos vemos en el [Día 49](day49.md)
