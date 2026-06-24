## El panorama: Contenedores

Empezamos nueva sección en la que nos centraremos en los contenedores. En particular vamos a mirar Docker para entender algunas claves de los contenedores.

Pasaremos por un poco de práctica para crear el contenedor que utilizaremos durante esta sección y en futuras secciones del desafío.

Como siempre, este primer post se va a centrar en el panorama general de cómo hemos llegado hasta aquí y lo que todo esto significa.

### ¿Por qué otra forma de ejecutar aplicaciones?

Lo primero que tenemos que plantearnos es por qué necesitamos otra forma de ejecutar software o aplicaciones. La elección es grande, podemos ejecutar nuestras aplicaciones de muchas formas diferentes, por ejemplo, podríamos ver las aplicaciones desplegadas en hardware físico con un sistema operativo y una sola aplicación desplegada allí, y podríamos verlas en máquina virtual o instancias IaaS basadas en la nube ejecutando nuestra aplicación, que luego se integran en una base de datos de nuevo en una MV o como PaaS ofreciendo en la nube pública. O, como vamos a ver, podemos ver nuestras aplicaciones ejecutándose en contenedores.

Ninguna de las opciones anteriores es incorrecta o correcta, pero cada una tiene sus razones para escogerla y seguramente ninguna de ellas desaparecezca. He visto una gran cantidad de contenido sobre Contenedores vs Máquinas Virtuales y realmente no debería haber un argumento, ya que son como argumentos de manzanas vs peras en el que ambos son frutas (formas de ejecutar nuestras aplicaciones), pero no son lo mismo.

También diría que si estás empezando y estás desarrollando una aplicación deberías inclinarte hacia los contenedores, más adelante lo veremos en detalle pero en resumen se trata de eficiencia, velocidad y tamaño. Pero esto tiene un precio, si no tienes ni idea de contenedores va a ser una curva de aprendizaje demasiado inclinada, te obligará a entender el por qué y entrar en esta mentalidad. Si has desarrollado tus aplicaciones de una forma particular o no estás en un entorno totalmente nuevo, puede que tengas que enfrentarte a más problemas incluso antes de considerar los contenedores.

Tenemos muchas opciones diferentes a la hora de descargar un determinado software, hay una variedad de sistemas operativos diferentes que podríamos utilizar. E instrucciones específicas sobre lo que tenemos que hacer para instalar nuestras aplicaciones.

![](Images/Day42_Containers1.png)

Cada vez nos encontramos más que las aplicaciones que una vez podrían haber necesitado un sistema operativo de servidor completo (una MV, instancia física o en la nube), ahora están lanzando versiones su software basadas en contenedores. Es interesante ya que esto abre las posibilidades de los contenedores y de Kubernetes a todo el mundo, y no sólo un enfoque en los desarrolladores de aplicaciones.

![](Images/Day42_Containers2.png)

Podrás deducir por lo que se ha dicho antes que no se va defender los contenedores como el mejor sistema para desplegar aplicaciones, pero desde luego existe el espacio para el debate.

![](Images/Day42_Containers4.png)

Hemos tenido la tecnología de contenedores durante mucho tiempo, así que ¿por qué en los últimos 10 años se ha hecho tan popular? Diría que incluso mucho más populares en los últimos 5. 

Hemos tenido contenedores durante décadas. ¿Se reduce al desafío de los contenedores o debería decir imágenes también? Porque si sólo tenemos la tecnología de contenedores, entonces todavía tendremos muchos de los mismos problemas que hemos tenido con la gestión de software.

Si pensamos en Docker como una herramienta, la razón por la que despegó es al ecosistema de imágenes que son fáciles de conseguir y de poner en marcha. Una parte importante de esto es la coherencia en todo el espacio, de todos estos diferentes desafíos que enfrentamos con el software. No importa si es MongoDB o nodeJS, el proceso para poner en marcha cualquiera de ellos será el mismo. El proceso para detener cualquiera de ellos es el mismo. Todos estos problemas seguirán existiendo, pero lo bueno es que, cuando unimos una buena tecnología de contenedores e imágenes conseguiremos un único conjunto de herramientas para ayudarnos a abordar todos estos problemas diferentes. Algunos de esos problemas son:

