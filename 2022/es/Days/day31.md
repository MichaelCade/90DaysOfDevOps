## Modelos de computación de Microsoft Azure

Siguiendo con los conceptos básicos sobre los modelos de seguridad dentro de Microsoft Azure vamos a ver los diferentes servicios de computación disponibles en Azure.

### Opciones de Disponibilidad de Servicio

Esta sección es muy importante el autor por su papel en la gestión de datos. Al igual que en el on-premises, es crítico asegurar la disponibilidad de tus servicios.

- Alta Disponibilidad (Protección dentro de una región)
- Recuperación ante desastres (protección entre regiones)
- Copia de seguridad (Recuperación desde un punto en el tiempo)

Microsoft despliega múltiples regiones dentro de una frontera geopolítica. Dos conceptos con Azure para la Disponibilidad de Servicios con los conjuntos y zonas:
- **Conjuntos de Disponibilidad** - Proporcionan resiliencia dentro de un centro de datos.
- **Zonas de Disponibilidad** - Proporcionan resiliencia entre centros de datos dentro de una región.

### Máquinas virtuales

Muy probablemente es el punto de partida para cualquier persona en la nube pública.

- Proporciona una variedad de series y tamaños de MV con diferentes capacidades (Algunas abrumadoras). [Tamaños para máquinas virtuales en Azure](https://docs.microsoft.com/en-us/azure/virtual-machines/sizes)
- Hay muchas opciones y enfoques diferentes para MVs desde alto rendimiento, y baja latencia hasta MVs con opciones de alta memoria.
- También tenemos un tipo de MV burstable que se puede encontrar bajo la Serie B. Esto es ideal para las cargas de trabajo en las que puede tener un bajo requerimiento de CPU en su mayor parte, pero pueden requerir una vez al mes el requisito de un pico de rendimiento.
- Las máquinas virtuales se colocan en una red virtual que puede proporcionar conectividad a cualquier red.
- Compatibilidad con sistemas operativos invitados como Windows y Linux.
- También hay kernels ajustados a Azure cuando se trata de distribuciones específicas de Linux. [Azure Tuned Kernals](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/endorsed-distros#azure-tuned-kernels)

### Plantillas

Ya se ha mencionado antes que todo lo que hay detrás o debajo de Microsoft Azure es JSON.

Hay varios portales de gestión y consolas diferentes que podemos utilizar para crear nuestros recursos, la ruta preferida va a ser a través de plantillas JSON.

Despliegues idempotentes en modo incremental o completo, es decir, estado deseado repetible.

Hay una gran selección de plantillas que pueden exportar definiciones de recursos desplegados. Me gusta pensar en esta característica de plantillas como AWS CloudFormation o podría ser Terraform para una opción multi-nube. Cubriremos más sobre la potencia de Terraform en la sección de Infraestructura como código.

### Escalado

El escalado automático es una gran característica de la nube pública, siendo capaz de reducir los recursos no utilizados o aumentarlos cuando se necesiten.

En Azure, existe algo llamado Virtual Machine Scale Sets (VMSS) para IaaS. Esto permite la creación automática y la escala de una imagen estándar de oro basado en horarios y métricas.

Esto es ideal para actualizar ventanas de modo que pueda actualizar sus imágenes y desplegarlas con el menor impacto.

Otros servicios como Azure App Services tienen autoescalado integrado.

### Contenedores

No hemos cubierto los contenedores como un caso de uso, ni qué ni cómo deben ser necesarios en nuestro viaje de aprendizaje DevOps, pero tenemos que mencionar que Azure tiene algunos servicios específicos centrados en contenedores que son dignos de mención.
- **Azure Kubernetes Service (AKS)** - Proporciona una solución Kubernetes gestionada, sin necesidad de preocuparse por el plano de control o la gestión de clústeres subyacentes. También veremos más sobre Kubernetes más adelante.
- **Azure Container Instances** - Contenedores como servicio con facturación por segundos. Ejecute una imagen e intégrela con su red virtual, sin necesidad de Container Orchestration.
- **Service Fabric** - Tiene muchas capacidades pero incluye orquestación para instancias de contenedor.

Azure también tiene el Container Registry que proporciona un registro privado para Docker Images, Helm charts, OCI Artifacts e imágenes. Más sobre esto de nuevo cuando lleguemos a la sección correspondiente de contenedores.

También debemos mencionar que muchos de los servicios de contenedores también pueden aprovechar los contenedores bajo el capó, pero esto se abstrae de su necesidad de gestionar.

Estos servicios centrados en contenedores mencionados también encontramos servicios similares en todas las demás nubes públicas.

### Servicios de aplicaciones

- Azure Application Services ofrece una solución de alojamiento de aplicaciones que proporciona un método sencillo para establecer servicios.
- Despliegue y escalado automáticos.
- Admite soluciones basadas en Windows y Linux.
- Los servicios se ejecutan en un App Service Plan que tiene un tipo y un tamaño.
- Número de servicios diferentes que incluyen aplicaciones web, aplicaciones API y aplicaciones móviles.
- Soporte para ranuras de Despliegue para pruebas y promoción fiables.

### Computación serverless

El objetivo con serverless es que sólo pagamos por el tiempo de ejecución de la función y no tenemos que tener máquinas virtuales o aplicaciones PaaS en ejecución todo el tiempo. Simplemente ejecutamos nuestra función cuando la necesitamos y luego desaparece.

**Azure Functions** - Proporciona código serverless. Si nos remontamos a nuestro primer vistazo a la nube pública recordaremos la capa de abstracción de la gestión, con funciones serverless sólo vas a estar gestionando el código.

**Event-Driven** con escala masiva. Proporciona enlace de entrada y salida a muchos servicios de Azure y de terceros. 

Soporta muchos lenguajes de programación diferentes. (C#, NodeJS, Python, PHP, batch, bash, Golang y Rust. Cualquier ejecutable)

**Azure Event Grid** permite disparar la lógica desde servicios y eventos.

**Azure Logic App** proporciona workflows e integración basado en gráficos.

También podemos echar un vistazo a Azure Batch, que puede ejecutar trabajos a gran escala en nodos Windows y Linux con una gestión y programación coherentes.

## Recursos

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

Nos vemos en el [Día 32](day32.md).
