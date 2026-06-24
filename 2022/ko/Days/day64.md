---
title: '#90DaysOfDevOps - Ansible: Getting Started - Day 64'
published: false
description: '90DaysOfDevOps - Ansible: Getting Started'
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048765
---

## Ansible: 시작하기

어제 [Day 63](day63.md)에서 Ansible이 무엇인지에 대해 조금 다루었지만 여기서는 이에 대해 조금 더 자세히 알아보겠습니다. 먼저 Ansible은 RedHat에서 제공합니다. 둘째, 에이전트가 없으며 SSH를 통해 연결하고 명령을 실행합니다. 셋째, 크로스 플랫폼(Linux 및 macOS, WSL2) 및 오픈 소스(엔터프라이즈용 유료 옵션도 있음)이며 다른 모델에 비해 Ansible은 구성을 push합니다.

### Ansible 설치

여러분이 상상할 수 있듯이, RedHat과 Ansible 팀은 Ansible을 문서화하는 데 환상적인 작업을 해왔습니다. 이는 일반적으로 [여기](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)에서 찾을 수 있는 설치 단계부터 시작됩니다. Ansible은 에이전트가 없는 자동화 도구이며, 이 도구는 "Control Node"라고 하는 시스템에 배포되고 이 Control Node에서 SSH를 통해 머신 및 기타 장치(네트워크일 수 있음)를 관리한다고 말씀드린 것을 기억하세요.

위에 링크된 문서에는 Windows OS를 Control Node로 사용할 수 없다고 명시되어 있습니다.

Control Node 및 적어도 이 데모에서는 [Linux 섹션](day20.md)에서 만든 Linux VM을 Control Node로 사용하겠습니다.

이 시스템은 우분투를 실행 중이며 설치 단계는 다음 명령만 있으면 됩니다.

```Shell
sudo apt update
sudo apt install software-properties-common
sudo add-apt-repository --yes --update ppa:ansible/ansible
sudo apt install ansible
```

이제 Control Node에 ansible이 설치되어 있어야 하며, `ansible --version`을 실행하여 확인할 수 있으며 아래와 비슷한 내용이 표시됩니다.

![](/2022/Days/Images/Day64_config1.png)

