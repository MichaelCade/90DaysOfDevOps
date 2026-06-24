## El panorama general: Gestión de registros

Continuando con los desafíos y soluciones de monitoreo de infraestructura, la gestión de registros es otra pieza del rompecabezas de la observabilidad en general.

### Gestión y agregación de registros

Hablemos de dos conceptos fundamentales, el primero de los cuales es la agregación de registros. Es una forma de recopilar y etiquetar los registros de aplicaciones de muchos servicios diferentes en un único panel de control que se pueda buscar fácilmente.

Uno de los primeros sistemas que se deben implementar en un sistema de gestión del rendimiento de aplicaciones es la agregación de registros. La gestión del rendimiento de aplicaciones es la parte del ciclo de vida de DevOps en la que se han construido e implementado las aplicaciones, y es necesario asegurarse de que funcionen de manera continua, tengan suficientes recursos asignados y no se muestren errores a los usuarios. En la mayoría de las implementaciones en producción, muchos eventos relacionados emiten registros a través de servicios. En Google, una única búsqueda puede llegar a diez servicios diferentes antes de ser devuelta al usuario. Si obtienes resultados de búsqueda inesperados, eso podría significar un problema lógico en cualquiera de los diez servicios, y la agregación de registros ayuda a empresas como Google a diagnosticar problemas en producción. Han creado un panel de control único en el que pueden asignar cada solicitud a un identificador único. Entonces, cada vez que una solicitud pasa por un servicio diferente, ese servicio asocia ese identificador a lo que está haciendo actualmente.

Esta es la esencia de una buena plataforma de agregación de registros: recopila eficientemente los registros de todos los servicios que los emiten y los hace fácilmente buscables en caso de una falla.

### Aplicación de ejemplo

Nuestra aplicación de ejemplo es una aplicación web que tiene un frontend y un backend típicos, y almacena nuestros datos críticos en una base de datos MongoDB.

Si un usuario nos informa que la página se puso completamente en blanco y se mostró un mensaje de error, nos costaría diagnosticar el problema con nuestra pila actual. El usuario tendría que enviarnos manualmente el error y tendríamos que buscar los registros relevantes en los otros tres servicios.

### ELK

Veamos ELK, una popular pila de agregación de registros de código abierto que recibe su nombre de sus tres componentes: Elasticsearch, Logstash y Kibana. Si lo instalamos en el mismo entorno que nuestra aplicación de ejemplo.

La aplicación web se conectaría al frontend, que a su vez se conecta al backend. El backend enviaría los registros a Logstash y así es como funcionan estos tres componentes.

### Los componentes de ELK

Elasticsearch, Logstash y Kibana son los componentes de ELK. Todos los servicios envían los registros a Logstash, que toma esos registros, que son texto emitido por la aplicación. Por ejemplo, en la aplicación web, cuando visitas una página web, la página web podría registrar el acceso del visitante a esa página en ese momento. Ese sería un ejemplo de un mensaje de registro. Esos registros se enviarían a Logstash.

Luego, Logstash extraería información de esos registros. Por ejemplo, para el mensaje de registro "El usuario hizo algo a las hora", Logstash extraería la hora, el mensaje y el usuario, y los incluiría como etiquetas. De esta manera, el mensaje sería un objeto con etiquetas y un mensaje, lo que permitiría buscarlos fácilmente y encontrar todas las solicitudes realizadas por un usuario específico. Sin embargo, Logstash no almacena los registros en sí mismo, los almacena en Elasticsearch, que es una base de datos eficiente para consultas de texto. Elasticsearch expone los resultados a Kibana, que es un servidor web que se conecta a Elasticsearch y permite a los administradores, como la persona encargada de DevOps o a otros miembros de tu equipo, ver los registros en producción cuando ocurre una falla importante. Como administrador, te conectarías a Kibana y Kibana consultaría Elasticsearch en busca de registros que coincidan con lo que deseas. 

Podrías decirle a Kibana en la barra de búsqueda: "quiero encontrar errores", y Kibana le diría a Elasticsearch que encuentre los mensajes que contengan la cadena "error", y luego Elasticsearch devolvería los resultados que Logstash ha registrado desde todos los demás servicios.

### ¿Cómo usaríamos ELK para diagnosticar un problema en producción?

Un usuario dice: "Vi el código de error 1234567 cuando intenté hacer esto". Con la configuración de ELK, tendríamos que ir a Kibana, ingresar 1234567 en la barra de búsqueda, presionar Enter, y eso nos mostraría los registros correspondientes a ese código. Uno de los registros podría decir "error interno del servidor, devolviendo 1234567", y veríamos que el servicio que emitió ese registro fue el backend, y también veríamos la hora en que se emitió ese registro. Así podríamos ir a esa hora en ese registro y ver los mensajes anteriores y posteriores a él en el backend, y así podríamos tener una imagen más clara de lo que sucedió en la solicitud del usuario. Podríamos repetir este proceso para otros servicios hasta encontrar la causa del problema para el usuario.

### Seguridad y acceso a los registros

Una pieza importante del rompecabezas es asegurarse de que solo los administradores (o los usuarios y grupos que necesitan acceder) puedan ver los registros. Los registros pueden contener información sensible como tokens, y solo los usuarios autenticados deberían tener acceso a ellos. No querrías exponer Kibana a Internet sin algún tipo de autenticación.

### Ejemplos de herramientas de gestión de registros

Algunos ejemplos de plataformas de gestión de registros son:

- Elasticsearch
- Logstash
- Kibana
- Fluentd: elección popular de código abierto
- Datadog: oferta alojada, comúnmente utilizada en grandes empresas
- LogDNA: oferta alojada
- Splunk

Los proveedores de servicios en la nube también ofrecen servicios de registro, como AWS CloudWatch Logs, Microsoft Azure Monitor y Google Cloud Logging.

La gestión de registros es un aspecto clave de la observabilidad general de tus aplicaciones y entornos de infraestructura para diagnosticar problemas en producción. Es relativamente sencillo instalar una solución llave en mano como ELK o CloudWatch, y facilita significativamente el diagnóstico y la solución de problemas en producción.

## Recursos

- [The Importance of Monitoring in DevOps](https://www.devopsonline.co.uk/the-importance-of-monitoring-in-devops/)
- [Understanding Continuous Monitoring in DevOps?](https://medium.com/devopscurry/understanding-continuous-monitoring-in-devops-f6695b004e3b)
- [DevOps Monitoring Tools](https://www.youtube.com/watch?v=Zu53QQuYqJ0)
- [Top 5 - DevOps Monitoring Tools](https://www.youtube.com/watch?v=4t71iv_9t_4)
- [How Prometheus Monitoring works](https://www.youtube.com/watch?v=h4Sl21AKiDg)
- [Introduction to Prometheus monitoring](https://www.youtube.com/watch?v=5o37CGlNLr8)
- [Promql cheat sheet with examples](https://www.containiq.com/post/promql-cheat-sheet-with-examples)
- [Log Management for DevOps | Manage application, server, and cloud logs with Site24x7](https://www.youtube.com/watch?v=J0csO_Shsj0)
- [Log Management what DevOps need to know](https://devops.com/log-management-what-devops-teams-need-to-know/)
- [What is ELK Stack?](https://www.youtube.com/watch?v=4X0WLg05ASw)
- [Fluentd simply explained](https://www.youtube.com/watch?v=5ofsNyHZwWE&t=14s)

Nos vemos en el [Día 80](day80.md)
