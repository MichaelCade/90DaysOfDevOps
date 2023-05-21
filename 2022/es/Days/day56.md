## El panorama: IaC

¡Los humanos cometen errores! ¡La automatización es el camino a seguir!

¿Cómo construyes tus sistemas hoy en día?

¿Cuál sería tu plan si hoy perdieras todo, máquinas físicas, máquinas virtuales, VMs en la nube, PaaS en la nube, etc.?

¿Cuánto tiempo te tomaría reemplazar todo?

La infraestructura como código proporciona una solución para poder hacer esto, al mismo tiempo que también podemos probarlo. No debemos confundir esto con la copia de seguridad y recuperación, pero en términos de tu infraestructura y entornos, tus plataformas deberían poder ser creadas y tratadas como ganado en lugar de mascotas.

En resumen, podemos usar código para reconstruir todo nuestro entorno.

Si también recordamos desde el principio, dijimos que DevOps en general es una forma de eliminar barreras para entregar sistemas de manera segura y rápida.

La infraestructura como código nos ayuda a entregar los sistemas. Hemos hablado mucho sobre procesos y herramientas. IaC nos brinda más herramientas para familiarizarnos y habilitar esta parte del proceso.

Nos vamos a centrar en la infraestructura como código en esta sección. Es posible que también escuches que se menciona como infraestructura desde el código o configuración como código. Creo que el término más conocido es probablemente infraestructura como código.

### Mascotas vs Ganado

Si echamos un vistazo al tiempo anterior a DevOps, si teníamos el requisito de construir una nueva aplicación, tendríamos que preparar nuestros servidores manualmente en su mayor parte.

- Implementar VMs | Servidores físicos e instalar el sistema operativo.
- Configurar la red.
- Crear tablas de enrutamiento.
- Instalar software y actualizaciones.
- Configurar el software.
- Instalar la base de datos.

Este sería un proceso manual realizado por administradores de sistemas. Cuanto más grande sea la aplicación y más recursos y servidores se requieran, más esfuerzo manual se necesitaría para poner en funcionamiento esos sistemas. Esto requeriría una gran cantidad de esfuerzo y tiempo humano, pero también como negocio tendrías que pagar por esos recursos para construir este entorno. Como comencé la sección con "¡Los humanos cometen errores! ¡La automatización es el camino a seguir!".

Después de la fase inicial de configuración mencionada anteriormente, también tendrías el mantenimiento de estos servidores.

- Actualizar versiones.
- Implementar nuevas versiones.
- Administración de datos.
- Recuperación de aplicaciones.
- Agregar, eliminar y escalar servidores.
- Configuración de red.

Agrega la complejidad de tener múltiples entornos de prueba y desarrollo.

Aquí es donde entra en juego la infraestructura como código. Lo mencionado anteriormente era un tiempo en el que nos preocupábamos por esos servidores como si fueran mascotas, incluso les dábamos nombres cariñosos o al menos los nombrábamos de alguna manera porque iban a estar allí por un tiempo, iban a ser parte de la "familia" durante un tiempo.

Con la infraestructura como código, podemos automatizar todas estas tareas de principio a fin. La infraestructura como código es un concepto y algunas herramientas llevan a cabo esta provisión automatizada de infraestructura. En este punto, si algo malo le sucede a un servidor, lo desechas y creas uno nuevo. Este proceso es automatizado y el servidor es exactamente como se define en el código. En este punto, no nos importa cómo se llamen, están allí en el campo cumpliendo su propósito hasta que ya no estén en el campo y tengamos otro para reemplazarlo, ya sea debido a una falla o porque actualizamos parte o la totalidad de nuestra aplicación.

Esto se puede utilizar en casi todas las plataformas, cargas de trabajo basadas en virtualización, en la nube e infraestructuras nativas de la nube como Kubernetes y contenedores.

### Provisión de Infraestructura

