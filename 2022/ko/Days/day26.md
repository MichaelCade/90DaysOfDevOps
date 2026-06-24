---
title: '#90DaysOfDevOps - Building our Lab - Day 26'
published: false
description: 90DaysOfDevOps - Building our Lab
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048762
---

## 실험실 구축

EVE-NG를 사용하여 에뮬레이트된 네트워크를 계속 설정한 다음 몇 가지 디바이스를 배포하고 이러한 디바이스의 구성을 자동화할 수 있는 방법에 대해 생각해 보겠습니다. [Day 25](day25.md)에서는 VMware Workstation을 사용하여 머신에 EVE-NG를 설치하는 방법을 다루었습니다.

### EVE-NG 클라이언트 설치하기

장치에 SSH를 할 때 사용할 애플리케이션을 선택할 수 있는 클라이언트 팩도 있습니다. 또한 링크 간 패킷 캡처를 위해 Wireshark를 설정합니다. 사용 중인 OS(윈도우, 맥OS, 리눅스)에 맞는 클라이언트 팩을 다운로드할 수 있습니다.

[EVE-NG Client Download](https://www.eve-ng.net/index.php/download/)

![](/2022/Days/Images/Day26_Networking1.png)

빠른 팁: Linux를 클라이언트로 사용하는 경우 이 [client pack](https://github.com/SmartFinn/eve-ng-integration)이 있습니다.

설치는 간단하며, 기본값을 그대로 두는 것이 좋습니다.

### 네트워크 Images 가져오기

이 단계는 어려웠습니다. 마지막에 링크할 몇 가지 동영상을 통해 라우터 및 스위치 이미지를 업로드하는 방법과 위치를 알려주면서 몇 가지 리소스와 다운로드 링크를 제공합니다.

저는 모든 것을 교육 목적으로 사용한다는 점에 유의하는 것이 중요합니다. 네트워크 공급업체에서 공식 이미지를 다운로드하는 것이 좋습니다.

[Blog & Links to YouTube videos](https://loopedback.com/2019/11/15/setting-up-eve-ng-for-ccna-ccnp-ccie-level-studies-includes-multiple-vendor-node-support-an-absolutely-amazing-study-tool-to-check-out-asap/)

[How To Add Cisco VIRL vIOS image to Eve-ng](https://networkhunt.com/how-to-add-cisco-virl-vios-image-to-eve-ng/)

전반적으로 단계가 조금 복잡하고 훨씬 더 쉬울 수도 있지만 위의 블로그와 동영상은 EVE-NG 박스에 이미지를 추가하는 과정을 안내합니다.

저는 FileZilla를 사용하여 SFTP를 통해 qcow2를 VM으로 전송했습니다.

실습에는 Cisco vIOS L2(스위치)와 Cisco vIOS(라우터)가 필요합니다.

### 실습 만들기

EVE-NG 웹 인터페이스 내에서 새로운 네트워크 Topology를 생성하겠습니다. 외부 네트워크에 대한 게이트웨이 역할을 할 스위치 4개와 라우터 1개가 있습니다.

| Node    | IP Address   |
| ------- | ------------ |
| Router  | 10.10.88.110 |
| Switch1 | 10.10.88.111 |
| Switch2 | 10.10.88.112 |
| Switch3 | 10.10.88.113 |
| Switch4 | 10.10.88.114 |

#### EVE-NG에 노드 추가하기

EVE-NG에 처음 로그인하면 아래와 같은 화면이 표시되며, 첫 번째 랩을 생성하는 것부터 시작하겠습니다.

![](/2022/Days/Images/Day26_Networking2.png)

랩 이름을 지정하고 다른 필드는 선택 사항입니다.

![](/2022/Days/Images/Day26_Networking3.png)

그러면 네트워크 생성을 시작할 수 있는 빈 캔버스가 표시됩니다. 캔버스를 마우스 오른쪽 버튼으로 클릭하고 노드 추가를 선택합니다.

여기에서 긴 노드 옵션 목록이 표시되며, 위의 과정을 따라 했다면 아래 표시된 파란색 두 개가 선택되어 있고 나머지는 회색으로 표시되어 선택할 수 없습니다.

![](/2022/Days/Images/Day26_Networking4.png)

실습에 다음을 추가하려고 합니다:

- Cisco vIOS 라우터 1대
- Cisco vIOS 스위치 4개

간단한 마법사를 실행하여 실습에 추가하면 다음과 같은 화면이 나타납니다.

![](/2022/Days/Images/Day26_Networking5.png)

#### 노드 연결하기

이제 라우터와 스위치 사이에 연결을 추가해야 합니다. 장치 위로 마우스를 가져가면 아래와 같이 연결 아이콘이 표시되고 이를 연결하려는 장치에 연결하면 아주 쉽게 연결할 수 있습니다.

![](/2022/Days/Images/Day26_Networking6.png)

환경 연결이 완료되면 마우스 오른쪽 클릭 메뉴에서 찾을 수 있는 상자나 원을 사용하여 물리적 경계나 위치를 정의하는 방법을 추가할 수도 있습니다. 실험실의 이름이나 IP 주소를 정의하고 싶을 때 유용한 텍스트를 추가할 수도 있습니다.

저는 아래와 같이 실험실을 만들었습니다.

![](/2022/Days/Images/Day26_Networking7.png)

위의 실습은 모두 전원이 꺼져 있으며, 모든 것을 선택하고 마우스 오른쪽 버튼을 클릭한 후 선택된 시작을 선택하면 실습을 시작할 수 있습니다.

![](/2022/Days/Images/Day26_Networking8.png)

실습을 시작하고 실행하면 각 장치에 콘솔을 연결할 수 있으며, 이 단계에서는 아무 구성도 하지 않은 상태로 매우 멍청한 상태임을 알 수 있습니다. 각 터미널에서 직접 복사하거나 생성하여 각 노드에 구성을 추가할 수 있습니다.

참조를 위해 리포지토리의 네트워킹 폴더에 구성을 남겨두겠습니다.

| Node    | Configuration                    |
| ------- | -------------------------------- |
| Router  | [R1](/2022/Days/Networking/R1)   |
| Switch1 | [SW1](/2022/Days/Networking/SW1) |
| Switch2 | [SW2](/2022/Days/Networking/SW2) |
| Switch3 | [SW3](/2022/Days/Networking/SW3) |
| Switch4 | [SW4](/2022/Days/Networking/SW4) |

## 자료

- [Free Course: Introduction to EVE-NG](https://www.youtube.com/watch?v=g6B0f_E0NMg)
- [EVE-NG - Creating your first lab](https://www.youtube.com/watch?v=9dPWARirtK8)
- [3 Necessary Skills for Network Automation](https://www.youtube.com/watch?v=KhiJ7Fu9kKA&list=WL&index=122&t=89s)
- [Computer Networking full course](https://www.youtube.com/watch?v=IPvYjXCsTg8)
- [Practical Networking](http://www.practicalnetworking.net/)
- [Python Network Automation](https://www.youtube.com/watch?v=xKPzLplPECU&list=WL&index=126)

저는 네트워크 엔지니어가 아니기 때문에 여기서 사용하는 대부분의 예제는 무료는 아니지만 네트워크 자동화를 이해하는 데 도움이 되는 시나리오 중 일부를 이 방대한 책에서 가져온 것입니다.

- [Hands-On Enterprise Automation with Python (Book)](https://www.packtpub.com/product/hands-on-enterprise-automation-with-python/9781788998512)

[Day 27](day27.md)에서 봐요!
