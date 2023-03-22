## El panorama general: DevOps y el aprendizaje de un lenguaje de programación

Es justo decir que para tener éxito a largo plazo como ingeniero DevOps tienes que saber al menos un lenguaje de programación a nivel básico. En esta primera sesión de la nueva sección "Aprender un lenguaje de programación", se explorará por qué es una habilidad tan imprescindible tener conocimientos de los lenguajes de programación. Al final de esta semana o sección, adquiriremos una mejor comprensión del por qué, cómo y qué hacer para avanzar en este viaje del aprendizaje.

Si preguntáramos en las redes sociales si es necesario tener conocimientos de programación para los puestos relacionados con DevOps, la respuesta sería un SÍ rotundo. Pero entonces surge una pregunta más relevante, y he aquí donde no podremos obtener una respuesta precisa: **¿qué lenguaje de programación?** La respuesta más común suele ser Python, el pico de aprendizaje quizá sea menor. También, cada vez más a menudo, estamos viendo que Golang o Go está cogiendo relevancia entre los lenguajes de iniciación.

Para tener éxito en DevOps se debe tener un buen conocimiento de las habilidades de programación. Ahora, tenemos que entender por qué lo necesitamos para poder escoger el camino correcto.

## Entender por qué necesitas aprender un lenguaje de programación.

La razón por la que Python y Go se recomiendan tan a menudo para los ingenieros DevOps es que en muchas ocasiones las herramientas DevOps están escritas con estos lenguajes, lo cual tiene sentido si vas a construir herramientas DevOps o extender alguna de ellas. Ahora bien, esto es importante, ya que esto determinará realmente lo que se debe aprender y cuál será el más beneficioso. 
- Si vas a construir herramientas DevOps o te unes a un equipo que lo hace tendrá sentido aprender ese mismo lenguaje.
- Si vas a estar muy involucrado en Kubernetes o Contenedores entonces es más que probable que quieras elegir Go como tu lenguaje de programación. 
 
En el caso del autor, Michael Cade, la empresa para la que trabaja (Kasten by Veeam) está en el ecosistema Cloud-Native centrado en la gestión de datos para Kubernetes y todo está escrito en Go.

Puede que estos no sean unos razonamientos claros como para elegir, puede que estés aún estudiando o en transición por alguna carrera profesional sin una decisión real tomada. En esta situación deberás asegurarte de escoger el lenguaje de programación que encaje mejor con las aplicaciones que estás buscando para trabajar. Al fin y al cabo, si entiendes bien los conceptos de los lenguajes de programación, una vez aprendido bien cualquiera de ellos podrás aprender más rápido el resto.

Recuerda que en el camino de la filosofía DevOps no estas buscando convertirte en un desarrollador de software, lo que necesitas es entender un poco más sobre el lenguaje de programación para que puedas leer y entender lo que esas herramientas están haciendo. Esto nos ayudará a mejorar las cosas para proyectos concretos con necesidades especiales.

