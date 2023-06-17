## Uso de Roles e Implementación de un Balanceador de Carga

En la última sesión, cubrimos los roles y utilizamos el comando `ansible-galaxy` para crear las estructuras de carpetas para algunos roles que vamos a utilizar. Terminamos con un repositorio de trabajo mucho más ordenado para nuestro código de configuración, ya que todo está oculto en nuestras carpetas de roles.

Sin embargo, solo hemos utilizado el rol apache2 y tenemos un playbook3.yaml funcional para manejar nuestros servidores web.

En este momento, si solo has ejecutado `vagrant up web01 web02`, es hora de ejecutar `vagrant up loadbalancer`. Esto iniciará otro sistema Ubuntu que utilizaremos como nuestro balanceador de carga/proxy.

Ya hemos definido esta nueva máquina en nuestro archivo de hosts, pero no tenemos configurada la clave SSH hasta que esté disponible, así que también necesitamos ejecutar `ssh-copy-id loadbalancer` cuando el sistema esté listo.

### Rol common

Al final de la sesión de ayer, creé el rol common, que se utilizará en todos nuestros servidores, mientras que los otros roles son específicos para casos de uso. Ahora, las aplicaciones que voy a instalar son tan comunes como espurias y no veo muchas razones para que esto sea así, pero esto muestra el objetivo. En la estructura de carpetas del rol common, navega hasta la carpeta tasks y tendrás un archivo main.yml. En este archivo YAML, debemos apuntar a nuestro archivo install_tools.yml y lo hacemos agregando la línea `- import_tasks: install_tools.yml`. Anteriormente, esto se hacía con `include`, pero esto se depreciará lo suficiente pronto, por lo que estamos usando import_tasks.

```Yaml
- name: "Install Common packages"
  apt: name={{ item }} state=latest
  with_items:
   - neofetch
   - tree
   - figlet
```

En nuestro playbook, luego agregamos el rol common para cada bloque de host.

```Yaml
- hosts: webservers
  become: yes
  vars:
    http_port: 8000
    https_port: 4443
    html_welcome_msg: "Hello 90DaysOfDevOps - Welcome to Day 66!"
  roles:
    - common
    - apache2
```

### nginx

La siguiente fase es instalar y configurar nginx en nuestra máquina de balanceo de carga (loadbalancer). Al igual que la estructura de carpetas common, tenemos la carpeta nginx basada en la sesión anterior.

Primero, vamos a agregar un bloque de host a nuestro playbook. Este bloque incluirá nuestro rol common y luego nuestro nuevo rol nginx.

El playbook se puede encontrar aquí: [playbook4.yml](Days/../Configmgmt/ansible-scenario4/playbook4.yml)

```Yaml
- hosts: webservers
  become: yes
  vars:
    http_port: 8000
    https_port: 4443
    html_welcome_msg: "Hello 90DaysOfDevOps - Welcome to Day 66!"
  roles:
    - common
    - apache2

- hosts: proxy
  become: yes
  roles:
    - common
    - nginx
```

Para que esto signifique algo, debemos definir las tareas que deseamos ejecutar. De la misma manera, modificaremos el archivo main.yml en tasks para que apunte a dos archivos en esta ocasión, uno para la instalación y otro para la configuración.

Hay algunos otros archivos que he modificado en función del resultado que deseamos. Echa un vistazo a la carpeta [ansible-scenario4](Days/Configmgmt/ansible-scenario4) para ver todos los archivos modificados. Debes revisar las carpetas tasks, handlers y templates en la carpeta nginx, y encontrarás esos cambios y archivos adicionales.

### Ejecutar el playbook actualizado

Desde ayer, hemos agregado el rol common, que ahora instalará algunos paquetes en nuestro sistema, y también hemos agregado nuestro rol nginx, que incluye la instalación y configuración.

Ejecutemos nuestro playbook4.yml usando `ansible-playbook playbook4.yml`

![](Images/Day67_config1.png)

Ahora que hemos configurado nuestros servidores web y nuestro balanceador de carga, deberíamos poder acceder a http://192.168.169.134/, que es la dirección IP de nuestro balanceador de carga.

![](Images/Day67_config2.png)

Si estás siguiendo los pasos y no tienes este resultado, podría deberse a las direcciones IP de los servidores que tienes en tu entorno. El archivo se puede encontrar en `templates\mysite.j2` y se ve similar al siguiente: Debes actualizarlo con las direcciones IP de tus servidores web.

```J2
    upstream webservers {
        server 192.168.169.131:8000;
        server 192.168.169.132:8000;
    }

    server {
        listen 80;

        location / {
                proxy_pass http://webservers;
        }
    }
```

Estoy bastante seguro de que lo que hemos instalado está todo bien, pero vamos a usar un comando ad-hoc con ansible para verificar la instalación de estas herramientas comunes.

`ansible loadbalancer -m command -a neofetch`

![](Images/Day67_config3.png)

## Recursos

- [What is Ansible](https://www.youtube.com/watch?v=1id6ERvfozo)
- [Ansible 101 - Episode 1 - Introduction to Ansible](https://www.youtube.com/watch?v=goclfp6a2IQ)
- [NetworkChuck - You need to learn Ansible right now!](https://www.youtube.com/watch?v=5hycyr-8EKs&t=955s)
- [Your complete guide to Ansible](https://www.youtube.com/playlist?list=PLnFWJCugpwfzTlIJ-JtuATD2MBBD7_m3u)
- [Chef vs Puppet vs Ansible vs Saltstack](https://vergaracarmona.es/chef-vs-puppet-vs-ansible-vs-saltstack/)

La lista de reproducción final mencionada anteriormente es de donde proviene gran parte del código y las ideas para esta sección, es un recurso excelente y una guía en formato de video.

Nos vemos en el [Día 68](day68.md)
