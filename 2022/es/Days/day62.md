## Pruebas, herramientas y alternativas

Al concluir esta sección sobre la infraestructura como código, debemos mencionar la importancia de probar nuestro código, las diversas herramientas disponibles y algunas alternativas a Terraform para lograr esto. Como mencioné al inicio de la sección, mi enfoque fue en Terraform porque es gratuito y de código abierto, además de ser multiplataforma y agnóstico en cuanto a entornos. Sin embargo, existen alternativas que deben considerarse, pero el objetivo general es crear conciencia de que esta es la forma de implementar la infraestructura.

### Deterioro del código

La primera área que quiero abordar en esta sesión es el deterioro del código. A diferencia del código de aplicaciones, el código de la infraestructura como código puede ser utilizado y luego quedar sin uso durante mucho tiempo. Tomemos como ejemplo el uso de Terraform para implementar nuestro entorno de máquinas virtuales en AWS. Funciona perfectamente la primera vez y tenemos nuestro entorno, pero este entorno no cambia con mucha frecuencia, por lo que el código queda en ese estado, posiblemente almacenado en un lugar central, pero sin cambios.

¿Qué sucede si algo cambia en la infraestructura? Pero se realiza fuera del flujo normal, o si otras cosas cambian en nuestro entorno.

- Cambios externos al flujo normal.
- Versiones no fijadas.
- Dependencias obsoletas.
- Cambios no aplicados.

### Testing

Otra área importante que se deriva del deterioro del código, en general, es la capacidad de probar nuestro código de infraestructura como código y asegurarnos de que todas las áreas funcionen como deberían.

Aquí hay algunos comandos de prueba incorporados que podemos utilizar:

| ComandO              | Descripción                                                                                       |
| -------------------- | ------------------------------------------------------------------------------------------------- |
| `terraform fmt`      | Reescribe los archivos de configuración de Terraform en un formato y estilo canónicos.            |
| `terraform validate` | Valida los archivos de configuración en un directorio, refiriéndose solo a la configuración.      |
| `terraform plan`     | Crea un plan de ejecución que te permite previsualizar los cambios que Terraform planea realizar. |
| Custom validation    | Validación de las variables de entrada para asegurarte de que coincidan con lo esperado.          |

También existen herramientas de prueba externas a Terraform:

- [tflint](https://github.com/terraform-linters/tflint)

- Encuentra posibles errores.
- Advierte sobre sintaxis obsoletas y declaraciones no utilizadas.
- Aplica las mejores prácticas y convenciones de nomenclatura.

Herramientas de escaneo:

- [checkov](https://www.checkov.io/) - escanea las configuraciones de infraestructura en la nube en busca de configuraciones incorrectas antes de implementarlas.
- [tfsec](https://aquasecurity.github.io/tfsec/v1.4.2/) - analizador de seguridad de análisis estático para tu código de Terraform.
- [terrascan](https://github.com/accurics/terrascan) - analizador de código estático para infraestructura como código.
- [terraform-compliance](https://terraform-compliance.com/) - un marco de prueba ligero centrado en la seguridad y el cumplimiento para Terraform, que permite la capacidad de prueba negativa para tu infraestructura como código.
- [snyk](https://docs.snyk.io/products/snyk-infrastructure-as-code/scan-terraform-files/scan-and-fix-security-issues-in-terraform-files) - escanea tu código de Terraform en busca de configuraciones incorrectas y problemas de seguridad.

Oferta de Cloud gestionada:

- [Terraform Sentinel](https://www.terraform.io/cloud-docs/sentinel) - marco de política como código integrado con los productos empresariales de HashiCorp. Permite decisiones de políticas lógicas detalladas y se puede ampliar para usar información de fuentes externas.

Pruebas automatizadas:

- [Terratest](https://terratest.gruntwork.io/) - Terratest es una biblioteca de Go que proporciona patrones y funciones auxiliares para probar la infraestructura.

Mencionado aparte:

- [Terraform Cloud](https://cloud.hashicorp.com/products/terraform) - Terraform Cloud es el servicio administrado de HashiCorp. Elimina la necesidad de herramientas y documentación innecesarias para que los profesionales, equipos y organizaciones utilicen Terraform en producción.

- [Terragrunt](https://terragrunt.gruntwork.io/) - Terragrunt es una capa delgada que proporciona herramientas adicionales para mantener las configuraciones DRY, trabajar con varios módulos de Terraform y gestionar el estado remoto.

- [Atlantis](https://www.runatlantis.io/) - Automatización de solicitudes de extracción de Terraform.

### Alternativas

Mencionamos en el Día 57 al comenzar esta sección que existen algunas alternativas, y tengo la intención de explorar esto más a fondo después de este desafío.

| Específico de la nube           | Agnóstico de la nube |
| ------------------------------- | -------------------- |
| AWS CloudFormation              | Terraform            |
| Azure Resource Manager          | Pulumi               |
| Google Cloud Deployment Manager |                      |

He utilizado AWS CloudFormation probablemente más que cualquier otra de las opciones anteriores, ya que es nativo de AWS, pero no he usado las demás alternativas aparte de Terraform. Como puedes imaginar, las versiones específicas de la nube son muy buenas en esa nube en particular, pero si tienes múltiples entornos en la nube, te resultará difícil migrar esas configuraciones o tendrás múltiples planos de administración para tus esfuerzos de infraestructura como código.

Creo que el siguiente paso interesante para mí es dedicar tiempo a aprender más sobre [Pulumi](https://www.pulumi.com/)

Desde una comparación de Pulumi en su sitio web:

> "Tanto Terraform como Pulumi ofrecen el modelo de infraestructura como código de estado deseado, donde el código representa el estado deseado de la infraestructura y el motor de implementación compara este estado deseado con el estado actual del stack y determina qué recursos deben crearse, actualizarse o eliminarse".

La diferencia más notable que puedo ver es que, a diferencia del lenguaje de configuración de HashiCorp (HCL), Pulumi permite utilizar lenguajes de propósito general como Python, TypeScript, JavaScript, Go y .NET.

Una descripción general rápida: [Introducción a Pulumi: Infraestructura moderna como código.](https://www.youtube.com/watch?v=QfJTJs24-JM) Me gusta la facilidad y las opciones que se ofrecen, y quiero explorar esto un poco más.

Esto concluye la sección de infraestructura como código y a continuación pasaremos a esa pequeña superposición con la administración de la configuración, en particular, a medida que avanzamos más allá de la imagen general de la administración de la configuración, utilizaremos Ansible para algunas de esas tareas y demostraciones.

## Recursos

A continuación, he enumerado muchos recursos y creo que este tema se ha cubierto muchas veces. Si tienes recursos adicionales, asegúrate de enviar una solicitud de extracción con tus recursos y estaré encantado de revisarlos y agregarlos a la lista.

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
- [Pulumi - IaC in your favorite programming language!](https://www.youtube.com/watch?v=vIjeiDcsR3Q&t=51s)
- [Herramientas para Terraform](https://vergaracarmona.es/herramientas-para-terraform/)

Nos vemos en el [Día 63](day63.md)
