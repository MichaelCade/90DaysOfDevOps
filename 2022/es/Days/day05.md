## [Planificar](#planificar) > [Código](#código) > [Construir](#construir) > [Testing](#testing) > [Lanzamiento](#lanzamiento) > [Deploy](#deploy) > [Operar](#operar) > [Monitorización](#monitorización) > [Aclarar y repetir](#aclarar-y-repetir)

Hoy vamos a centrarnos en los pasos individuales de principio a fin en el ciclo continuo de una Aplicación desde la mirada DevOps.

![DevOps](Images/Day5_DevOps8.png)

### Planificar

Todo comienza con el proceso de planificación, que es donde el equipo de desarrollo se reúne y determina qué tipo de características y correcciones de errores van a lanzar en su próximo sprint. Como Ingeniero DevOps también te debes involucrar en este punto para así conocer qué tipo de cosas vendrán más adelante y poder influir en las decisiones para ayudar a trabajar con la infraestructura que mejor se puede adaptar o dirigirlos hacia algo que funcionará mejor tanto para ellos como para ti. Algo clave a señalar aquí es que tanto los desarrolladores como el equipo de ingeniería de software son los clientes de los ingenieros DevOps, por lo que es importante trabajar con los clientes antes de que vayan por un mal camino.

### Código

Ahora, una vez que la sesión de planificación ha terminado, van a empezar a escribir el código y es importante ya estar involucrado, ya que cuando se escribe el código la filosofía DevOps puede ayudar a entender mejor la infraestructura si se sabe qué servicios se requieren. Estar pendiente del inicio es la mejor manera de escoger los servicios adecuados anticipándose, para así empezar a fusionar ese código en los repositorios cuanto antes.

### Construir

