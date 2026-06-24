---
title: '#90DaysOfDevOps - The Big Picture: Log Management - Day 79'
published: false
description: 90DaysOfDevOps - The Big Picture Log Management
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049057
---

## 큰 그림: 로그 관리

인프라 모니터링 과제와 솔루션의 연속선상에 있는 로그 관리는 전체 통합 가시성 퍼즐의 또 다른 퍼즐 조각입니다.

### 로그 관리 및 집계

두 가지 핵심 개념에 대해 이야기해 보겠습니다. 첫 번째는 로그 집계로, 다양한 서비스에서 애플리케이션 로그를 수집하고 태그를 지정하여 쉽게 검색할 수 있는 단일 대시보드에 저장하는 방법입니다.

애플리케이션 성능 관리 시스템에서 가장 먼저 구축해야 하는 시스템 중 하나는 로그 집계입니다. 애플리케이션 성능 관리는 애플리케이션이 빌드 및 배포된 후에도 지속적으로 작동하는지 확인하여 리소스가 충분히 할당되고 사용자에게 오류가 표시되지 않도록 해야 하는 데브옵스 라이프사이클의 일부입니다. 대부분의 프로덕션 배포에서 많은 관련 이벤트가 Google의 여러 서비스에서 로그를 생성합니다. 하나의 검색이 사용자에게 반환되기 전에 10개의 다른 서비스에 도달할 수 있으며, 10개의 서비스 중 하나에 논리 문제가 있을 수 있는 예기치 않은 검색 결과가 표시되는 경우 로그 집계는 Google과 같은 회사가 프로덕션에서 문제를 진단하는 데 도움이 되며, 모든 요청을 고유 ID에 매핑할 수 있는 단일 대시보드를 구축하여 검색하면 검색에 고유 ID가 부여되고 검색이 다른 서비스를 통과할 때마다 해당 서비스가 현재 수행 중인 작업과 해당 ID가 연결되도록 합니다.

이것이 바로 로그를 생성하는 모든 곳에서 로그를 효율적으로 수집하고 장애가 다시 발생할 경우 쉽게 검색할 수 있도록 하는 좋은 로그 수집 플랫폼의 핵심입니다.

### 예제 앱

예제 애플리케이션은 웹 앱으로, 일반적인 프론트엔드와 백엔드가 중요한 데이터를 MongoDB 데이터베이스에 저장하고 있습니다.

사용자가 페이지가 모두 하얗게 변하고 오류 메시지가 인쇄되었다고 말하면 현재 스택의 문제를 진단하기 위해 사용자가 수동으로 오류를 보내야 하고 다른 세 가지 서비스의 관련 로그와 일치시켜야 합니다.

### Elk

예제 앱과 동일한 환경에 설치한 경우 세 가지 구성 요소인 Elasticsearch, Logstash, Kibana의 이름을 딴 인기 있는 오픈 소스 로그 집계 스택인 Elk에 대해 살펴보겠습니다.

웹 애플리케이션은 프론트엔드에 연결되고, 프론트엔드는 백엔드에 연결되며, 백엔드는 Logstash로 로그를 전송하고, 이 세 가지 구성 요소가 작동하는 방식은 다음과 같습니다.

### Elk의 구성 요소

Elasticsearch, Logstash, Kibana의 구성 요소는 모든 서비스가 Logstash로 로그를 전송하고, Logstash는 애플리케이션이 방출하는 텍스트인 이 로그를 가져온다는 것입니다. 예를 들어, 웹 애플리케이션에서 사용자가 웹 페이지를 방문할 때 웹 페이지는 이 방문자가 현재 이 페이지에 액세스한 것을 기록할 수 있으며, 이것이 로그 메시지의 예입니다.

그러면 Logstash는 이 로그 메시지에서 사용자가 **시간**에 **무엇**을 했다는 내용을 추출합니다. 시간을 추출하고 메시지를 추출하고 사용자를 추출하고 그것들을 모두 태그로 포함시켜서 메시지가 태그와 메시지의 객체가 되어 특정 사용자의 모든 요청을 쉽게 검색할 수 있도록 하지만 Logstash는 자체적으로 저장하는 것이 아니라 텍스트 쿼리를 위한 효율적인 데이터베이스인 Elasticsearch에 저장하고 Elasticsearch는 Kibana로 결과를 노출하고 Kibana는 Elasticsearch에 연결하는 웹 서버로서 개발자나 팀의 다른 사람들이 관리자를 허용합니다, 대기 중인 엔지니어가 주요 장애가 발생할 때마다 프로덕션에서 로그를 볼 수 있습니다. 관리자는 Kibana에 연결하고, Kibana는 사용자가 원하는 것과 일치하는 로그를 찾기 위해 Elasticsearch를 쿼리합니다.

