---
title: '#90DaysOfDevOps - Data Services - Day 85'
published: false
description: 90DaysOfDevOps - Data Services
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048781
---

## 데이터 서비스

데이터베이스는 우리 환경에서 가장 흔히 접하는 데이터 서비스가 될 것입니다. 이번 세션에서는 다양한 유형의 데이터베이스와 각 데이터베이스의 사용 사례에 대해 살펴보고자 합니다. 이번 챌린지를 통해 우리가 사용하고 본 몇 가지 사례를 살펴보겠습니다.

애플리케이션 개발 관점에서 올바른 데이터 서비스 또는 데이터베이스를 선택하는 것은 애플리케이션의 성능과 확장성 측면에서 매우 중요한 결정이 될 것입니다.

https://www.youtube.com/watch?v=W2Z7fbCLSTw

### Key-value

Key-value 데이터베이스는 간단한 Key-value 방식을 사용하여 데이터를 저장하는 비관계형 데이터베이스의 한 유형입니다. Key-value 데이터베이스는 키가 고유 식별자 역할을 하는 Key-value 쌍의 모음으로 데이터를 저장합니다. 키와 값은 단순한 객체부터 복잡한 복합 객체에 이르기까지 무엇이든 될 수 있습니다. Key-value 데이터베이스는 고도로 파티셔닝이 가능하며 다른 유형의 데이터베이스에서는 달성할 수 없는 규모로 수평 확장이 가능합니다.

Key-value 데이터베이스의 예로는 Redis가 있습니다.

_Redis는 인메모리 데이터 구조 저장소로, 분산형 인메모리 Key-value 데이터베이스, 캐시 및 메시지 브로커로 사용되며 내구성을 옵션으로 선택할 수 있습니다. Redis는 문자열, 목록, 맵, 집합, 정렬된 집합, HyperLogLogs, 비트맵, 스트림, 공간 인덱스 등 다양한 종류의 추상 데이터 구조를 지원합니다._

![](/2022/Days/Images/Day85_Data1.png)

Redis에 대한 설명에서 알 수 있듯이 이것은 데이터베이스가 빠르다는 것을 의미하지만, 그 대가로 공간에 제한이 있다는 것을 의미합니다. 또한 쿼리나 조인이 없으므로 데이터 모델링 옵션이 매우 제한적입니다.

Best for:

- 캐싱
- Pub/Sub
- 리더보드
- 쇼핑 카트

일반적으로 다른 영구 데이터 레이어 위에 캐시로 사용됩니다.

### Wide-column

Wide-column 데이터베이스는 데이터 저장소를 여러 서버 또는 데이터베이스 노드에 분산할 수 있는 유연한 컬럼으로 구성하고, 다차원 매핑을 사용하여 컬럼, 행, 타임스탬프로 데이터를 참조하는 NoSQL 데이터베이스입니다.

_Cassandra는 무료 오픈소스, 분산형, 광역 열 저장소, NoSQL 데이터베이스 관리 시스템으로, 여러 상품 서버에서 대량의 데이터를 처리하도록 설계되어 단일 장애 지점 없이 고가용성을 제공합니다._

![](/2022/Days/Images/Day85_Data2.png)

스키마가 없어 비정형 데이터를 처리할 수 없지만 일부 워크로드에는 이점으로 작용할 수 있습니다.

Best for:

- 시계열
- 과거 기록
- 많은 쓰기, 적은 읽기

### 문서

문서 데이터베이스(문서 지향 데이터베이스 또는 문서 저장소라고도 함)는 문서에 정보를 저장하는 데이터베이스입니다.

_MongoDB는 소스 사용이 가능한 크로스 플랫폼 문서 지향 데이터베이스 프로그램입니다. NoSQL 데이터베이스 프로그램으로 분류되는 MongoDB는 선택적 스키마와 함께 JSON과 유사한 문서를 사용합니다. MongoDB는 MongoDB Inc.에서 개발했으며 서버 사이드 퍼블릭 라이선스에 따라 라이선스가 부여됩니다._

![](/2022/Days/Images/Day85_Data3.png)

NoSQL 문서 데이터베이스를 사용하면 복잡한 SQL 코드를 사용하지 않고도 간단한 데이터를 저장할 수 있습니다. 안정성 저하 없이 빠르게 저장할 수 있습니다.

Best for:

- 대부분의 애플리케이션
- 게임
- 사물 인터넷

### 관계형

데이터베이스를 처음 사용하시지만 데이터베이스에 대해 알고 계신다면 관계형 데이터베이스를 접해 보셨을 것입니다.

관계형 데이터베이스는 1970년 E. F. Codd가 제안한 데이터의 관계형 모델에 기반한 디지털 데이터베이스입니다. 관계형 데이터베이스를 유지 관리하는 데 사용되는 시스템은 관계형 데이터베이스 관리 시스템입니다. 많은 관계형 데이터베이스 시스템에는 데이터베이스를 쿼리하고 유지 관리하기 위해 SQL을 사용할 수 있는 옵션이 있습니다.

_MySQL은 오픈소스 관계형 데이터베이스 관리 시스템입니다. 이 이름은 공동 창립자 Michael Widenius의 딸 이름인 'My'와 구조화된 쿼리 언어의 약자인 'SQL'을 조합한 것입니다._

MySQL은 관계형 데이터베이스의 한 예이며 다른 많은 옵션이 있습니다.

![](/2022/Days/Images/Day85_Data4.png)