이제 환경의 다른 Node를 제어하는 방법을 살펴보기 전에 로컬 머신에 대해 `ansible localhost -m ping` 명령을 실행하여 ansible의 기능을 확인할 수도 있습니다. 이 명령은 [Ansible 모듈](https://docs.ansible.com/ansible/2.9/user_guide/modules_intro.html)을 사용하며, 여러 시스템에서 하나의 작업을 빠르게 수행할 수 있는 방법이기도 합니다. 로컬 호스트만으로는 그다지 재미있지 않지만, 무언가를 얻거나 모든 시스템이 가동 중이고 1000개 이상의 서버와 디바이스가 있다고 상상해 보세요.

![](/2022/Days/Images/Day64_config2.png)

또는 실제 모듈의 실제 사용법은 `ansible webservers -m service -a "name=httpd state=started"`와 같이 모든 웹서버에 httpd 서비스가 실행 중인지 여부를 알려주는 것일 수 있습니다. 이 명령에 사용된 웹서버 용어에 대해 간략히 설명했습니다.

### 호스트

위에서 localhost를 사용하여 시스템에 대한 간단한 핑 모듈을 실행하는 방법으로는 네트워크에 다른 컴퓨터를 지정할 수 없습니다. 예를 들어 VirtualBox가 실행 중인 Windows 호스트에는 IP 10.0.0.1의 네트워크 어댑터가 있지만 아래에서 볼 수 있듯이 핑으로 연결할 수 있지만 ansible을 사용하여 해당 작업을 수행할 수 없습니다.

![](/2022/Days/Images/Day64_config3.png)

이러한 작업으로 자동화할 호스트 또는 Node를 지정하려면 이를 정의해야 합니다. 시스템의 /etc/ansible 디렉토리로 이동하여 정의할 수 있습니다.

![](/2022/Days/Images/Day64_config4.png)

편집하려는 파일은 호스트 파일이며, 텍스트 편집기를 사용하여 호스트를 정의할 수 있습니다. 호스트 파일에는 파일을 사용하고 수정하는 방법에 대한 많은 훌륭한 지침이 포함되어 있습니다. 아래로 스크롤하여 [windows]라는 새 그룹을 만들고 해당 호스트에 대한 `10.0.0.1` IP 주소를 추가하겠습니다. 파일을 저장합니다.

![](/2022/Days/Images/Day64_config5.png)

그러나 Ansible이 시스템에 연결하려면 SSH를 사용할 수 있어야 한다고 말씀드린 것을 기억하세요. 아래에서 볼 수 있듯이 `ansible windows -m ping`을 실행하면 SSH를 통한 연결에 실패하여 연결할 수 없다는 메시지가 표시됩니다.

![](/2022/Days/Images/Day64_config6.png)

이제 인벤토리에 호스트를 추가하기 시작했는데, 모든 장치를 정의하는 곳이기 때문에 이 파일의 다른 이름인 네트워크 장치, 스위치 및 라우터도 여기에 추가하고 그룹화할 수 있습니다. 하지만 호스트 파일에는 Linux 시스템 그룹에 액세스하기 위한 자격 증명도 추가했습니다.

![](/2022/Days/Images/Day64_config7.png)

이제 `ansible Linux -m ping`을 실행하면 아래와 같이 성공합니다.

![](/2022/Days/Images/Day64_config8.png)

이제 구성을 자동화하려는 대상 시스템인 Node 요구 사항이 있습니다. 이 시스템에는 Ansible을 위한 어떤 것도 설치하지 않습니다(소프트웨어를 설치할 수는 있지만 필요한 Ansible의 클라이언트는 없습니다). Ansible은 SSH를 통해 연결하고 SFTP를 통해 모든 것을 전송합니다. (원하는 경우 SSH를 구성한 경우 SCP 대 SFTP를 사용할 수 있습니다.)

### Ansible 명령

리눅스 머신에 대해 `ansible Linux -m ping`을 실행하고 응답을 얻을 수 있는 것을 보셨겠지만, 기본적으로 Ansible을 사용하면 많은 adhoc 명령을 실행할 수 있습니다. 하지만 이 명령을 시스템 그룹에 대해 실행하여 해당 정보를 다시 가져올 수 있습니다. [adhoc 명령](https://docs.ansible.com/ansible/latest/user_guide/intro_adhoc.html)

명령을 반복하거나 이러한 명령을 실행하기 위해 개별 시스템에 로그인해야 하는 경우 Ansible이 도움이 될 수 있습니다. 예를 들어, 아래의 간단한 명령은 Linux 그룹에 추가하는 모든 시스템에 대한 모든 운영 체제 세부 정보를 출력합니다.  
`ansible linux -a "cat /etc/os-release"`

다른 사용 사례로는 시스템 재부팅, 파일 복사, 패커 및 사용자 관리 등이 있습니다. adhoc 명령과 Ansible 모듈을 결합할 수도 있습니다.

adhoc 명령은 선언적 모델을 사용하여 지정된 최종 상태에 도달하는 데 필요한 작업을 계산하고 실행합니다. adhoc 명령은 시작하기 전에 현재 상태를 확인하고 현재 상태가 지정된 최종 상태와 다르지 않으면 아무 작업도 수행하지 않음으로써 일종의 무임승차를 하는 셈입니다.

## 자료

- [What is Ansible](https://www.youtube.com/watch?v=1id6ERvfozo)
- [Ansible 101 - Episode 1 - Introduction to Ansible](https://www.youtube.com/watch?v=goclfp6a2IQ)
- [NetworkChuck - You need to learn Ansible right now!](https://www.youtube.com/watch?v=5hycyr-8EKs&t=955s)

[Day 65](day65.md)에서 봐요!
