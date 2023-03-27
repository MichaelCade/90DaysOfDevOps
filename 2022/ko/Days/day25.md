---
title: '#90DaysOfDevOps - Python for Network Automation - Day 25'
published: false
description: 90DaysOfDevOps - Python for Network Automation
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049038
---

## 네트워크 자동화를 위한 Python

Python은 자동화된 네트워크 운영에 사용되는 표준 언어입니다.

네트워크 자동화뿐만 아니라 리소스를 찾을 때 어디에나 있는 것 같고, 앞서 언급했듯이 파이썬이 아니라면 일반적으로 파이썬으로도 작성된 Ansible입니다.

이미 언급한 것 같지만 '프로그래밍 언어 배우기' 섹션에서 제가 파이썬 대신 Go를 선택한 이유는 회사에서 Go로 개발 중이어서 배울 수 있는 좋은 이유였지만, 그렇지 않았다면 파이썬을 배우는 데 시간이 더 걸렸을 것입니다.

- 가독성 및 사용 편의성 - 파이썬은 그냥 말이 되는 것 같습니다. 코드에서 블록을 시작하고 끝내기 위해 `{}`에 대한 요구 사항이 없는 것 같습니다. 여기에 VS Code와 같은 강력한 IDE를 결합하면 파이썬 코드를 실행할 때 매우 쉽게 시작할 수 있습니다.

여기서 언급할 만한 또 다른 IDE는 Pycharm일 수 있습니다.

