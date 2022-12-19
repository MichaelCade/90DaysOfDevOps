## Automatización de redes

### Fundamentos de la automatización de redes

Principales impulsores de la automatización de redes:

- Lograr agilidad
- Reducir costes
- Eliminar errores
- Garantizar el cumplimiento
- Gestión centralizada

El proceso de adopción de la automatización es específico de cada empresa. No hay una talla única para todos cuando se trata de desplegar la automatización, la capacidad de identificar y adoptar el enfoque que mejor funciona para su organización es fundamental para avanzar hacia el mantenimiento o la creación de un entorno más ágil, el foco debe estar siempre en el valor de negocio y la experiencia del usuario final. (Ya se comentó unos días atrás algo parecido sobre la relación de DevOps y el cambio cultural y el proceso automatizado que conlleva).

Para desglosar esto tendrías que identificar cómo la tarea o el proceso que estás tratando de automatizar va a lograr mejorar la experiencia del usuario final o el valor de negocio mientras sigues un enfoque sistemático paso a paso.

"Si no sabes adónde vas, cualquier camino te llevará allí".

Tenga un marco o una estructura de diseño que intente alcanzar, sepa cuál es su objetivo final y trabaje paso a paso para conseguirlo, midiendo el éxito de la automatización en las distintas fases en función de los resultados empresariales.

No hay necesidad de diseñar los conceptos de automatización en una burbuja porque deben aplicarse a su aplicación, su servicio y su infraestructura, así que empiece a construir los conceptos y a modelarlos en torno a su infraestructura y sus aplicaciones actuales.

### Enfoque de la automatización de redes

Debemos **identificar las tareas y las solicitudes de cambio de red** para que tener los problemas más comunes localizados para su posterior automatizar de la solución:

- Haga una lista de todas las solicitudes de cambio y flujos de trabajo que actualmente se abordan manualmente.
- Determine las actividades más comunes, que consumen más tiempo y son más propensas a errores.
- Priorice las solicitudes adoptando un enfoque orientado a la empresa.
- Este es el marco para construir un proceso de automatización, qué debe automatizarse y qué no.

A continuación, hay que **dividir las tareas y analizar** cómo funcionan e interactúan entre sí las distintas funciones de la red:

- El equipo de infraestructura/red recibe tickets de cambio en múltiples capas para desplegar aplicaciones.
- Basándose en los servicios de red, divídalos en diferentes áreas y comprenda cómo interactúan entre sí.
  - Optimización de aplicaciones
  - ADC - Application Delivery Controller (controlador de entrega de aplicaciones)
  - Firewalls
  - DDI (DNS, DHCP, IPAM, etc.)
  - Routing
  - Otros
- Identificar las distintas dependencias para abordar las diferencias empresariales y culturales e introducir la colaboración entre equipos.

**Políticas reutilizables**: Definir y simplificar tareas, procesos y entradas/salidas de servicios reutilizables.

- Definir ofertas para diversos servicios, procesos y entradas/salidas.
- Simplificar el proceso de implantación reducirá el tiempo de comercialización tanto de las cargas de trabajo nuevas como de las existentes.
- Una vez que se dispone de un proceso estándar, se puede secuenciar y alinear con las solicitudes individuales para un enfoque y una entrega multihilo.

**Combine las políticas con actividades específicas de la empresa**. ¿Cómo ayuda a la empresa la aplicación de esta política? ¿Ahorra tiempo? ¿Ahorra dinero? ¿Proporciona un mejor resultado empresarial?

- Asegúrese de que las tareas de servicio son interoperables.
- Asociar las tareas de servicio incrementales de modo que se alineen para crear servicios de negocio.
- Permitir la flexibilidad para asociar y reasociar tareas de servicio bajo demanda.
- Desplegar capacidades de autoservicio y allanar el camino para mejorar la eficiencia operativa.
- Permitir que los múltiples conjuntos de habilidades tecnológicas sigan contribuyendo con la supervisión y el cumplimiento.

**Iterar** en las políticas y procesos, añadiendo y mejorando al tiempo que se mantiene la disponibilidad y el servicio:

- Empezar poco a poco con la automatizando las tareas existentes.
- Familiarícese con el proceso de automatización, de modo que pueda identificar otras áreas que puedan beneficiarse de la automatización.
- Itere sus iniciativas de automatización, añadiendo agilidad de forma incremental mientras mantiene la disponibilidad requerida.
- Adoptar un enfoque gradual allana el camino hacia el éxito.

**Orquestar** el servicio de red:

- La automatización del proceso de despliegue es necesaria para ofrecer aplicaciones con rapidez.
- La creación de un entorno de servicio ágil requiere la gestión de diferentes elementos a través de conjuntos de competencias tecnológicas.
- Prepárese para una orquestación de extremo a extremo que permita controlar la automatización y el orden de las implantaciones.