- Primero tenemos que encontrar software en Internet. 
- Luego tenemos que descargarlo.
- ¿Confiamos en la fuente?
- ¿Necesitamos una licencia? ¿Qué licencia?
- ¿Es compatible con diferentes plataformas?
- ¿Cuál es el paquete? ¿Binario? ¿Ejecutable? ¿Gestor de paquetes?
- ¿Cómo se configura el software?
- ¿Dependencias? ¿La descarga general nos ha cubierto o también las necesitamos?
- ¿Dependencias de dependencias?
- ¿Cómo iniciamos la aplicación?
- ¿Cómo detenemos la aplicación?
- ¿Se reiniciará automáticamente?
- ¿Se iniciará al arrancar?
- ¿Conflictos de recursos?
- ¿Conflictos de librerías?
- Conflictos de puertos
- ¿Seguridad del software?
- ¿Actualizaciones del software?
- ¿Cómo puedo eliminar el software?

Podemos dividir lo anterior en 3 áreas de la complejidad del software que los contenedores y las imágenes ayudan con estos.

| Distribución | Instalación    | Funcionamiento      |
| ------------ | -------------- | ------------------- |
| Buscar       | Instalar       | Iniciar             |
| Descargar    | Configuración  | Seguridad           |
| Licencia     | Desinstalación | Puertos             |
| Paquete      | Dependencias   | Conflictos Recursos |
| Confianza    | Plataforma     | Reinicio automático |
| Buscar       | Bibliotecas    | Actualizaciones     |

Los contenedores y las imágenes nos van a ayudar a eliminar algunos de estos retos que tenemos posiblemente con otro software y aplicaciones.

En un alto nivel podríamos mover la instalación y el funcionamiento en la misma lista, las imágenes nos van a ayudar desde el punto de vista de la distribución y los contenedores ayudan con la instalación y las operaciones.

Ok, probablemente suena muy bien y emocionante, pero todavía tenemos que entender lo que es un contenedor y ahora he mencionado imágenes así que vamos a verlo a continuación.

Otra cosa que habrás visto a menudo cuando hablamos de contenedores para el desarrollo de software es la analogía utilizada con los contenedores marítimos, los contenedores marítimos se utilizan para transportar diversas mercancías a través de los mares utilizando grandes buques.

![](Images/Day42_Containers5.png)

¿Qué tiene que ver con nuestro tema de los contenedores? Piensa en el código que escriben los desarrolladores de software, ¿cómo podemos enviar ese código concreto de una máquina a otra?

Si pensamos en lo que hemos tocado antes sobre la distribución de software, instalación y operaciones, pero ahora empezamos a construir esto en un entorno visual. Tenemos hardware y un sistema operativo donde se ejecutarán múltiples aplicaciones. Por ejemplo, nodejs tiene ciertas dependencias y necesita ciertas bibliotecas. Si a continuación, desea instalar MySQL entonces necesita sus bibliotecas y sus dependencias propias. Cada aplicación de software tendrá sus bibliotecas y dependencias. Podríamos tener mucha suerte y no tener ningún conflicto entre cualquiera de nuestras aplicaciones donde las bibliotecas y dependencias específicas están chocando causando problemas, pero cuantas más aplicaciones más posibilidades o riesgo de conflictos. Sin embargo, se trata de que un despliegue se producirá muchas veces, y no puede esperar a que todo este arreglado, porque en cada actualización se podrían introducir estos conflictos.

![](Images/Day42_Containers6.png)

Los contenedores pueden ayudar a resolver este problema. Los contenedores ayudan a **construir**, a **enviar**, a **desplegar** y a **escalar** estas aplicaciones con facilidad de forma independiente. echemos un vistazo a la arquitectura, tendrás hardware y sistema operativo y luego encima tendrás un motor de contenedores como docker que veremos en breve. El software de motor de contenedores ayuda a crear contenedores que empaquetan las bibliotecas y dependencias junto con él para que pueda mover este contenedor sin problemas de una máquina a otra sin preocuparse por las bibliotecas y dependencias, ya que vienen como parte de un paquete en el contenedor. Con lo cual, se puede mover a través de los sistemas sin preocuparse por las dependencias subyacentes que la aplicación necesita para funcionar.

