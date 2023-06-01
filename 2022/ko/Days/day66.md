---
title: '#90DaysOfDevOps - Ansible Playbooks Continued... - Day 66'
published: false
description: 90DaysOfDevOps - Ansible Playbooks Continued...
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048712
---

## Ansible Playbook (계속)

지난 섹션에서는 vagrant 파일을 사용하여 4대의 머신을 배포하는 작은 실험실을 만드는 것으로 시작했고, 이 섹션에서 만든 Linux 머신을 Ansible 제어 시스템으로 사용했습니다.

또한 Playbook의 몇 가지 시나리오를 실행했고 마지막에는 web01과 web02를 개별 웹 서버로 만드는 Playbook을 만들었습니다.

![](/2022/Days/Images/Day66_config1.png)

### 깔끔하게 정리하기

추가 자동화 및 배포를 시작하기 전에 Playbook을 간결하고 깔끔하게 유지하는 기능과 작업과 Handler를 하위 폴더로 분리하는 방법에 대해 다뤄야 합니다.

작업을 폴더 내의 해당 파일에 복사하겠습니다.

```Yaml
- name: ensure apache is at the latest version
  apt: name=apache2 state=latest

- name: write the apache2 ports.conf config file
  template:
    src=templates/ports.conf.j2
    dest=/etc/apache2/ports.conf
  notify: restart apache

- name: write a basic index.html file
  template:
    src: templates/index.html.j2
    dest: /var/www/html/index.html
  notify:
  - restart apache

- name: ensure apache is running
  service:
    name: apache2
    state: started
```

Handler도 마찬가지입니다.

```Yaml
- name: restart apache
  service:
    name: apache2
    state: restarted
```

Playbook의 이름을 `playbook2.yml`로 지정한 다음, 이 파일을 가리킵니다. 이 모든 파일은 [ansible-scenario2](/2022/Days/Configmgmt/ansible-scenario2/)에서 찾을 수 있습니다.

제어 머신에서 테스트할 수 있습니다. 리포지토리에서 파일을 복사한 경우 "write a basic index.html file"에서 변경된 사항을 발견했을 것입니다.

![](/2022/Days/Images/Day66_config2.png)

`curl web01:8000`을 사용하여 어떤 간단한 변경이 있었는지 알아봅시다.

![](/2022/Days/Images/Day66_config3.png)

방금 Playbook을 정리하고 규모에 따라 Playbook을 매우 압도적으로 만들 수 있는 영역을 분리하기 시작했습니다.

### Role과 Ansible Galaxy

현재 4개의 VM을 배포했으며 이 중 2개의 VM을 웹 서버로 구성했지만 데이터베이스 서버와 로드 밸런서 또는 프록시 등 좀 더 구체적인 기능이 있습니다. 이 작업을 수행하고 리포지토리를 정리하기 위해 Ansible 내에서 Role을 사용할 수 있습니다.

이를 위해 공유 리포지토리에서 Ansible Role을 관리하기 위해 존재하는 `ansible-galaxy` 명령을 사용합니다.

![](/2022/Days/Images/Day66_config4.png)

우리는 `ansible-galaxy`를 사용하여 웹서버에 대한 세부 정보를 넣을 apache2의 Role을 생성할 것입니다.

![](/2022/Days/Images/Day66_config5.png)

위의 명령 `ansible-galaxy init roles/apache2`는 위에 표시된 폴더 구조를 생성합니다. 다음 단계는 기존 작업과 template을 새 구조의 관련 폴더로 이동하는 것입니다.

![](/2022/Days/Images/Day66_config6.png)

복사하여 붙여넣으면 파일을 쉽게 옮길 수 있지만, tasks/main.yml을 변경하여 이 파일이 apache2_install.yml을 가리키도록 해야 합니다.

또한 새로운 Role을 참조하도록 Playbook을 변경해야 합니다. playbook1.yml과 playbook2.yml에서 작업과 Handler를 두 버전 간에 변경하면서 다른 방식으로 결정합니다. 아래와 같이 이 Role을 사용하도록 Playbook을 변경해야 합니다:

```Yaml
- hosts: webservers
  become: yes
  vars:
    http_port: 8000
    https_port: 4443
    html_welcome_msg: "Hello 90DaysOfDevOps - Welcome to Day 66!"
  roles:
    - apache2
```

![](/2022/Days/Images/Day66_config7.png)

이제 새 Playbook 이름인 `ansible-playbook playbook3.yml`로 Playbook을 다시 실행하면 deprecated가 발생한 것을 확인할 수 있으며, 다음에 수정할 수 있습니다.

![](/2022/Days/Images/Day66_config8.png)

Playbook이 실행되었지만, deprecated가 발생했으므로 이제 방법을 수정해야 합니다. 이를 위해 tasks/main.yml의 include 옵션을 아래와 같이 import_tasks로 변경했습니다.

![](/2022/Days/Images/Day66_config9.png)

이 파일은 [ansible-scenario3](/2022/Days/Configmgmt/ansible-scenario3)에서 찾을 수 있습니다.

또한 우리가 만들 `ansible-galaxy`를 사용하면서 몇 가지 Role을 더 만들 것입니다:

- common = 모든 서버용(`ansible-galaxy init roles/common`)
- nginx = 로드밸런서용(`ansible-galaxy init roles/nginx`)

![](/2022/Days/Images/Day66_config10.png)

여기서는 여기까지만 하고 다음 세션에서는 배포했지만, 아직 아무것도 하지 않은 다른 Node에 대한 작업을 시작하겠습니다.

## 자료

- [What is Ansible](https://www.youtube.com/watch?v=1id6ERvfozo)
- [Ansible 101 - Episode 1 - Introduction to Ansible](https://www.youtube.com/watch?v=goclfp6a2IQ)
- [NetworkChuck - You need to learn Ansible right now!](https://www.youtube.com/watch?v=5hycyr-8EKs&t=955s)
- [Your complete guide to Ansible](https://www.youtube.com/playlist?list=PLnFWJCugpwfzTlIJ-JtuATD2MBBD7_m3u)

위에 나열된 마지막 재생 목록은 이 섹션의 많은 코드와 아이디어가 나온 곳이며, 동영상 형식의 훌륭한 리소스이자 워크스루입니다.

[Day 67](day67.md)에서 봐요!
