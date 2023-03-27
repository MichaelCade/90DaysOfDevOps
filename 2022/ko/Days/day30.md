---
title: '#90DaysOfDevOps - Microsoft Azure Security Models - Day 30'
published: false
description: 90DaysOfDevOps - Microsoft Azure Security Models
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049039
---

## Microsoft Azure 보안 모델

Microsoft Azure 개요에 이어서 Azure 보안부터 시작하여 일상적인 업무에 어떤 도움이 될 수 있는지 살펴보겠습니다. 대부분의 경우 기본 제공 역할로도 충분하지만, 다양한 인증 및 구성 영역을 만들고 작업할 수 있다는 것을 알게 되었습니다. 다른 퍼블릭 클라우드에 비해 Microsoft Azure의 Active Directory 배경이 상당히 발전되어 있다는 것을 알았습니다.

Microsoft Azure가 다른 퍼블릭 클라우드 공급자와 다르게 작동하는 것처럼 보이는 영역 중 하나는 Azure에는 항상 Azure AD가 있다는 것입니다.

### 디렉토리 서비스

- Azure Active Directory는 Microsoft Azure 및 기타 Microsoft 클라우드 서비스에서 사용하는 보안 원칙을 호스팅합니다.
- 인증은 SAML, WS-Federation, OpenID Connect 및 OAuth2와 같은 프로토콜을 통해 수행됩니다.
- 쿼리는 Microsoft Graph API라는 REST API를 통해 수행됩니다.
- 테넌트는 tenant.onmicrosoft.com을 기본 이름으로 사용하지만 사용자 지정 도메인 이름을 사용할 수도 있습니다.
- 구독은 Azure Active Directory 테넌트와 연결됩니다.

AWS와 비교하기 위해 AWS를 생각해보면 여전히 매우 다르지만, AWS IAM(ID 및 액세스 관리)이 동등한 오퍼링이 될 것입니다.

Azure AD Connect는 AD에서 Azure AD로 계정을 복제하는 기능을 제공합니다. 여기에는 그룹과 때로는 개체가 포함될 수도 있습니다. 이는 세분화 및 필터링할 수 있습니다. 여러 포리스트 및 도메인을 지원합니다.

Microsoft Azure AD(Active Directory)에서 클라우드 계정을 만들 수 있지만 대부분의 조직은 이미 온-프레미스에 있는 자체 Active Directory에서 사용자를 계정화하고 있습니다.

Azure AD Connect를 사용하면 Windows AD 서버뿐만 아니라 다른 Azure AD, Google 및 기타 서버도 볼 수 있습니다. 또한 외부 사람 및 조직과 공동 작업할 수 있는 기능을 제공하는데, 이를 Azure B2B라고 합니다.

액티브 디렉토리 도메인 서비스와 Microsoft Azure 액티브 디렉토리 간의 인증 옵션은 암호 해시와 ID 동기화를 통해 모두 가능합니다.

![](/2022/Days/Images/Day30_Cloud1.png)

암호 해시 전달은 선택 사항이며, 암호 해시를 사용하지 않는 경우 통과 인증이 필요합니다.

아래에 링크된 비디오에서 패스스루 인증에 대해 자세히 설명합니다.

