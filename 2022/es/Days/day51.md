## Despliegue del primer clúster de Kubernetes

En esta publicación, vamos a poner en marcha un clúster de Kubernetes en nuestra máquina local utilizando minikube. Esto nos proporcionará un clúster de Kubernetes básico para el resto de la sección de Kubernetes, aunque también veremos cómo implementar un clúster de Kubernetes en VirtualBox más adelante. La razón por la que elegimos este método en lugar de crear un clúster de Kubernetes administrado en la nube pública es que esto va a costar dinero incluso con el nivel gratuito. Sin embargo, compartí algunos blogs en la sección anterior [Día 50](day50.md) si deseas crear ese entorno.

### ¿Qué es Minikube?

> "Minikube configura rápidamente un clúster local de Kubernetes en macOS, Linux y Windows. Nos enfocamos en ayudar a los desarrolladores de aplicaciones y a los nuevos usuarios de Kubernetes".

Es posible que no encajes en la descripción anterior, pero he descubierto que Minikube es una excelente herramienta si solo quieres probar algo con Kubernetes. Puedes implementar fácilmente una aplicación y también tienen algunos complementos increíbles que también cubriré.

Para empezar, independientemente de tu sistema operativo de trabajo, puedes ejecutar Minikube. Primero, dirígete a la [página del proyecto aquí](https://minikube.sigs.k8s.io/docs/start/). La primera opción que tienes es elegir el método de instalación. No utilicé este método, pero puedes elegir usarlo en lugar del mío (mi método se explicará a continuación).

A continuación, se indica que debes tener "contenedores o administradores de máquinas virtuales, como Docker, HyperKit, Hyper-V, KVM, Parallels, Podman, VirtualBox o VMware". Aquí es donde se ejecutará Minikube, y la opción más sencilla, a menos que se indique lo contrario en el repositorio que estoy utilizando, es Docker. Puedes instalar Docker en tu sistema utilizando el siguiente [enlace](https://docs.docker.com/get-docker/).

![Day51_Kubernetes1](Images/Day51_Kubernetes1.png)

### Mi forma de instalar Minikube y otros requisitos previos...

He estado utilizando Arkade durante algún tiempo para obtener todas esas herramientas y CLIs de Kubernetes. Puedes ver los pasos de instalación en este [repositorio de GitHub](https://github.com/alexellis/arkade) para comenzar con Arkade. También he mencionado esto en otras publicaciones de blog donde necesitaba instalar algo. La simplicidad de simplemente ejecutar `arkade get` y luego verificar si tu herramienta o CLI está disponible es muy útil. En la sección de Linux, hablamos sobre el administrador de paquetes y el proceso para obtener nuestro software. Puedes pensar en Arkade como ese mercado de aplicaciones y CLIs para Kubernetes. Es una herramienta muy útil para tener en tus sistemas, escrita en Golang y multiplataforma.

![Day51_Kubernetes2](Images/Day51_Kubernetes2.png)

Como parte de la larga lista de aplicaciones disponibles en Arkade, Minikube es una de ellas. Por lo tanto, con un simple comando `arkade get minikube`, ahora estamos descargando el binario y listos para continuar.

![Day51_Kubernetes3](Images/Day51_Kubernetes3.png)

También necesitaremos kubectl como parte de nuestras herramientas, por lo que también puedes obtenerlo a través de Arkade o creo que la documentación de Minikube lo descarga mediante los comandos curl mencionados anteriormente. Hablaremos más sobre kubectl más adelante en la publicación.

### Poner en marcha un clúster de Kubernetes

En esta sección en particular, quiero cubrir las opciones disponibles para poner en marcha un clúster de Kubernetes en tu máquina local. Simplemente podríamos ejecutar el siguiente comando y se crearía un clúster para que lo utilices.

Minikube se utiliza en la línea de comandos y, simplemente, una vez que tienes todo instalado, puedes ejecutar `minikube start` para implementar tu primer clúster de Kubernetes. A continuación, verás que el controlador de Docker es el predeterminado para ejecutar nuestro nodo de virtualización anidado. Mencioné al principio de la publicación las otras opciones disponibles, las cuales te ayudarán cuando desees ampliar cómo debe verse este clúster de Kubernetes local.

Un clúster de Minikube consistirá en un solo contenedor de Docker en este caso, el cual tendrá el nodo del plano de control y el nodo de trabajo en una sola instancia. Por lo general, separaríamos esos nodos. Esto es algo que cubriremos en la siguiente sección, donde veremos entornos de Kubernetes tipo laboratorio doméstico, pero un poco más cercanos a la arquitectura de producción.

![Day51_Kubernetes4](Images/Day51_Kubernetes4.png)

He mencionado esto varias veces antes, me gusta Minikube debido a los complementos disponibles, la capacidad de implementar un clúster con un simple comando, incluidos todos los complementos requeridos desde el principio, me ayuda a implementar la misma configuración requerida cada vez.

A continuación, puedes ver una lista de esos complementos. Por lo general, utilizo los complementos `CSI-host path-driver` y `volumesnapshots`, pero puedes ver la lista completa a continuación. Estos complementos generalmente se pueden implementar utilizando Helm, algo que también cubriremos más adelante en la sección de Kubernetes, pero esto hace las cosas mucho más simples.

![Day51_Kubernetes5](Images/Day51_Kubernetes5.png)

También estoy definiendo en nuestro proyecto alguna configuración adicional. El servidor de API se establece en 6433 en lugar de un puerto API aleatorio, y también defino el tiempo de ejecución del contenedor como containerd. Sin embargo, Docker es el valor predeterminado y CRI-O también está disponible. También estoy estableciendo una versión específica de Kubernetes.

![Day51_Kubernetes6](Images/Day51_Kubernetes6.png)

Ahora estamos listos para implementar nuestro primer clúster de Kubernetes utilizando Minikube. Sin embargo, mencioné anteriormente que también necesitarás `kubectl` para interactuar con tu clúster. Puedes obtener kubectl instalado mediante Arkade con el comando `arkade get kubectl`

![Day51_Kubernetes7](Images/Day51_Kubernetes7.png)

o puedes descargarlo para varias plataformas desde los siguientes enlaces:

- [Linux](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux)
- [macOS](https://kubernetes.io/docs/tasks/tools/install-kubectl-macos)
- [Windows](https://kubernetes.io/docs/tasks/tools/install-kubectl-windows)

Una vez que hayas instalado kubectl, puedes interactuar con tu clúster con un simple comando como `kubectl get nodes`

![Day51_Kubernetes8](Images/Day51_Kubernetes8.png)

### ¿Qué es kubectl?

Ahora tenemos nuestro clúster minikube | Kubernetes en funcionamiento y te he pedido que instales tanto Minikube, donde he explicado al menos qué hace, pero no he explicado qué es kubectl y qué hace.

kubectl es una interfaz de línea de comandos que se utiliza para interactuar con clústeres de Kubernetes. Lo estamos utilizando aquí para interactuar con nuestro clúster minikube, pero también usaríamos kubectl para interactuar con nuestros clústeres empresariales en la nube pública.

Usamos kubectl para implementar aplicaciones, inspeccionar y gestionar recursos del clúster. Puedes encontrar una [Descripción general de kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) mucho más detallada aquí en la documentación oficial de Kubernetes.

kubectl interactúa con el servidor de API que se encuentra en el nodo de plano de control, del cual hablamos brevemente en una publicación anterior.

### Hoja de referencia de kubectl

Además de la documentación oficial, también he encontrado que siempre tengo abierta esta página cuando busco comandos de kubectl. [Kubernetes no oficial](https://unofficial-kubernetes.readthedocs.io/en/latest/)

| Listar recursos            |                                                                             |
| -------------------------- | --------------------------------------------------------------------------- |
| kubectl get nodes          | Listar todos los nodos en el clúster                                        |
| kubectl get namespaces     | Listar todos los espacios de nombres en el clúster                          |
| kubectl get pods           | Listar todas las pods en el clúster en el espacio de nombres predeterminado |
| kubectl get pods -n nombre | Listar todas las pods en el espacio de nombres "nombre"                     |

| Crear recursos                  |                                                     |
| ------------------------------- | --------------------------------------------------- |
| kubectl create namespace nombre | Crear un espacio de nombres llamado "nombre"        |
| kubectl create -f [archivo]     | Crear un recurso a partir de un archivo JSON o YAML |

| Editar recursos                  |                    |
| -------------------------------- | ------------------ |
| kubectl edit svc/nombre-servicio | Editar un servicio |

| Más detalles sobre los recursos |                                                                |
| ------------------------------- | -------------------------------------------------------------- |
| kubectl describe nodes          | mostrar el estado de cualquier cantidad de recursos en detalle |

| Eliminar recursos  |                                                                           |
| ------------------ | ------------------------------------------------------------------------- |
| kubectl delete pod | Eliminar recursos, puede ser desde la entrada estándar o desde un archivo |

Te encontrarás deseando conocer los nombres cortos para algunos de los comandos de kubectl. Por ejemplo, `-n` es el nombre corto para `namespace`, lo que facilita escribir un comando y, si estás escribiendo scripts, puedes tener un código más ordenado.

| Nombre corto | Nombre completo            |
| ------------ | -------------------------- |
| csr          | certificatesigningrequests |
| cs           | componentstatuses          |
| cm           | configmaps                 |
| ds           | daemonsets                 |
| deploy       | deployments                |
| ep           | endpoints                  |
| ev           | events                     |
| hpa          | horizontalpodautoscalers   |
| ing          | ingresses                  |
| limits       | limitranges                |
| ns           | namespaces                 |
| no           | nodes                      |
| pvc          | persistentvolumeclaims     |
| pv           | persistentvolumes          |
| po           | pods                       |
| pdb          | poddisruptionbudgets       |
| psp          | podsecuritypolicies        |
| rs           | replicasets                |
| rc           | replicationcontrollers     |
| quota        | resourcequotas             |
| sa           | serviceaccounts            |
| svc          | services                   |

Lo último que quiero añadir aquí es que he creado otro proyecto en torno a minikube para ayudarme a crear rápidamente entornos de demostración para mostrar servicios de datos y proteger esas cargas de trabajo con Kasten K10. Puedes encontrar [Project Pace](https://github.com/MichaelCade/project_pace) allí y me encantaría recibir tus comentarios o interactuar contigo. También muestra o incluye algunas formas automatizadas de implementar tus clústeres de minikube y crear diferentes aplicaciones de servicios de datos.

A continuación, pasaremos a implementar varios nodos en máquinas virtuales utilizando VirtualBox, pero vamos a utilizar la opción fácil como lo hicimos en la sección de Linux, donde usamos Vagrant para crear rápidamente las máquinas y desplegar nuestro software según lo deseemos.

Añadí esta lista a la publicación de ayer, que son blogs explicativos que he hecho sobre la implementación de diferentes clústeres de Kubernetes.


- [Kubernetes playground – How to choose your platform](https://vzilla.co.uk/vzilla-blog/building-the-home-lab-kubernetes-playground-part-1)
- [Kubernetes playground – Setting up your cluster](https://vzilla.co.uk/vzilla-blog/building-the-home-lab-kubernetes-playground-part-2)
- [Getting started with Amazon Elastic Kubernetes Service (Amazon EKS)](https://vzilla.co.uk/vzilla-blog/getting-started-with-amazon-elastic-kubernetes-service-amazon-eks)
- [Getting started with Microsoft Azure Kubernetes Service (AKS)](https://vzilla.co.uk/vzilla-blog/getting-started-with-microsoft-azure-kubernetes-service-aks)
- [Getting Started with Microsoft AKS – Azure PowerShell Edition](https://vzilla.co.uk/vzilla-blog/getting-started-with-microsoft-aks-azure-powershell-edition)
- [Getting started with Google Kubernetes Service (GKE)](https://vzilla.co.uk/vzilla-blog/getting-started-with-google-kubernetes-service-gke)
- [Kubernetes, How to – AWS Bottlerocket + Amazon EKS](https://vzilla.co.uk/vzilla-blog/kubernetes-how-to-aws-bottlerocket-amazon-eks)
- [Getting started with CIVO Cloud](https://vzilla.co.uk/vzilla-blog/getting-started-with-civo-cloud)
- [Minikube - Kubernetes Demo Environment For Everyone](https://vzilla.co.uk/vzilla-blog/project_pace-kasten-k10-demo-environment-for-everyone)

## Recursos

Si tienes recursos GRATUITOS que has utilizado, no dudes en agregarlos aquí a través de una solicitud de extracción (PR) al repositorio y estaré encantado de incluirlos.

- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [TechWorld with Nana - Kubernetes Tutorial for Beginners [FULL COURSE in 4 Hours]](https://www.youtube.com/watch?v=X48VuDVv0do)
- [TechWorld with Nana - Kubernetes Crash Course for Absolute Beginners](https://www.youtube.com/watch?v=s_o8dwzRlu4)
- [Kunal Kushwaha - Kubernetes Tutorial for Beginners | What is Kubernetes? Architecture Simplified!](https://www.youtube.com/watch?v=KVBON1lA9N8)

Nos vemos en el [Día 52](day52.md)
