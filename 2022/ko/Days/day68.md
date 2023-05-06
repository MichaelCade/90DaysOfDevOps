---
title: '#90DaysOfDevOps - Tags, Variables, Inventory & Database Server config - Day 68'
published: false
description: '90DaysOfDevOps - Tags, Variables, Inventory & Database Server config'
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048780
---

## Tag, Variable, Inventory 및 Database 서버 구성

### Tag

어제 세션에 playbook을 남겨두었으므로 모든 작업을 실행하고 해당 playbook 내에서 play해야 합니다. 즉, 웹서버와 로드밸런서 play 및 작업을 모두 실행해야 완료할 수 있습니다.

하지만 Tag를 사용하면 원하는 경우 이러한 작업을 분리할 수 있습니다. 이는 환경에 매우 크고 긴 playbook이 있는 경우 효율적인 방법이 될 수 있습니다.

이 경우 playbook 파일에서는 [ansible-scenario5](/2022/Days/Configmgmt/ansible-scenario5/playbook5.yml)를 사용하고 있습니다.

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
  tags: web

- hosts: proxy
  become: yes
  roles:
    - common
    - nginx
  tags: proxy
```

그런 다음 `ansible-playbook playbook5.yml --list-tags`를 사용하여 이를 확인할 수 있으며, list Tag는 우리가 playbook에서 정의한 Tag의 윤곽을 나타냅니다.

![](/2022/Days/Images/Day68_config1.png)

이제 프록시만 타깃팅하려면 `ansible-playbook playbook5.yml --tags proxy`를 실행하면 아래에서 볼 수 있듯이 프록시에 대해서만 playbook을 실행할 수 있습니다.

![](/2022/Days/Images/Day68_config2.png)

Tag는 작업 수준에서도 추가할 수 있으므로 원하는 위치와 작업을 세분화할 수 있습니다. 예를 들어 애플리케이션 중심 Tag가 될 수도 있고, 작업을 살펴보고 설치, 구성 또는 제거에 따라 작업에 Tag를 지정할 수도 있습니다. 사용할 수 있는 또 다른 매우 유용한 Tag는 다음과 같습니다.

`tag: always` 이것은 명령에 어떤 --tags를 사용하든 상관없이 항상 값으로 Tag가 지정된 항목이 있으면 ansible-playbook 명령을 실행할 때 항상 실행되도록 보장합니다.

Tag를 사용하여 여러 Tag를 함께 묶을 수도 있으며, `ansible-playbook playbook5.yml --tags proxy,web`을 실행하도록 선택하면 해당 Tag가 있는 모든 항목이 실행됩니다. 물론 이 예제에서는 playbook을 실행하는 것과 같은 의미이지만, 다른 play가 여러 개 있는 경우에는 이 방법이 의미가 있을 것입니다.

둘 이상의 Tag를 정의할 수도 있습니다.

### Variable

Ansible에는 크게 두 가지 유형의 Variable이 있습니다.

- User created
- Ansible Facts

### Ansible Facts

playbook을 실행할 때마다 "Gathering facts"라는 정의되지 않은 작업이 있었는데, 이러한 Variable 또는 fact를 사용하여 자동화 작업을 수행할 수 있습니다.

![](/2022/Days/Images/Day68_config3.png)

다음 `ansible proxy -m setup` 명령을 실행하면 JSON 형식의 많은 출력을 볼 수 있습니다. 이를 사용하려면 터미널에 많은 정보가 있어야 하므로 `ansible proxy -m setup >> facts.json`을 사용하여 파일로 출력하고 싶습니다. [여기](/2022/Days/Configmgmt/ansible-scenario5/facts.json)에서 이 파일을 볼 수 있습니다.

![](/2022/Days/Images/Day68_config4.png)

이 파일을 열면 명령에 대한 모든 종류의 정보를 볼 수 있습니다. IP 주소, 아키텍처, 바이오스 버전을 확인할 수 있습니다. 이 정보를 활용하여 playbook에 사용하려는 경우 유용한 정보가 많이 있습니다.

한 가지 아이디어는 웹서버의 IP 주소를 하드코딩한 nginx template mysite.j2 내에서 이러한 Variable 중 하나를 잠재적으로 사용하는 것입니다. mysite.j2에 for 루프를 생성하면 [webservers] 그룹을 순환하여 2개 이상의 웹서버를 자동으로 동적으로 생성하거나 이 로드 밸런서 구성에 추가할 수 있습니다.

```
#Dynamic Config for server {{ ansible_facts['nodename'] }}
    upstream webservers {
  {% for host in groups['webservers'] %}
        server {{ hostvars[host]['ansible_facts']['nodename'] }}:8000;
    {% endfor %}
    }

    server {
        listen 80;

        location / {
                proxy_pass http://webservers;
        }
    }
