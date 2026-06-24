## Python para la Automatización de Redes

Python es el lenguaje estándar utilizado para las operaciones de red automatizadas.

Si bien no es sólo para la automatización de la red, realmente parece estar en todas partes cuando buscas recursos y, como se mencionó anteriormente, si no es Python generalmente es Ansible, el cuál también en Python.

- Legibilidad y facilidad de uso - Parece que Python simplemente tiene sentido. No parece haber requisitos en torno a `{}` en el código para iniciar y finalizar bloques. Si unimos esto a un IDE potente como [VS Code](https://code.visualstudio.com/), tendremos un comienzo bastante fácil cuando queramos ejecutar código python.

[Pycharm](https://www.jetbrains.com/pycharm/) podría ser otro IDE digno de mención aquí.

- Librerías - La extensibilidad de Python es la verdadera mina de oro aquí, como se ha comentado no es sólo para Automatización de Redes pero de hecho, hay librerías de sobra para todo tipo de dispositivos y configuraciones. Puedes ver la gran cantidad aquí [PyPi](https://pypi.python.org/pypi)

Cuando quieres descargar la librería a tu workstation, entonces usas una herramienta llamada `pip` para conectarte a PyPI y descargarla localmente. Vendedores de redes como Cisco, Juniper y Arista desarrollaron librerías para facilitar el acceso a sus dispositivos.

- Potente y Eficiente - ¿Recuerdas que durante los días de Go pasé por el escenario "Hola Mundo" y recorrimos creo que 6 líneas de código? En Python es

```
print('hello world')
```

Ponga todos los puntos anteriores juntos y debería ser fácil ver por qué Python se menciona generalmente como la herramienta de-facto cuando se trabaja en la automatización.

Creo que es importante tener en cuenta que es posible que hace varios años había scripts que podrían haber interactuado con sus dispositivos de red para tal vez automatizar la copia de seguridad de la configuración o para recopilar registros y otras ideas en sus dispositivos. La automatización de la que estamos hablando aquí es un poco diferente y eso se debe a que el panorama general de las redes también ha cambiado para adaptarse mejor a esta forma de pensar y ha permitido una mayor automatización.

- Red definida por software: los controladores SDN asumen la responsabilidad de proporcionar la configuración del plano de control a todos los dispositivos de la red, lo que significa un único punto de contacto para cualquier cambio en la red, ya no es necesario conectarse por telnet o SSH a cada dispositivo y depender de humanos para hacerlo, lo que tiene una probabilidad repetible de fallos o errores de configuración.

- Orquestación de alto nivel: suba un nivel desde los controladores SDN y esto permite la orquestación de los niveles de servicio, luego está la integración de esta capa de orquestación en sus plataformas de elección, VMware, Kubernetes, nubes públicas, etc.

- Gestión basada en políticas: ¿Qué desea tener? ¿Cuál es el estado deseado? Usted describe esto y el sistema tiene todos los detalles sobre cómo resolverlo para convertirse en el estado deseado.

## Configuración del entorno de laboratorio

No todo el mundo tiene acceso a routers físicos, switches y otros dispositivos de red.

Algunas herramientas pueden ayudar para aprender a automatizar la configuración de nuestras redes. Algunas de ellas son:

- [GNS3 VM](https://www.gns3.com/software/download-vm)
- [Eve-ng](https://www.eve-ng.net/)
- [Unimus](https://unimus.net/) No es un entorno de laboratorio pero es un concepto interesante.

Vamos a construir nuestro laboratorio utilizando [Eve-ng](https://www.eve-ng.net/) pero también se puede utilizar un dispositivo físico. Una ventaja de utilizar un entorno virtual es que podemos tener un entorno de sandbox para probar muchos escenarios diferentes. Además de ser capaz de jugar con diferentes dispositivos y topologías que podrían resultar de interés.

Vamos a hacer todo en EVE-NG con la edición de la comunidad.

### Introducción

La edición de la comunidad viene en formatos ISO y OVF para [descargar](https://www.eve-ng.net/index.php/download/).

Vamos a utilizar la descarga OVF, pero con la ISO existe la opción de construir en un servidor bare metal sin necesidad de un hipervisor.

![](Images/Day25_Networking1.png)

Para este tutorial se utilizará un VMware Workstation, con una licencia a través de vExpert, pero también se puede utilizar VMware Player o cualquiera de las otras opciones mencionadas en la [documentación](https://www.eve-ng.net/index.php/documentation/installation/system-requirement/). Desafortunadamente, ¡no podemos utilizar el Virtualbox utilizado anteriormente! 

Aquí es también donde hubo un problema con GNS3 con Virtual Box a pesar de ser compatible.

[Descargar VMware Workstation Player - GRATIS](https://www.vmware.com/uk/products/workstation-player.html)

[VMware Workstation PRO](https://www.vmware.com/uk/products/workstation-pro.html)  ¡También tiene un periodo de evaluación gratuito!

### Instalación en VMware Workstation PRO

Ahora tenemos el software de hipervisor descargado e instalado, y tenemos el OVF de EVE-NG descargado. Si estás usando VMware Player por favor hazme saber si este proceso es el mismo.

Ahora estamos listos para configurar las cosas. Abre VMware Workstation y selecciona `file` y `open`.

![](Images/Day25_Networking2.png)

Cuando descargues la imagen OVF de EVE-NG estará dentro de un archivo comprimido. Extraiga el contenido en su carpeta para que se vea así.

![](Images/Day25_Networking3.png)

Navegue hasta la ubicación donde descargaste la imagen OVF de EVE-NG y comience la importación.

Dale un nombre reconocible y guarda la máquina virtual en algún lugar de tu sistema.

![](Images/Day25_Networking4.png)

Cuando finalice la importación, aumente el número de procesadores a 4 y la memoria asignada a 8 GB. (Esto debería ser así después de la importación con la última versión, si no es así, edite la configuración de la máquina virtual).

Además, asegúrese de que la casilla Virtualizar Intel VT-x/EPT o AMD-V/RVI está activada. Esta opción indica a VMware workstation que pase los flags de virtualización al sistema operativo invitado (virtualización anidada). Este era el problema que tenía con GNS3 con Virtual Box aunque el CPU lo permite.

![](Images/Day25_Networking5.png)

### Encendido y acceso

Nota al margen y Rabbit hole: ¿Recuerdas que mencioné que esto no funcionaría con VirtualBox? Bueno, sí, tuve el mismo problema con VMware Workstation y EVE-NG, pero no fue culpa de la plataforma de virtualización.

Con WSL2 ejecutándose en Windows y esto parece eliminar la capacidad de poder ejecutar cualquier cosa anidada dentro del entorno. Estoy confundido en cuanto a por qué el Ubuntu VM se ejecuta, ya que parece sacar el Intel VT-d aspecto de virtualización de la CPU cuando se utiliza WSL2.

Para resolver esto podemos ejecutar el siguiente comando en nuestra máquina Windows y reiniciar el sistema, tenga en cuenta que mientras esto está apagado no serás capaz de utilizar WSL2.

`bcdedit /set hypervisorlaunchtype off`

Cuando quieras volver a utilizar WSL2 tendrás que ejecutar este comando y reiniciar.

`bcdedit /set hypervisorlaunchtype auto`

Ambos comandos deben ejecutarse como administrador.

Ahora deberías tener una máquina encendida en VMware Workstation y un prompt similar a este.

![](Images/Day25_Networking6.png)

En el prompt de arriba puedes usar:

username = root
password = eve

A continuación, se le pedirá que proporcione la contraseña de root de nuevo, que se utilizará para SSH en el host más adelante.

A continuación, podemos cambiar el nombre de host.

![](Images/Day25_Networking7.png)

A continuación, definimos un Nombre de Dominio DNS, he utilizado el de abajo, pero no estoy seguro de si esto tendrá que ser cambiado más adelante.

![](Images/Day25_Networking8.png)

Luego configuramos la red, selecciono estática para que la dirección IP sea persistente después de reiniciar.

![](Images/Day25_Networking9.png)

El paso final, proporcionar una dirección IP estática de una red que sea alcanzable desde su estación de trabajo.

![](Images/Day25_Networking10.png)

Hay algunos pasos adicionales aquí donde tendrá que proporcionar una máscara de subred para su red, puerta de enlace predeterminada y DNS.

Una vez terminado se reiniciará, cuando vuelva a funcionar puedes tomar tu dirección IP estática y ponerla en tu navegador.

![](Images/Day25_Networking11.png)

El nombre de usuario por defecto para la GUI es `admin` y la contraseña es `eve` mientras que el nombre de usuario por defecto para SSH es `root` y la contraseña es `eve` pero esto se puede haber cambiado durante la configuración.

![](Images/Day25_Networking12.png)

Elegí HTML5 para la consola vs nativa ya que esto abrirá una nueva pestaña en tu navegador cuando estés navegando a través de diferentes consolas.

A continuación vamos a:

- Instalar el paquete cliente de EVE-NG
- Cargar algunas imágenes de red en EVE-NG
- Construir una topología de red
- Añadir Nodos
- Conectar Nodos
- Empezar a construir Scripts Python
- Mira telnetlib, Netmiko, Paramiko y Pexpect

## Recursos

- [Free Course: Introduction to EVE-NG](https://www.youtube.com/watch?v=g6B0f_E0NMg)
- [EVE-NG - Creating your first lab](https://www.youtube.com/watch?v=9dPWARirtK8)
- [3 Necessary Skills for Network Automation](https://www.youtube.com/watch?v=KhiJ7Fu9kKA&list=WL&index=122&t=89s)
- [Computer Networking full course](https://www.youtube.com/watch?v=IPvYjXCsTg8)
- [Practical Networking](http://www.practicalnetworking.net/)
- [Python Network Automation](https://www.youtube.com/watch?v=xKPzLplPECU&list=WL&index=126)

Nos vemos el [Día 26](day26.md).
