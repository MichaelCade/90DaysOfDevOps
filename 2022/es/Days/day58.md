## HashiCorp Configuration Language (HCL)

Antes de empezar a hacer cosas con Terraform, tenemos que adentrarnos un poco en el lenguaje de configuración de HashiCorp (HCL). Hasta ahora, durante nuestro desafío, hemos visto varios lenguajes de scripting y programación, y aquí tenemos otro. Hemos mencionado el [lenguaje de programación Go](day07.md), luego los [scripts bash](day19.md) e incluso un poco de Python cuando se trata de [automatización de redes](day27.md).

Ahora debemos cubrir el lenguaje de configuración de HashiCorp (HCL). Si esta es la primera vez que ves este lenguaje, puede parecer un poco intimidante, pero es bastante simple y muy potente.

A medida que avanzamos en esta sección, vamos a utilizar ejemplos que podemos ejecutar localmente en nuestro sistema, independientemente del sistema operativo que estés usando. Vamos a utilizar VirtualBox, aunque no sea la plataforma de infraestructura que normalmente se usa con Terraform. Sin embargo, ejecutar esto localmente es gratuito y nos permitirá lograr lo que buscamos en esta publicación. También podríamos extender este concepto a Docker o Kubernetes.

En general, sin embargo, deberías usar Terraform para implementar tu infraestructura en la nube pública (AWS, Google, Microsoft Azure), pero también en tus entornos de virtualización, como VMware, Microsoft Hyper-V, Nutanix AHV. En la nube pública, Terraform nos permite hacer mucho más que simplemente implementar máquinas virtuales. Podemos crear toda la infraestructura necesaria, como cargas de trabajo de PaaS y todos los activos de red requeridos, como VPC y grupos de seguridad.

Hay dos aspectos importantes en Terraform: el código, que vamos a explorar en esta publicación, y el estado. Ambos juntos podrían denominarse el núcleo de Terraform. Luego tenemos el entorno con el que deseamos interactuar e implementar, que se ejecuta mediante proveedores de Terraform, mencionados brevemente en la última sesión. Pero hay muchos proveedores disponibles, como AWS, Azure, etc. Hay cientos.

### Uso básico de Terraform

Echemos un vistazo a un archivo `.tf` de Terraform para ver cómo se componen. El primer ejemplo que veremos será código para implementar recursos en AWS, lo que también requeriría tener la CLI de AWS instalada en tu sistema y configurada para tu cuenta.

### Proveedores

