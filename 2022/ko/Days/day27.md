---
title: '#90DaysOfDevOps - Getting Hands-On with Python & Network - Day 27'
published: false
description: 90DaysOfDevOps - Getting Hands-On with Python & Network
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048735
---

## 파이썬 및 네트워크 실습하기

네트워킹 기초의 마지막 섹션에서는 [Day 26](day26.md)에 만든 실습 환경으로 몇 가지 자동화 작업과 도구를 다뤄보겠습니다.

클라이언트 vs telnet에서 장치에 연결하기 위해 SSH 터널을 사용할 것입니다. 클라이언트와 기기 사이에 생성된 SSH 터널은 암호화됩니다. [Day 18](day18.md)의 Linux 섹션에서도 SSH에 대해 다루었습니다.

## 가상 에뮬레이트 환경에 액세스하기

스위치와 상호 작용하려면 EVE-NG 네트워크 내부에 워크스테이션이 필요하거나 Python이 설치된 Linux 박스를 배포하여 자동화를 수행할 수 있습니다([Resource for setting up Linux inside EVE-NG](https://www.youtube.com/watch?v=3Qstk3zngrY)) 또는 저처럼 워크스테이션에서 액세스할 수 있는 클라우드를 정의할 수 있습니다.

![](/2022/Days/Images/Day27_Networking3.png)

이를 위해 캔버스를 마우스 오른쪽 버튼으로 클릭하고 네트워크를 선택한 다음 "Management(Cloud0)"를 선택하면 홈 네트워크에 연결됩니다.

![](/2022/Days/Images/Day27_Networking4.png)

하지만 이 네트워크 내부에는 아무것도 없으므로 새 네트워크에서 각 장치로 연결을 추가해야 합니다. (제 네트워킹 지식이 더 필요한데, 다음 단계로 최상위 라우터에 이 작업을 수행한 다음, 이 하나의 케이블을 통해 나머지 네트워크에 연결할 수 있을 것 같나요?)

그런 다음 각 장치에 로그온하고 클라우드가 들어오는 위치에 적용되는 인터페이스에 대해 다음 명령을 실행했습니다.

```
enable
config t
int gi0/0
IP add DHCP
no sh
exit
exit
sh ip int br
```

마지막 단계에서는 홈 네트워크의 DHCP 주소를 제공합니다. 내 장치 네트워크 목록은 다음과 같습니다:

| Node    | IP Address   | Home Network IP |
| ------- | ------------ | --------------- |
| Router  | 10.10.88.110 | 192.168.169.115 |
| Switch1 | 10.10.88.111 | 192.168.169.178 |
| Switch2 | 10.10.88.112 | 192.168.169.193 |
| Switch3 | 10.10.88.113 | 192.168.169.125 |
| Switch4 | 10.10.88.114 | 192.168.169.197 |

### 네트워크 장치에 SSH

위의 설정이 완료되었으므로 이제 워크스테이션을 사용하여 홈 네트워크의 장치에 연결할 수 있습니다. 저는 Putty를 사용하고 있지만 장치에 SSH를 할 수 있는 git bash와 같은 다른 터미널에도 액세스할 수 있습니다.

아래에서 라우터 장치에 대한 SSH 연결을 확인할 수 있습니다. (R1)

![](/2022/Days/Images/Day27_Networking5.png)

### 파이썬을 사용하여 장치에서 정보 수집하기

Python을 활용하는 첫 번째 예는 모든 장치에서 정보를 수집하는 것이며, 특히 각 장치에 연결하여 간단한 명령을 실행하여 인터페이스 구성 및 설정을 제공할 수 있기를 원합니다. 이 스크립트를 여기에 저장해 두었습니다. [netmiko_con_multi.py](/2022/Days/Networking/netmiko_con_multi.py)

이제 이 스크립트를 실행하면 모든 장치에서 각 포트 구성을 볼 수 있습니다.

![](/2022/Days/Images/Day27_Networking6.png)

다양한 장치를 많이 사용하는 경우 이 스크립트를 작성하면 한 곳에서 모든 구성을 중앙에서 빠르게 제어하고 이해할 수 있어 편리합니다.

### Python을 사용하여 디바이스 구성하기

위의 내용은 유용하지만 Python을 사용하여 장치를 구성하는 것은 어떻습니까? 시나리오에서 `SW1`과 `SW2` 사이에 트렁크 포트가 있습니다. 이 작업을 여러 개의 동일한 스위치에서 자동화하고 각 스위치에 수동으로 연결하여 구성을 변경할 필요가 없다고 다시 상상해 보겠습니다.

이를 위해 [netmiko_sendchange.py](/2022/Days/Networking/netmiko_sendchange.py)를 사용할 수 있습니다. 이렇게 하면 SSH를 통해 연결되고 `SW1`에서 해당 변경을 수행하여 `SW2`로 변경됩니다.

![](/2022/Days/Images/Day27_Networking7.png)

이제 코드를 보면 `sending configuration to device`이라는 메시지가, 이것이 발생했는지 확인되지 않은 경우 스크립트에 추가 코드를 추가하여 스위치에서 해당 확인 및 유효성 검사를 수행하거나 이전에 스크립트를 수정하여 이를 표시할 수 있습니다. [netmiko_con_multi_vlan.py](/2022/Days/Networking/netmiko_con_multi_vlan.py)

![](/2022/Days/Images/Day27_Networking8.png)

### 디바이스 구성 백업하기

또 다른 사용 사례는 네트워크 구성을 캡처하여 백업하는 것이지만, 네트워크에 있는 모든 장치에 연결하고 싶지는 않으므로 [backup.py](/2022/Days/Networking/backup.py)를 사용하여 이 작업을 자동화할 수도 있습니다. 또한 백업하려는 IP 주소로 [backup.txt](/2022/Days/Networking/backup.txt)를 채워야 합니다.

스크립트를 실행하면 아래와 같은 내용이 표시됩니다.

![](/2022/Days/Images/Day27_Networking9.png)

파이썬으로 간단한 인쇄 스크립트를 작성했기 때문에 백업 파일도 함께 보여드릴 수 있습니다.

![](/2022/Days/Images/Day27_Networking10.png)

### Paramiko

SSH를 위해 널리 사용되는 Python 모듈입니다. 자세한 내용은 [공식 GitHub 링크](https://github.com/paramiko/paramiko)에서 확인할 수 있습니다.

이 모듈은 `pip install paramiko` 명령어를 사용하여 설치할 수 있습니다.

![](/2022/Days/Images/Day27_Networking1.png)

파이썬 셸에 들어가서 Paramiko 모듈을 가져와서 설치가 완료되었는지 확인할 수 있습니다.

![](/2022/Days/Images/Day27_Networking2.png)

### Netmiko

Netmiko 모듈은 네트워크 디바이스만을 대상으로 하지만, Paramiko는 SSH 연결 전반을 처리하는 더 광범위한 도구입니다.

위에서 Paramiko와 함께 사용한 Netmiko는 `pip install netmiko`를 사용하여 설치할 수 있습니다.

Netmiko는 다양한 네트워크 벤더와 디바이스를 지원하며, 지원 디바이스 목록은 [GitHub Page](https://github.com/ktbyers/netmiko#supports)에서 확인할 수 있습니다.

### 기타 모듈

우리가 살펴볼 기회가 없었지만, 네트워크 자동화와 관련하여 훨씬 더 많은 기능을 제공하는 몇 가지 다른 모듈도 언급할 가치가 있습니다.

`netaddr`은 IP 주소 작업 및 조작에 사용되며, 설치는 `pip install netaddr`로 간단합니다.

스위치 구성을 엑셀 스프레드시트에 많이 저장하고 싶을 때, `xlrd`를 사용하면 스크립트가 엑셀 통합 문서를 읽고 행과 열을 행렬로 변환할 수 있습니다. `pip install xlrd`를 실행하면 모듈을 설치할 수 있습니다.

제가 미처 살펴보지 못한 네트워크 자동화를 사용할 수 있는 더 많은 사용 사례는 [여기](https://github.com/ktbyers/pynet/tree/master/presentations/dfwcug/examples)에서 확인할 수 있습니다.

네트워킹은 제가 한동안 다루지 않았던 분야이고 다뤄야 할 내용이 훨씬 더 많지만, 제가 작성한 노트와 이 글에서 공유한 리소스가 누군가에게 도움이 되기를 바랍니다.

## 자료

- [Free Course: Introduction to EVE-NG](https://www.youtube.com/watch?v=g6B0f_E0NMg)
- [EVE-NG - Creating your first lab](https://www.youtube.com/watch?v=9dPWARirtK8)
- [3 Necessary Skills for Network Automation](https://www.youtube.com/watch?v=KhiJ7Fu9kKA&list=WL&index=122&t=89s)
- [Computer Networking full course](https://www.youtube.com/watch?v=IPvYjXCsTg8)
- [Practical Networking](http://www.practicalnetworking.net/)
- [Python Network Automation](https://www.youtube.com/watch?v=xKPzLplPECU&list=WL&index=126)

저는 네트워크 엔지니어가 아니기 때문에 여기서 사용하는 대부분의 예제는 무료는 아니지만 네트워크 자동화를 이해하는 데 도움이 되는 시나리오 중 일부를 이 방대한 책에서 가져온 것입니다.

- [Hands-On Enterprise Automation with Python (Book)](https://www.packtpub.com/product/hands-on-enterprise-automation-with-python/9781788998512)

클라우드 컴퓨팅에 대해 알아보고 주제에 대한 이해와 기초 지식을 쌓을 수 있는 [Day 28](day28.md)에서 봐요!