También es importante saber cómo se interactúa con las herramientas DevOps que podrían ser [Kasten K10](https://www.kasten.io/) o podría ser [Terraform](https://www.terraform.io/) y [HCL](https://github.com/hashicorp/hcl). Esto es lo que llamaremos archivos de configuración y es cómo se interactúa con esas herramientas DevOps para hacer que las cosas sucedan. Comúnmente los ficheros serán YAML. (Más adelante veremos más detalles de los ficheros YAML)

## ¿Acabo de convencerme de que no debo aprender un lenguaje de programación?

La mayor parte del tiempo, o dependiendo del rol, estarás ayudando a los equipos de ingeniería a implementar DevOps en su flujo de trabajo, muchas pruebas alrededor de la aplicación y asegurándote de que el flujo de trabajo que se construye se alinea con esos principios de DevOps que mencionamos en los primeros días. En realidad, esto va a ser una gran parte de la solución de problemas de rendimiento de la aplicación o algo por el estilo. Esto vuelve al razonamiento original, el lenguaje de programación que necesitas saber es el código que está escrito en los proyectos o trabajos en los que estés o que te interesan. Si la aplicación está escrita en NodeJS no ayudará mucho si tienes veinte certificaciones de Go o Python.

## ¿Por qué Go?

Porque Golang es el próximo lenguaje de programación para DevOps, Go se ha convertido en un lenguaje de programación muy popular en los últimos años. Según la encuesta de StackOverflow del 2021, Go quedó en cuarto lugar entre los lenguajes de programación, scripting y markup más buscados, con Python a la cabeza:

- [StackOverflow 2021 Developer Survey – Most Wanted Link](https://insights.stackoverflow.com/survey/2021#section-most-loved-dreaded-and-wanted-programming-scripting-and-markup-languages)


Como también se ha mencionado, algunas de las herramientas y plataformas DevOps más conocidas están escritas en Go como Kubernetes, Docker, Grafana y Prometheus.

¿Cuáles son algunas de las características de Go que lo hacen preferible para DevOps?

## Construcción y despliegue de programas en Go

Una ventaja de usar un lenguaje interpretado como Python en un rol de DevOps es que no necesitas compilar python antes de ejecutar el programa. Especialmente para las tareas de automatización más pequeñas, no te ralentiza el proceso de construcción que requiere la compilación. A pesar de que Go es un lenguaje de programación compilado, **Go compila directamente en código máquina**. Go es conocido también por sus rápidos tiempos de compilación.

## Go vs Python para DevOps

Los programas Go están enlazados estáticamente. Esto significa que cuando se compila un programa Go todo está incluido en un solo ejecutable binario, y no se requieren dependencias externas que tendrían que ser instaladas en la máquina remota. Esto facilita el despliegue de los programas Go, en comparación con el programa Python que utiliza bibliotecas externas y tienes que asegurarte tener todas esas bibliotecas instaladas en la máquina remota en la que se desea ejecutar.

Go es un lenguaje independiente de la plataforma, lo que significa que puede producir ejecutables binarios para todos los sistemas operativos (Linux, Windows, macOS, etc), además de ser muy fácil de hacer. Con Python, no es tan fácil crear estos ejecutables binarios para determinados sistemas operativos.

Go es un lenguaje muy eficiente, tiene una compilación rápida y un tiempo de ejecución rápido con un menor uso de recursos como la CPU y la memoria, especialmente en comparación con Python. Se han implementado numerosas optimizaciones en el lenguaje Go que lo hace muy eficiente. (Más abajo, en recursos puedes ver más detalles)

A diferencia de Python, que a menudo requiere el uso de bibliotecas de terceros para implementar un programa de Python en particular, Go incluye una biblioteca estándar que tiene la mayoría de las funcionalidades que se necesitan para DevOps construidas directamente en ella. Esto incluye la funcionalidad de procesamiento de archivos, servicios web HTTP, procesamiento JSON, soporte nativo para la concurrencia y el paralelismo, así como pruebas incorporadas. 👀

Con esto no se quiere decir de ninguna manera que debemos arrojar Python bajo un autobús, sólo se están dando razones subjetivas que hacen elegir Go al autor. No es el anterior "Go vs Python" lo que le hace escoger Go como lenguaje favorito, es porque tiene sentido para el trabajo en la empresa en la que trabaja y donde se desarrolla software en Go.

Como se ha comentado, una vez que se aprende el primer lenguaje de programación se hace más fácil de entender y aprender otros lenguajes. Probablemente nunca vas a tener un solo trabajo, en cualquier empresa o en cualquier proyecto, donde evites lidiar con la gestión, la arquitectura, la orquestación y la depuración de aplicaciones JavaScript y Node JS.

## Recursos

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

En los próximos 6 días de esta sección la intención es trabajar a través de algunos de los recursos mencionados anteriormente y documentar con notas cada día. Generalmente son alrededor de 3 horas para un curso completo, pues se ha compartido la lista completa que ha considerado el autor para avanzar y trabajar a través de cada uno si el tiempo lo permite.

Nos vemos en el [Día 8](day08.md).
