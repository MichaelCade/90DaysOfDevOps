---
title: '#90DaysOfDevOps - Using Roles & Deploying a Loadbalancer - Day 67'
published: false
description: '90DaysOfDevOps - Using Roles & Deploying a Loadbalancer'
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048713
---

## Role 사용 및 로드밸런서 배포하기

지난 세션에서는 Role에 대해 다루고 `ansible-galaxy` 명령을 사용하여 앞으로 사용할 몇 가지 Role에 대한 폴더 구조를 만들었습니다. 모든 것이 Role 폴더에 숨겨져 있기 때문에 구성 코드의 작업 저장소가 훨씬 더 깔끔해졌습니다.

하지만 지금까지는 apache2 Role만 사용했고 웹서버를 처리하기 위해 작동하는 playbook3.yaml이 있습니다.

이 시점에서 `vagrant up web01 web02`만 사용했다면 이제 `vagrant up loadbalancer`를 실행하면 로드 밸런서/프록시로 사용할 다른 Ubuntu 시스템이 나타납니다.

호스트 파일에 이 새 시스템을 이미 정의했지만, 사용할 수 있을 때까지 ssh 키가 구성되지 않았으므로 시스템이 가동되고 준비되면 `ssh-copy-id loadbalancer`도 실행해야 합니다.

### Common Role

어제 세션 마지막에 `common` Role을 만들었는데, Common은 모든 서버에서 사용되는 반면 다른 Role은 사용 사례에 따라 다르지만 이제 설치하려는 애플리케이션은 가짜처럼 일반적이며 이것이 왜 그런지 많은 이유를 알 수는 없지만 그 목적을 보여줍니다. Common Role 폴더 구조에서 작업 폴더로 이동하면 main.yml이 있습니다. 이 YAML에서 이 파일을 install_tools.yml 파일로 가리켜야 하며, `- import_tasks: install_tools.yml` 줄을 추가하여 이를 수행합니다. 이전에는 `include`였으나 곧 사용되지 않을 예정이므로 import_tasks를 사용합니다.

```Yaml
- name: "Install Common packages"
  apt: name={{ item }} state=latest
  with_items:
   - neofetch
   - tree
   - figlet
```

그런 다음 playbook에서 각 호스트 블록에 대한 Common Role을 추가합니다.

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

### Nginx

다음 단계는 로드밸런서 VM에 nginx를 설치하고 구성하는 것입니다. 일반적인 폴더 구조와 마찬가지로, 마지막 세션을 기반으로 nginx를 구성합니다.

먼저 playbook에 호스트 블록을 추가하겠습니다. 이 블록에는 Common Role과 새로운 nginx Role이 포함됩니다.

playbook은 여기에서 찾을 수 있습니다. [playbook4.yml](/2022/Days/Configmgmt/ansible-scenario4/playbook4.yml)

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

이것이 의미가 있으려면 실행할 작업을 정의해야 하며, 같은 방식으로 이번에는 설치용 파일과 구성용 파일 두 개를 가리키도록 작업의 main.yml을 수정하겠습니다.

원하는 결과에 따라 수정한 다른 파일이 몇 개 더 있는데, [ansible-scenario4](/2022/Days/Configmgmt/ansible-scenario4) 폴더에서 변경된 모든 파일을 살펴보세요. nginx 폴더의 작업, Handler 및 template 폴더를 확인하면 추가 변경 사항과 파일을 찾을 수 있습니다.

### 업데이트된 playbook 실행

어제부터 시스템에 일부 패키지를 설치하는 Common Role을 추가한 데 이어 설치 및 구성을 포함하는 nginx Role도 추가했습니다.

`anible-playbook playbook4.yml`을 사용하여 playbook4.yml을 실행해 보겠습니다.

![](/2022/Days/Images/Day67_config1.png)

이제 웹 서버와 로드밸런서가 구성되었으므로 이제 로드밸런서의 IP 주소인 http://192.168.169.134/ 로 이동할 수 있어야 합니다.

![](/2022/Days/Images/Day67_config2.png)

이 과정을 따르고 있는데도 이 상태가 나타나지 않는다면 사용 중인 환경의 서버 IP 주소 때문일 수 있습니다. 이 파일은 `templates\mysite.j2`에서 찾을 수 있으며 아래와 유사하게 보입니다: 웹 서버 IP 주소로 업데이트해야 합니다.

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

우리가 설치한 것이 모두 정상이라고 확신하지만, ansible을 사용하여 임시 명령을 사용하여 이러한 일반적인 도구 설치를 확인해 보겠습니다.

`ansible loadbalancer -m command -a neofetch`

![](/2022/Days/Images/Day67_config3.png)

## 자료

- [What is Ansible](https://www.youtube.com/watch?v=1id6ERvfozo)
- [Ansible 101 - Episode 1 - Introduction to Ansible](https://www.youtube.com/watch?v=goclfp6a2IQ)
- [NetworkChuck - You need to learn Ansible right now!](https://www.youtube.com/watch?v=5hycyr-8EKs&t=955s)
- [Your complete guide to Ansible](https://www.youtube.com/playlist?list=PLnFWJCugpwfzTlIJ-JtuATD2MBBD7_m3u)

위에 나열된 마지막 재생 목록은 이 섹션의 많은 코드와 아이디어가 나온 곳이며, 동영상 형식의 훌륭한 리소스이자 워크스루입니다.

[Day 68](day68.md)에서 봐요!
