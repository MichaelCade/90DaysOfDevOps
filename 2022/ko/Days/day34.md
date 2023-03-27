---
title: '#90DaysOfDevOps - Microsoft Azure Hands-On Scenarios - Day 34'
published: false
description: 90DaysOfDevOps - Microsoft Azure Hands-On Scenarios
tags: 'DevOps, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048763
---

## Microsoft Azure 실습 시나리오

지난 6일 동안 Microsoft Azure와 퍼블릭 클라우드 전반에 대해 집중적으로 살펴봤는데요, Azure의 빌딩 블록을 이해하기 위해서는 많은 이론이 필요했지만, 다른 주요 클라우드 제공업체에도 잘 적용될 수 있을 것입니다.

처음에 퍼블릭 클라우드에 대한 기초 지식을 얻고 최소한 한 공급자를 선택하는 것에 대해 언급했는데, 여러 클라우드를 얕게 다룰 경우 길을 잃기 쉽지만, 한 공급자를 선택한 후 기본 사항을 이해하면 다른 클라우드로 이동하여 학습을 가속화하는 것이 매우 쉽다고 생각됩니다.

이번 마지막 세션에서는 Microsoft에서 만든 참고 자료로 [AZ-104 Microsoft Azure Administrator](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/) 시험 준비에 사용되는 이 페이지에서 실습 시나리오를 골라보겠습니다.

여기에는 컨테이너와 쿠버네티스와 같이 아직 자세히 다루지 않은 내용도 있으므로 아직은 여기에 뛰어들지 않겠습니다.

이전 포스팅에서 Module 1,2,3의 대부분을 만들었습니다.

### 가상 네트워킹

[Module 04](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/Instructions/Labs/LAB_04-Implement_Virtual_Networking.html) 따라 하기:

위의 내용을 살펴보고 #90DaysOfDevOps의 이름을 몇 가지 변경했습니다. 또한 클라우드 셸을 사용하는 대신 Windows 머신에서 Azure CLI를 사용하여 전날 생성한 새 사용자로 로그인했습니다.

브라우저를 열고 계정에 인증할 수 있는 `az login`을 사용하여 이 작업을 수행할 수 있습니다.

그런 다음 아래 작업 중 일부를 빌드하는 데 사용할 PowerShell 스크립트와 Module에서 몇 가지 참조를 만들었습니다. 관련 파일은 이 폴더에서 찾을 수 있습니다.
(/2022/Days/Cloud/01VirtualNetworking)

스크립트의 파일 위치를 사용자 환경에 맞게 변경해야 합니다.

첫 번째 단계에서는 환경에 가상 네트워크나 가상 머신이 생성되어 있지 않고 리소스 그룹에 클라우드 셸 스토리지 위치만 구성되어 있습니다.

먼저 [PowerShell 스크립트](/2022/Days/Cloud/01VirtualNetworking/Module4_90DaysOfDevOps.ps1)를 실행합니다.

![](/2022/Days/Images/Day34_Cloud1.png)

- Task 1: 가상 네트워크 생성 및 구성

![](/2022/Days/Images/Day34_Cloud2.png)

- Task 2: 가상 네트워크에 가상 머신 배포

![](/2022/Days/Images/Day34_Cloud3.png)

- Task 3: Azure VM의 프라이빗 및 퍼블릭 IP 주소 구성

![](/2022/Days/Images/Day34_Cloud4.png)

- Task 4: 네트워크 보안 그룹 구성

![](/2022/Days/Images/Day34_Cloud5.png)
![](/2022/Days/Images/Day34_Cloud6.png)

- Task 5: 내부 이름 확인을 위한 Azure DNS 구성

![](/2022/Days/Images/Day34_Cloud7.png)
![](/2022/Days/Images/Day34_Cloud8.png)

### 네트워크 트래픽 관리

[Module 06](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/Instructions/Labs/LAB_06-Implement_Network_Traffic_Management.html)에 이어서

다음 단계에서는 지난번 리소스 그룹으로 이동하여 리소스를 삭제했는데, 저처럼 사용자 계정을 해당 리소스 그룹에만 접근하도록 설정하지 않았다면 Module 이름을 '90Days\*'로 변경하여 모든 리소스와 리소스 그룹을 삭제할 수 있습니다. 이것이 다음 실습 각각에 대한 제 프로세스입니다.

이 실습에서는 아래 작업 중 일부를 구축하는 데 사용할 Module의 일부 참조와 PowerShell 스크립트도 만들었습니다. 관련 파일은 이 폴더에서 찾을 수 있습니다.
(/2022/Days/Cloud/02TrafficManagement)

- Task 1: 실습 환경 프로비저닝

먼저 [PowerShell 스크립트](/2022/Days/Cloud/02TrafficManagement/Mod06_90DaysOfDevOps.ps1)를 실행합니다.

![](/2022/Days/Images/Day34_Cloud9.png)

- Task 2: 허브 및 스포크 네트워크 Topology 구성

![](/2022/Days/Images/Day34_Cloud10.png)

- Task 3: 가상 네트워크 피어링의 전이성 테스트