En la parte superior de nuestra estructura de archivo `.tf`, generalmente llamada `main.tf`, al menos hasta que las cosas se vuelvan más complejas, vamos a definir los proveedores que mencionamos anteriormente. Nuestra fuente del proveedor de AWS, como puedes ver, es `hashicorp/aws`, lo que significa que el proveedor es mantenido o ha sido publicado por HashiCorp ellos mismos. Por defecto, haremos referencia a proveedores disponibles en el [Terraform Registry](https://registry.terraform.io/). También puedes escribir tus propios proveedores y utilizarlos localmente, o publicarlos en el Terraform Registry.


```
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}
```

We might also add in a region as well here to determine which AWS region we would like to provision to we can do this by adding the following:

```
provider "aws" {
  region = "ap-southeast-1" //region where resources need to be deployed
}
```

### Terraform Resources

- Another important component of a terraform config file which describes one or more infrastructure objects like EC2, Load Balancer, VPC, etc.

- A resource block declares a resource of a given type ("aws_instance") with a given local name ("90daysofdevops").

- The resource type and name together serve as an identifier for a given resource.

```
resource "aws_instance" "90daysofdevops" {
  ami               = data.aws_ami.instance_id.id
  instance_type     = "t2.micro"
  availability_zone = "us-west-2a"
  security_groups   = [aws_security_group.allow_web.name]
  user_data         = <<-EOF
                #! /bin/bash
                sudo yum update
                sudo yum install -y httpd
                sudo systemctl start httpd
                sudo systemctl enable httpd
                echo "
<h1>Deployed via Terraform</h1>

" | sudo tee /var/www/html/index.html
        EOF
  tags = {
    Name = "Created by Terraform"
  }
}
```

Puedes ver que también estamos ejecutando una actualización de `yum` e instalando `httpd` en nuestra instancia de EC2.

Si ahora observamos el archivo main.tf completo, podría lucir algo así.


```
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }

  required_version = ">= 0.14.9"
}

provider "aws" {
  profile = "default"
  region  = "us-west-2"
}

resource "aws_instance" "90daysofdevops" {
  ami           = "ami-830c94e3"
  instance_type = "t2.micro"
  availability_zone = "us-west-2a"
    user_data         = <<-EOF
                #! /bin/bash
                sudo yum update
                sudo yum install -y httpd
                sudo systemctl start httpd
                sudo systemctl enable httpd
                echo "
<h1>Deployed via Terraform</h1>

" | sudo tee /var/www/html/index.html
        EOF
  tags = {
    Name = "Created by Terraform"

  tags = {
    Name = "ExampleAppServerInstance"
  }
}
```

El código anterior desplegará un servidor web muy sencillo como una instancia EC2 en AWS. Lo genial de esto y de cualquier otra configuración similar es que podemos repetirlo y obtener la misma salida cada vez. Aparte de la posibilidad de que haya cometido un error en el código, no hay interacción humana con lo anterior.

Podemos echar un vistazo a un ejemplo super sencillo, uno que probablemente nunca usarás, pero vamos a hacerle honor de todos modos. Como en todos los buenos lenguajes de scripting y programación, deberíamos comenzar con un escenario de "Hola mundo".

```
terraform {
  # This module is now only being tested with Terraform 0.13.x. However, to make upgrading easier, we are setting
  # 0.12.26 as the minimum version, as that version added support for required_providers with source URLs, making it
  # forwards compatible with 0.13.x code.
  required_version = ">= 0.12.26"
}

# website::tag::1:: The simplest possible Terraform module: it just outputs "Hello, World!"
output "hello_world" {
  value = "Hello, 90DaysOfDevOps from Terraform"
}
```

Encontrarás este archivo en la carpeta IAC dentro de hello-world, pero de entrada, esto no funcionará simplemente, hay algunos comandos que debemos ejecutar para usar nuestro código de Terraform.

En tu terminal, navega hasta la carpeta donde se haya creado el archivo main.tf, esto puede ser desde este repositorio o puedes crear uno nuevo usando el código anterior.

Cuando estés en esa carpeta, ejecutaremos `terraform init`.

Necesitamos realizar esto en cualquier directorio donde tengamos, o antes de ejecutar cualquier código de Terraform. La inicialización de un directorio de configuración descarga e instala los proveedores definidos en la configuración. En este caso, no tenemos proveedores, pero en el ejemplo anterior, esto descargaría el proveedor de AWS para esta configuración.

![Day58_IAC1](Images/Day58_IAC1.png)

El siguiente comando será `terraform plan`.

El comando `terraform plan` crea un plan de ejecución que te permite previsualizar los cambios que Terraform planea realizar en tu infraestructura.

Puedes ver simplemente a continuación que con nuestro ejemplo de hello-world veremos la salida. Si esto fuera una instancia EC2 de AWS, veríamos todos los pasos que se crearán.

![Day58_IAC2](Images/Day58_IAC2.png)

En este punto, hemos inicializado nuestro repositorio y hemos descargado los proveedores donde sea necesario. Hemos realizado una prueba para asegurarnos de que esto es lo que queremos ver, así que ahora podemos ejecutar y desplegar nuestro código.

`terraform apply` nos permite hacer esto. Hay una medida de seguridad incorporada en este comando, y nuevamente te mostrará un plan de lo que va a suceder, lo cual requiere una respuesta de tu parte para decir que sí, para continuar.

![Day58_IAC3](Images/Day58_IAC3.png)

Cuando escribimos "yes" para ingresar un valor, nuestro código se despliega. No es muy emocionante, pero puedes ver que tenemos la salida que definimos en nuestro código.

![Day58_IAC4](Images/Day58_IAC4.png)

Ahora, no hemos desplegado nada, no hemos agregado, cambiado o destruido nada, pero si lo hiciéramos, también lo veríamos indicado en lo anterior. Sin embargo, si hubiéramos desplegado algo y quisiéramos eliminar todo lo que hemos creado, podemos usar el comando `terraform destroy`. Nuevamente, este tiene esa medida de seguridad donde tienes que escribir "yes". Aunque puedes usar `--auto-approve` al final de tus comandos `apply` y `destroy` para evitar esa intervención manual. Pero te aconsejaría que solo uses este atajo cuando estés aprendiendo y probando, ya que todo desaparecerá a veces más rápido de lo que se construyó.

A partir de esto, hemos cubierto 4 comandos de la CLI de Terraform.

- `terraform init` = prepara la carpeta del proyecto con los proveedores.
- `terraform plan` = muestra lo que se va a crear y cambiar durante el próximo comando, según nuestro código.
- `terraform apply` = desplegará los recursos definidos en nuestro código.
- `terraform destroy` = destruirá los recursos que hemos creado en nuestro proyecto.

También hemos cubierto dos aspectos importantes de nuestros archivos de código.

- Proveedores = cómo Terraform se comunica con la plataforma final a través de APIs.
- Recursos = lo que queremos desplegar con código.

Otra cosa a tener en cuenta al ejecutar `terraform init` es observar el árbol de la carpeta antes y después para ver qué sucede y dónde almacenamos los proveedores y los módulos.

### Estado de Terraform

También debemos tener en cuenta el archivo de estado que se crea dentro de nuestro directorio. Para este ejemplo de hello world, nuestro archivo de estado es simple. Se trata de un archivo JSON que representa el estado del mundo según Terraform. El estado mostrará felizmente tus datos confidenciales, así que ten cuidado y, como buena práctica, agrega tus archivos `.tfstate` a tu archivo `.gitignore` antes de subirlos a GitHub.

Por defecto, el archivo de estado, como puedes ver, se encuentra dentro del mismo directorio que tu código de proyecto, pero también se puede almacenar de forma remota como opción. En un entorno de producción, es probable que esto se almacene en un lugar compartido, como un bucket de S3.

Otra opción podría ser Terraform Cloud, que es un servicio gestionado de pago (gratuito hasta 5 usuarios).

Las ventajas de almacenar el estado en un lugar remoto son:

- Datos confidenciales encriptados.
- Colaboración.
- Automatización.

Sin embargo, esto podría aumentar la complejidad.

```JSON
{
  "version": 4,
  "terraform_version": "1.1.6",
  "serial": 1,
  "lineage": "a74296e7-670d-0cbb-a048-f332696ca850",
  "outputs": {
    "hello_world": {
      "value": "Hello, 90DaysOfDevOps from Terraform",
      "type": "string"
    }
  },
  "resources": []
}
```

## Recursos

A continuación, he enumerado muchos recursos y creo que este tema se ha cubierto muchas veces en otros lugares. Si tienes recursos adicionales, asegúrate de enviar una solicitud de extracción (PR) con tus recursos y estaré encantado de revisarlos y agregarlos a la lista.

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

Nos vemos en el [Día 59](day59.md)
