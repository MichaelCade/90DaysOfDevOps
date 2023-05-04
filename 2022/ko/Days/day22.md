---
title: '#90DaysOfDevOps - The OSI Model - The 7 Layers - Day 22'
published: false
description: 90DaysOfDevOps - The OSI Model - The 7 Layers
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049037
---

아래 콘텐츠는 대부분 실용적인 네트워킹의 [Networking Fundamentals series](https://www.youtube.com/playlist?list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi)에서 가져온 것입니다. 이 콘텐츠를 동영상으로 보고 싶으시다면 다음 두 동영상을 확인해 보세요:

- [The OSI Model: A Practical Perspective - Layers 1 / 2 / 3](https://www.youtube.com/watch?v=LkolbURrtTs&list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi&index=3)
- [The OSI Model: A Practical Perspective - Layers 4 / 5+](https://www.youtube.com/watch?v=0aGqGKrRE0g&list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi&index=4)

## OSI 모델 - 7개의 Layers

산업으로서 네트워킹의 전반적인 목적은 두 호스트가 데이터를 공유할 수 있도록 하는 것입니다. 네트워킹 이전에는 이 호스트에서 이 호스트로 데이터를 가져오려면 이 호스트에 무언가를 꽂고 다른 호스트에 가서 다른 호스트에 꽂아야 했습니다.

네트워킹을 사용하면 호스트가 유선을 통해 데이터를 자동으로 공유할 수 있으므로 이 작업을 자동화할 수 있으며, 이 작업을 수행하려면 호스트가 일련의 규칙을 따라야 합니다.

이는 모든 언어와 다르지 않습니다. 영어에는 두 명의 영어 사용자가 따라야 하는 일련의 규칙이 있습니다. 스페인어에도 고유한 규칙이 있습니다. 프랑스어에는 고유한 규칙이 있으며, 네트워킹에도 고유한 규칙이 있습니다.

네트워킹에 대한 규칙은 7가지 계층으로 나뉘며, 이러한 계층을 OSI 모델이라고 합니다.

### OSI 모델 소개

OSI 모델(개방형 시스템 상호 연결 모델)은 네트워킹 시스템의 기능을 설명하는 데 사용되는 프레임워크입니다. OSI 모델은 서로 다른 제품과 소프트웨어 간의 상호 운용성을 지원하기 위해 컴퓨팅 기능을 보편적인 규칙과 요구 사항으로 특성화합니다. OSI 참조 모델에서 컴퓨팅 시스템 간의 통신은 7가지 추상화 계층으로 나뉩니다: **물리적, 데이터 링크, 네트워크, 전송, 세션, 프레젠테이션 및 애플리케이션**.

![](/2022/Days/Images/Day22_Networking1.png)

### 물리적

OSI 모델에서 Layer 1은 물리적이라고 하며, 물리적 케이블 등 수단을 통해 한 호스트에서 다른 호스트로 데이터를 전송할 수 있다는 전제하에 이 Layer에서도 Wi-Fi를 고려할 수 있습니다. 또한 한 호스트에서 다른 호스트로 데이터를 전송하기 위해 허브와 중계기 주변에서 더 많은 레거시 하드웨어를 볼 수 있습니다.

![](/2022/Days/Images/Day22_Networking2.png)

### 데이터 링크

Layer 2의 데이터 링크는 데이터가 프레임으로 패키지화되는 노드 간 전송을 가능하게 합니다. 또한 물리 계층에서 발생했을 수 있는 오류를 수정하는 수준도 있습니다. 또한 이 단계에서 MAC 주소가 도입되거나 처음 등장합니다.

네트워킹 첫날인 [Day 21](day21.md)
에서 다룬 스위치에 대한 첫 번째 언급도 여기에서 볼 수 있습니다.

![](/2022/Days/Images/Day22_Networking3.png)

### 네트워크

Layer 3 스위치 또는 Layer 2 스위치라는 용어를 들어보셨을 것입니다. OSI 모델인 Layer 3에서 네트워크는 End to End 전송을 목표로 하며, 첫날 개요에서 언급한 IP 주소가 있는 곳입니다.

라우터와 호스트는 Layer 3에 존재하며, 라우터는 여러 네트워크 간에 라우팅하는 기능이라는 점을 기억하세요. IP가 있는 모든 것은 Layer 3으로 간주할 수 있습니다.

![](/2022/Days/Images/Day22_Networking4.png)

그렇다면 왜 Layer 2와 3 모두에 주소 체계가 필요할까요? (MAC 주소와 IP 주소)

한 호스트에서 다른 호스트로 데이터를 전송하는 것을 생각해 보면, 각 호스트에는 IP 주소가 있지만 그사이에는 여러 개의 스위치와 라우터가 있습니다. 각 장치에는 Layer 2 MAC 주소가 있습니다.

Layer 2 MAC 주소는 호스트에서 스위치/라우터로만 이동하며 hop에 중점을 두는 반면, Layer 3 IP 주소는 데이터 패킷이 최종 호스트에 도달할 때까지 해당 패킷을 유지합니다. (End to End)

IP 주소 - Layer 3 = End to End 전송

MAC 주소 - Layer 2 = Hop to Hop 전송

이제 Layer3와 Layer2 주소를 연결하는 ARP(주소 확인 프로토콜)라는 네트워크 프로토콜이 있지만 오늘은 다루지 않겠습니다.

### 전송

서비스 간 전송, Layer 4는 데이터 스트림을 구분하기 위해 존재합니다. Layer 3과 Layer 2에 주소 체계가 있는 것과 마찬가지로 Layer 4에는 포트가 있습니다.

![](/2022/Days/Images/Day22_Networking5.png)

### 세션, 프레젠테이션, 애플리케이션

Layers 5,6,7의 구분이 다소 모호해졌습니다.

보다 최근의 이해를 돕기 위해 [TCP IP Model](https://www.geeksforgeeks.org/tcp-ip-model/)을 살펴보는 것이 좋습니다.

이제 이 네트워킹 스택을 사용하여 호스트가 서로 통신할 때 어떤 일이 일어나는지 설명해 보겠습니다. 이 호스트에는 다른 호스트로 전송할 데이터를 생성하는 애플리케이션이 있습니다.

소스 호스트는 캡슐화 프로세스라고 하는 과정을 거치게 됩니다. 이 데이터는 먼저 Layer 4로 전송됩니다.

Layer 4는 해당 데이터에 헤더를 추가하여 서비스 간 전송이라는 Layer 4의 목표를 용이하게 할 수 있습니다. 이 헤더는 TCP 또는 UDP를 사용하는 포트가 될 것입니다. 또한 소스 포트와 대상 포트도 포함됩니다.

이를 세그먼트(데이터 및 포트)라고도 합니다.

이 세그먼트는 OSI 스택을 통해 네트워크 계층인 Layer 3으로 전달되고, 네트워크 계층은 이 데이터에 또 다른 헤더를 추가합니다.
이 헤더는 End to End 전송이라는 Layer 3의 목표를 달성하기 위한 것으로, 이 헤더에는 소스 IP 주소와 대상 IP가 있으며, 헤더와 데이터를 합쳐 패킷이라고도 합니다.

그러면 Layer 3이 해당 패킷을 받아 Layer 2로 전달하고, Layer 2는 다시 한번 해당 데이터에 또 다른 헤더를 추가하여 Layer 2의 목표인 Hop to Hop 전송을 달성하게 되는데, 이 헤더에는 소스 및 대상 맥 주소가 포함됩니다.
Layer 2 헤더와 데이터가 있을 때 이를 프레임이라고 합니다.

그런 다음 이 프레임은 1과 0으로 변환되어 Layer 1 물리적 케이블 또는 Wi-Fi를 통해 전송됩니다.

![](/2022/Days/Images/Day22_Networking6.png)

위에서 헤더와 데이터의 각 Layer에 대한 네이밍에 대해 언급했지만, 이것도 그려보기로 했습니다.

![](/2022/Days/Images/Day22_Networking7.png)

데이터를 전송하는 애플리케이션은 어딘가로 데이터를 전송하고 있으므로 수신은 스택을 거슬러 올라가 수신 호스트에 전달하기 위해 다소 역방향으로 이루어집니다.

![](/2022/Days/Images/Day22_Networking8.png)

## 자료

- [Networking Fundamentals](https://www.youtube.com/playlist?list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi)

* [Computer Networking full course](https://www.youtube.com/watch?v=IPvYjXCsTg8)

[Day 23](day23.md)에서 봐요!
