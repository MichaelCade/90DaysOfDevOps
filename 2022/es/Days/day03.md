## Ciclo de vida DevOps - Centrado en la aplicación

A medida que continuemos vamos a encontrarnos al 100% con el Desarrollo Continuo, testing, Despliegue, Monitorización, entre otros conceptos, una y otra vez. Si te diriges hacia el rol de Ingeniero DevOps entonces la repetibilidad será algo a lo que te acostumbrarás. Mejorar constantemente es lo que mantiene todo esto interesante.

En esta hora vamos a echar un vistazo a la vista de alto nivel de la aplicación desde el principio hasta el final, y luego de nuevo como un bucle constante.

### Desarrollo

Tomemos un ejemplo de una aplicación. Para empezar no tenemos nada creado, tal vez como desarrollador tengamos que discutir con el cliente o el usuario final los requisitos y llegar a algún tipo de plan o requisitos para la aplicación. A continuación, tenemos que crear a partir de estos requisitos nuestra nueva aplicación.

En lo que respecta a las herramientas en esta etapa, no hay ningún requisito real, aparte de elegir la IDE y el lenguaje de programación con el que se desea construir la aplicación.

Como ingeniero DevOps seguramente no crees este plan ni codificará la aplicación para el usuario final, esto lo hará un desarrollador experto. Pero no estará de más que seas capaz de leer parte del código para poder tomar las mejores decisiones de infraestructura de cara a la aplicación.

Anteriormente se mencionó que la aplicación puede ser escrita en cualquier lenguaje. Es importante que se mantenga un sistema de control de versiones, esto es algo que se cubrirá en detalle más adelante y en particular, nos sumergiremos en **Git**.

También es probable que no sea un solo desarrollador el que trabaje en este proyecto, aunque este podría ser el caso, aún así las mejores prácticas requerirían un repositorio de código para almacenar y colaborar en el código, esto podría ser privado o público y podría ser alojado o desplegado de forma privada o no. En general se suelen utilizar **GitHub o GitLab** como un repositorio de código para este menester. De todo ello hablaremos en la sección sobre **Git** más adelante.

### Testing

En esta fase, tenemos nuestros requisitos y tenemos nuestra aplicación en desarrollo. Pero tenemos que asegurarnos de que estamos probando nuestro código en los diferentes entornos que tenemos a nuestra disposición o específicamente para el lenguaje de programación elegido.

En esta fase es el departamento de control de calidad quien prueba si existen posibles errores. Cada vez más vemos que se utilizan contenedores para simular el entorno de pruebas, lo que en general puede suponer un gran ahorro en los gastos generales de la infraestructura física o de la nube.

Esta fase también es probable que se automatice como parte de la siguiente área, que es la integración continua.

La capacidad de automatizar estas pruebas frente a los 10, 100 o incluso 1000 ingenieros de control de calidad que tienen que hacer esto manualmente habla por sí mismo, estos ingenieros pueden centrarse en otras cosas, por ejemplo, asegurando que la pila avanza rápidamente o para el desarrollo de más funcionalidad frente a las pruebas de errores y software. Este tiende ser el retraso en la mayoría de los lanzamientos de software tradicionales que utilizan una metodología en cascada.

### Integración

La integración se encuentra en el centro del ciclo de vida DevOps. Es la práctica en la que los desarrolladores requieren confirmar los cambios en el código fuente con mayor frecuencia. Puede ser diaria o semanalmente.

Con cada confirmación, su aplicación puede pasar por las fases de pruebas automatizadas y esto permite la detección temprana de problemas o errores antes de la siguiente fase.

No te preocupes, muchas empresas hacen esto y continuarán haciéndolo, y será el proveedor de software el que se concentre en las 3 fases anteriores, pero es posible que desee adoptar la fase final, ya que esto permitirá una ejecución más rápida y eficiente de los despliegues.

Tener este conocimiento es muy importante, ya que puede comprar el software de la plataforma hoy, pero ¿qué pasará mañana o en el próximo trabajo?

### Despliegue

Bien, tenemos nuestra aplicación construida y probada, con los requisitos de nuestro usuario final. Ahora tenemos que seguir adelante y desplegar en producción para que nuestros usuarios finales la puedan utilizar.