## Herramientas de Automatización de Red

Una buena noticia es que, en su mayor parte, las herramientas que utilizamos para la automatización de redes son generalmente las mismas que utilizaremos para otras áreas de automatización. Lo que ya hemos cubierto hasta ahora o lo que cubriremos en futuras sesiones.

**Sistema Operativo** - Este reto se centra en hacer la mayor parte del aprendizaje con un sistema operativo Linux. Las razones se dieron en la sección de Linux, pero casi todas las herramientas que tocaremos, aunque muchas ya sean multi-plataforma, la mayoría comenzaron como aplicaciones basadas en Linux.

**Entorno de Desarrollo Integrado (IDE)** - De nuevo, no hay mucho que decir aquí aparte de que yo sugeriría Visual Studio Code como IDE, por sus gran cantidad de plugins disponibles para diferentes lenguajes.

**Gestión de la configuración** - Todavía no llegamos a la sección de gestión de la configuración, pero está muy claro que Ansible es uno de los favoritos en esta área para la gestión y automatización de configuraciones. Ansible está escrito en Python, pero no es necesario saber Python.

- Agentless (Sin agente)
- Sólo requiere SSH
- Gran comunidad de soporte
- Muchos módulos de red
- Modelo sólo push
- Configurado con YAML
- ¡Código abierto!

[Link to Ansible Network Modules](https://docs.ansible.com/ansible/2.9/modules/list_of_network_modules.html)

Tocaremos **Ansible Tower** en la sección de gestión de la configuración, es una interfaz gráfica de Ansible.

**CI/CD** - Vamos a cubrir más aspectos acerca de los conceptos y herramientas en torno a la integración y el despliegue continuo, pero es importante al menos mencionarlo aquí, ya que esto abarca la creación de redes y mucho más, como todo el aprovisionamiento de servicios y plataformas.

En particular, Jenkins es una herramienta popular para la automatización de redes.

- Monitorea el repositorio git en busca de cambios y luego los inicia.

**Control de versiones** - De nuevo algo en lo que profundizaremos más adelante.

- Git proporciona control de versiones de el código en tu dispositivo local y en remoto. 
- Multiplataforma
- GitHub, GitLab, Gitea, BitBucket, etc. Son sitios web en línea donde puedes guardar tus repositorios y subir código. También se conocen como forjas.

**Lenguaje | Scripting** - Un lenguaje de programación que no cubriremos es Python. Se escogió Go en su lugar como el lenguaje de programación por una visión subjetiva de las posibles necesidades a cubrir, pero fue una elección reñida de Golang con Python, ya que Python está muy presente en todas partes y además parece ser el ganador para la Automatización de Redes.

- Nornir merece tener se mención, es un marco de automatización escrito en Python. Este parece tomar el rol de Ansible pero específicamente alrededor de la Automatización de Redes. Echa  un vistazo a la [Documentación de Nornir](https://nornir.readthedocs.io/en/latest/)

**Analizar APIs** - Postman es una gran herramienta para analizar APIs RESTful. Ayuda a construir, probar y modificar APIs.

- POST >>> Para crear objetos de recursos.
- GET >>> Para recuperar un recurso.
- PUT >>> Para crear o reemplazar los recursos.
- PATCH >>> Para crear o actualizar un objeto de recurso.
- Delete >>> Para borrar un recurso

[Descarga de herramientas Postman](https://www.postman.com/downloads/)

### Otras herramientas a mencionar

- [Cisco NSO (Network Services Orchestrator - Orquestador de servicios de red)](https://www.cisco.com/c/en/us/products/cloud-systems-management/network-services-orchestrator/index.html)
- [NetYCE - Simplificar la Automatización de Redes](https://netyce.com/)
- [Automatización de pruebas de red](https://pubhub.devnetcloud.com/media/genie-feature-browser/docs/#/)

Durante los próximos 3 días nos pondremos manos a la obra con algunas de teorías que hemos cubierto sobre la red, así podremos trabajar un poco en torno a Python y automatización de red.

No hemos cubierto todos los temas de redes, tan solo se ha repasado la teoría lo justo y necesario como para tener una pequeña base de conceptos. Es aconsejable ampliar los conocimientos con los recursos que se agregan a continuación.

## Recursos

- [3 Necessary Skills for Network Automation](https://www.youtube.com/watch?v=KhiJ7Fu9kKA&list=WL&index=122&t=89s)
- [Computer Networking full course](https://www.youtube.com/watch?v=IPvYjXCsTg8)
- [Practical Networking](http://www.practicalnetworking.net/)
- [Python Network Automation](https://www.youtube.com/watch?v=xKPzLplPECU&list=WL&index=126)

Nos vemos el [Día 25](day25.md).
