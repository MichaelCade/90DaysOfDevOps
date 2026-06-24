## Elegir tu plataforma Kubernetes

En esta sesión se va a desglosar algunas de las plataformas o tal vez distribuciones es un mejor término a utilizar aquí. Una cosa que ha sido un reto en el mundo Kubernetes es eliminar la complejidad.

Kubernetes the hard way (el camino difícil) explica cómo construir desde la nada un clúster Kubernetes completamente funcional. Esto es extremo, pero cada vez más, al menos las personas con las que hablo, quieren eliminar esa complejidad y ejecutar un clúster Kubernetes gestionado. El problema es que cuesta más dinero, pero las ventajas podrían ser que, si se utiliza un servicio gestionado, es necesario conocer la arquitectura de los nodos y lo que ocurre desde el punto de vista del control plane, algo a lo que generalmente no se tiene acceso.

Luego tenemos las distribuciones de desarrollo local que nos permiten utilizar nuestros sistemas y ejecutar una versión local de Kubernetes para que los desarrolladores puedan tener el entorno de trabajo completo para ejecutar sus aplicaciones en la plataforma para la que están destinadas.

La base general de todos estos conceptos es que todos son una variante de Kubernetes, lo que significa que deberíamos poder migrar y mover libremente nuestras cargas de trabajo donde las necesitemos para adaptarlas a nuestras necesidades.

Gran parte de nuestra elección dependerá también de las inversiones realizadas. También he mencionado la experiencia de los desarrolladores, pero algunos de esos entornos locales de Kubernetes que ejecutan nuestros portátiles son estupendos para familiarizarse con la tecnología sin gastar dinero.

### Clusters Bare-Metal

Una opción para muchos podría ser ejecutar su sistema operativo Linux directamente en varios servidores físicos para crear nuestro clúster, también podría ser Windows pero no he oído hablar mucho de la tasa de adopción en entornos Windows, puede ser anecdótica. 

Si en tu empresa se ha tomado la decisión de CAPEX para comprar sus servidores físicos, entonces esto podría ser una forma en la que construir la arquitectura, un clúster Kubernetes. La gestión y administración de lado significa que la vas a tener que construir tu mismo y gestionarlo todo desde el principio.

### Virtualización

Independientemente de si se trata de entornos de prueba y aprendizaje o de clústeres Kubernetes listos para la funcionar, la virtualización es una gran opción, ya que permite crear máquinas virtuales que actúan como nodos y, a continuación, agruparlos en clústeres. Dispones de la arquitectura de apoyo, la eficiencia y la velocidad de la virtualización, además de aprovechar el gasto existente. VMware, por ejemplo, ofrece una gran solución tanto para máquinas virtuales como para Kubernetes en varias versiones.

Mi primer clúster Kubernetes se construyó sobre la base de la virtualización utilizando Microsoft Hyper-V en un servidor antiguo que tenía que era capaz de ejecutar algunas máquinas virtuales como mis nodos. Esto me permitió construir un clúster Kubernetes y empezar a jugar con él.

### Opciones de escritorio local

Hay varias opciones cuando se trata de ejecutar un clúster local de Kubernetes en su ordenador de sobremesa o portátil. Esto, como se dijo anteriormente, ofrece a los desarrolladores la posibilidad de ver cómo se verá su aplicación sin tener que tener múltiples clústeres costosos o complejos. Personalmente, he utilizado mucho minikube. Tiene una gran funcionalidad y complementos que cambian la forma de poner algo en marcha.

### Servicios gestionados Kubernetes

He mencionado la virtualización, y esto se puede lograr con hipervisores a nivel local, pero sabemos por secciones anteriores que también podríamos aprovechar las máquinas virtuales en la nube pública para actuar como nuestros nodos. De lo que estoy hablando aquí con los servicios gestionados Kubernetes son las ofertas que vemos de los grandes hiperescaladores, pero también de los MSP eliminando capas de gestión y control lejos del usuario final, esto podría ser eliminar el plano de control del usuario final, esto es lo que sucede con Amazon EKS, Microsoft AKS y Google Kubernetes Engine. (GKE)

### Elección abrumadora

Quiero decir que la elección es grande, pero hay un punto en el que las cosas se vuelven abrumadoras y esto no es una mirada en profundidad a todas las opciones dentro de cada categoría enumerada anteriormente. 

No podemos olvidar que además de lo anterior también tenemos OpenShift que es de Red Hat. Esta opción se puede ejecutar a través de las opciones anteriores en todos los principales proveedores de nube y, probablemente, hoy en día ofrece la mejor facilidad de uso general para los administradores, independientemente de donde se despliegan los clusters.

Por lo tanto, ¿por dónde empezar desde la perspectiva del aprendizaje? Como he dicho yo empecé por la ruta de la virtualización, pero eso fue porque tenía acceso a un servidor físico que podía utilizar para este fin, lo agradezco muchísimo.

Mi consejo actual sería utilizar Minikube como primera opción o Kind (Kubernetes en Docker), pero Minikube nos da algunos beneficios adicionales que casi abstrae la complejidad, ya que sólo podemos utilizar complementos y hacer las cosas rápidamente y luego podemos volarlo cuando hayamos terminado, podemos ejecutar múltiples clusters, podemos ejecutarlo casi en cualquier lugar, multiplataforma y agnóstico de hardware.

He estado a través de un poco de un viaje con mi aprendizaje en torno a Kubernetes así que voy a dejar la elección de la plataforma y los detalles aquí para enumerar las opciones que he probado para darme una mejor comprensión de Kubernetes la plataforma y donde se puede ejecutar. Lo que podría hacer con las siguientes entradas de blog es echar otro vistazo a estos actualizarlos y traerlos más aquí frente a ellos están vinculados a las entradas del blog.

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

- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [TechWorld with Nana - Kubernetes Tutorial for Beginners [FULL COURSE in 4 Hours]](https://www.youtube.com/watch?v=X48VuDVv0do)
- [TechWorld with Nana - Kubernetes Crash Course for Absolute Beginners](https://www.youtube.com/watch?v=s_o8dwzRlu4)
- [Kunal Kushwaha - Kubernetes Tutorial for Beginners | What is Kubernetes? Architecture Simplified!](https://www.youtube.com/watch?v=KVBON1lA9N8)

Nos vemos en el [Día 51](day51.md)