- 라이브러리 - 파이썬의 확장성은 네트워크 자동화뿐만 아니라 실제로 모든 종류의 장치와 구성을 위한 라이브러리가 많이 있다는 점을 앞서 언급했습니다. 여기에서 방대한 양을 확인할 수 있습니다. [PyPi](https://pypi.python.org/pypi)

라이브러리를 워크스테이션에 다운로드하려면 `pip`라는 도구를 사용하여 PyPI에 연결하고 로컬로 다운로드하면 됩니다. Cisco, 주니퍼, 아리스타와 같은 네트워크 벤더는 자사 디바이스에 쉽게 액세스할 수 있도록 라이브러리를 개발했습니다.

- 강력하고 효율적 - Go 파트에서 "Hello World" 시나리오를 진행하면서 6줄 정도의 코드를 작성했던 것을 기억하시나요? Python에서는

```python
print('hello world')
```

위의 모든 사항을 종합하면 자동화 작업 시 파이썬이 언급되는 이유를 쉽게 알 수 있을 것입니다.

몇 년 전에도 네트워크 장치와 상호 작용하여 구성 백업을 자동화하거나 장치에 대한 로그 및 기타 인사이트를 수집하는 스크립트가 있었을 가능성이 있다는 점에 유의하는 것이 중요하다고 생각합니다. 여기서 말하는 자동화는 조금 다른데, 이는 전반적인 네트워킹 환경이 이러한 사고방식에 더 잘 맞도록 변화하여 더 많은 자동화를 가능하게 했기 때문입니다.

- 소프트웨어 정의 네트워크 - SDN 컨트롤러는 네트워크의 모든 디바이스에 컨트롤 플레인 구성을 제공하는 책임을 지므로 네트워크 변경에 대한 단일 접점만 있으면 더 이상 모든 디바이스에 텔넷 또는 SSH로 접속할 필요가 없으며, 사람이 이 작업을 수행해야 하므로 실패 또는 구성 오류의 가능성이 반복적으로 발생할 수 있습니다.

- 높은 수준의 오케스트레이션 - 이러한 SDN 컨트롤러에서 한 단계 더 올라가면 서비스 수준을 오케스트레이션 할 수 있으며, 이 오케스트레이션 계층을 원하는 플랫폼(VMware, Kubernetes, 퍼블릭 클라우드 등)에 통합할 수 있습니다.

- 정책 기반 관리 - 무엇을 원하시나요? 원하는 상태는 무엇인가요? 이를 설명하면 시스템에서 원하는 상태가 되기 위한 모든 세부 사항을 파악합니다.

## 실습 환경 설정하기

모든 사람이 실제 라우터, 스위치 및 기타 네트워킹 장치에 액세스할 수 있는 것은 아닙니다.

저는 앞서 언급한 몇 가지 도구를 살펴보면서 네트워크 구성을 자동화하는 방법을 직접 실습하고 배울 수 있도록 하고 싶었습니다.

옵션에 관해서는 몇 가지를 선택할 수 있습니다.

- [GNS3 VM](https://www.gns3.com/software/download-vm)
- [Eve-ng](https://www.eve-ng.net/)
- [Unimus](https://unimus.net/) 실험실 환경은 아니지만 흥미로운 개념입니다.

앞서 언급했듯이 [Eve-ng](https://www.eve-ng.net/)를 사용하여 랩을 구축할 예정입니다. 가상 환경은 다양한 시나리오를 테스트할 수 있는 샌드박스 환경을 가질 수 있다는 것을 의미하며, 다양한 디바이스와 Topology로 실행 수 있다는 점도 흥미로울 수 있습니다.

저희는 커뮤니티 에디션을 통해 EVE-NG의 모든 것을 보여드리려고 합니다.

### 시작하기

[커뮤니티 에디션](https://www.eve-ng.net/index.php/download/)은 ISO 및 OVF 형식으로 제공됩니다.

저희는 OVF 다운로드를 사용할 예정이지만, ISO를 사용하면 하이퍼바이저 없이 베어메탈 서버에서 빌드할 수 있습니다.

![](/2022/Days/Images/Day25_Networking1.png)

여기서는 vExpert를 통해 라이선스가 있는 VMware Workstation을 사용하지만, [documentation](https://www.eve-ng.net/index.php/documentation/installation/system-requirement/)에 언급된 다른 옵션이나 VMware Player도 동일하게 사용할 수 있습니다. 안타깝게도 이전에 사용하던 가상 박스는 사용할 수 없습니다!

가상 박스가 지원되기는 하지만 GNS3에서 문제가 발생했던 부분이기도 합니다.

[Download VMware Workstation Player - FREE](https://www.vmware.com/uk/products/workstation-player.html)

[VMware Workstation PRO](https://www.vmware.com/uk/products/workstation-pro.html) 역시 무료 평가판 기간이 있다고 합니다!

### VM웨어 워크스테이션 프로에 설치하기

이제 하이퍼바이저 소프트웨어 다운로드 및 설치가 완료되었고, EVE-NG OVF가 다운로드되었습니다. VMware Player를 사용 중이시라면 이 과정이 동일한지 알려주시기 바랍니다.

이제 구성할 준비가 되었습니다.

VMware 워크스테이션을 열고 `file`과 `open`를 선택합니다.

![](/2022/Days/Images/Day25_Networking2.png)

EVE-NG OVF 이미지를 다운로드하면 압축 파일 안에 들어 있습니다. 해당 폴더에 내용을 압축을 풀면 다음과 같은 모양이 됩니다.

![](/2022/Days/Images/Day25_Networking3.png)

EVE-NG OVF 이미지를 다운로드한 위치로 이동하여 가져오기를 시작합니다.

알아볼 수 있는 이름을 지정하고 시스템 어딘가에 가상 머신을 저장합니다.

![](/2022/Days/Images/Day25_Networking4.png)

가져오기가 완료되면 프로세서 수를 4개로 늘리고 할당된 메모리를 8GB로 늘립니다. (최신 버전으로 가져온 후 그렇지 않은 경우 VM 설정을 편집해야 합니다.)

또한 Intel VT-x/EPT 또는 AMD-V/RVI 가상화 확인란이 활성화되어 있는지 확인합니다. 이 옵션은 VMware 워크스테이션이 가상화 플래그를 게스트 OS에 전달하도록 지시합니다(중첩 가상화). CPU가 이를 허용함에도 불구하고 가상 박스가 있는 GNS3에서 이 문제가 발생했습니다.

![](/2022/Days/Images/Day25_Networking5.png)

### 전원 켜기 및 액세스

사이드노트 & Rabbit hole: 가상박스에서는 작동하지 않는다고 말씀드린 것을 기억하세요! 네, VM웨어 워크스테이션과 EVE-NG에서도 같은 문제가 발생했지만, 가상화 플랫폼의 잘못은 아니었습니다!

Windows 머신에서 WSL2를 실행 중이며 이로 인해 환경 내부에 중첩된 모든 것을 실행할 수 있는 기능이 제거된 것 같습니다. WSL2를 사용할 때 CPU의 인텔 VT-d 가상화 측면이 제거되는 것처럼 보이는데 왜 우분투 VM이 실행되는지 혼란스럽습니다.

이 문제를 해결하려면 Windows 시스템에서 다음 명령을 실행하고 시스템을 재부팅 할 수 있으며, 이 명령이 꺼져 있는 동안에는 WSL2를 사용할 수 없다는 점에 유의하세요.

`bcdedit /set hypervisorlaunchtype off`

다시 돌아가서 WSL2를 사용하려면 이 명령을 실행하고 재부팅 해야 합니다.

`bcdedit /set hypervisorlaunchtype auto`

이 두 명령은 모두 관리자 권한으로 실행해야 합니다!

다시 돌아와서, 이제 VMware Workstation에 전원이 켜진 머신이 있어야 하며 다음과 유사한 프롬프트가 표시되어야 합니다.

![](/2022/Days/Images/Day25_Networking6.png)

위의 프롬프트에서 다음을 사용할 수 있습니다:

username = root  
password = eve

그러면 루트 password를 다시 입력하라는 메시지가 표시되며, 이 password는 나중에 호스트에 SSH할 때 사용됩니다.

그런 다음 호스트 이름을 변경할 수 있습니다.

![](/2022/Days/Images/Day25_Networking7.png)

다음으로 DNS 도메인 이름을 정의합니다. 저는 아래 이름을 사용했지만, 나중에 변경해야 할지는 아직 모르겠습니다.

![](/2022/Days/Images/Day25_Networking8.png)

그런 다음 네트워킹을 구성하고, 재부팅 후에도 주어진 IP 주소가 영구적으로 유지되도록 정적을 선택합니다.

![](/2022/Days/Images/Day25_Networking9.png)

마지막 단계로 워크스테이션에서 연결할 수 있는 네트워크의 고정 IP 주소를 제공합니다.

![](/2022/Days/Images/Day25_Networking10.png)

여기에는 네트워크, 기본 게이트웨이 및 DNS에 대한 서브넷 마스크를 제공해야 하는 몇 가지 추가 단계가 있습니다.

완료되면 재부팅되고 백업이 완료되면 고정 IP 주소를 가져와 브라우저에 입력할 수 있습니다.

![](/2022/Days/Images/Day25_Networking11.png)

GUI의 기본 username은 `admin`이고 password는 `eve`이며, SSH의 기본 username은 `root`이고 password는 `eve`이지만 설정 중에 변경한 경우 변경되었을 수 있습니다.

![](/2022/Days/Images/Day25_Networking12.png)

다른 콘솔을 탐색할 때 브라우저에 새 탭이 열리므로 콘솔과 네이티브에 대해 HTML5를 선택했습니다.

다음 단계는 다음과 같습니다:

- EVE-NG 클라이언트 팩 설치
- 일부 네트워크 이미지를 EVE-NG에 로드합니다.
- 네트워크 Topology 빌드
- 노드 추가하기
- 노드 연결하기
- 파이썬 스크립트 빌드 시작
- telnetlib, Netmiko, Paramiko, Pexpect 살펴보기

## 자료

- [Free Course: Introduction to EVE-NG](https://www.youtube.com/watch?v=g6B0f_E0NMg)
- [EVE-NG - Creating your first lab](https://www.youtube.com/watch?v=9dPWARirtK8)
- [3 Necessary Skills for Network Automation](https://www.youtube.com/watch?v=KhiJ7Fu9kKA&list=WL&index=122&t=89s)
- [Computer Networking full course](https://www.youtube.com/watch?v=IPvYjXCsTg8)
- [Practical Networking](http://www.practicalnetworking.net/)
- [Python Network Automation](https://www.youtube.com/watch?v=xKPzLplPECU&list=WL&index=126)

[Day 26](day26.md)에서 봐요!
