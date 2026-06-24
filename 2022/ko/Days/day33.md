---
title: '#90DaysOfDevOps - Microsoft Azure Networking Models + Azure Management - Day 33'
published: false
description: 90DaysOfDevOps - Microsoft Azure Networking Models + Azure Management
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048706
---

## Microsoft Azure 네트워킹 모델 + Azure 관리

오늘이 Microsoft Azure의 12번째 생일을 기념하는 날인 것 같습니다! (2022년 2월 1일) 아무튼, 이번 포스팅에서는 Microsoft Azure 내의 네트워킹 모델과 Azure 관리 옵션에 대해 알아보려고 합니다. 지금까지는 Azure 포털만 사용했지만, 플랫폼 내에서 리소스를 구동하고 생성하는 데 사용할 수 있는 다른 영역에 대해서도 언급했습니다.

## Azure 네트워크 모델

### 가상 네트워크

- 가상 네트워크는 Azure에서 생성된 구조입니다.
- 가상 네트워크에는 하나 이상의 IP 범위가 할당됩니다.
- 가상 네트워크는 지역 내 구독 내에 존재합니다.
- 가상 서브넷은 네트워크 범위를 분할하기 위해 가상 네트워크에 생성됩니다.
- 가상 머신은 가상 서브넷에 배치됩니다.
- 가상 네트워크 내의 모든 가상 머신은 통신할 수 있습니다.
- 가상 네트워크당 65,536개의 프라이빗 IP.
- 리전으로부터의 송신 트래픽에 대해서만 비용을 지불합니다. (리전을 벗어나는 데이터)
- IPv4 및 IPv6 지원.
  - 퍼블릭 대면 및 가상 네트워크 내 IPv6.

Azure 가상 네트워크는 AWS VPC에 비유할 수 있습니다. 하지만 몇 가지 주의해야 할 차이점이 있습니다:

- AWS에서는 기본 VNet이 생성되지만 Microsoft Azure에서는 그렇지 않으므로 요구 사항에 맞게 첫 번째 가상 네트워크를 생성해야 합니다.
- Azure의 모든 가상 머신은 기본적으로 인터넷에 대한 NAT 액세스 권한을 갖습니다. AWS에 따른 NAT 게이트웨이가 없습니다.
- Microsoft Azure에는 사설 또는 공용 서브넷의 개념이 없습니다.
- 공용 IP는 vNIC 또는 로드 밸런서에 할당할 수 있는 리소스입니다.
- 가상 네트워크와 서브넷에는 서브넷 수준 위임을 가능하게 하는 자체 ACL이 있습니다.
- AWS에서는 가용 영역별로 서브넷이 있는 반면, 가상 네트워크에서는 가용 영역 전체에 걸쳐 서브넷이 있습니다.

또한 가상 네트워크 피어링도 있습니다. 이를 통해 테넌트 및 리전 전반의 가상 네트워크를 Azure 백본을 사용하여 연결할 수 있습니다. 전이되지는 않지만, 허브 가상 네트워크의 Azure 방화벽을 통해 활성화할 수 있습니다. 게이트웨이 트랜짓을 사용하면 피어링된 가상 네트워크가 연결된 네트워크에 연결할 수 있으며, 그 예로 ExpressRoute to On-Premises를 들 수 있습니다.

### 액세스 제어

- Azure는 네트워크 보안 그룹을 활용하며, 이는 stateful합니다.
- 규칙을 만든 다음 네트워크 보안 그룹에 할당할 수 있습니다.
- 네트워크 보안 그룹은 서브넷 또는 VM에 적용됩니다.
- 서브넷에 적용되는 경우에도 가상 머신 NIC에서는 여전히 "Edge" 디바이스가 아닌 것으로 적용됩니다.

![](/2022/Days/Images/Day33_Cloud1.png)

- 규칙은 네트워크 보안 그룹에 결합됩니다.
- 우선순위에 따라 유연한 구성이 가능합니다.
- 우선순위 숫자가 낮을수록 우선순위가 높습니다.
- 대부분의 로직은 IP 주소로 구축되지만, 일부 태그와 레이블도 사용할 수 있습니다.

| Description      | Priority | Source Address     | Source Port | Destination Address | Destination Port | Action |
| ---------------- | -------- | ------------------ | ----------- | ------------------- | ---------------- | ------ |
| Inbound 443      | 1005     | \*                 | \*          | \*                  | 443              | Allow  |
| ILB              | 1010     | Azure LoadBalancer | \*          | \*                  | 10000            | Allow  |
| Deny All Inbound | 4000     | \*                 | \*          | \*                  | \*               | DENY   |

또한 애플리케이션 보안 그룹(ASG)도 있습니다.

