---
title: '#90DaysOfDevOps - Ansible Playbooks - Day 65'
published: false
description: 90DaysOfDevOps - Ansible Playbooks
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049054
---

### Ansible Playbook

이 섹션에서는 적어도 Ansible의 경우 하나의 명령으로 여러 서버를 실행하여 긴 서버 목록을 재부팅하는 것과 같은 간단한 명령을 수행하고 각 서버에 개별적으로 연결해야 하는 번거로움을 줄일 수 있다는 점이 가장 큰 장점이라고 할 수 있습니다.

하지만 실제로 베어 운영 체제를 가져와서 해당 시스템에서 실행할 소프트웨어와 서비스를 선언하고 모두 원하는 상태로 실행되도록 하는 것은 어떨까요?

이것이 바로 Ansible Playbook이 필요한 이유입니다. Playbook을 사용하면 서버 그룹을 가져와서 해당 그룹에 대해 구성 및 설치 작업을 수행할 수 있습니다.

### Playbook 형식

Playbook > Play > Task

스포츠에 관심이 있는 분이라면 Playbook이라는 용어를 들어보셨을 텐데요, Playbook은 다양한 Play와 Task로 구성된 Play 방법을 팀에게 알려주는 것으로, Play를 스포츠나 게임의 세트 피스라고 생각하면 각 Play에 Task가 연관되어 있고, Play를 구성하는 여러 Task가 있을 수 있으며, Playbook에는 여러 가지 다른 Play가 있을 수 있습니다.

이러한 Playbook은 YAML(YAML은 마크업 언어가 아님)로 작성되어 있으며, 지금까지 다룬 많은 섹션, 특히 컨테이너와 Kubernetes에서 YAML 형식의 구성 파일을 찾을 수 있습니다.

playbook.yml이라는 간단한 Playbook을 살펴보겠습니다.

```Yaml
- name: Simple Play
  hosts: localhost
  connection: local
  tasks:
    - name: Ping me
      ping:
    - name: print os
      debug:
        msg: "{{ ansible_os_family }}"
```

위의 파일 [simple_play](/2022/Days/Configmgmt/simple_play.yml)을 찾을 수 있습니다. 그런 다음 `ansible-playbook simple_play.yml` 명령을 사용하면 다음 단계를 수행합니다.

![](/2022/Days/Images/Day65_config1.png)

