---
title: '#90DaysOfDevOps - Microsoft Azure Fundamentals - Day 29'
published: false
description: 90DaysOfDevOps - Microsoft Azure Fundamentals
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048705
---

## Microsoft Azure 기초

시작하기 전에 트위터 투표의 승자는 Microsoft Azure였으며, 따라서 페이지의 제목도 그렇게 정해졌습니다. 박빙의 승부가 펼쳐졌고 24시간 동안의 결과가 꽤 흥미로웠습니다.

![](/2022/Days/Images/Day29_Cloud1.png)

이 주제를 다루면서 Microsoft Azure에서 사용할 수 있는 서비스에 대해 더 잘 이해하고 업데이트할 수 있다는 측면에서 저는 오늘 하루는 Amazon AWS에 기대고 있습니다. 하지만 세 가지 주요 클라우드 제공업체 모두에 대해 준비해 둔 리소스를 남겨 두었습니다.

설문조사에는 이 세 가지 서비스만 포함되었고 특히 오라클 클라우드에 대한 의견도 몇 개 있었습니다. 실제 사용 중인 다른 클라우드 제공업체에 대한 의견을 더 듣고 싶습니다.

### 기본 사항

- 퍼블릭 클라우드 서비스 제공
- 지리적으로 분산되어 있음(전 세계 60개 이상의 지역)
- 인터넷 및/또는 개인 연결을 통해 액세스
- 멀티 테넌트 모델
- 사용량 기반 청구 - (종량제 및 사용량 증가에 따른 과금)
- 다양한 요구 사항에 맞는 다양한 서비스 유형 및 제품.