```

위의 결과는 지금과 동일하게 보이지만 웹 서버를 더 추가하거나 제거하면 프록시 구성이 동적으로 변경됩니다. 이 기능을 사용하려면 이름 확인을 구성해야 합니다.

### User created

User created Variable은 우리가 직접 만든 Variable입니다. playbook을 살펴보면 `vars:`가 있고 거기에 3개의 Variable 목록이 있는 것을 볼 수 있습니다.

```Yaml
- hosts: webservers
  become: yes
  vars:
    http_port: 8000
    https_port: 4443
    html_welcome_msg: "Hello 90DaysOfDevOps - Welcome to Day 68!"
  roles:
    - common
    - apache2
  tags: web

- hosts: proxy
  become: yes
  roles:
    - common
    - nginx
  tags: proxy
```

그러나 Variable을 해당 파일로 이동하여 playbook에 Variable이 없도록 할 수 있습니다. 이 작업을 수행하되, [ansible-scenario6](/2022/Days/Configmgmt/ansible-scenario6) 폴더로 이동하겠습니다. 해당 폴더의 루트에 group_vars 폴더를 만들겠습니다. 그런 다음 all이라는 또 다른 폴더를 만듭니다(모든 그룹이 이 Variable을 가져옵니다). 이 폴더에 `common_variables.yml`이라는 파일을 만들고 playbook의 Variable을 이 파일에 복사합니다. Variable과 함께 playbook에서 Variable을 제거합니다.

```Yaml
http_port: 8000
https_port: 4443
html_welcome_msg: "Hello 90DaysOfDevOps - Welcome to Day 68!"
```

이 Variable을 전역 Variable로 연결하기 때문에 여기에 NTP 및 DNS 서버도 추가할 수 있습니다. Variable은 우리가 만든 폴더 구조에서 설정됩니다. 이제 playbook이 얼마나 깔끔해졌는지 아래에서 확인할 수 있습니다.

```Yaml
- hosts: webservers
  become: yes
  roles:
    - common
    - apache2
  tags: web

- hosts: proxy
  become: yes
  roles:
    - common
    - nginx
  tags: proxy
```

그 Variable 중 하나는 http_port로, 아래와 같이 mysite.j2 내의 for 루프에서 이 Variable을 다시 사용할 수 있습니다:

```J2
#Dynamic Config for server {{ ansible_facts['nodename'] }}
    upstream webservers {
  {% for host in groups['webservers'] %}
        server {{ hostvars[host]['ansible_facts']['nodename'] }}:{{ http_port }};
    {% endfor %}
    }

    server {
        listen 80;

        location / {
                proxy_pass http://webservers;
        }
    }
```

또한, 어떤 웹서버를 사용하고 있는지 파악할 수 있도록 roles/apache2/templates/index.HTML.j2 파일에 분석 가능한 사실을 정의할 수도 있습니다.

```J2
<html>

<h1>{{ html_welcome_msg }}! I'm webserver {{ ansible_facts['nodename'] }} </h1>