- NSG가 성장하는 환경에서 유지 관리가 어려울 수 있는 IP 주소 범위에 초점을 맞추는 경우.
- ASG를 사용하면 다양한 애플리케이션 역할(웹서버, DB 서버, WebApp1 등)에 대한 실제 이름(Monikers)을 정의할 수 있습니다.
- 가상 머신 NIC는 하나 이상의 ASG의 멤버가 됩니다.

그런 다음 ASG는 네트워크 보안 그룹의 일부인 규칙에서 통신 흐름을 제어하는 데 사용할 수 있으며 서비스 태그와 같은 NSG 기능을 계속 사용할 수 있습니다.

| Action | Name               | Source     | Destination | Port         |
| ------ | ------------------ | ---------- | ----------- | ------------ |
| Allow  | AllowInternettoWeb | Internet   | WebServers  | 443(HTTPS)   |
| Allow  | AllowWebToApp      | WebServers | AppServers  | 443(HTTPS)   |
| Allow  | AllowAppToDB       | AppServers | DbServers   | 1443 (MSSQL) |
| Deny   | DenyAllinbound     | Any        | Any         | Any          |

### 로드 밸런싱

Microsoft Azure에는 두 가지 로드 밸런싱 솔루션이 있습니다. (퍼스트 파티, Azure 마켓플레이스에서 사용할 수 있는 타사도 있습니다.) 두 솔루션 모두 외부 또는 내부 엔드포인트에서 작동할 수 있습니다.

- Load Balancer (Layer 4)는 해시 기반 배포 및 포트 포워딩을 지원합니다.
- App Gateway (Layer 7)는 SSL offload, 쿠키 기반 세션 선호도, URL 기반 콘텐츠 라우팅과 같은 기능을 지원합니다.

또한 앱 게이트웨이를 사용하면 선택적으로 웹 애플리케이션 방화벽 구성 요소를 사용할 수 있습니다.

## Azure 관리 도구

대부분의 이론 시간을 Azure 포털을 살펴보는 데 할애했는데, 데브옵스 문화를 따르고 이러한 작업, 특히 프로비저닝과 관련된 많은 작업을 처리할 때 API 또는 커맨드라인 도구를 통해 수행한다고 제안하고 싶습니다. Azure 환경의 프로비저닝을 자동화할 때 이를 알아야 하므로 사용할 수 있는 다른 관리 도구 몇 가지를 언급하고 싶었습니다.

### Azure 포털

Microsoft Azure Portal은 커맨드라인 도구의 대안을 제공하는 웹 기반 콘솔입니다. Azure 포털에서 구독을 관리할 수 있습니다. 간단한 웹 앱부터 복잡한 클라우드 배포까지 모든 것을 빌드, 관리 및 모니터링할 수 있습니다. 포털 내에서 찾을 수 있는 또 다른 것은 이러한 이동 경로이며, 앞서 언급했듯이 JSON은 모든 Azure 리소스의 기반이며, 포털에서 시작하여 기능, 서비스 및 기능을 이해한 다음 나중에 자동화된 workflow에 통합하기 위해 그 아래에 있는 JSON을 이해할 수 있습니다.

![](/2022/Days/Images/Day33_Cloud2.png)

Azure Preview Portal도 있으며, 이 포털을 사용하여 새로운 서비스 및 향후 개선 사항을 보고 테스트할 수 있습니다.

![](/2022/Days/Images/Day33_Cloud3.png)

### PowerShell

Azure PowerShell을 살펴보기 전에 PowerShell에 대해 먼저 소개할 필요가 있습니다. PowerShell은 작업 자동화 및 구성 관리 프레임워크, 커맨드라인 셸 및 스크립팅 언어입니다. Linux 섹션에서 셸 스크립팅에 대해 다룬 것과 비슷하다고 감히 말할 수 있습니다. PowerShell은 Windows OS에서 처음 발견되었지만, 이제는 크로스 플랫폼입니다.

Azure PowerShell은 PowerShell 커맨드라인에서 직접 Azure 리소스를 관리하기 위한 cmdlets의 집합입니다.

아래에서 PowerShell 명령 `Connect-AzAccount`를 사용하여 구독에 연결할 수 있음을 확인할 수 있습니다.

![](/2022/Days/Images/Day33_Cloud4.png)

그런 다음 Azure VM과 관련된 특정 명령을 찾으려면 다음 명령을 실행하면 됩니다. 이 PowerShell 프로그래밍 언어에 대해 더 많이 배우고 이해하는 데 몇 시간을 투자할 수 있습니다.

![](/2022/Days/Images/Day33_Cloud5.png)