- [Microsoft Azure 글로벌 인프라](https://infrastructuremap.microsoft.com/explore)

SaaS 및 하이브리드 클라우드에 대해 많이 이야기했지만 여기서는 해당 주제를 다루지 않을 계획입니다.

시작하고 따라 하는 가장 좋은 방법은 링크를 클릭하여 [Microsoft Azure 무료 계정](https://azure.microsoft.com/en-gb/free/)을 생성하는 것입니다.

### 지역

위의 인터랙티브 지도를 링크했는데, 아래 이미지에서 전 세계 Microsoft Azure 플랫폼에서 제공되는 다양한 지역을 확인할 수 있습니다.

![](/2022/Days/Images/Day29_Cloud2.png)
[Microsoft 문서](https://docs.microsoft.com/en-us/azure/networking/microsoft-global-network)에서 가져온 이미지 - _01/05/2021_

또한 다른 지역과 연결되지 않거나 다른 지역과 대화할 수 없음을 의미하는 여러 "sovereign" 클라우드를 볼 수 있는데, 예를 들어 `AzureUSGovernment`, `AzureChinaCloud` 등과 같은 정부와 관련된 클라우드가 있습니다.

Microsoft Azure 내에서 서비스를 배포할 때 거의 모든 것에 대해 지역을 선택합니다. 하지만 모든 지역에서 모든 서비스를 이용할 수 있는 것은 아니라는 점에 유의해야 합니다. 이 글을 쓰는 시점에 [지역별 사용 가능한 제품](https://azure.microsoft.com/en-us/global-infrastructure/services/?products=all)을 보면 미국 중서부 지역에서는 Azure Databricks을 사용할 수 없다는 것을 알 수 있습니다.

또한 위에서 "거의 모든 것"이라고 언급했는데, Azure Bot Services, Bing Speech, Azure Virtual Desktop, Static Web Apps 등과 같이 지역과 연결된 특정 서비스가 있습니다.

이면에서는 region이 둘 이상의 데이터 센터로 구성될 수 있습니다. 이를 가용 영역이라고 합니다.

아래 이미지에서 볼 수 있듯이 이 이미지는 Microsoft 공식 문서에서 가져온 것으로, region이란 무엇이며 region이 가용 영역으로 어떻게 구성되는지 설명합니다. 그러나 모든 지역에 가용 영역이 여러 개 있는 것은 아닙니다.

![](/2022/Days/Images/Day29_Cloud3.png)

Microsoft 설명서는 매우 훌륭하며 [여기](https://docs.microsoft.com/en-us/azure/availability-zones/az-overview)에서 지역 및 가용 영역에 대한 자세한 내용을 읽을 수 있습니다.

### 구독

Microsoft Azure는 모든 주요 클라우드 공급자가 이 모델을 따르는 소비 모델 클라우드라고 언급했음을 기억하세요.

엔터프라이즈인 경우 회사에서 이러한 Azure 서비스를 사용할 수 있도록 Microsoft와 엔터프라이즈 계약을 원하거나 이미 설정했을 수 있습니다.

저와 같이 교육용으로 Microsoft Azure를 사용하는 경우 몇 가지 다른 옵션이 있습니다.

일반적으로 일정 기간 동안 Azure에서 사용할 수 있는 몇 가지 무료 클라우드 크레딧을 제공하는 [Microsoft Azure 무료 계정](https://azure.microsoft.com/en-gb/free/)이 있습니다.

또한 Visual Studio 연간 구독과 함께 매달 몇 가지 무료 크레딧을 제공하는 [Visual Studio](https://azure.microsoft.com/en-us/pricing/member-offers/credit-for-visual-studio-subscribers/) 구독을 사용할 수도 있습니다(몇 년 전에는 일반적으로 MSDN으로 알려짐).

마지막으로 신용카드를 건네고 [종량제](https://azure.microsoft.com/en-us/pricing/purchase-options/pay-as-you-go/) 모델을 사용하는 방법이 있습니다.

구독은 잠재적으로 비용 중심이지만 완전히 다른 환경을 가진 서로 다른 구독 간의 경계로 볼 수 있습니다. 구독은 리소스가 생성되는 곳입니다.

### 관리 그룹

관리 그룹을 사용하면 Azure AD(Active Directory) 또는 테넌트 환경 전반에서 제어를 분리할 수 있습니다. 관리 그룹을 통해 정책, RBAC(역할 기반 액세스 제어) 및 예산을 제어할 수 있습니다.

구독은 이러한 관리 그룹에 속하므로 Azure AD 테넌트에 많은 구독을 보유할 수 있으며, 이러한 구독은 정책, RBAC 및 예산도 제어할 수 있습니다.

### 리소스 관리자 및 리소스 그룹

#### Azure 리소스 관리자

- 리소스 공급자를 기반으로 구축된 JSON 기반 API입니다.
- 리소스는 리소스 그룹에 속하며 공통 수명 주기를 공유합니다.
- 병렬 처리
- JSON 기반 배포는 선언적이고 비독립적이며 리소스 간의 종속성을 이해하여 생성 및 순서를 관리합니다.

#### 리소스 그룹

- 모든 Azure Resource Manager 리소스는 하나의 리소스 그룹에만 존재합니다!
- 리소스 그룹은 지역 외부의 리소스를 포함할 수 있는 지역에서 만들어집니다.
- 리소스 그룹 간에 리소스 이동 가능
- 리소스 그룹은 다른 리소스 그룹과 차단되지 않으며, 리소스 그룹 간에 통신할 수 있습니다.
- 리소스 그룹은 정책, RBAC 및 예산도 제어할 수 있습니다.

### 실습

이제 연결하여 **Subscription**을 사용할 수 있는지 확인해 보겠습니다. 기본 제공되는 간단한 **Management Group**을 확인한 다음, 원하는 **Region**에 새로운 전용 **Resource Group**을 만들 수 있습니다.

[Azure 포털](https://portal.azure.com/#home)에 처음 로그인하면 상단에 리소스, 서비스 및 문서를 검색할 수 있는 기능이 표시됩니다.

![](/2022/Days/Images/Day29_Cloud4.png)

먼저 구독을 살펴보겠습니다. 여기에서는 매달 무료 크레딧을 제공하는 Visual Studio Professional 구독을 사용하고 있음을 알 수 있습니다.

![](/2022/Days/Images/Day29_Cloud5.png)

더 넓은 화면으로 이동하여 구독으로 어떤 일이 일어나고 있는지 또는 무엇을 할 수 있는지 살펴보면 왼쪽에 IAM 액세스 제어를 정의할 수 있는 제어 기능과 함께 청구 정보를 볼 수 있으며 더 아래로 내려가면 더 많은 리소스를 사용할 수 있습니다.

![](/2022/Days/Images/Day29_Cloud6.png)

여러 개의 구독이 있고 이를 모두 하나의 구독으로 관리하려는 시나리오가 있을 수 있는데, 이때 관리 그룹을 사용하여 책임 그룹을 분리할 수 있습니다. 아래 제 경우에는 구독이 있는 테넌트 루트 그룹만 있는 것을 볼 수 있습니다.

또한 이전 이미지에서 상위 관리 그룹이 테넌트 루트 그룹에 사용된 아이디와 동일한 것을 볼 수 있습니다.

![](/2022/Days/Images/Day29_Cloud7.png)

다음으로 리소스 그룹이 있는데, 여기에서 리소스를 결합하여 한 곳에서 쉽게 관리할 수 있습니다. 저는 다른 여러 프로젝트를 위해 몇 개를 만들었습니다.

![](/2022/Days/Images/Day29_Cloud8.png)

앞으로 며칠 동안 할 작업을 위해 리소스 그룹을 만들고 싶습니다. 이 콘솔에서 이전 이미지의 생성 옵션을 누르면 쉽게 할 수 있습니다.

![](/2022/Days/Images/Day29_Cloud9.png)

유효성 검사 단계가 수행된 후 생성한 내용을 검토한 다음 생성할 수 있습니다. 또한 하단에 "Download a template for automation"가 표시되는데, 이를 통해 JSON 형식을 가져와서 나중에 원할 경우, 이 간단한 작업을 자동화 방식으로 수행할 수 있으며, 이에 대해서도 나중에 다룰 것입니다.

![](/2022/Days/Images/Day29_Cloud10.png)

생성 버튼을 누르면 리소스 그룹 목록에 다음 세션에서 수행할 작업을 위한 "90DaysOfDevOps" 그룹이 준비됩니다.

![](/2022/Days/Images/Day29_Cloud11.png)

## 자료

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

[Day 30](day30.md)에서 봐요!
