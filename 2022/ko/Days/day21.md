---
title: '#90DaysOfDevOps - The Big Picture: DevOps and Networking - Day 21'
published: false
description: 90DaysOfDevOps - The Big Picture DevOps and Networking
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048761
---

## 큰 그림: 데브옵스와 네트워킹

모든 섹션과 마찬가지로 공개 및 무료 교육 자료를 사용하고 있으며 많은 콘텐츠가 다른 사람의 저작물일 수 있습니다. 네트워킹 섹션의 경우, 표시된 콘텐츠의 대부분은 [Practical Networking](https://www.practicalnetworking.net/)의 무료 [Networking Fundamentals series](https://www.youtube.com/playlist?list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi)에서 가져온 것입니다. 리소스에도 링크와 함께 언급되어 있지만 커뮤니티의 관점에서 특정 기술 분야에 대한 이해를 돕기 위해 이 과정을 활용했기 때문에 이 점을 강조하는 것이 적절합니다. 이 저장소는 제 메모를 저장하고 커뮤니티가 이 과정과 나열된 리소스로부터 혜택을 받을 수 있도록 하는 저장소입니다.

21일 차에 오신 것을 환영합니다! 앞으로 7일 동안 네트워킹과 데브옵스가 가장 중요한 주제이지만, 네트워킹의 기본에 대해서도 살펴볼 필요가 있을 것입니다.

궁극적으로 앞서 말씀드렸듯이 데브옵스는 조직 내 문화와 프로세스 변화에 관한 것으로, 앞서 논의한 것처럼 가상 머신, 컨테이너 또는 Kubernetes가 될 수도 있지만 네트워크가 될 수도 있습니다. 데브옵스 관점에서 네트워크를 더 많이 포함해야 하는 인프라에 이러한 데브옵스 원칙을 사용한다면 사용 가능한 다양한 Topology 및 네트워킹 도구와 스택에서와 같이 네트워크에 대해서도 알아야 합니다.

저는 코드로 인프라스트럭처의 네트워킹 디바이스를 구성하고 가상 머신처럼 모든 것을 자동화해야 한다고 주장하지만, 그렇게 하려면 무엇을 자동화할지 잘 이해해야 합니다.

### 넷데브옵스 | 네트워크 데브옵스란 무엇인가요?

네트워크 데브옵스 또는 넷데브옵스라는 용어를 들어보셨을 수도 있습니다. 이미 네트워크 엔지니어이고 인프라 내의 네트워크 구성 요소를 잘 파악하고 있다면 DHCP, DNS, NAT 등과 같은 네트워킹에 사용되는 요소를 이해하고 있을 것입니다. 또한 하드웨어 또는 소프트웨어 정의 네트워킹 옵션, 스위치, 라우터 등에 대해서도 잘 이해하고 있을 것입니다.

하지만 네트워크 엔지니어가 아니라면 네트워크 데브옵스의 최종 목표를 이해하려면 이러한 영역 중 일부에 대한 전반적인 기초 지식이 필요할 수 있습니다.

그러나 이러한 용어와 관련하여 NetDevOps 또는 네트워크 데브옵스를 네트워크에 데브옵스 원칙과 관행을 적용하고 버전 제어 및 자동화 도구를 네트워크 생성, 테스트, 모니터링 및 배포에 적용하는 것으로 생각할 수 있습니다.

네트워크 데브옵스를 자동화가 필요한 것으로 생각한다면, 앞서 데브옵스가 팀 간의 사일로를 허무는 것에 대해 언급했습니다. 네트워킹 팀이 유사한 모델과 프로세스로 변경하지 않으면 병목 현상이 발생하거나 전체적으로 장애가 발생할 수 있습니다.

프로비저닝, 구성, 테스트, 버전 제어 및 배포와 관련된 자동화 원칙을 사용하는 것이 좋은 출발점입니다. 자동화는 전반적으로 배포 속도, 네트워킹 인프라의 안정성, 지속적인 개선을 가능하게 할 뿐만 아니라 테스트가 완료되면 여러 환경에서 프로세스를 공유할 수 있습니다. 예를 들어 한 환경에서 완전히 테스트 된 네트워크 정책은 이전에는 수동으로 작성된 프로세스와 달리 코드로 되어 있기 때문에 다른 위치에서도 빠르게 사용할 수 있습니다.
이러한 사고에 대한 좋은 관점과 개요는 여기에서 찾을 수 있습니다. [Network DevOps](https://www.thousandeyes.com/learning/techtorials/network-devops)

## 네트워킹 기본 사항

여기서는 데브옵스 측면은 잠시 접어두고 이제 네트워킹의 기본 사항에 대해 간략히 살펴보겠습니다.

### 네트워크 장치

이 콘텐츠를 동영상으로 보고 싶으시다면 Practical Networking에서 다음 동영상을 확인하세요:

- [Network Devices - Hosts, IP Addresses, Networks - Networking Fundamentals - Lesson 1a](https://www.youtube.com/watch?v=bj-Yfakjllc&list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi&index=1)
- [Network Devices - Hub, Bridge, Switch, Router - Networking Fundamentals - Lesson 1b](https://www.youtube.com/watch?v=H7-NR3Q3BeI&list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi&index=2)

**호스트**는 트래픽을 보내거나 받는 모든 장치입니다.

![](/2022/Days/Images/Day21_Networking1.png)

**IP 주소**는 각 호스트의 신원입니다.

![](/2022/Days/Images/Day21_Networking2.png)

**네트워크**는 호스트 간에 트래픽을 전송하는 역할을 합니다. 네트워크가 없다면 데이터를 수동으로 많이 이동해야 할 것입니다!

비슷한 연결이 필요한 호스트의 논리적 그룹입니다.

![](/2022/Days/Images/Day21_Networking3.png)

**스위치**는 네트워크 내에서 통신을 촉진합니다. 스위치는 호스트 간에 데이터 패킷을 전달합니다. 스위치는 호스트에 직접 패킷을 보냅니다.

- 네트워크: 유사한 연결이 필요한 호스트들의 그룹.
- 네트워크의 호스트는 동일한 IP 주소 공간을 공유합니다.

![](/2022/Days/Images/Day21_Networking4.png)

**라우터**는 네트워크 간의 통신을 용이하게 합니다. 앞서 스위치가 네트워크 내의 통신을 관리한다고 말했듯이 라우터는 이러한 네트워크를 함께 연결하거나 최소한 허용되는 경우 서로에 대한 액세스 권한을 부여할 수 있게 해줍니다.

라우터는 트래픽 제어 지점(보안, 필터링, 리디렉션)을 제공할 수 있습니다. 최근에는 이러한 기능 중 일부를 제공하는 스위치도 점점 더 많아지고 있습니다.

라우터는 자신이 연결된 네트워크를 학습합니다. 이를 라우트라고 하며 라우팅 테이블은 라우터가 알고 있는 모든 네트워크입니다.

라우터는 연결된 네트워크에 IP 주소를 가지고 있습니다. 이 IP는 게이트웨이라고도 하는 각 호스트가 로컬 네트워크에서 나가는 통로가 되기도 합니다.

라우터는 또한 앞서 언급한 네트워크에서 계층 구조를 생성합니다.

![](/2022/Days/Images/Day21_Networking5.png)

## 스위치 대 라우터

**라우팅**은 네트워크 간에 데이터를 이동하는 프로세스입니다.

- 라우터는 라우팅을 주요 목적으로 하는 장치입니다.

**스위칭**은 네트워크 내에서 데이터를 이동하는 프로세스입니다.

- 스위치는 스위칭을 주목적으로 하는 장치입니다.

다음과 같이 다양한 네트워크 장치가 있다는 것을 알고 있으므로 이는 장치에 대한 기본적인 개요입니다:

- 액세스 포인트
- 방화벽
- 로드 밸런서
- 레이어 3 스위치
- IDS/IPS
- 프록시
- 가상 스위치
- 가상 라우터

이러한 모든 장치는 라우팅 및/또는 스위칭을 수행합니다.

앞으로 며칠 동안 이 목록에 대해 조금 더 자세히 알아보도록 하겠습니다.

- OSI 모델
- 네트워크 프로토콜
- DNS(도메인 이름 시스템)
- NAT
- DHCP
- 서브넷

## 자료

- [Networking Fundamentals](https://www.youtube.com/playlist?list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi)
- [Computer Networking full course](https://www.youtube.com/watch?v=IPvYjXCsTg8)

[Day 22](day22.md)에서 봐요!