이 작업의 경우 90DaysOfDevOps 그룹은 권한 때문에 네트워크 감시자에 액세스할 수 없었는데, 이는 네트워크 감시자가 리소스 그룹에 묶여 있지 않은 리소스 중 하나이기 때문인 것으로 예상됩니다(이 사용자에 대해 RBAC이 적용된 곳). 미국 동부 네트워크 감시자 기여자 역할을 90DaysOfDevOps 그룹에 추가했습니다.

![](/2022/Days/Images/Day34_Cloud11.png)
![](/2022/Days/Images/Day34_Cloud12.png)
![](/2022/Days/Images/Day34_Cloud13.png)

두 개의 스포크 가상 네트워크가 서로 피어링되지 않기 때문에 예상되는 현상입니다(가상 네트워크 피어링은 전이적이지 않음).

- Task 4: 허브 및 스포크 Topology에서 라우팅 구성하기

여기서 제 계정이 90DaysOfDevOps 그룹 내에서 제 사용자로 스크립트를 실행할 수 없는 또 다른 문제가 있었는데, 잘 모르겠으므로 다시 기본 관리자 계정으로 이동했습니다. 90DaysOfDevOps 그룹은 90DaysOfDevOps 리소스 그룹의 모든 소유자이므로 VM 내에서 명령을 실행할 수 없는 이유를 알고 싶습니다.

![](/2022/Days/Images/Day34_Cloud14.png)
![](/2022/Days/Images/Day34_Cloud15.png)

그런 다음 michael.cade@90DaysOfDevOps.com 계정으로 돌아가서 이 섹션을 계속할 수 있었습니다. 여기서 동일한 테스트를 다시 실행하고 있지만 이제 결과에 도달할 수 있습니다.

![](/2022/Days/Images/Day34_Cloud16.png)

- Task 5: Azure 로드밸런싱 장치 구현하기

![](/2022/Days/Images/Day34_Cloud17.png)
![](/2022/Days/Images/Day34_Cloud18.png)

- Task 6: Azure 애플리케이션 게이트웨이 구현

![](/2022/Days/Images/Day34_Cloud19.png)
![](/2022/Days/Images/Day34_Cloud20.png)

### Azure 스토리지

[Module 07](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/Instructions/Labs/LAB_07-Manage_Azure_Storage.html)에 이어서 진행합니다:

이 실습에서는 아래 작업 중 일부를 빌드하는 데 사용할 PowerShell 스크립트와 Module의 일부 참조도 만들었습니다. 관련 파일은 이 폴더에서 찾을 수 있습니다.
(/2022/Days/Cloud/03Storage)

- Task 1: 실습 환경 프로비저닝

먼저 [PowerShell 스크립트](/2022/Days/Cloud/03Storage/Mod07_90DaysOfDeveOps.ps1)를 실행합니다.

![](/2022/Days/Images/Day34_Cloud21.png)

- Task 2: Azure Storage 계정 만들기 및 구성

![](/2022/Days/Images/Day34_Cloud22.png)

- Task 3: 블롭 스토리지 관리

![](/2022/Days/Images/Day34_Cloud23.png)

- Task 4: Azure 스토리지에 대한 인증 및 권한 부여 관리

![](/2022/Days/Images/Day34_Cloud24.png)
![](/2022/Days/Images/Day34_Cloud25.png)

이 기능이 허용되기를 기다리는 동안 조금 조바심이 났지만 결국 작동했습니다.

![](/2022/Days/Images/Day34_Cloud26.png)

- Task 5: Azure 파일 공유 만들기 및 구성하기

실행 명령에서 michael.cade@90DaysOfDevOps.com 에서는 작동하지 않아서 승격된 계정을 사용했습니다.

![](/2022/Days/Images/Day34_Cloud27.png)
![](/2022/Days/Images/Day34_Cloud28.png)
![](/2022/Days/Images/Day34_Cloud29.png)

- Task 6: Azure 스토리지에 대한 네트워크 액세스 관리

![](/2022/Days/Images/Day34_Cloud30.png)

### 서버리스(웹 앱 구현)

[Module 09a](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/Instructions/Labs/LAB_09a-Implement_Web_Apps.html)를 따릅니다:

- Task 1: Azure 웹 앱 만들기

![](/2022/Days/Images/Day34_Cloud31.png)

- Task 2: 스테이징 배포 슬롯 만들기

![](/2022/Days/Images/Day34_Cloud34.png)

- Task 3: 웹 앱 배포 설정 구성

![](/2022/Days/Images/Day34_Cloud33.png)

- Task 4: 스테이징 배포 슬롯에 코드 배포

![](/2022/Days/Images/Day34_Cloud32.png)

- Task 5: 스테이징 슬롯 교체

![](/2022/Days/Images/Day34_Cloud35.png)

- Task 6: Azure 웹 앱의 자동 확장 구성 및 테스트

제가 사용하는 이 스크립트는 (/2022/Days/Cloud/04Serverless)에서 찾을 수 있습니다.

![](/2022/Days/Images/Day34_Cloud36.png)

이것으로 Microsoft Azure와 퍼블릭 클라우드 전반에 대한 섹션을 마무리합니다. 이 시나리오를 작업하는 것이 매우 즐거웠다고 말씀드리고 싶습니다.

## 자료

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

다음으로 버전 관리 시스템, 특히 git과 코드 저장소 개요에 대해 살펴볼 것이며, 제가 선호하는 옵션인 GitHub를 선택하겠습니다.

[Day 35](day35.md)에서 봐요!