관계형 데이터베이스를 연구하는 동안 **ACID**라는 용어 또는 약어가 많이 언급되었는데, (원자성, 일관성, 격리성, 내구성)은 오류, 정전 및 기타 사고에도 불구하고 데이터 유효성을 보장하기 위한 데이터베이스 트랜잭션의 일련의 속성입니다. 데이터베이스의 맥락에서 ACID 속성을 충족하는 일련의 데이터베이스 작업(데이터에 대한 단일 논리적 작업으로 인식될 수 있음)을 트랜잭션이라고 합니다. 예를 들어, 한 은행 계좌에서 다른 은행 계좌로 자금을 이체하는 경우, 한 계좌에서 인출하고 다른 계좌에 입금하는 등 여러 가지 변경 사항이 포함되더라도 이는 단일 트랜잭션에 해당합니다.

Best for:

- 대부분의 애플리케이션(수년 동안 사용되어 왔지만, 최고라는 의미는 아님)

비정형 데이터에는 적합하지 않으며, 특정 워크로드에 대해 더 나은 확장 기능을 제공하는 다른 NoSQL이 언급되는 경우 확장 기능이 이상적이지 않습니다.

### 그래프

그래프 데이터베이스는 테이블이나 문서 대신 노드와 관계를 저장합니다. 화이트보드에 아이디어를 스케치하는 것처럼 데이터가 저장됩니다. 데이터는 사전 정의된 모델에 제한되지 않고 저장되므로 매우 유연한 방식으로 데이터를 사고하고 사용할 수 있습니다.

_Neo4j는 Neo4j, Inc.에서 개발한 그래프 데이터베이스 관리 시스템입니다. 개발자는 네이티브 그래프 저장 및 처리 기능을 갖춘 ACID 호환 트랜잭션 데이터베이스라고 설명합니다._

Best for:

- 그래프
- 지식 그래프
- 추천 엔진

### 검색 엔진

지난 섹션에서는 Elasticsearch 방식으로 검색 엔진 데이터베이스를 사용했습니다.

검색 엔진 데이터베이스는 데이터 콘텐츠 검색 전용으로 사용되는 비관계형 데이터베이스의 한 유형입니다. 검색 엔진 데이터베이스는 인덱스를 사용해 데이터 간에 유사한 특성을 분류하고 검색 기능을 용이하게 합니다.

_Elasticsearch는 Lucene 라이브러리를 기반으로 하는 검색 엔진입니다. HTTP 웹 인터페이스와 스키마가 없는 JSON 문서를 갖춘 분산형 멀티테넌트 지원 전체 텍스트 검색 엔진을 제공합니다._

Best for:

- 검색 엔진
- Typeahead
- 로그 검색

### Multi-model

Multi-model 데이터베이스는 단일 통합 백엔드에 대해 여러 데이터 모델을 지원하도록 설계된 데이터베이스 관리 시스템입니다. 반면, 대부분의 데이터베이스 관리 시스템은 데이터를 구성, 저장 및 조작하는 방법을 결정하는 단일 데이터 모델을 중심으로 구성됩니다. 문서, 그래프, 관계형, Key-value 모델은 Multi-model 데이터베이스에서 지원할 수 있는 데이터 모델의 예입니다.

_Fauna는 네이티브 GraphQL을 통해 안전하고 확장 가능한 클라우드 API로 제공되는 유연하고 개발자 친화적인 트랜잭션 데이터베이스입니다._

Best for:

- 데이터 모델 선택에 얽매이지 않는 경우
- ACID 준수
- 빠름
- 프로비저닝 오버헤드 없음
- 데이터를 어떻게 소비하고 클라우드에 무거운 작업을 맡기고 싶으신가요?

어떤 산업에 종사하든 데이터베이스의 한 영역을 접하게 될 것이므로, 이것으로 데이터베이스 개요 세션을 마무리하겠습니다. 그런 다음 이 섹션의 뒷부분에서 몇 가지 예를 들어 데이터 관리, 특히 이러한 데이터 서비스의 보호 및 저장에 대해 살펴볼 것입니다.

아래에 링크한 수많은 리소스를 통해 모든 데이터베이스 유형과 이에 수반되는 모든 것을 심도 있게 살펴보는 데 90년이 걸릴 수도 있습니다.

## 자료

- [Redis Crash Course - the What, Why and How to use Redis as your primary database](https://www.youtube.com/watch?v=OqCK95AS-YE)
- [Redis: How to setup a cluster - for beginners](https://www.youtube.com/watch?v=GEg7s3i6Jak)
- [Redis on Kubernetes for beginners](https://www.youtube.com/watch?v=JmCn7k0PlV4)
- [Intro to Cassandra - Cassandra Fundamentals](https://www.youtube.com/watch?v=YjYWsN1vek8)
- [MongoDB Crash Course](https://www.youtube.com/watch?v=ofme2o29ngU)
- [MongoDB in 100 Seconds](https://www.youtube.com/watch?v=-bt_y4Loofg)
- [What is a Relational Database?](https://www.youtube.com/watch?v=OqjJjpjDRLc)
- [Learn PostgreSQL Tutorial - Full Course for Beginners](https://www.youtube.com/watch?v=qw--VYLpxG4)
- [MySQL Tutorial for Beginners [Full Course]](https://www.youtube.com/watch?v=7S_tz1z_5bA)
- [What is a graph database? (in 10 minutes)](https://www.youtube.com/watch?v=REVkXVxvMQE)
- [What is Elasticsearch?](https://www.youtube.com/watch?v=ZP0NmfyfsoM)
- [FaunaDB Basics - The Database of your Dreams](https://www.youtube.com/watch?v=2CipVwISumA)
- [Fauna Crash Course - Covering the Basics](https://www.youtube.com/watch?v=ihaB7CqJju0)

[Day 86](day86.md)에서 봐요!
