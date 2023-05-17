## Configuración de un clúster de Kubernetes multinodo

Quería que este título fuera "Configuración de un clúster de Kubernetes multinodo con Vagrant", pero pensé que podría ser un poco largo.

En la sesión de ayer, utilizamos un proyecto interesante para implementar nuestro primer clúster de Kubernetes y tener un poco de práctica con la herramienta CLI más importante que encontrarás al usar Kubernetes (`kubectl`).

Aquí vamos a utilizar VirtualBox como base, pero como mencioné la última vez que hablamos de Vagrant en la sección de Linux, podemos usar cualquier hipervisor o herramienta de virtualización compatible. Fue [Día 14](day14.md) cuando pasamos por el proceso de implementar una máquina Ubuntu para la sección de Linux.

### Un breve repaso sobre Vagrant

Vagrant es una utilidad de línea de comandos que gestiona el ciclo de vida de tus máquinas virtuales. Podemos usar Vagrant para crear y eliminar máquinas virtuales en diferentes plataformas, incluyendo vSphere, Hyper-V, VirtualBox y también Docker. Tiene otros proveedores, pero nos quedaremos con el que estamos usando aquí, que es VirtualBox.

Voy a utilizar como referencia el [blog y repositorio](https://devopscube.com/kubernetes-cluster-vagrant/) para guiarme en la configuración. Sin embargo, te recomendaría que si es la primera vez que implementas un clúster de Kubernetes, también investigues cómo hacerlo manualmente, así al menos sabrás cómo se ve el proceso. Aunque debo decir que las operaciones del Día 0 y el esfuerzo se están volviendo más eficientes con cada versión de Kubernetes. Esto me recuerda mucho a los días de VMware y ESX, cuando necesitábamos al menos un día para implementar 3 servidores ESX, y ahora podemos tener todo funcionando en una hora. Nos estamos acercando a ese nivel en lo que respecta a Kubernetes.

### Entorno de laboratorio de Kubernetes

He subido en la [carpeta de Kubernetes](Kubernetes) el archivo `vagrantfile` que usaremos para construir nuestro entorno. Descárgalo y navega hasta este directorio en tu terminal. En mi caso, estoy usando Windows, así que usaré PowerShell para ejecutar los comandos de mi estación de trabajo con Vagrant. Si no tienes Vagrant instalado, puedes usar arkade, que cubrimos ayer al instalar minikube y otras herramientas. Un simple comando `arkade get vagrant` debería descargar e instalar la última versión de Vagrant.

Cuando estés en el directorio, simplemente ejecuta `vagrant up` y, si todo está configurado correctamente, deberías ver que se inicia el proceso en tu terminal.

![Day52_Kubernetes1.png](https://example.com/Day52_Kubernetes1.png)

En la terminal, verás que se ejecutan varios pasos, pero mientras tanto, echemos un vistazo a lo que estamos construyendo aquí.

![Day52_Kubernetes2.png](https://example.com/Day52_Kubernetes2.png)

Desde lo anterior, puedes ver que vamos a crear 3 máquinas virtuales. Tendremos un nodo de control y luego dos nodos de trabajo. Si vuelves al [Día 49](day49.md), verás más descripciones de estas áreas que vemos en la imagen.

Además, en la imagen indicamos que nuestro acceso a `kubectl` vendrá desde fuera del clúster y llegará al `kube-apiserver`. En realidad, como parte de la provisión de Vagrant, estamos implementando `kubectl` en cada uno de estos nodos para poder acceder al clúster desde dentro de cada uno de ellos.

El proceso de construcción de este laboratorio puede tardar entre 5 y 30 minutos, según la configuración que tengas.

Voy a cubrir los scripts en breve, pero notarás que, si examinas el archivo Vagrant, estamos llamando a 3 scripts como parte de la implementación, y es aquí donde se crea realmente el clúster. Hemos visto lo fácil que es usar Vagrant para implementar nuestras máquinas virtuales e instalaciones del sistema operativo mediante cajas de Vagrant, pero la capacidad de ejecutar un script de shell como parte del proceso de implementación es lo que resulta interesante para automatizar la creación de estos laboratorios.

Una vez completado, podemos acceder a uno de nuestros nodos con `vagrant ssh master` desde la terminal. El nombre de usuario y contraseña predeterminados son `vagrant/vagrant`.

También puedes usar `vagrant ssh node01` y `vagrant ssh node02` para acceder a los nodos de trabajo si así lo deseas.

![Day52_Kubernetes3.png](https://example.com/Day52_Kubernetes3.png)

Ahora estamos en uno de los nodos anteriores de nuestro nuevo clúster y podemos ejecutar `kubectl get nodes` para mostrar nuestro clúster de 3 nodos y su estado.

![Day52_Kubernetes4.png](https://example.com/Day52_Kubernetes4.png)

En este punto, tenemos un clúster de 3 nodos en funcionamiento, con 1 nodo de control y 2 nodos de trabajo.

### Recorrido por el Vagrantfile y los scripts de shell

Si echamos un vistazo a nuestro Vagrantfile, verás que estamos definiendo varios nodos de trabajo, direcciones IP de red para la red puente dentro de VirtualBox y algunos nombres. Otra cosa que notarás es que también estamos llamando a algunos scripts que queremos ejecutar en hosts específicos.

```
NUM_WORKER_NODES=2
IP_NW="10.0.0."
IP_START=10

Vagrant.configure("2") do |config|
    config.vm.provision "shell", inline: <<-SHELL
        apt-get update -y
        echo "$IP_NW$((IP_START))  master-node" >> /etc/hosts
        echo "$IP_NW$((IP_START+1))  worker-node01" >> /etc/hosts
        echo "$IP_NW$((IP_START+2))  worker-node02" >> /etc/hosts
    SHELL
    config.vm.box = "bento/ubuntu-21.10"
    config.vm.box_check_update = true

    config.vm.define "master" do |master|
      master.vm.hostname = "master-node"
      master.vm.network "private_network", ip: IP_NW + "#{IP_START}"
      master.vm.provider "virtualbox" do |vb|
          vb.memory = 4048
          vb.cpus = 2
          vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
      end
      master.vm.provision "shell", path: "scripts/common.sh"
      master.vm.provision "shell", path: "scripts/master.sh"
    end

    (1..NUM_WORKER_NODES).each do |i|
      config.vm.define "node0#{i}" do |node|
        node.vm.hostname = "worker-node0#{i}"
        node.vm.network "private_network", ip: IP_NW + "#{IP_START + i}"
        node.vm.provider "virtualbox" do |vb|
            vb.memory = 2048
            vb.cpus = 1
            vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
        end
        node.vm.provision "shell", path: "scripts/common.sh"
        node.vm.provision "shell", path: "scripts/node.sh"
      end
    end
  end
```

Vamos a analizar esos scripts que se están ejecutando. Tenemos tres scripts enumerados en el VAGRANTFILE anterior que se ejecutarán en nodos específicos.

`master.vm.provision "shell", path: "scripts/common.sh"`

Este script se centrará en preparar los nodos, se ejecutará en los 3 de nuestros nodos y eliminará cualquier componente de Docker existente y volverá a instalar Docker y ContainerD, así como kubeadm, kubelet y kubectl. Este script también actualizará los paquetes de software existentes en el sistema.

`master.vm.provision "shell", path: "scripts/master.sh"`

El script master.sh solo se ejecutará en el nodo del plano de control, este script creará el clúster de Kubernetes utilizando comandos de kubeadm. También preparará el contexto de configuración para acceder a este clúster, lo cual veremos a continuación.

`node.vm.provision "shell", path: "scripts/node.sh"`

Esto simplemente tomará la configuración creada por el nodo maestro y unirá nuestros nodos al clúster de Kubernetes. Este proceso de unión también utiliza kubeadm y otro script que se puede encontrar en la carpeta de configuración.

### Acceso al clúster de Kubernetes

Ahora que tenemos dos clústeres desplegados, el clúster minikube que desplegamos en la sección anterior y el nuevo clúster de 3 nodos que acabamos de desplegar en VirtualBox.

Además, ese archivo de configuración al que también tendrás acceso en la máquina desde la que ejecutaste Vagrant, consiste en cómo podemos acceder a nuestro clúster desde nuestro puesto de trabajo.

Antes de mostrar eso, permíteme hablar sobre el contexto.

![](Images/Day52_Kubernetes5.png)

El contexto es importante, se requiere la capacidad de acceder a tu clúster de Kubernetes desde tu escritorio o portátil. Hay muchas opciones diferentes y las personas utilizan diferentes sistemas operativos como sus controladores diarios.

Por defecto, el cliente de línea de comandos de Kubernetes (kubectl) utiliza el archivo C:\Users\nombredeusuario.kube\config para almacenar los detalles del clúster de Kubernetes, como el punto final y las credenciales. Si has desplegado un clúster, podrás ver este archivo en esa ubicación. Pero si hasta ahora has estado utilizando quizás el nodo maestro para ejecutar todos tus comandos de kubectl a través de SSH u otros métodos, esta publicación espero que te ayude a familiarizarte con la capacidad de conectarte desde tu estación de trabajo.

Luego debemos obtener el archivo kubeconfig del clúster, o también podemos obtenerlo de nuestro archivo de configuración una vez desplegado, obtener el contenido de este archivo ya sea mediante SCP o simplemente abrir una sesión de consola en tu nodo maestro y copiarlo a la máquina local de Windows.

![](Images/Day52_Kubernetes6.png)

Luego queremos hacer una copia de ese archivo de configuración y moverlo a nuestra ubicación $HOME/.kube/config.

![](Images/Day52_Kubernetes7.png)

Ahora, desde tu estación de trabajo local, podrás ejecutar `kubectl cluster-info` y `kubectl get nodes` para validar que tienes acceso a tu clúster.

![](Images/Day52_Kubernetes8.png)

Esto no solo permite la conectividad y el control desde tu máquina con Windows, sino que también nos permite realizar un reenvío de puertos para acceder a ciertos servicios desde nuestra máquina con Windows.

Si te interesa saber cómo gestionar múltiples clústeres en tu estación de trabajo, tengo un tutorial más detallado [aquí](https://vzilla.co.uk/vzilla-blog/building-the-home-lab-kubernetes-playground-part-6).

He añadido esta lista de blogs explicativos que he realizado sobre el despliegue de diferentes clústeres de Kubernetes.


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

Si tienes recursos GRATUITOS que has utilizado, no dudes en añadirlos aquí mediante un PR (Pull Request) al repositorio y estaré encantado de incluirlos.

- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [TechWorld with Nana - Kubernetes Tutorial for Beginners [FULL COURSE in 4 Hours]](https://www.youtube.com/watch?v=X48VuDVv0do)
- [TechWorld with Nana - Kubernetes Crash Course for Absolute Beginners](https://www.youtube.com/watch?v=s_o8dwzRlu4)
- [Kunal Kushwaha - Kubernetes Tutorial for Beginners | What is Kubernetes? Architecture Simplified!](https://www.youtube.com/watch?v=KVBON1lA9N8)

Nos vemos en el [Día 53](day53.md)
