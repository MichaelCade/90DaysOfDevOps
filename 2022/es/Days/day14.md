## El panorama: DevOps & Linux

Linux y DevOps comparten culturas y perspectivas muy similares; ambos se centran en la personalización y la escalabilidad. Estos dos aspectos de Linux son de especial importancia para DevOps.

Muchas tecnologías parten de Linux, especialmente si están relacionadas con el desarrollo de software o la gestión de infraestructuras.

Además, muchos proyectos de código abierto, especialmente las herramientas DevOps, se diseñaron para funcionar en Linux desde el principio.

Desde el punto de vista de DevOps o de cualquier función de operaciones, te encontrarás con Linux. Hay un lugar para WinOps pero la mayor parte del tiempo vas a estar administrando y desplegando servidores Linux. 🐧

Michael Cade ha estado usando Linux diariamente durante varios años, aun siendo su equipo Desktop de macOS o Windows. Sin embargo, cuando empezó el rol de Cloud Native en el que está actualmente, se aseguró de que su portátil estuviera completamente basado en Linux y fuera el Sistema Operativo utilizado a diario, mientras que todavía necesitaba Windows por aplicaciones para el trabajo que no son multiplataforma y por incompatibilidad de equipos de audio y vídeo que no se ejecutan en Linux por secuestro corporativo de Windows o Mac. Quería ejecutar un Linux Desktop a tiempo completo para obtener una mejor comprensión de muchas de las cosas que vamos a ver en los próximos 7 días.

Por otra parte, el traductor, lleva más de 10 años con Sistemas Operativos Linux, solo utilizando Windows o MacOS en momentos puntuales por imperativos de entornos sociales o laborales. Tan solo echo de menos Windows si me apetece jugar a un juego de vanguardia o quiero probar software, en el primer caso tengo un disco duro preparado por si fuera necesario; y en el segundo caso, para simplemente probar, tengo una virtualización de Windows 10. Windows 11, hasta la fecha es un verdadero desastre. Y tengo que decir que Microsoft se puso mucho las pilas con la versión 10 y, aunque no cumple mis necesidades operativas y de privacidad, tengo que admitir que es bastante robusto comparado a sus anteriores fiascos. 

## Empecemos

No se pretende conseguir que el lector siga los mismos pasos descritos anteriormente, ni mucho menos, pero sí que puede ser interesante conocer las opiniones anteriores. Hay opciones más sencillas, por supuesto, los Sistemas Operativos de Microsoft o Apple son un "llave en mano", pero hay que destacar que dar el paso de utilizar a tiempo completo alguna distribución Linux te obliga a aprender más rápido cómo hacer que las cosas funcionen, familiarizandote con la shell y, al fin y al cabo, te permite conocer como funcionan los Sistemas Operativos, la red y su relación con el hardware.

