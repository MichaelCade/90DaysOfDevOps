---
title: '#90DaysOfDevOps - Microsoft Azure Storage Models - Day 32'
published: false
description: 90DaysOfDevOps - Microsoft Azure Storage Models
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048775
---

## Microsoft Azure 스토리지 모델

### 스토리지 서비스

- Azure 스토리지 서비스는 스토리지 계정으로 제공됩니다.
- 스토리지 계정은 주로 REST API를 통해 액세스합니다.
- 스토리지 계정은 DNS 이름(`<Storage Account name>.core.windows.net`)의 일부인 고유한 이름을 가져야 합니다.
- 다양한 복제 및 암호화 옵션.
- 리소스 그룹 내에 위치

Azure Portal 상단의 검색 창에서 저장소 그룹을 검색하여 저장소 그룹을 만들 수 있습니다.

![](/2022/Days/Images/Day32_Cloud1.png)

그런 다음 이 이름은 고유해야 하며 공백 없이 모두 소문자이어야 하지만 숫자를 포함할 수 있다는 점을 기억하고 스토리지 계정을 만들기 위한 단계를 실행할 수 있습니다.

![](/2022/Days/Images/Day32_Cloud2.png)

또한 스토리지 계정과 여기에 저장하는 모든 항목에 대해 원하는 중복성 수준을 선택할 수 있습니다. 목록에서 아래로 내려갈수록 옵션은 더 비싸지만, 데이터는 더 많이 분산됩니다.

기본 중복 옵션조차도 데이터의 사본 3개를 제공합니다.