Esta es la etapa en la que el código se despliega en los servidores de producción. Ahora es cuando las cosas se ponen muy interesantes y es donde tendremos que sumergirnos el resto de nuestros 86 días. Diferentes aplicaciones requieren diferentes hardware y/o configuraciones. Aquí es donde la **Gestión de la Configuración de la Aplicación** y la **Infraestructura como Código** podrían jugar un papel clave en el ciclo de vida DevOps. Puede ser que la aplicación esté **contenerizada** pero también disponible para ejecutarse en una máquina virtual. Esto también nos lleva a plataformas como **Kubernetes**, que orquestaría nuestros contenedores y aseguraría el estado deseado.

De estos atrevidos temas, entraremos en más detalle en las próximas semanas para tener un mejor conocimiento fundacional de lo que son y cuándo usarlos.

### Monitorizar

Las cosas se mueven rápidamente aquí y tenemos nuestra aplicación que estamos actualizando continuamente con nuevas características y funcionalidades. Tenemos el testing asegurándose de que no se hay gremlins por ahí. Tenemos la aplicación funcionando en nuestro entorno de producción y puede estar continuamente mantenida la configuración y el rendimiento requeridos.

Pero ahora necesitamos estar seguros de que nuestros usuarios finales están obteniendo la experiencia que requieren. Aquí tenemos que asegurarnos de que el rendimiento de nuestra aplicación está siendo continuamente monitorizado. Esta fase va a permitir a los desarrolladores tomar las mejores decisiones sobre las mejoras en las futuras versiones de la aplicación.

Esta sección es también donde vamos a capturar el feedback sobre las características que se han implementado y cómo los usuarios finales les gustaría mejorar la aplicación para ellos mismos.

La fiabilidad es un factor clave aquí, al fin y al cabo queremos que nuestra aplicación esté disponible todo el tiempo que se necesite, tener alta disponibilidad. Esto nos lleva a otras áreas de **observabilidad, seguridad y gestión de datos** que deben ser monitorizados continuamente. El feedback siempre puede ser utilizado para mejorar, actualizar y liberar la aplicación continuamente.

En este proceso continuo también debemos tener involucrados los equipos de [FinOps](https://www.ibm.com/es-es/topics/finops). Las aplicaciones y los datos se ejecutan y se almacenan en algún lugar que debe tener seguimiento financiero para asegurarse de que si las cosas cambian desde un punto de vista de recursos económicos, que no sean causa de algún dolor financiero importante en las facturas del cloud.

Al hablar con otros en la comunidad, el título de "Ingeniero DevOps" no debería ser la meta para nadie porque realmente cualquier posición debería adoptar los procesos DevOps y la cultura explicada aquí. DevOps debería ser usado en muchas posiciones diferentes como ingeniero/arquitecto Cloud-Native, administrador de virtualización, arquitecto/ingeniero cloud, y administrador de infraestructura. Esto es por nombrar algunos, pero la razón de usar Ingeniero DevOps arriba fue realmente para resaltar el alcance del proceso utilizado por cualquiera de las posiciones anteriores.

## Recursos

Siempre se dejarán recursos adicionales en los archivos readme, ya que estamos aquí para aprender.

Lo mejor para el seguimiento completo de estos documentos es que veas los siguientes vídeos y leas los artículos para hacerte una idea general de lo que hemos visto hoy y puedas profundizar un poco más.

- [Continuous Development](https://www.youtube.com/watch?v=UnjwVYAN7Ns) 
- [Continuous Testing - IBM YouTube](https://www.youtube.com/watch?v=RYQbmjLgubM)
- [Continuous Integration - IBM YouTube](https://www.youtube.com/watch?v=1er2cjUq1UI)
- [Continuous Monitoring](https://www.youtube.com/watch?v=Zu53QQuYqJ0)
- [The Remote Flow](https://www.notion.so/The-Remote-Flow-d90982e77a144f4f990c135f115f41c6)
- [FinOps Foundation - What is FinOps](https://www.finops.org/introduction/what-is-finops/)
- [**NOT FREE** The Phoenix Project: A Novel About IT, DevOps, and Helping Your Business Win](https://www.amazon.com/Phoenix-Project-DevOps-Helping-Business/dp/1942788290/)

Si has llegado hasta aquí entonces sabrás si es aquí donde quieres estar o no. Nos vemos en el [Día 4](day04.md).