Durante la mayor parte de estos 7 días, desplegaremos una máquina virtual en Virtual Box con una versiones Desktop de una distribución de Linux. Muchos de los servidores Linux que administrarás probablemente serán servidores que vienen sin interfaz gráfica de usuario ([GUI](https://es.wikipedia.org/wiki/Interfaz_gr%C3%A1fica_de_usuario)) y todo está basado en el shell ([CLI](https://es.wikipedia.org/wiki/Interfaz_de_l%C3%ADnea_de_comandos)). Sin embargo, como dije al principio, muchas de las herramientas que cubrimos a lo largo de estos 90 días empezaron en Linux, es aconsejable que te sumerjas en la ejecución de distribuciones Linux, sobre todo en CLI, para tener también esa experiencia de aprendizaje que te servirá u mucho, incluso para entender mejor otros Sistemas Operativos como Windows o MacOS.

Para el resto de este post, vamos a centrarnos en conseguir una máquina virtual de Ubuntu Desktop en nuestro entorno de Virtual Box. Antes de seguir descarga de las páginas oficiales [Virtual Box](https://www.virtualbox.org/) y la última versión de [Ubuntu ISO](https://ubuntu.com/download). 

Otra buena razón para utilizar distribuciones Linux es que son gratuitas y de [código abierto](https://es.wikipedia.org/wiki/C%C3%B3digo_abierto), e incluso algunas de [software libre](https://es.wikipedia.org/wiki/Software_libre). Elegimos Ubuntu, ya que actualmente es la distribución más utilizada, sin pensar en los smartphones y los servidores empresariales RedHat Enterprise. Esto es una visión subjetiva, puede ser errónea, pero con lo sucedido en los últimos tiempos con CentOS puede que Ubuntu se haya puesto en lo alto de la lista.

## Presentación de HashiCorp Vagrant

Vagrant es una utilidad CLI que gestiona el ciclo de vida de tus máquinas virtuales. Podemos utilizar vagrant para hacer girar y bajar las máquinas virtuales a través de muchas plataformas diferentes, incluyendo vSphere, Hyper-v, Virtual Box y también Docker. Tiene otros proveedores, pero nos quedaremos con Virtual Box, así que estamos listos para continuar.

Lo primero que tenemos que hacer es instalar Vagrant en nuestra máquina, cuando vayas a la página de descargas verás todos los sistemas operativos listados para tu elección. [HashiCorp Vagrant](https://www.vagrantup.com/downloads).

Lo siguiente que tenemos que hacer es instalar [Virtual Box](https://www.virtualbox.org/wiki/Downloads) Se puede instalar en muchos sistemas operativos diferentes y una buena razón para elegir esto y vagrant es que podrás utilizarlo tanto en Windows, macOS, o Linux.

Ambas instalaciones son bastante sencillas y ambas tienen una gran comunidad alrededor de ellas, así que siéntete libre de contactar con ellos si tienes problemas. Sus páginas de Documentación:
- [Virtual Box](https://www.virtualbox.org/wiki/Documentation)
- [Ubuntu](https://help.ubuntu.com/)
- [Vagrant](https://developer.hashicorp.com/vagrant/docs)

## Nuestro primer VAGRANTFILE

El VAGRANTFILE describe el tipo de máquina que queremos desplegar. También define la configuración y el aprovisionamiento para esta máquina.

Para guardar máquinas virtuales, los VAGRANTFILEs se ordenan en carpetas en distintos espacios de trabajo. Puedes ver a continuación cómo se ve esto en un Windows. Espero que después de esto juegues con Vagrant y veas la facilidad de hacer girar diferentes sistemas, también es genial para ese agujero de conejo conocido como salto de distro para Linux Desktop.

![](Images/Day14_Linux1.png)

Echemos un vistazo a ese VAGRANTFILE y veamos lo que estamos construyendo.

```
Vagrant.configure("2") do |config|

  config.vm.box = "chenhan/ubuntu-desktop-20.04"

  config.vm.provider :virtualbox do |v|

   v.memory  = 8096

   v.cpus    = 4

   v.customize ["modifyvm", :id, "--vram", "128"]

end

end
```

Este es un VAGRANTFILE muy simple. Estamos diciendo que queremos una "box" específica, siendo una imagen pública o una construcción privada del sistema que estás buscando. Puedes encontrar una larga lista de "boxes" disponibles públicamente aquí en el [catálogo público de boxes Vagrant](https://app.vagrantup.com/boxes/search)

En la siguiente línea decimos que queremos usar un proveedor específico y en este caso es `VirtualBox`. También definimos la memoria de nuestra máquina a `8GB` y el número de CPUs a `4`. Mi experiencia me dice que es posible que quieras añadir también la siguiente línea si experimentas problemas de visualización. Esto establecerá la memoria de vídeo a lo que usted quiere, yo aumentaría esto hasta `128MB` pero depende de su sistema.

```
v.customize ["modifyvm", :id, "--vram", ""]
```

También hay que colocar una copia de este archivo específico de vagrant en la [Carpeta Linux](Linux/VAGRANTFILE)

## Aprovisionando nuestro Linux Desktop

Ya estamos listos para poner en marcha nuestra primera máquina, en el terminal de nuestra estación de trabajo. Más abajo puedes ver un ejemplo en PowerShell de Windows. Navega hasta tu carpeta de proyectos y allí encontrarás tu VAGRANTFILE. Una vez allí puedes escribir el comando `vagrant up` y si todo está bien verás algo así.

![](Images/Day14_Linux2.png)

Otra cosa que hay que añadir aquí es que la red estará configurada como `NAT` en la máquina virtual. En esta etapa no necesitamos saber sobre NAT ya que tendremos una sesión completa hablando sobre ello en la sesión de Redes. Sepa que es el botón fácil cuando se trata de conseguir una máquina en su red doméstica, también es el modo de red por defecto en Virtual Box. Puedes encontrar más información en la [documentación de Virtual Box](https://www.virtualbox.org/manual/ch06.html#network_nat).

Una vez que `vagrant up` se ha completado podemos usar `vagrant ssh` para saltar directamente a la terminal de nuestra nueva VM.

![](Images/Day14_Linux3.png)

Aquí es donde vamos a hacer la mayor parte de nuestra exploración en los próximos días, pero también nos sumergiremos en algunas personalizaciones para su estación de trabajo de desarrollador y que harán su vida mucho más simple cuando ejecuta esto a diario, y por supuesto, ¿estás realmente en DevOps solamente cuando tengas el buen terminal fresco?

Para confirmar en Virtual Box, deberías ver el prompt de login cuando seleccionas tu VM.

![](Images/Day14_Linux4.png)

Ah y si has llegado hasta aquí y te has preguntado **¿QUÉ NOMBRE DE USUARIO Y CONTRASEÑA?** 🤔
- Username = vagrant
- Password = vagrant

Mañana veremos algunos comandos y lo que hacen. La terminal va a ser el lugar para hacer que todo suceda y acabará siendo tu amiga.

## Recursos

- [Learn the Linux Fundamentals - Part 1](https://www.youtube.com/watch?v=kPylihJRG70)
- [Linux for hackers (don't worry you don't need to be a hacker!)](https://www.youtube.com/watch?v=VbEx7B_PTOE)
- [Webminal](https://www.webminal.org/) 

Habrá muchos recursos que encontrarás a medida que vayamos avanzando y, al igual que los recursos de Go, generalmente mantendré los que tengan contenido GRATUITO.

Nos vemos en el [día 15](day15.md)