[Azure 스토리지 이중화](https://docs.microsoft.com/en-us/azure/storage/common/storage-redundancy)

위 링크의 요약은 아래와 같습니다:

- **Locally-redundant storage** - 기본 지역의 단일 데이터 센터 내에서 데이터를 세 번 복제합니다.
- **Geo-redundant storage** - LRS를 사용하여 기본 지역의 단일 물리적 위치 내에서 데이터를 세 번 동기식으로 복사합니다.
- **Zone-redundant storage** - 기본 지역의 Azure 가용성 영역 3개에 걸쳐 Azure 스토리지 데이터를 동기식으로 복제합니다.
- **Geo-zone-redundant storage** - 가용성 영역 간 중복으로 제공되는 고가용성과 지리적 복제를 통해 제공되는 지역 중단으로부터의 보호를 결합합니다. GZRS 스토리지 계정의 데이터는 기본 지역의 Azure 가용 영역 3개에 복사되며, 지역 재해로부터 보호하기 위해 두 번째 지리적 영역에도 복제됩니다.

![](/2022/Days/Images/Day32_Cloud3.png)

다시 성능 옵션으로 이동합니다. 스탠다드와 프리미엄 중에서 선택할 수 있습니다. 여기서는 스탠다드를 선택했지만, 프리미엄은 몇 가지 구체적인 옵션을 제공합니다.

![](/2022/Days/Images/Day32_Cloud4.png)

드롭다운에서 다음 세 가지 옵션 중에서 선택할 수 있습니다.

![](/2022/Days/Images/Day32_Cloud5.png)

스토리지 계정에 사용할 수 있는 고급 옵션은 훨씬 더 많지만, 지금은 이러한 영역에 대해 자세히 설명할 필요가 없습니다. 이러한 옵션은 암호화와 데이터 보호에 관한 것입니다.

### 관리 디스크

몇 가지 방법으로 스토리지에 액세스할 수 있습니다.

인증된 액세스를 통한 액세스:

- 전체 제어를 위한 공유키.
- 위임된 세분화된 액세스를 위한 공유 액세스 서명.
- Azure Active Directory(사용 가능한 경우)

공용 액세스:

- 공용 액세스를 부여하여 HTTP를 통한 익명 액세스를 활성화할 수도 있습니다.
- 예를 들어 기본 콘텐츠와 파일을 블록 블롭에 호스팅하여 브라우저가 이 데이터를 보고 다운로드할 수 있도록 할 수 있습니다.

다른 Azure 서비스에서 스토리지에 액세스하는 경우 트래픽은 Azure 내에 유지됩니다.

스토리지 성능에는 두 가지 유형이 있습니다:

- **Standard** - 최대 IOPS 수
- **Premium** - 보장된 IOPS 수

IOPS => 초당 입/출력 작업 수.

작업에 적합한 스토리지를 선택할 때 고려해야 할 비관리형 디스크와 관리형 디스크의 차이점도 있습니다.

### 가상 머신 스토리지

- 가상 머신 OS 디스크는 일반적으로 persistent 스토리지에 저장됩니다.
- 일부 상태 비저장 워크로드에는 persistent 스토리지가 필요하지 않으며 레이턴시 감소가 더 큰 이점이 됩니다.
- 노드-로컬 스토리지에 생성되는 임시 OS 관리 디스크를 지원하는 VM이 있습니다.
  - 이러한 디스크는 VM 스케일 세트와 함께 사용할 수도 있습니다.

관리 디스크는 Azure Virtual 머신과 함께 사용할 수 있는 내구성 있는 블록 스토리지입니다. Ultra Disk Storage, Premium SSD, Standard SSD 또은 Standard HDD를 사용할 수 있습니다. 또한 몇 가지 특징이 있습니다.

- 스냅샷 및 이미지 지원
- SKU 간 간단한 이동
- 가용성 세트와 결합하여 가용성 향상
- 사용한 스토리지가 아닌 디스크 크기를 기준으로 요금이 청구됩니다.

## 아카이브 스토리지

- **Cool Tier** - 블롭을 차단하고 추가할 수 있는 쿨 티어 스토리지를 사용할 수 있습니다.
  - 스토리지 비용 절감
  - 트랜잭션 비용이 더 높습니다.
- **Archive Tier** - 블록 블롭에 아카이브 스토리지를 사용할 수 있습니다.
  - 이것은 블롭 단위로 구성됩니다.
  - 더 저렴한 비용, 더 긴 데이터 검색 대기 시간.
  - 일반 Azure 스토리지와 동일한 데이터 내구성.
  - 필요에 따라 사용자 지정 데이터 계층화를 사용하도록 설정할 수 있습니다.

### 파일 공유

위에서 스토리지 계정을 만들었으므로 이제 파일 공유를 만들 수 있습니다.

![](/2022/Days/Images/Day32_Cloud6.png)

이렇게 하면 Azure에서 SMB2.1 및 3.0 파일 공유가 제공됩니다.

Azure 내부에서 사용할 수 있으며 인터넷에 연결된 SMB3 및 포트 445를 통해 외부에서도 사용할 수 있습니다.

Azure에서 공유 파일 스토리지를 제공합니다.

REST API 외에 표준 SMB 클라이언트를 사용하여 매핑할 수 있습니다.

[Azure NetApp Files](https://vzilla.co.uk/vzilla-blog/azure-netapp-files-how)(SMB 및 NFS)도 참조하세요.

### 캐싱 및 미디어 서비스

Azure 콘텐츠 전송 네트워크는 전 세계에 위치한 정적 웹 콘텐츠의 캐시를 제공합니다.

Azure 미디어 서비스는 재생 서비스 외에 미디어 트랜스코딩 기술을 제공합니다.

## Microsoft Azure 데이터베이스 모델

지난 [Day 28](day28.md)에서 다양한 서비스 옵션에 대해 다뤘습니다. 그중 하나는 많은 양의 인프라와 운영 체제를 추상화하고 애플리케이션 또는 이 경우 데이터베이스 모델만 제어할 수 있는 PaaS(서비스형 플랫폼)였습니다.

### 관계형 데이터베이스

Azure SQL 데이터베이스는 Microsoft SQL Server를 기반으로 하는 서비스로서의 관계형 데이터베이스를 제공합니다.

이는 특정 기능 버전이 필요한 경우 데이터베이스 호환성 수준을 사용할 수 있는 최신 SQL 브랜치를 실행하는 SQL입니다.

이를 구성하는 방법에는 몇 가지 옵션이 있는데, 인스턴스에서 하나의 데이터베이스를 제공하는 단일 데이터베이스를 제공할 수 있으며, 탄력적 풀을 사용하면 용량 풀을 공유하고 집합적으로 확장하는 여러 데이터베이스를 사용할 수 있습니다.

이러한 데이터베이스 인스턴스는 일반 SQL 인스턴스처럼 액세스할 수 있습니다.

MySQL, PostgreSQL 및 MariaDB를 위한 추가 관리형 제품.

![](/2022/Days/Images/Day32_Cloud7.png)

### NoSQL 솔루션

Azure Cosmos DB는 스키마에 구애받지 않는 NoSQL 구현입니다.

99.99% SLA

자동 유도를 통해 전 세계 어디서나 99번째 백분위수에서 한 자릿수 지연 시간을 제공하는 전 세계에 분산된 데이터베이스입니다.

데이터의 partitioning/sharding/distribution에 파티션 키를 활용합니다.

다양한 데이터 모델 지원(documents, key-value, graph, column-friendly)

다양한 API 지원(DocumentDB SQL, MongoDB, Azure Table Storage, Gremlin)

![](/2022/Days/Images/Day32_Cloud9.png)

[CAP 정리](https://en.wikipedia.org/wiki/CAP_theorem)를 기반으로 다양한 정합성 모델을 제공합니다.

![](/2022/Days/Images/Day32_Cloud8.png)

### 캐싱

Redis와 같은 캐싱 시스템에 대한 자세한 설명은 생략하고 Microsoft Azure에는 Azure Cache for Redis라는 서비스가 있다는 점을 말씀드리고 싶었습니다.

Azure Cache for Redis는 Redis 소프트웨어를 기반으로 하는 인메모리 데이터 저장소를 제공합니다.

- 이 서비스는 오픈 소스 Redis Cache를 구현한 것입니다.
  - 호스팅된 보안 Redis 캐시 인스턴스입니다.
  - 다양한 티어를 사용할 수 있습니다.
  - 캐시를 활용하려면 애플리케이션을 업데이트해야 합니다.
  - 쓰기에 비해 읽기 요구 사항이 높은 애플리케이션을 대상으로 합니다.
  - 키-값 저장소 기반.

![](/2022/Days/Images/Day32_Cloud10.png)

지난 며칠 동안 Microsoft Azure에 대해 많은 메모를 하고 이론을 공부했지만, 이러한 구성 요소가 어떻게 결합되고 작동하는지에 대한 실무적인 측면으로 들어가기 전에 기본 구성 요소를 다루고 싶었습니다.

시나리오 기반 서비스 배포를 시작하고 실행하기 전에 네트워킹과 관련된 이론이 한 가지 더 남아 있습니다. 또한 지금까지 사용하던 포털을 사용할 때와 Microsoft Azure와 상호 작용할 수 있는 몇 가지 다른 방법도 살펴보고자 합니다.

## 자료

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

[Day 33](day33.md)에서 봐요!