Aquí es donde vamos a iniciar el primero de nuestros procesos de automatización porque vamos a tomar el código de la aplicación y vamos a construir las dependencias del lenguaje que estén usando, podría ser interpretado o compilado, o se puede crear una imagen docker con ese código. De cualquier manera vamos a hacer este proceso usando desde el inicio una pipeline [CI/CD](https://es.wikipedia.org/wiki/CI/CD).

## Testing

Una vez que tengamos construido nuestro contenedor, realizaremos las primeras pruebas. El equipo de desarrollo suele escribir las pruebas adecuadas para el software específico, pero nosotros tenemos que realizarlas teniendo en cuenta que las pruebas son una forma de intentar minimizar la introducción de errores en la producción. No lo garantiza, pero tenemos que asegurar lo máximo que podamos para poder ofrecer la mayor garantía de que no tengamos problemas posteriormente y de que no se rompan cosas que solían funcionar. En producción no vale decir "En mi máquina funcionaba".

## Lanzamiento

Una vez pasadas las pruebas seguiremos con el proceso de lanzar nuestro código y, de nuevo, dependiendo del tipo de aplicación con la que trabajamos, tenemos que hacerlo siguiendo los pasos adecuados. Por ejemplo, el código puede estar en un repositorio GitHub o cualquier repositorio git privado, o podría ser un proceso de correr un código compilado o una imagen docker que ya tienes construida y simplemente subirla a un registro o repositorio donde sea accesible para los servidores de producción, para que el proceso de despliegue y del control de versiones sea lo más sencillo posible.

## Deploy

Aquí es la parte final del ciclo de vida, porque con los despliegues ponemos el código en producción y es aquí cuando nuestro negocio se da cuenta del valor de todo el esfuerzo y del tiempo invertido por los equipos de ingeniería de software.

## Operar

Ahora vamos a operar con nuestro despliegue. Esto puede implicar algo así como empezar a recibir llamadas de clientes molestos porque el sitio o la aplicación funciona lento. En este supuesto, necesitaremos averiguar que sucede lo más rápido posible a través de las métricas. Posiblemente deberemos construir un auto-escalado para manejar el aumento del número de recursos disponibles durante los períodos con grandes picos de tráfico y viceversa, disminuir los recursos durante los períodos de baja concurrencia. Esta es tan solo una de las operativas que deberemos tener en cuenta. 

Otra operativa es incluir como un bucle de retroalimentación de la producción al equipo de operaciones, para permitirle así saber de primera mano de los eventos clave que ocurren en producción. Tales como un despliegue erróneo o la necesidad de un paso atrás en la versión. Esto puede ser o no automatizado, siempre en función del entorno. El objetivo es que siempre que sea posible automatizar pero hay algunos entornos en los que el auto-escalado puede ser un inconveniente para los costes, por ejemplo. En algunos entornos se debe preparar con algunos pasos antes de estar listo para auto-escalar o quizá la producción ya es estable y no lo necesite. 

Lo ideal es que cualquier despliegue previsible este dentro del proceso de automatización. Y cuando lo hagas se deben incluir todos los pasos operativos, incluyendo algún tipo de notificación para que el equipo de operaciones sepa los cambios automáticos asi como todos los eventos importantes. Por supuesto, los despliegues.

## Monitorización

Todas las partes anteriores conducen a este paso final: la monitorización. Especialmente en torno a la resolución de problemas del auto-escalado que hemos comentado. ¿Cómo si no sabrías
que hay un problema si no tienes una monitorización que te lo diga?

Así que algunas de las cosas que se pueden controlar a través de la monitorización son:
- La utilización de la memoria.
- La utilización de la CPU.
- El espacio en disco.
- El punto final de la API.
- El tiempo de respuesta
- La rapidez con la que ese punto final está respondiendo.
- El ancho del tráfico.

En una gran parte de todo esto te ayudarán los logs, que dan la capacidad a los desarrolladores de ver lo que está sucediendo sin tener que acceder a los sistemas de producción.

## Aclarar y repetir

Una vez en este punto se vuelve al principio, a la etapa de planificación, y se repite todo el proceso en una segunda, tercera, cuarta iteración... El trabajo de la ingeniería de software e infraestructura es seguir iterando siempre para ofrecer la mejor aplicación a los clientes/usuarios finales. 

## Continuo

Muchas herramientas nos ayudan a lograr el proceso continuo mencionado, todo el código y el objetivo final de tenerlo todo completamente automatizado. La infraestructura cloud o cualquier entorno se describe a menudo como **Integración Continua / Entrega Continua / Despliegue Continuo**. "CI/CD" para abreviar. En estos 90 días, una semana entera estará dedicada a el CI/CD con ejemplos y recorridos para comprender e interiorizar los fundamentos.

### Entrega continua

**Entrega continua** = Planificar > Codificar > Construir > Probar 

### Integración continua

Esto es efectivamente el resultado de las fases de **Entrega Continua** anteriores más el resultado de la fase de Lanzamiento, tanto de los fracasos como de los éxitos. Pero esto se retroalimenta en la **entrega continua** o se traslada al **Despliegue Continuo**.

**Integración continua** = Planificar > Codificar > Construir > Probar > Lanzar

### Despliegue continuo

Con un lanzamiento exitoso de la **integración continua** ya podemos pasar al **Despliegue Continuo** que trae las siguientes fases:

Integración continua con lanzamiento exitoso = **Despliegue continuo** = Desplegar > Operar > Monitorizar

Estas tres nociones continuas son el flujo por las fases del ciclo de vida DevOps.

Esta última parte fue una recapitulación para el Día 3, en un intento de hacer las cosas lo más claras posibles.

### Recursos

- [DevOps for Developers – Software or DevOps Engineer?](https://www.youtube.com/watch?v=a0-uE3rOyeU)
- [Techworld with Nana -DevOps Roadmap 2022 - How to become a DevOps Engineer? What is DevOps?](https://www.youtube.com/watch?v=9pZ2xmsSDdo&t=125s)
- [How to become a DevOps Engineer in 2021 - DevOps Roadmap](https://www.youtube.com/watch?v=5pxbp6FyTfk)

Si has llegado hasta aquí entonces sabrás si es aquí donde quieres estar o no.

Nos vemos en el [Día 6](day06.md).