[Azure Active Directory 패스스루 인증을 사용한 사용자 로그인](https://docs.microsoft.com/en-us/azure/active-directory/hybrid/how-to-connect-pta)

![](/2022/Days/Images/Day30_Cloud2.png)

### 페더레이션

Microsoft 365, Microsoft Dynamics 및 온-프레미스 Active Directory를 사용하는 경우 페더레이션을 위해 Azure AD를 이해하고 통합하는 것은 매우 쉽습니다. 하지만 Microsoft 에코시스템 외부의 다른 서비스를 사용하고 있을 수도 있습니다.

Azure AD는 이러한 다른 타사 앱 및 기타 디렉토리 서비스에 대한 페더레이션 브로커 역할을 할 수 있습니다.

이는 Azure 포털에서 다양한 옵션이 있는 엔터프라이즈 애플리케이션으로 표시됩니다.

![](/2022/Days/Images/Day30_Cloud3.png)

엔터프라이즈 애플리케이션 페이지에서 아래로 스크롤 하면 추천 애플리케이션의 긴 목록을 볼 수 있습니다.

![](/2022/Days/Images/Day30_Cloud4.png)

이 옵션을 사용하면 개발 중인 애플리케이션이나 갤러리 이외의 애플리케이션을 "직접 가져와서" 통합할 수도 있습니다.

이전에 이 기능을 살펴본 적은 없지만 다른 클라우드 제공업체 및 기능과 비교할 때 상당한 기능 집합이라는 것을 알 수 있습니다.

### 역할 기반 액세스 제어

여기서 다룰 범위는 이미 [Day 29](day29.md)에서 다뤘으며, 이러한 영역 중 하나에 따라 역할 기반 액세스 제어를 설정할 수 있습니다.

- 구독
- 관리 그룹
- 리소스 그룹
- 리소스

역할은 세 가지로 나눌 수 있으며, Microsoft Azure에는 많은 기본 제공 역할이 있습니다. 이 세 가지 역할은 다음과 같습니다:

- 소유자
- 기여자
- 리더

소유자와 기여자는 범위가 매우 유사하지만, 소유자가 권한을 변경할 수 있습니다.

다른 역할은 특정 유형의 Azure 리소스 및 사용자 지정 역할에 따라 다릅니다.

그룹과 사용자에 대한 사용 권한 할당에 중점을 두어야 합니다.

사용 권한은 상속됩니다.

다시 돌아가서 우리가 만든 "90DaysOfDevOps" 리소스 그룹을 살펴보고 그 안의 액세스 제어(IAM)를 확인하면 기여자 목록과 고객 사용자 액세스 관리자 목록이 있고 소유자 목록이 있는 것을 볼 수 있습니다(하지만 이 목록은 표시할 수 없습니다).

![](/2022/Days/Images/Day30_Cloud5.png)

또한 여기에서 할당된 역할이 빌트인 역할인지 여부와 해당 역할이 어떤 카테고리에 속하는지 확인할 수 있습니다.

![](/2022/Days/Images/Day30_Cloud6.png)

이 리소스 그룹에 대해 계정을 확인하고 해당 액세스 권한을 갖고자 하는 계정에 올바른 권한이 있는지 확인하거나 사용자가 너무 많은 액세스 권한을 갖고 있는지 확인하려는 경우 액세스 확인 탭을 사용할 수도 있습니다.

![](/2022/Days/Images/Day30_Cloud7.png)

### 클라우드용 Microsoft Defender

- 클라우드용 Microsoft Defender(이전의 Azure Security Center)는 전체 Azure 환경의 보안에 대한 인사이트를 제공합니다.

- 단일 대시보드를 통해 모든 Azure 및 비 Azure 리소스의 전반적인 보안 상태를 파악할 수 있으며(Azure Arc를 통해) 보안 강화 지침을 제공합니다.

- 프리티어에는 지속적인 평가 및 보안 권장 사항이 포함됩니다.

- 보호되는 리소스 유형(예: Servers, AppService, SQL, Storage, Containers, KeyVault)에 대한 유료 플랜.

다른 구독으로 전환하여 Azure 보안 센터를 확인했으며, 여기에서 몇 가지 리소스를 기반으로 한 곳에서 몇 가지 권장 사항을 확인할 수 있습니다.

![](/2022/Days/Images/Day30_Cloud8.png)

### Azure 정책

- Azure Policy는 조직 표준을 적용하고 대규모로 규정 준수를 평가하는 데 도움이 되는 Azure 기본 서비스입니다.

- 클라우드용 Microsoft Defender에 통합됩니다. Azure Policy는 규정을 준수하지 않는 리소스를 감사하고 수정 사항을 적용합니다.

- 리소스 일관성, 규정 준수, 보안, 비용 및 관리 표준을 관리하는 데 일반적으로 사용됩니다.

- JSON 형식을 사용하여 평가 로직을 저장하고 리소스의 규정 준수 여부와 규정 미준수 시 취할 조치(예: 감사, AuditIfNotExists, 거부, 수정, DeployIfNotExists)를 결정합니다.

- 무료로 사용할 수 있습니다. 단, Azure 정책 게스트 구성 사용량에 대해 서버/월당 청구되는 Azure Arc 연결 리소스는 예외입니다.

### 실습

www.90DaysOfDevOps.com 도메인을 구입했는데 이 도메인을 내 Azure Active Directory 포털에 추가하고 싶습니다. [Azure Active Directory 포털을 사용하여 사용자 지정 도메인 이름 추가](https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/add-custom-domain)

![](/2022/Days/Images/Day30_Cloud9.png)

이제 새 Active Directory 도메인에 새 사용자를 만들 수 있습니다.

![](/2022/Days/Images/Day30_Cloud10.png)

이제 모든 새로운 90DaysOfDevOps 사용자를 하나의 그룹으로 묶는 그룹을 만들고 싶습니다. 아래와 같이 그룹을 만들 수 있으며, Azure AD가 사용자 계정을 쿼리하여 동적으로 추가하는 "Dynamic User"와 수동으로 사용자를 그룹에 추가하는 "Assigned"를 사용하고 있음을 알 수 있습니다.

![](/2022/Days/Images/Day30_Cloud11.png)

쿼리를 작성할 때 많은 옵션이 있지만, 저는 단순히 대표 이름을 찾아 이름에 @90DaysOfDevOps.com이 포함되어 있는지 확인하려고 합니다.

![](/2022/Days/Images/Day30_Cloud12.png)

이제 michael.cade@90DaysOfDevOps.com 에 대한 사용자 계정을 이미 만들었으므로 규칙이 작동하는지 확인할 수 있습니다. 비교를 위해 다른 도메인에 연결한 다른 계정도 여기에 추가했는데, 이 규칙으로 인해 사용자가 이 그룹에 속하지 않는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day30_Cloud13.png)

