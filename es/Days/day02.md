## Responsabilidades de un Ingeniero DevOps

Con suerte, usted has llegado aquí a través de los recursos y la publicación del [Día 1 de #90DaysOfDevOps](day01.md).

Se tocó brevemente el concepto DevOps en el primer día, pero ahora debemos profundizar y entender que hay dos partes principales cuando se crea una aplicación:
- La parte de **Desarrollo**, donde los desarrolladores de software programan la aplicación y la prueban
- La parte de **Operaciones**, donde la aplicación se despliega y se mantiene en un servidor.

## DevOps es el vínculo entre Desarrollo y Operaciones

Para entender el funcionamiento de DevOps o las tareas que un ingeniero con mentalidad DevOps lleva a cabo, tenemos que entender las herramientas, el proceso y la visión general del desarrollo y las operaciones.

Todo comienza con la aplicación, todo gira en torno a ella cuando se trata de DevOps.

Los desarrolladores crearán una aplicación, esto puede hacerse con muchas pilas de tecnología diferentes y vamos a dejar eso a la imaginación por ahora, ya que entramos en esto más adelante. Esto también puede implicar muchos lenguajes de programación diferentes, herramientas de construcción, repositorios de código, etc.

Como ingeniero DevOps no estarás programando la aplicación, pero tener una buena comprensión de los conceptos de cómo trabaja un desarrollador y los sistemas, herramientas y procesos que están utilizando es clave para el éxito.

En un nivel muy alto, usted va a necesitar saber cómo la aplicación está configurada para hablar con todos sus servicios requeridos, servicios de datos y luego también espolvorear un requisito de cómo esto debe ser probado.

La aplicación tendrá que ser desplegada en algún lugar, vamos a mantenerlo simple y hacer esto en un servidor, no importa dónde, pero un servidor. Se espera que el cliente o el usuario final puedan acceder a él, dependiendo de la aplicación que se haya creado.

Este servidor debe ejecutarse en algún lugar, en las instalaciones, en un cloud público, serverless (Ok, esto es demasiado lejos, no vamos a cubrir serverless, pero es una opción y más y más empresas se dirigen en esta dirección)... Alguien tiene que crear y configurar estos servidores y prepararlos para que la aplicación se ejecute. Ahora bien, este elemento puede recaer en ti como ingeniero DevOps para desplegar y configurar estos servidores.

Estos servidores ejecutan un sistema operativo y, en general, será una distribución Linux, más adelante tenemos toda una sección (una semana) en la que cubrimos algunos de los conocimientos fundamentales que se deben tener de Linux.

También es probable que necesitemos comunicarnos con otros servicios en nuestra red o entorno, por lo que también necesitamos tener ese nivel de conocimiento en torno a la red y su configuración, esto podría caer en cierta medida en los pies del ingeniero DevOps. De nuevo se ve esto con más detalle en una sección dedicada a hablar de todo lo relacionado con DNS, DHCP, equilibrio de carga, etc.

## El que todo lo sabe, no lo sabe.

En este punto, sin embargo, diré que no es necesario ser un especialista en redes o en infraestructura, sino que se necesita un conocimiento básico de cómo hacer que las cosas funcionen y se comuniquen entre sí, al igual que tener un conocimiento básico de un lenguaje de programación, pero no es necesario ser un desarrollador. Sin embargo, es posible que te incorpores a esto como especialista en un área y eso es una gran base para adaptarte a otras áreas.

También es muy probable que no te hagas cargo de la gestión de estos servidores o de la aplicación a diario.

Hemos estado hablando de servidores, pero lo más probable es que su aplicación se desarrolle para ejecutarse en contenedores, que están en un servidor en su mayor parte donde se ejecutan. No solo se necesita una comprensión de la virtualización, sino también de la infraestructura del Cloud como Servicio (IaaS), también de la contenerización, etc. El enfoque en estos 90 días será más atendido hacia los contenedores.

## Visión general de alto nivel

Por un lado tenemos a nuestros desarrolladores creando nuevas características y funcionalidades (así como correcciones de errores) para la aplicación.

Por otro lado, tenemos algún tipo de entorno, infraestructura o servidores que están configurados y gestionados para ejecutar esta aplicación y comunicarse con todos sus servicios necesarios.

La gran pregunta es ¿cómo incorporamos esas funciones y correcciones de errores a nuestros productos y las ponemos a disposición de los usuarios finales?

¿Cómo lanzamos la nueva versión de la aplicación? Esta es una de las principales tareas de un ingeniero de DevOps, y lo importante aquí no es sólo averiguar cómo hacer esto una vez, sino que tenemos que hacerlo de forma continua y de una manera automatizada y eficiente, y además debe incluir las pruebas oportunas.

Aquí es donde vamos a terminar este día de aprendizaje, hay que coger fuerzas para continuar. En los próximos días, se profundizará un poco más en algunas áreas de DevOps para prepararnos a la entrada de las secciones especializadas en las herramientas y procesos y los beneficios de estos.

## Recursos

Siempre se dejarán recursos adicionales en los archivos readme, ya que estamos aquí para aprender.

Lo mejor para el seguimiento completo de estos documentos es que veas los siguientes vídeos y leas los artículos para hacerte una idea general de lo que hemos visto hoy y puedas profundizar un poco más.

- [What is DevOps? - TechWorld with Nana (ENG)](https://www.youtube.com/watch?v=0yWAtQ6wYNM)
- [What is DevOps? - GitHub YouTube (ENG)](https://www.youtube.com/watch?v=kBV8gPVZNEE)
- [What is DevOps? - IBM YouTube (ENG)](https://www.youtube.com/watch?v=UbtB4sMaaNM)
- [What is DevOps? - AWS (ENG)](https://aws.amazon.com/devops/what-is-devops/)
- [What is DevOps? - Microsoft (ENG)](https://docs.microsoft.com/en-us/devops/what-is-devops)

Si has llegado hasta aquí, sabrás si es aquí donde quieres estar o no. Nos vemos en el [día 3](./day03.md).
