---
title: '#90DaysOfDevOps - The Big Picture: IaC - Day 56'
published: false
description: 90DaysOfDevOps - The Big Picture IaC
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048709
---

## 큰 그림: IaC

인간은 실수를 합니다! 자동화를 도입해야 합니다!

현재 시스템을 어떻게 구축하고 계신가요?

오늘 물리 머신, 가상 머신, 클라우드 VM, 클라우드 PaaS 등 모든 것을 잃게 된다면 어떤 계획을 세우시겠습니까?

모든 것을 교체하는 데 얼마나 걸리나요?

IaC는 이를 테스트할 수 있는 동시에 이를 수행할 수 있는 솔루션을 제공하며, 이를 백업 및 복구와 혼동해서는 안 되지만 인프라와 환경, 플랫폼 측면에서 이를 스핀업하여 가축과 반려동물로 취급할 수 있어야 합니다.

결론은 코드를 사용하여 전체 환경을 재구축할 수 있다는 것입니다.

처음부터 데브옵스에 대해 일반적으로 장벽을 허물어 시스템을 안전하고 신속하게 프로덕션에 배포하는 방법이라고 말씀드린 것을 기억하세요.

IaC는 시스템을 제공하는 데 도움이 되며, 우리는 많은 프로세스와 도구에 대해 이야기했습니다. IaC는 프로세스의 이 부분을 구현하기 위해 익숙한 더 많은 도구를 제공합니다.

이 섹션에서는 IaC에 대해 집중적으로 살펴보겠습니다. 이를 코드의 인프라 또는 코드로서의 구성이라고도 합니다. 가장 잘 알려진 용어는 아마도 IaC가 아닐까 생각합니다.

### 반려동물 대 가축

데브옵스 이전을 살펴보면, 새로운 애플리케이션을 빌드해야 하는 경우 대부분의 경우 서버를 수동으로 준비해야 합니다.

- 가상 머신 배포 | 물리적 서버 및 운영 체제 설치
- 네트워킹 구성
- 라우팅 테이블 생성
- 소프트웨어 및 업데이트 설치
- 소프트웨어 구성
- 데이터베이스 설치

이는 시스템 관리자가 수동으로 수행하는 프로세스입니다. 애플리케이션의 규모가 클수록 더 많은 리소스와 서버가 필요하며, 이러한 시스템을 가동하는 데 더 많은 수작업이 필요합니다. 이 작업에는 엄청난 인력과 시간이 소요될 뿐만 아니라 기업 입장에서도 이러한 환경을 구축하기 위해 리소스에 대한 비용을 지불해야 합니다. "인간은 실수를 합니다! 자동화를 도입해야 합니다!"로 섹션을 시작한 이유입니다.

위의 초기 설정 단계에 이어서 이러한 서버를 유지 관리해야 합니다.

- 버전 업데이트
- 새 릴리스 배포
- 데이터 관리
- 애플리케이션 복구
- 서버 추가, 제거 및 확장
- 네트워크 구성

여러 테스트 및 개발 환경의 복잡성을 추가하세요.

위에서 언급한 서버를 마치 반려동물처럼 돌보던 시절에는 서버에 애칭을 붙이거나 적어도 한동안 '가족'의 일원이 되기를 바라며 이름을 지어주기도 했습니다.

IaC를 사용하면 이러한 모든 작업을 처음부터 끝까지 자동화할 수 있습니다. IaC는 인프라를 자동으로 프로비저닝하는 개념으로, 일부 툴은 서버에 문제가 발생하면 서버를 폐기하고 새 서버를 스핀업하는 작업을 수행합니다. 이 프로세스는 자동화되어 있으며 서버는 코드에 정의된 대로 정확하게 작동합니다. 이 시점에서는 장애가 발생하거나 애플리케이션의 일부 또는 전부를 업데이트하여 더 이상 현장에 존재하지 않고 이를 대체할 다른 서버가 있을 때까지는 서버의 명칭이 무엇이든 상관없습니다.

이는 거의 모든 플랫폼, 가상화, 클라우드 기반 워크로드, 그리고 Kubernetes 및 컨테이너와 같은 클라우드 네이티브 인프라에서 사용할 수 있습니다.