그 후 user1@90DaysOfDevOps.com 을 새로 추가했고 그룹을 확인하면 회원을 볼 수 있습니다.

![](/2022/Days/Images/Day30_Cloud14.png)

이 요구 사항이 많아진다면, 콘솔에서 이 모든 작업을 수행하지 않고 대량 옵션을 사용하여 사용자를 만들고, 초대하고, 삭제하거나 PowerShell을 사용하여 자동화된 확장 방식을 사용하고 싶을 것입니다.

이제 리소스 그룹으로 이동하여 90DaysOfDevOps 리소스 그룹에서 소유자가 방금 만든 그룹이 되도록 지정할 수 있습니다.

![](/2022/Days/Images/Day30_Cloud15.png)

마찬가지로 여기로 이동하여 리소스 그룹에 대한 할당 액세스를 거부할 수도 있습니다.

이제 새 사용자 계정으로 Azure Portal에 로그인하면 액세스 권한이 없기 때문에 이전 그림에서 볼 수 있는 다른 리소스 그룹에는 액세스 권한이 없고 90DaysOfDevOps 리소스 그룹에만 액세스 권한이 있는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day30_Cloud16.png)

모든 사용자가 포털을 알고 있을 필요는 없지만, 액세스 권한을 확인하기 위해 [앱 포털](https://myapps.microsoft.com/)을 사용하여 테스트할 수 있는 싱글 사인온 포털입니다.

![](/2022/Days/Images/Day30_Cloud17.png)

이 포털을 브랜딩으로 사용자 지정할 수 있으며, 이 부분은 나중에 다시 다룰 수 있습니다.

## 자료

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

[Day 31](day31.md)에서 봐요!