</html>
```

Variable을 변경하여 `ansible-playbook playbook6.yml` 명령을 실행한 결과, 로드밸런서에 도달하면 그룹에 있는 웹서버 중 하나에 도달한 것을 볼 수 있습니다.

![](/2022/Days/Images/Day68_config5.png)

또한 host_vars라는 폴더를 추가하고 web01.yml을 생성하여 원하는 경우 특정 메시지를 표시하거나 호스트별로 표시되는 내용을 변경할 수 있습니다.

### Inventory 파일

지금까지 호스트를 결정하기 위해 /etc/ansible 폴더에 있는 기본 호스트 파일을 사용했습니다. 그러나 프로덕션 및 스테이징과 같이 환경마다 다른 파일을 사용할 수 있습니다. 더 많은 환경을 만들지는 않겠습니다. 하지만 호스트 파일은 만들 수 있습니다.

서버와 노드의 다양한 Inventory에 대해 여러 개의 파일을 만들 수 있습니다. 우리는 `ansible-playbook -i dev playbook.yml`을 사용하여 이를 호출합니다. 호스트 파일 내에 Variable을 정의한 다음 이를 출력하거나 playbook의 다른 곳에서 해당 Variable을 활용할 수도 있습니다. 예를 들어 아래 예제 및 교육 과정에서는 호스트 파일에서 생성된 환경 Variable을 로드밸런서 웹 페이지 template에 추가하여 웹 페이지 메시지의 일부로 환경을 표시했습니다.

### Database 서버 배포

아직 전원을 켜고 구성하지 않은 머신이 하나 더 있습니다. vagrant 파일이 있는 곳에서 `vagrant up db01`을 사용하여 이 작업을 수행할 수 있습니다. 전원이 켜지고 액세스할 수 있게 되면 `ssh-copy-id db01`을 사용하여 SSH 키를 복사하여 액세스할 수 있도록 해야 합니다.

여기서는 [ansible-scenario7](/2022/Days/Configmgmt/ansible-scenario7) 폴더에서 작업할 것입니다.

그런 다음 `ansible-galaxy init roles/mysql`을 사용하여 "MySQL"이라는 새 Role에 대한 새 폴더 구조를 만들어 보겠습니다.

playbook에서 Database 구성을 위한 새로운 play 블록을 추가하겠습니다. /etc/ansible/hosts 파일에 그룹 Database가 정의되어 있습니다. 그런 다음 Database 그룹에 공통 Role과 이전 단계에서 생성한 MySQL이라는 새 Role을 갖도록 지시합니다. 또한 Database 그룹에 Database Tag를 지정하고 있는데, 이는 앞서 설명한 것처럼 원하는 경우 이러한 Tag에 대해서만 실행하도록 선택할 수 있음을 의미합니다.

```Yaml
- hosts: webservers
  become: yes
  roles:
    - common
    - apache2
  tags:
    web

- hosts: proxy
  become: yes
  roles:
    - common
    - nginx
  tags:
    proxy

- hosts: database
  become: yes
  roles:
    - common
    - mysql
  tags: database
```

이제 Role 폴더 구조 내에서 트리가 자동으로 생성되었으므로 다음을 채워야 합니다:

Handlers - main.yml

```Yaml
# handlers file for roles/mysql
- name: restart mysql
  service:
    name: mysql
    state: restarted
```

Tasks - install_mysql.yml, main.yml & setup_mysql.yml

install_mysql.yml - 이 작업은 MySQL을 설치하고 서비스가 실행 중인지 확인하기 위해 수행됩니다.

```Yaml
- name: "Install Common packages"
  apt: name={{ item }} state=latest
  with_items:
   - python3-pip
   - mysql-client
   - python3-mysqldb
   - libmysqlclient-dev

- name: Ensure mysql-server is installed latest version
  apt: name=mysql-server state=latest

- name: Installing python module MySQL-python
  pip:
    name: PyMySQL

- name: Ensure mysql-server is running
  service:
    name: mysql
    state: started
```

main.yml은 이러한 파일에서 작업을 가져오도록 제안하는 포인터 파일입니다.

```Yaml
# tasks file for roles/mysql
- import_tasks: install_mysql.yml
- import_tasks: setup_mysql.yml
```

setup_mysql.yml - 이 작업은 Database와 Database 사용자를 생성합니다.

```Yaml
- name: Create my.cnf configuration file
  template: src=templates/my.cnf.j2 dest=/etc/mysql/conf.d/mysql.cnf
  notify: restart mysql

- name: Create database user with name 'devops' and password 'DevOps90' with all database privileges
  community.mysql.mysql_user:
    login_unix_socket: /var/run/mysqld/mysqld.sock
    login_user: "{{ mysql_user_name }}"
    login_password: "{{ mysql_user_password }}"
    name: "{{db_user}}"
    password: "{{db_pass}}"
    priv: '*.*:ALL'
    host: '%'
    state: present