### 인프라 프로비저닝

이 섹션에서 사용할 도구는 아래의 처음 두 가지 영역만 다루고 있습니다. Terraform은 우리가 다룰 도구이며, 이를 통해 무에서 시작하여 인프라의 모양을 코드로 정의한 다음 배포할 수 있으며, 인프라를 관리하고 애플리케이션을 처음 배포할 수도 있지만 그 시점에서는 애플리케이션을 추적할 수 없게 되므로 다음 섹션에서 구성 관리 도구로서 Ansible과 같은 것이 이 부분에서 더 잘 작동할 수 있습니다.

너무 앞서 나가지 않고 초기 애플리케이션 설정을 처리한 다음 해당 애플리케이션과 그 구성을 관리하는 데는 chef, puppet 및 ansible과 같은 도구가 가장 적합합니다.

소프트웨어의 초기 설치 및 구성

- 새 서버 스핀업
- 네트워크 구성
- 로드 밸런서 생성
- 인프라 수준에서 구성

### 프로비저닝된 인프라 구성

- 서버에 애플리케이션 설치
- 애플리케이션을 배포할 서버를 준비합니다.

### 애플리케이션 배포

- 애플리케이션 배포 및 관리
- 유지 관리 단계
- 소프트웨어 업데이트
- 재구성

### IaC 도구의 차이점

선언적 방식과 절차적 방식

절차적

- 단계별 지침
- 서버 생성 > 서버 추가 > 이 변경하기

선언적

- 결과 선언
- 서버 2개

변경 가능(반려동물) vs 변경 불가(가축)

변경 가능

- 대체 대신 변경
- 일반적으로 수명이 길다

변경 불가

- 변경 대신 교체
- 수명이 짧을 수 있음

이것이 바로 IaC를 위한 다양한 옵션이 있는 이유입니다. 모든 것을 지배하는 하나의 도구가 없기 때문입니다.

저희는 주로 Terraform을 사용하며 직접 실습해볼 예정인데, 이것이 IaC가 실제로 작동할 때 그 이점을 확인할 수 있는 가장 좋은 방법이기 때문입니다. 실습은 코드 작성에 필요한 기술을 익힐 수 있는 가장 좋은 방법이기도 합니다.

다음에는 직접 사용해 보기 전에 101을 통해 Terraform에 대해 살펴보겠습니다.

## 자료

아래에 많은 리소스를 나열했으며 이 주제는 이미 여러 번 다루어졌다고 생각합니다. 추가 리소스가 있는 경우 리소스와 함께 PR을 올리면 기꺼이 검토하여 목록에 추가해 드리겠습니다.

- [What is Infrastructure as Code? Difference of Infrastructure as Code Tools](https://www.youtube.com/watch?v=POPP2WTJ8es)
- [Terraform Tutorial | Terraform Course Overview 2021](https://www.youtube.com/watch?v=m3cKkYXl-8o)
- [Terraform explained in 15 mins | Terraform Tutorial for Beginners](https://www.youtube.com/watch?v=l5k1ai_GBDE)
- [Terraform Course - From BEGINNER to PRO!](https://www.youtube.com/watch?v=7xngnjfIlK4&list=WL&index=141&t=16s)
- [HashiCorp Terraform Associate Certification Course](https://www.youtube.com/watch?v=V4waklkBC38&list=WL&index=55&t=111s)
- [Terraform Full Course for Beginners](https://www.youtube.com/watch?v=EJ3N-hhiWv0&list=WL&index=39&t=27s)
- [KodeKloud - Terraform for DevOps Beginners + Labs: Complete Step by Step Guide!](https://www.youtube.com/watch?v=YcJ9IeukJL8&list=WL&index=16&t=11s)
- [Terraform Simple Projects](https://terraform.joshuajebaraj.com/)
- [Terraform Tutorial - The Best Project Ideas](https://www.youtube.com/watch?v=oA-pPa0vfks)
- [Awesome Terraform](https://github.com/shuaibiyy/awesome-terraform)

[Day 57](day57.md)에서 봐요!