검색 창에 "오류를 찾고 싶어요"라고 말하면, Kibana는 문자열 오류가 포함된 메시지를 찾으라고 말하고, 그러면 Elasticsearch는 Logstash가 채운 결과를 반환합니다. Logstash는 다른 모든 서비스에서 이러한 결과를 전송받았을 것입니다.

### Elk를 사용해 프로덕션 문제를 진단하는 방법

한 사용자가 Elk 설정으로 이 작업을 수행하려고 할 때 오류 코드 1234567을 보았다고 말합니다. 검색 창에 1234567을 입력하고 엔터를 누르면 그에 해당하는 로그가 표시되고 로그 중 하나에서 내부 서버 오류가 1234567을 반환한다고 표시될 수 있으며, 그 로그를 생성한 서비스가 로그를 생성한 서비스가 백엔드임을 알 수 있고, 해당 로그가 생성된 시간을 확인할 수 있으므로 해당 로그의 시간으로 이동하여 백엔드에서 그 위와 아래의 메시지를 보면 사용자의 요청에 대해 어떤 일이 발생했는지 더 잘 파악할 수 있으며, 이 과정을 다른 서비스에서도 반복하여 사용자에게 문제가 발생한 원인을 찾을 때까지 반복할 수 있습니다.

### 보안 및 로그 액세스

퍼즐의 중요한 조각은 로그가 관리자(또는 액세스 권한이 필요한 사용자 및 그룹)에게만 표시되도록 하는 것입니다. 로그에는 인증된 사용자만 액세스해야 하는 토큰과 같은 민감한 정보가 포함될 수 있으며, 인증 방법 없이 Kibana를 인터넷에 노출하고 싶지 않을 것입니다.

### 로그 관리 도구의 예

로그 관리 플랫폼의 예는 다음과 같습니다.

- Elasticsearch
- Logstash
- Kibana
- Fluentd - 인기 있는 오픈 소스 선택
- Datadog - 대기업에서 일반적으로 사용되는 호스팅 제품,
- LogDNA - 호스트형 제품
- Splunk

클라우드 제공업체는 AWS CloudWatch Logs, Microsoft Azure Monitor, Google Cloud Logging과 같은 로깅도 제공합니다.

로그 관리는 프로덕션 환경의 문제를 진단하기 위한 애플리케이션 및 인프라 환경의 전반적인 통합 가시성의 핵심 요소입니다. Elk 또는 CloudWatch와 같은 즌비된 솔루션을 설치하는 것은 비교적 간단하며 프로덕션 환경의 문제를 진단하고 분류하는 것이 훨씬 쉬워집니다.

## 자료

- [The Importance of Monitoring in DevOps](https://www.devopsonline.co.uk/the-importance-of-monitoring-in-devops/)
- [Understanding Continuous Monitoring in DevOps?](https://medium.com/devopscurry/understanding-continuous-monitoring-in-devops-f6695b004e3b)
- [DevOps Monitoring Tools](https://www.youtube.com/watch?v=Zu53QQuYqJ0)
- [Top 5 - DevOps Monitoring Tools](https://www.youtube.com/watch?v=4t71iv_9t_4)
- [How Prometheus Monitoring works](https://www.youtube.com/watch?v=h4Sl21AKiDg)
- [Introduction to Prometheus monitoring](https://www.youtube.com/watch?v=5o37CGlNLr8)
- [Promql cheat sheet with examples](https://www.containiq.com/post/promql-cheat-sheet-with-examples)
- [Log Management for DevOps | Manage application, server, and cloud logs with Site24x7](https://www.youtube.com/watch?v=J0csO_Shsj0)
- [Log Management what DevOps need to know](https://devops.com/log-management-what-devops-teams-need-to-know/)
- [What is Elk Stack?](https://www.youtube.com/watch?v=4X0WLg05ASw)
- [Fluentd simply explained](https://www.youtube.com/watch?v=5ofsNyHZwWE&t=14s)

[Day 80](day80.md)에서 봐요!