- name: Create a new database with name '90daysofdevops'
  mysql_db:
    login_user: "{{ mysql_user_name }}"
    login_password: "{{ mysql_user_password }}"
    name: "{{ db_name }}"
    state: present
```

위에서 비밀번호, 사용자 이름 및 Database와 같은 일부 구성을 결정하기 위해 몇 가지 Variable을 사용하고 있음을 알 수 있으며, 이 Variable은 모두 group_vars/all/common_variables.yml 파일에 저장되어 있습니다.

```Yaml
http_port: 8000
https_port: 4443
html_welcome_msg: "Hello 90DaysOfDevOps - Welcome to Day 68!"

mysql_user_name: root
mysql_user_password: "vagrant"
db_user: devops
db_pass: DevOps90
db_name: 90DaysOfDevOps
```

또한 template 폴더에 아래와 같은 my.cnf.j2 파일이 있습니다:

```J2
[mysql]
bind-address = 0.0.0.0
```

### playbook 실행

이제 VM이 실행 중이고 구성 파일이 준비되었으므로 이제 playbook을 실행할 준비가 되었습니다. 다음 `ansible-playbook playbook7.yml`을 실행하면 이전에 수행한 모든 작업이 포함된 playbook을 실행할 수 있고, `ansible-playbook playbook7.yml --tags database` 명령으로 Database 그룹에 배포하여 새 구성 파일만 실행하도록 선택할 수도 있습니다.

Database Tag에 대해서만 실행했지만, 오류가 발생했습니다. 이 오류는 pip3(파이썬)가 설치되어 있지 않다는 것을 알려줍니다. 이 오류는 일반 작업에 추가하여 해결할 수 있습니다.

![](/2022/Days/Images/Day68_config6.png)

위의 문제를 수정하고 playbook을 다시 실행했더니 성공적으로 변경되었습니다.

![](/2022/Days/Images/Day68_config7.png)

이제 새로 구성한 db01 서버에서 모든 것이 우리가 원하는 대로 작동하는지 확인해야 합니다. Control Node에서 `ssh db01` 명령을 사용하여 이 작업을 수행할 수 있습니다.

MySQL에 연결하기 위해 `sudo /usr/bin/mysql -u root -p`를 사용하고 프롬프트에서 root에 대한 vagrant 암호를 제공했습니다.

연결이 완료되면 먼저 DevOps라는 사용자가 생성되었는지 확인합니다. `select user, host from mysql.user;`

![](/2022/Days/Images/Day68_config8.png)

이제 `SHOW DATABASES;` 명령을 실행하여 생성된 새 Database를 확인할 수 있습니다.

![](/2022/Days/Images/Day68_config9.png)

루트를 사용하여 연결했지만 이제 `sudo /usr/bin/MySQL -u devops -p`를 사용하여 동일한 방법으로 DevOps 계정으로 로그인할 수도 있지만 여기서 암호는 DevOps90입니다.

제가 발견한 한 가지는 `setup_mysql.yml`에 `login_unix_socket: /var/run/mysqld/mysqld.sock` 줄을 추가해야 db01 MySQL 인스턴스에 성공적으로 연결할 수 있었고 이제 이 작업을 실행할 때마다 사용자를 만들 때 변경 사항을 보고하므로 어떤 제안이라도 대단히 감사하겠습니다.

## 자료

- [What is Ansible](https://www.youtube.com/watch?v=1id6ERvfozo)
- [Ansible 101 - Episode 1 - Introduction to Ansible](https://www.youtube.com/watch?v=goclfp6a2IQ)
- [NetworkChuck - You need to learn Ansible right now!](https://www.youtube.com/watch?v=5hycyr-8EKs&t=955s)
- [Your complete guide to Ansible](https://www.youtube.com/playlist?list=PLnFWJCugpwfzTlIJ-JtuATD2MBBD7_m3u)

위에 나열된 마지막 재생 목록은 이 섹션의 많은 코드와 아이디어가 나온 곳이며, 동영상 형식의 훌륭한 리소스이자 워크스루입니다.

[Day 69](day69.md)에서 봐요!
