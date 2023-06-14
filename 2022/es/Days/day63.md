## El panorama general: Gestión de la configuración

Continuando directamente desde la sección que cubre la Infraestructura como Código, es probable que haya cierta superposición al hablar de la Gestión de la Configuración o la Gestión de la Configuración de Aplicaciones.

La Gestión de la Configuración es el proceso de mantener las aplicaciones, sistemas y servidores en el estado deseado. La superposición con la Infraestructura como Código es que IaC se asegurará de que tu infraestructura esté en el estado deseado, pero después de eso, especialmente Terraform, no se encargará del estado deseado de la configuración de tu sistema operativo o aplicación, y ahí es donde entran en juego las herramientas de Gestión de la Configuración. Aseguran que el sistema y las aplicaciones funcionen como se espera a medida que se producen cambios con el tiempo.

La gestión de la configuración te evita realizar cambios pequeños o grandes que no estén documentados.

### Escenario: ¿Por qué querrías usar la Gestión de la Configuración?

El escenario o por qué querrías usar la Gestión de la Configuración. Conoce a Dean, nuestro administrador del sistema, y Dean es un feliz y competente trabajador que se ocupa de todos los sistemas en su entorno.

¿Qué sucede si su sistema falla, si hay un incendio, si un servidor se cae? Dean sabe exactamente qué hacer, puede solucionar ese problema fácilmente. Sin embargo, los problemas se vuelven difíciles para Dean cuando varios servidores comienzan a fallar, especialmente cuando tienes entornos grandes y en expansión. Por eso Dean necesita tener una herramienta de gestión de la configuración. Las herramientas de Gestión de la Configuración pueden ayudar a que Dean parezca un experto, todo lo que tiene que hacer es configurar los códigos correctos que le permitan enviar las instrucciones sobre cómo configurar cada uno de los servidores rápida y efectivamente, y a gran escala.

### Herramientas de Gestión de la Configuración

Existen diversas herramientas de gestión de la configuración disponibles, y cada una tiene características específicas que la hacen más adecuada para algunas situaciones que otras.

![](Images/Day63_config1.png)

En esta etapa, echaremos un vistazo rápido a las opciones en la imagen anterior antes de tomar nuestra decisión sobre cuál utilizaremos y por qué.

- **Chef**

  - Chef se asegura de que la configuración se aplique de manera consistente en cada entorno, a cualquier escala, con automatización de infraestructura.
  - Chef es una herramienta de código abierto desarrollada por OpsCode escrita en Ruby y Erlang.
  - Chef es más adecuado para organizaciones que tienen una infraestructura heterogénea y buscan soluciones maduras.
  - Las recetas y libros de cocina determinan el código de configuración para tus sistemas.
  - Ventajas: hay una amplia colección de recetas disponible.
  - Ventajas: se integra bien con Git, lo que proporciona un sólido control de versiones.
  - Desventajas: tiene una curva de aprendizaje pronunciada, se requiere una cantidad considerable de tiempo.
  - Desventajas: el servidor principal no tiene mucho control.
  - Arquitectura: Servidor / Clientes
  - Facilidad de configuración: Moderada
  - Lenguaje: Procedimental: especifica cómo hacer una tarea

- **Puppet**

  - Puppet es una herramienta de gestión de configuración que admite implementación automática.
  - Puppet está construido en Ruby y utiliza un DSL para escribir manifiestos.
  - Puppet también funciona bien con infraestructuras heterogéneas donde el enfoque se centra en la escalabilidad.
  - Ventajas: gran comunidad de soporte.
  - Ventajas: mecanismo de informes bien desarrollado.
  - Desventajas: las tareas avanzadas requieren conocimientos del lenguaje Ruby.
  - Desventajas: el servidor principal no tiene mucho control.
  - Arquitectura: Servidor / Clientes
  - Facilidad de configuración: Moderada
  - Lenguaje: Declarativo: especifica solo qué hacer

- **Ansible**

  - Ansible es una herramienta de automatización de TI que automatiza la gestión de la configuración, la provisión en la nube, la implementación y la orquestación.
  - El núcleo de los libros de jugadas de Ansible está escrito en YAML. (Deberíamos hacer una sección sobre YAML, ya que lo hemos visto varias veces).
  - Ansible funciona bien cuando hay entornos que se centran en poner las cosas en marcha rápidamente.
  - Trabaja con libros de jugadas que proporcionan instrucciones a tus servidores.
  - Ventajas: no se necesitan agentes en los nodos remotos.
  - Ventajas: YAML es fácil de aprender.
  - Desventajas: la velocidad de rendimiento suele ser menor que la de otras herramientas (más rápido que hacerlo   manualmente Dean).
  - Desventajas: YAML no es tan poderoso como Ruby, pero tiene una curva de aprendizaje menor.
  - Arquitectura: Solo cliente  
  - Facilidad de configuración: Muy fácil
  - Lenguaje: Procedimental: especifica cómo hacer una tarea

- **SaltStack**

  - SaltStack es una herramienta basada en CLI que automatiza la gestión de la configuración y la ejecución remota.
  - SaltStack está basado en Python, mientras que las instrucciones se escriben en YAML o su DSL.
  - Perfecto para entornos con prioridad en la escalabilidad y la resiliencia.
  - Ventajas: fácil de usar una vez que está en funcionamiento.
  - Ventajas: buen mecanismo de informes.
  - Desventajas: la fase de configuración es difícil.
  - Desventajas: nueva interfaz web menos desarrollada que las demás.
  - Arquitectura: Servidor / Clientes
  - Facilidad de configuración: Moderada
  - Lenguaje: Declarativo: especifica solo qué hacer

### Ansible vs Terraform

La herramienta que utilizaremos para esta sección será Ansible (fácil de usar y requiere nociones básicas del lenguaje).

Creo que es importante mencionar algunas diferencias entre Ansible y Terraform antes de adentrarnos un poco más en las herramientas.

|                  | Ansible                                                                 | Terraform                                                                   |
| ---------------- | ----------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| Tipo             | Ansible es una herramienta de gestión de configuración                  | Terraform es una herramienta de orquestación                                |
| Infraestructura  | Ansible admite infraestructura mutable                                  | Terraform admite infraestructura inmutable                                  |
| Lenguaje         | Ansible sigue un lenguaje procedural                                    | Terraform sigue un lenguaje declarativo                                     |
| Provisionamiento | Ansible proporciona aprovisionamiento parcial (VM, red, almacenamiento) | Terraform proporciona un amplio aprovisionamiento (VM, red, almacenamiento) |
| Empaquetado      | Ansible ofrece soporte completo para empaquetado y plantillas           | Terraform ofrece soporte parcial para empaquetado y plantillas              |
| Ciclo de vida    | Ansible no tiene gestión del ciclo de vida                              | Terraform depende en gran medida del ciclo de vida y la gestión del estado  |


## Recursos

- [What is Ansible](https://www.youtube.com/watch?v=1id6ERvfozo)
- [Ansible 101 - Episode 1 - Introduction to Ansible](https://www.youtube.com/watch?v=goclfp6a2IQ)
- [NetworkChuck - You need to learn Ansible right now!](https://www.youtube.com/watch?v=5hycyr-8EKs&t=955s)
- [Chef vs Puppet vs Ansible vs Saltstack](https://vergaracarmona.es/chef-vs-puppet-vs-ansible-vs-saltstack/)

Nos vemos en el [Día 64](day64.md)