No todas las herramientas de IaC cubren todo lo siguiente. Encontrarás que la herramienta que vamos a utilizar en esta sección solo cubre realmente las primeras 2 áreas a continuación; Terraform es esa herramienta que vamos a cubrir y nos permite partir de la nada y definir en código cómo debería ser nuestra infraestructura y luego implementarla. También nos permitirá administrar esa infraestructura y también implementar inicialmente una aplicación, pero en ese momento perderá el seguimiento de la aplicación, es ahí donde entra la siguiente sección y algo como Ansible como una herramienta de gestión de configuración podría funcionar mejor en ese sentido.

Sin adelantarnos, herramientas como Chef, Puppet y Ansible son las más adecuadas para lidiar con la configuración inicial de la aplicación y luego administrar esas aplicaciones y su configuración.

Instalación y configuración inicial del software:

- Puesta en marcha de nuevos servidores.
- Configuración de red.
- Creación de balanceadores de carga.
- Configuración a nivel de infraestructura.

### Configuración de la infraestructura provisionada

- Instalación de aplicaciones en servidores.
- Preparar los servidores para implementar tu aplicación.

### Despliegue de la aplicación

- Implementar y gestionar la aplicación.
- Fase de mantenimiento.
- Actualizaciones de software.
- Reconfiguración.

### Diferencia entre herramientas de IaC

Declarativo vs Procedimental

Procedimental:

- Instrucciones paso a paso.
- Crear un servidor > Agregar un servidor > Hacer este cambio.

Declarativo:

- Declarar el resultado.
- 2 servidores.

Mutable (mascotas) vs Inmutable (ganado)

Mutable:

- Cambiar en lugar de reemplazar.
- Generalmente tienen una larga vida útil.

Inmutable:

- Reemplazar en lugar de cambiar.
- Posiblemente de corta duración.

Esta es la razón por la cual tenemos muchas opciones diferentes para la infraestructura como código, porque no hay una herramienta que las gobierne todas.

Vamos a utilizar principalmente Terraform y a trabajar prácticamente, ya que esta es la mejor manera de comenzar a ver los beneficios de la infraestructura como código cuando está en acción. Trabajar prácticamente también es la mejor manera de adquirir las habilidades necesarias para escribir código.

Continuaremos analizando Terraform con una introducción antes de poner manos a la obra y acostumbrarnos a su uso.

## Recursos

A continuación, he enumerado muchos recursos y creo que este tema ha sido cubierto tantas veces, que si tienes recursos adicionales, asegúrate de enviar una solicitud de extracción (PR) con tus recursos y estaré encantado de revisarlos y agregarlos a la lista.


- [What is Infrastructure as Code? Difference of Infrastructure as Code Tools](https://www.youtube.com/watch?v=POPP2WTJ8es)
- [Terraform Tutorial | Terraform Course Overview 2021](https://www.youtube.com/watch?v=m3cKkYXl-8o)
- [Terraform explained in 15 mins | Terraform Tutorial for Beginners](https://www.youtube.com/watch?v=l5k1ai_GBDE)
- [Terraform Course - From BEGINNER to PRO!](https://www.youtube.com/watch?v=7xngnjfIlK4&list=WL&index=141&t=16s)
- [HashiCorp Terraform Associate Certification Course](https://www.youtube.com/watch?v=V4waklkBC38&list=WL&index=55&t=111s)
- [Terraform Full Course for Beginners](https://www.youtube.com/watch?v=EJ3N-hhiWv0&list=WL&index=39&t=27s)
- [KodeKloud - Terraform for DevOps Beginners + Labs: Complete Step by Step Guide!](https://www.youtube.com/watch?v=YcJ9IeukJL8&list=WL&index=16&t=11s)
- [Terraform Simple Projects](https://terraform.joshuajebaraj.com/)
- [Terraform Tutorial - The Best Project Ideas](https://www.youtube.com/watch?v=oA-pPa0vfks)
- [Awesome Terraform](https://github.com/shuaibiyy/awesome-terraform)
- [Herramientas para Terraform](https://vergaracarmona.es/herramientas-para-terraform/)

Nos vemos en el [Día 57](day57.md)