'gathering facts'라는 첫 번째 Task가 발생한 것을 볼 수 있지만, 우리가 트리거하거나 요청하지 않았나요? 이 모듈은 원격 호스트에 대한 유용한 변수를 수집하기 위해 Playbook에서 자동으로 호출됩니다. [ansible.builtin.setup](https://docs.ansible.com/ansible/latest/collections/ansible/builtin/setup_module.html)

두 번째 Task는 Ping을 설정하는 것이었는데, 이것은 ICMP Ping이 아니라 원격 또는 로컬 호스트에 대한 연결 성공 시 `pong`을 반환하는 파이썬 스크립트입니다. [ansible.builtin.ping](https://docs.ansible.com/ansible/latest/collections/ansible/builtin/ping_module.html)

그런 다음 첫 번째 Task로 정의한 세 번째 또는 두 번째 Task는 OS를 알려주는 메시지 인쇄를 비활성화하지 않는 한 실행됩니다. 이 Task에서는 조건문을 사용하고 있으므로 모든 유형의 운영 체제에 대해 이 Playbook을 실행할 수 있으며, 그러면 OS 이름이 반환됩니다. 편의상 이 출력은 단순히 메시지를 출력하고 있지만 다음과 같이 말하는 Task를 추가할 수 있습니다:

```Yaml
tasks:
  - name: "shut down Debian flavoured systems"
    command: /sbin/shutdown -t now
    when: ansible_os_family == "Debian"
```

### Vagrant로 환경 설정하기

Vagrant를 사용하여 Node 환경을 설정할 것입니다. 저는 이것을 합리적인 4 Node로 유지하려고 하지만 300 Node 또는 3000 Node가 될 수 있으며 이것이 서버를 구성할 수 있는 Ansible 및 기타 구성 관리 도구의 힘이라는 것을 알 수 있기를 바랍니다.

이 파일은 [여기](/2022/Days/Configmgmt/Vagrantfile)에서 찾을 수 있습니다.

```Vagrant
Vagrant.configure("2") do |config|
  servers=[
    {
      :hostname => "db01",
      :box => "bento/ubuntu-21.10",
      :ip => "192.168.169.130",
      :ssh_port => '2210'
    },
    {
      :hostname => "web01",
      :box => "bento/ubuntu-21.10",
      :ip => "192.168.169.131",
      :ssh_port => '2211'
    },
    {
      :hostname => "web02",
      :box => "bento/ubuntu-21.10",
      :ip => "192.168.169.132",
      :ssh_port => '2212'
    },
    {
      :hostname => "loadbalancer",
      :box => "bento/ubuntu-21.10",
      :ip => "192.168.169.134",
      :ssh_port => '2213'
    }

  ]

config.vm.base_address = 600

  servers.each do |machine|

    config.vm.define machine[:hostname] do |node|
      node.vm.box = machine[:box]
      node.vm.hostname = machine[:hostname]

      node.vm.network :public_network, bridge: "Intel(R) Ethernet Connection (7) I219-V", ip: machine[:ip]
      node.vm.network "forwarded_port", guest: 22, host: machine[:ssh_port], id: "ssh"

      node.vm.provider :virtualbox do |v|
        v.customize ["modifyvm", :id, "--memory", 2048]
        v.customize ["modifyvm", :id, "--name", machine[:hostname]]
      end
    end
  end

end
```

`vagrant up` 명령을 사용하여 VirtualBox에서 이러한 머신을 스핀업하면 메모리를 더 추가할 수 있고 각 머신에 대해 다른 private_network 주소를 정의할 수도 있지만 제 환경에서는 이 방법이 작동합니다. 컨트롤 박스는 리눅스 섹션에서 배포한 우분투 데스크톱이라는 점을 기억하세요.

리소스가 제한되어 있는 경우 `vagrant up web01 web02`를 실행하여 여기서 사용 중인 웹서버만 불러올 수도 있습니다.

### Ansible 호스트 구성

이제 환경이 준비되었으므로 Ansible을 확인할 수 있으며, 이를 위해 우분투 데스크톱(이 데스크톱을 사용할 수도 있지만 아래 네트워크에 액세스하는 네트워크에 있는 모든 Linux 기반 머신을 동일하게 사용할 수 있음)을 컨트롤로 사용하고, ansible 호스트 파일에서 새 Node를 그룹에 추가해 보겠습니다, 이 파일을 인벤토리로 생각할 수 있으며, 이에 대한 대안으로 `-i filename`을 사용하여 ansible 명령의 일부로 호출되는 또 다른 인벤토리 파일을 사용할 수 있습니다. 이는 프로덕션, 테스트 및 스테이징 환경마다 다른 파일을 가질 수 있으므로 호스트 파일을 사용하는 것보다 유용할 수 있습니다. 기본 호스트 파일을 사용하고 있으므로 이 파일을 기본으로 사용하므로 지정할 필요가 없습니다.

기본 호스트 파일에 다음을 추가했습니다.

```Text
[control]
ansible-control

[proxy]
loadbalancer

[webservers]
web01
web02

[database]
db01

```

![](/2022/Days/Images/Day65_config2.png)

계속 진행하기 전에 Node에 대해 명령을 실행할 수 있는지 확인하기 위해 `ansible nodes -m command -a hostname`을 실행해 보겠습니다. 이 간단한 명령은 연결이 있는지 테스트하고 호스트 이름을 다시 보고합니다.

또한, 연결을 보장하기 위해 /etc/hosts 파일 내의 Ubuntu 제어 Node에 이러한 Node와 IP를 추가했습니다. 우분투 상자에서 각 Node에 대해 SSH 구성을 수행해야 할 수도 있습니다.

```Text
192.168.169.140 ansible-control
192.168.169.130 db01
192.168.169.131 web01
192.168.169.132 web02
192.168.169.133 loadbalancer
```

![](/2022/Days/Images/Day65_config3.png)

이 단계에서는 제어 Node와 서버 Node 사이에 SSH 키를 설정하는 과정을 진행하겠습니다. 다음 단계에서는 호스트의 파일에 변수를 추가하여 사용자 이름과 비밀번호를 제공하는 방법을 사용할 수 있습니다. 이 방법은 결코 모범 사례가 될 수 없으므로 권장하지 않습니다.

SSH를 설정하고 Node 간에 공유하려면 아래 단계를 따르세요. 비밀번호(`vagrant`)를 묻는 메시지가 표시되면 `y`를 몇 번 눌러야 수락할 수 있습니다.

`ssh-keygen`

![](/2022/Days/Images/Day65_config5.png)

`ssh-copy-id localhost`

![](/2022/Days/Images/Day65_config6.png)

이제 모든 VM이 켜져 있다면 `ssh-copy-id web01 && ssh-copy-id web02 && ssh-copy-id loadbalancer && ssh-copy-id db01`을 실행하면 비밀번호를 입력하라는 메시지가 표시됩니다(이 경우 비밀번호는 `vagrant`입니다).

모든 VM을 실행하지 않고 웹서버만 실행하고 있으므로 `ssh-copy-id web01 && ssh-copy-id web02`를 발급했습니다.

![](/2022/Days/Images/Day65_config7.png)

Playbook을 실행하기 전에 그룹과 간단하게 연결되었는지 확인하고 싶어서 `ansible webservers -m ping`을 실행하여 연결을 테스트했습니다.

![](/2022/Days/Images/Day65_config4.png)

### 첫 번째 "진짜" Ansible Playbook

첫 번째 Ansible Playbook은 웹 서버를 구성하는 것으로, 호스트 파일에서 [webservers] 그룹 아래에 웹 서버를 그룹화했습니다.

Playbook을 실행하기 전에 web01과 web02에 apache가 설치되어 있지 않은 것을 확인할 수 있습니다. 아래 스크린샷 상단은 이 Playbook을 실행하기 위해 ansible 컨트롤 내에서 생성한 폴더 및 파일 레이아웃을 보여줍니다. `playbook1.yml`이 있고, template 폴더에는 `index.html.j2`와 `ports.conf.j2` 파일이 있습니다. 이 파일은 위에 나열된 리포지토리 폴더에서 찾을 수 있습니다.

그런 다음 web01에 SSH로 접속하여 apache가 설치되어 있는지 확인합니다.

![](/2022/Days/Images/Day65_config8.png)

위에서 web01에 apache가 설치되어 있지 않다는 것을 알 수 있으므로 아래 Playbook을 실행하여 이 문제를 해결할 수 있습니다.

```Yaml
- hosts: webservers
  become: yes
  vars:
    http_port: 8000
    https_port: 4443
    html_welcome_msg: "Hello 90DaysOfDevOps"
  tasks:
  - name: ensure apache is at the latest version
    apt:
      name: apache2
      state: latest

  - name: write the apache2 ports.conf config file
    template:
      src: templates/ports.conf.j2
      dest: /etc/apache2/ports.conf
    notify:
    - restart apache

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

  handlers:
    - name: restart apache
      service:
        name: apache2
        state: restarted
```

위의 Playbook을 분석합니다:

- `host: webserver`는 이 Playbook을 실행할 그룹이 웹서버라는 그룹이라는 것을 의미합니다.
- `become: yes`는 Playbook을 실행하는 사용자가 원격 시스템에서 루트가 된다는 의미입니다. 루트 비밀번호를 입력하라는 메시지가 표시됩니다.
- 그런 다음 `vars`가 있는데, 이는 웹서버 전체에서 원하는 몇 가지 환경 변수를 정의합니다.

그런 다음 Task를 시작합니다,

- Task 1은 apache가 최신 버전을 실행 중인지 확인하는 것입니다.
- Task 2는 template 폴더에 있는 소스에서 ports.conf 파일을 작성하는 것입니다.
- Task 3은 기본 index.html 파일을 생성하는 것입니다.
- Task 4는 apache가 실행 중인지 확인하는 것입니다.

마지막으로 Handler 섹션인 [Handler: 변경에 대한 Task 실행](https://docs.ansible.com/ansible/latest/user_guide/playbooks_handlers.html)이 있습니다.

"때로는 컴퓨터에서 변경이 이루어질 때만 Task가 실행되기를 원할 때가 있습니다. 예를 들어, 태스크가 해당 서비스의 구성을 업데이트하지만, 구성이 변경되지 않은 경우 서비스를 다시 시작하고 싶을 수 있습니다. Ansible은 이 사용 사례를 해결하기 위해 Handler를 사용합니다. Handler는 알림을 받을 때만 실행되는 태스크입니다. 각 Handler는 전 세계적으로 고유한 이름을 가져야 합니다."

이 단계에서는 5개의 VM을 배포했다고 생각할 수 있습니다(Ansible 컨트롤 역할을 하는 Ubuntu 데스크톱 머신 포함). 다른 시스템은 이 섹션의 나머지 부분에서 다루게 될 것입니다.

### Playbook 실행

이제 Node에 대해 Playbook을 실행할 준비가 되었습니다. Playbook을 실행하려면 `ansible-playbook playbook1.yml`을 사용하면 됩니다. Playbook 내에서 Playbook이 실행될 호스트를 정의했으며, 정의한 Task를 안내합니다.

명령이 완료되면 Play와 Task를 보여주는 출력이 표시되며, 아래 이미지에서 원하는 상태를 설치하는 데 시간이 걸리는 것을 확인할 수 있습니다.

![](/2022/Days/Images/Day65_config9.png)

그런 다음 Node로 이동하여 Node에 소프트웨어가 설치되었는지 확인하여 이를 다시 확인할 수 있습니다.

![](/2022/Days/Images/Day65_config10.png)

위와 같이 두 개의 독립형 웹서버를 배포했으므로 이제 정의한 각각의 IP로 이동하여 새 웹 사이트를 가져올 수 있습니다.

![](/2022/Days/Images/Day65_config11.png)

이 섹션의 나머지 부분을 진행하면서 이 Playbook을 기반으로 구축할 것입니다. 또한 Ubuntu 데스크톱을 가져와서 Ansible을 사용하여 애플리케이션과 구성을 부트스트랩할 수 있는지 살펴보고 싶어서 이 부분도 다뤄볼 수 있을 것 같습니다. 예를 들어 명령에서 로컬 호스트를 사용하여 로컬 호스트에 대해 Playbook을 실행할 수 있다는 것을 보셨습니다.

여기에 추가해야 할 또 다른 사항은 우리는 실제로 Ubuntu VM으로만 작업하고 있지만 Ansible은 대상 시스템에 구애받지 않는다는 것입니다. 시스템을 관리하기 위해 이전에 언급했던 대안은 서버별로 서버를 관리할 수 있습니다(서버 수가 많을 경우 확장성이 떨어지고 Node가 3개일 때도 문제가 있음) Linux 섹션에서 다시 다룬 셸 스크립팅을 사용할 수도 있지만 이러한 Node는 잠재적으로 다를 수 있으므로 가능하지만 누군가 스크립트를 유지 및 관리해야 합니다. Ansible은 무료이며 전문화된 스크립트가 필요 없는 대신 간편한 버튼을 누릅니다.

## 자료

- [What is Ansible](https://www.youtube.com/watch?v=1id6ERvfozo)
- [Ansible 101 - Episode 1 - Introduction to Ansible](https://www.youtube.com/watch?v=goclfp6a2IQ)
- [NetworkChuck - You need to learn Ansible right now!](https://www.youtube.com/watch?v=5hycyr-8EKs&t=955s)
- [Your complete guide to Ansible](https://www.youtube.com/playlist?list=PLnFWJCugpwfzTlIJ-JtuATD2MBBD7_m3u)

위에 나열된 마지막 재생 목록은 이 섹션의 많은 코드와 아이디어가 나온 곳이며, 동영상 형식의 훌륭한 리소스이자 워크스루입니다.

[Day 66](day66.md)에서 봐요!