![](Images/Day42_Containers7.png)

### Las ventajas de estos contenedores

- Los contenedores ayudan a empaquetar todas las dependencias dentro del contenedor y aislarlas.

- Es fácil gestionar.

- Se puede pasar de un sistema a otro.

- Los contenedores ayudan a empaquetar el software y se puede enviar fácilmente sin duplicar esfuerzos.

- Los contenedores son fácilmente escalables.

Usando contenedores puedes escalar contenedores independientes y usar un equilibrador de carga
o un servicio que ayude a dividir el tráfico y puedas escalar las aplicaciones horizontalmente. Los contenedores ofrecen mucha flexibilidad y facilidad a la hora de gestionar las aplicaciones.

### ¿Qué es un contenedor?

Cuando ejecutamos aplicaciones en nuestro ordenador, podría ser en el navegador o en el VScode que quizá estés utilizando para leer este post, esa aplicación se ejecuta como un proceso. En nuestros portátiles o sistemas tendemos a ejecutar múltiples aplicaciones o, como hemos dicho, procesos. Cuando abrimos una nueva aplicación o hacemos clic en el icono de la aplicación, a veces podría ser un servicio que sólo queremos ejecutar en segundo plano, nuestro sistema operativo está lleno de servicios que se ejecutan en segundo plano y que proporcionan la experiencia de usuario final.

Ese icono de aplicación representa un enlace a un ejecutable en algún lugar de tu sistema de archivos, el sistema operativo entonces carga ese ejecutable en la memoria. Curiosamente, a veces nos referimos a ese ejecutable como una imagen cuando hablamos de un proceso.

Un contenedor es una unidad estándar de software que empaqueta el código y todas sus dependencias para que la aplicación se ejecute de forma rápida y fiable de un entorno informático a otro como lo hacen las aplicaciones.

El software en contenedores siempre se ejecutará igual, independientemente de la infraestructura. Los contenedores aíslan el software en su entorno mínimo viable y garantizan que funcione de manera uniforme a pesar de las diferencias, por ejemplo, entre el desarrollo y la puesta en escena.

He mencionado las imágenes en la última sección cuando se trata de cómo y por qué la combinación de contenedores e imágenes ha hecho que los contenedores sean populares en nuestro ecosistema.

### ¿Qué es una imagen?

Una imagen de contenedor es un paquete de software ligero, independiente y ejecutable que incluye todo lo necesario para ejecutar una aplicación: código, tiempo de ejecución, herramientas del sistema, bibliotecas del sistema y configuraciones. Las imágenes de contenedor se convierten en contenedores en tiempo de ejecución.

## Recursos

- [TechWorld with Nana - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
- [Programming with Mosh - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=pTFZFxd4hOI)
- [Docker Tutorial for Beginners - What is Docker? Introduction to Containers](https://www.youtube.com/watch?v=17Bl31rlnRM&list=WL&index=128&t=61s)
- [Introduction to Container By Red Hat](https://www.redhat.com/en/topics/containers)
- [En español] En los [apuntes](https://vergaracarmona.es/apuntes/) del traductor:
  - [Preparación de entorno de pruebas local para docker](https://vergaracarmona.es/preparacion-de-entorno-de-pruebas-local-para-docker/)
  - [Uso básico de docker](https://vergaracarmona.es/uso-basico-de-docker/)
  - [Una breve historia sobre contenedores](https://vergaracarmona.es/breve-historia-de-contenedores/)
  - [Desplegar con docker-compose los servicios Traefik y Portainer](https://vergaracarmona.es/desplegar-con-docker-compose-los-servicios-traefik-y-portainer/)

Nos vemos en el [Día 43](day43.md)