[여기](https://docs.microsoft.com/en-us/powershell/azure/get-started-azureps?view=azps-7.1.0)에서 PowerShell에서 서비스를 시작하고 프로비저닝하는 방법에 대한 Microsoft의 훌륭한 빠른 시작 가이드를 확인하세요.

### Visual Studio Code

많은 분들께서 보셨듯이 제가 자주 사용하는 IDE는 Visual Studio Code입니다.

Visual Studio Code는 Microsoft에서 Windows, Linux 및 macOS용으로 만든 무료 소스 코드 편집기입니다.

아래에서 Visual Studio Code에 내장된 많은 통합 및 도구를 사용하여 Microsoft Azure 및 그 안의 서비스와 상호 작용할 수 있음을 확인할 수 있습니다.

![](/2022/Days/Images/Day33_Cloud6.png)

### Cloud Shell

Azure Cloud Shell은 Azure 리소스를 관리하기 위한 대화형, 인증된 브라우저 액세스 가능 셸입니다. 작업 방식에 가장 적합한 셸 환경을 선택할 수 있는 유연성을 제공합니다.

![](/2022/Days/Images/Day33_Cloud7.png)

아래에서 볼 수 있듯이 포털 내에서 Cloud Shell을 처음 실행하면 Bash와 PowerShell 중에서 선택할 수 있습니다.

![](/2022/Days/Images/Day33_Cloud8.png)

Cloud Shell을 사용하려면 구독에 약간의 저장 공간을 제공해야 합니다.

Cloud Shell을 사용하도록 선택하면 머신이 스핀업되며, 이러한 머신은 일시적이지만 파일은 디스크 이미지와 마운트된 파일 공유를 통해 두 가지 방식으로 유지됩니다.

![](/2022/Days/Images/Day33_Cloud9.png)

- Cloud Shell은 세션별, 사용자별로 제공되는 임시 호스트에서 실행됩니다.
- 대화형 활동이 없으면 20분 후에 Cloud Shell이 시간 초과됨
- Cloud Shell을 사용하려면 Azure 파일 공유를 마운트해야 함
- Cloud Shell은 Bash와 PowerShell 모두에 동일한 Azure 파일 공유를 사용합니다.
- Cloud Shell은 사용자 계정당 하나의 머신이 할당됨
- Cloud Shell은 파일 공유에 보관된 5GB 이미지를 사용하여 $HOME을 유지합니다.
- 권한은 Bash에서 일반 Linux 사용자로 설정됨

위 내용은 [Cloud Shell 개요](https://docs.microsoft.com/en-us/azure/cloud-shell/overview)에서 복사한 것입니다.

### Azure CLI

마지막으로 Azure CLI에 대해 알아보겠습니다. Azure CLI는 윈도우, 리눅스, 맥OS에서 설치할 수 있습니다. 설치가 완료되면 `az`를 입력한 후 다른 명령어를 입력해 Azure 리소스를 생성, 업데이트, 삭제, 조회할 수 있습니다.

처음 Azure 학습에 들어갔을 때 Azure PowerShell과 Azure CLI가 있어서 약간 혼란스러웠습니다.

이에 대한 커뮤니티의 피드백이 있으면 좋겠습니다. 하지만 제가 알기로 Azure PowerShell은 Windows PowerShell 또는 PowerShell Core에 추가된 모듈(다른 OS에서도 사용 가능하지만 전부는 아님)인 반면, Azure CLI는 Azure에 연결하여 해당 명령을 실행하는 크로스 플랫폼 커맨드라인 프로그램이라고 알고 있습니다.

이 두 옵션 모두 구문은 다르지만, 제가 보기에는 매우 유사한 작업을 수행할 수 있습니다.

예를 들어 PowerShell에서 가상 머신을 만들려면 `New-AzVM` cmdlet을 사용하는 반면, Azure CLI는 `az VM create`를 사용합니다.

앞서 시스템에 Azure PowerShell 모듈이 설치되어 있고 Windows 컴퓨터에서 PowerShell을 통해 호출할 수 있는 Azure CLI도 설치되어 있는 것을 보았습니다.

![](/2022/Days/Images/Day33_Cloud10.png)

이미 언급했듯이 여기서 중요한 점은 올바른 도구를 선택하는 것입니다. Azure는 자동화를 기반으로 실행됩니다. 포털 내에서 수행하는 모든 작업은 리소스를 읽거나, 만들거나, 수정하거나, 삭제하기 위해 실행되는 코드로 변환됩니다.

Azure CLI

- Windows, macOS, Linux에 설치할 수 있는 크로스 플랫폼 커맨드라인 인터페이스
- Windows PowerShell, Cmd, Bash 및 기타 Unix 셸에서 실행됩니다.

Azure PowerShell

- 크로스 플랫폼 PowerShell 모듈, Windows, macOS, Linux에서 실행됨
- Windows PowerShell 또는 PowerShell 필요

사용 중인 환경에서 PowerShell을 사용할 수 없지만 .mdor bash를 사용할 수 있는 경우 Azure CLI를 선택하는 것이 좋습니다.

다음으로 지금까지 살펴본 모든 이론을 바탕으로 몇 가지 시나리오를 만들어 Azure에서 실습해 보겠습니다.

## 자료

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

[Day 34](day34.md)에서 봐요!
