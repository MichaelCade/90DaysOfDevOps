---
title: '#90DaysOfDevOps - Network Protocols - Day 23'
published: false
description: 90DaysOfDevOps - Network Protocols
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048704
---

아래 콘텐츠는 대부분 실용적인 네트워킹의 [Networking Fundamentals series](https://www.youtube.com/playlist?list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi)에서 가져온 것입니다. 이 콘텐츠를 동영상으로 보고 싶으시다면 이 동영상을 확인해보세요:

- [Network Protocols - ARP, FTP, SMTP, HTTP, SSL, TLS, HTTPS, DNS, DHCP](https://www.youtube.com/watch?v=E5bSumTAHZE&list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi&index=12)

## 네트워크 프로토콜

인터넷 표준을 구성하는 일련의 규칙과 메시지입니다.

- ARP - 주소 확인 프로토콜

ARP에 대해 자세히 알고 싶으시다면 여기에서 인터넷 표준을 읽어보세요. [RFC 826](https://datatracker.ietf.org/doc/html/rfc826)

Layer 2 네트워크에서 IP 주소를 고정된 물리적 컴퓨터 주소(MAC 주소라고도 함)에 연결합니다.

![](/2022/Days/Images/Day23_Networking1.png)

- FTP - 파일 전송 프로토콜

소스에서 대상으로 파일을 전송할 수 있습니다. 일반적으로 이 프로세스는 인증되지만, 익명 액세스를 사용하도록 구성한 경우 사용할 수 있습니다. 이제 더 나은 보안을 위해 클라이언트에서 FTP 서버로 SSL/TLS 연결을 제공하는 FTPS를 더 자주 보게 될 것입니다. 이 프로토콜은 OSI 모델의 애플리케이션 계층에서 찾을 수 있습니다.

![](/2022/Days/Images/Day23_Networking2.png)

- SMTP - 단순 메일 전송 프로토콜

전자 메일 전송에 사용되는 메일 서버는 SMTP를 사용하여 메일 메시지를 보내고 받습니다. Microsoft 365를 사용하더라도 SMTP 프로토콜이 동일한 용도로 사용된다는 것을 알 수 있습니다.

![](/2022/Days/Images/Day23_Networking3.png)

- HTTP - 하이퍼 텍스트 전송 프로토콜

HTTP는 인터넷과 콘텐츠 탐색의 기초입니다. 우리가 즐겨 찾는 웹사이트에 쉽게 액세스할 수 있게 해줍니다. HTTP는 여전히 많이 사용되고 있지만, 대부분의 즐겨 찾는 사이트에서는 HTTPS가 더 많이 사용되고 있거나 사용되어야 합니다.

![](/2022/Days/Images/Day23_Networking4.png)

- SSL - 보안 소켓 계층 | TLS - 전송 계층 보안

SSL의 뒤를 이어 등장한 TLS는 네트워크를 통해 안전한 통신을 제공하는 **암호화 프로토콜**입니다. 메일, 인스턴트 메시징 및 기타 애플리케이션에서 찾을 수 있지만 가장 일반적으로 HTTPS를 보호하는 데 사용됩니다.

![](/2022/Days/Images/Day23_Networking5.png)

- HTTPS - SSL/TLS로 보호되는 HTTP

네트워크를 통한 보안 통신에 사용되는 HTTP의 확장인 HTTPS는 위에서 언급한 것처럼 TLS로 암호화됩니다. 호스트 간에 데이터가 교환되는 동안 인증, 개인 정보 보호 및 무결성을 제공하는 데 중점을 두었습니다.

![](/2022/Days/Images/Day23_Networking6.png)

- DNS - 도메인 이름 시스템

DNS는 인간 친화적인 도메인 이름을 매핑하는 데 사용됩니다. 예를 들어 우리가 모두 알고 있는 [google.com](https://google.com)은 브라우저를 열고 [8.8.8.8](https://8.8.8.8)을 입력하면 우리가 거의 알고 있는 Google이 표시됩니다. 하지만 모든 웹사이트의 IP 주소를 모두 기억하기란 쉽지 않으며, 일부 웹사이트는 정보를 찾기 위해 구글을 사용하기도 합니다.

호스트, 서비스 및 기타 리소스에 연결할 수 있도록 하는 DNS의 역할이 바로 여기에 있습니다.

모든 호스트에서 인터넷 연결이 필요한 경우 해당 도메인 이름을 확인할 수 있도록 DNS가 있어야 합니다. DNS는 며칠 또는 몇 년을 투자하여 학습할 수 있는 영역입니다. 또한 경험상 네트워킹과 관련된 모든 오류의 일반적인 원인은 대부분 DNS라고 말하고 싶습니다. 하지만 네트워크 엔지니어가 동의할지는 모르겠습니다.

![](/2022/Days/Images/Day23_Networking7.png)

- DHCP - 동적 호스트 구성 프로토콜

인터넷에 액세스하거나 서로 간에 파일을 전송하는 등 호스트가 작동하는 데 필요한 프로토콜에 대해 많이 논의했습니다.

이 두 가지 작업을 모두 수행하기 위해 모든 호스트에 필요한 4가지 사항이 있습니다.

- IP 주소
- 서브넷 마스크
- 기본 게이트웨이
- DNS

IP 주소는 호스트가 위치한 네트워크에서 호스트의 고유 주소로, 집 번호라고 생각하시면 됩니다.

서브넷 마스크는 곧 다루겠지만 우편번호나 우편번호로 생각하시면 됩니다.

기본 게이트웨이는 일반적으로 네트워크에서 Layer 3 연결을 제공하는 라우터의 IP입니다. 이를 우리 동네를 빠져나갈 수 있는 하나의 도로라고 생각할 수 있습니다.

그리고 방금 설명한 DNS는 복잡한 공인 IP 주소를 보다 적합하고 기억하기 쉬운 도메인 이름으로 변환하는 데 도움을 줍니다. 어쩌면 이것을 올바른 게시물을 얻기 위한 거대한 분류 사무실이라고 생각할 수 있습니다.

각 호스트마다 이 네 가지가 필요하다고 말씀드렸듯이 호스트가 1,000개 또는 10,000개라면 이 네 가지를 개별적으로 결정하는 데 매우 오랜 시간이 걸릴 것입니다. 이때 DHCP가 등장하여 네트워크의 범위를 결정하면 이 프로토콜이 네트워크에서 사용 가능한 모든 호스트에 배포됩니다.

또 다른 예로 커피숍에 가서 커피를 마시고 노트북이나 휴대폰을 들고 앉았다고 가정해 봅시다. 호스트를 커피숍 와이파이에 연결하면 인터넷에 접속할 수 있고, 메시지와 메일이 핑을 받기 시작하며, 웹 페이지와 소셜 미디어를 탐색할 수 있습니다. 커피숍 와이파이에 연결했을 때 컴퓨터는 전용 DHCP 서버 또는 DHCP를 처리하는 라우터에서 DHCP 주소를 선택했을 것입니다.

![](/2022/Days/Images/Day23_Networking8.png)

### 서브넷

서브넷은 IP 네트워크의 논리적 세분화입니다.

서브넷은 대규모 네트워크를 더 효율적으로 실행할 수 있는 더 작고 관리하기 쉬운 네트워크로 나눕니다.

각 서브넷은 더 큰 네트워크의 논리적 세분화입니다. 충분한 서브넷을 가진 연결된 장치는 공통 IP 주소 식별자를 공유하여 서로 통신할 수 있습니다.

라우터는 서브넷 간의 통신을 관리합니다.

서브넷의 크기는 연결 요구 사항과 사용되는 네트워크 기술에 따라 달라집니다.

조직은 사용 가능한 주소 공간의 한도 내에서 서브넷의 수와 크기를 결정할 책임이 있습니다.
사용 가능한 주소 공간의 한도 내에서 서브넷의 수와 크기를 결정할 책임이 있으며, 세부 사항은 해당 조직에 로컬로 유지됩니다. 서브넷은 지점 간 링크 또는 몇 개의 디바이스를 지원하는 서브네트워크와 같이 더 작은 서브넷으로 세분화할 수도 있습니다.

무엇보다도 대규모 네트워크를 서브넷으로 세분화하면
네트워크를 서브넷으로 분할하면 IP 주소
재할당을 가능하게 하고 네트워크 혼잡, 간소화, 네트워크 통신 및 효율성을 완화합니다.

서브넷은 네트워크 보안도 향상시킬 수 있습니다.
네트워크의 한 섹션이 손상되면 해당 섹션을 격리하여 악의적인 공격자가 더 큰 네트워크에서 이동하기 어렵게 만들 수 있습니다.

![](/2022/Days/Images/Day23_Networking9.png)

## 자료

- [Networking Fundamentals](https://www.youtube.com/playlist?list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi)
- [Subnetting Mastery](https://www.youtube.com/playlist?list=PLIFyRwBY_4bQUE4IB5c4VPRyDoLgOdExE)
- [Computer Networking full course](https://www.youtube.com/watch?v=IPvYjXCsTg8)

[Day 24](day24.md)에서 봐요!
