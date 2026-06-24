---
title: '#90DaysOfDevOps - An intro to Terraform - Day 57'
published: false
description: 90DaysOfDevOps - An intro to Terraform
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048710
---

## Terraform 소개

"Terraform은 인프라를 안전하고 효율적으로 구축, 변경, 버전 관리할 수 있는 도구입니다."

위의 인용문은 Terraform의 개발사인 HashiCorp에서 인용한 것입니다.

"Terraform은 수백 개의 클라우드 서비스를 관리하기 위한 일관된 CLI workflow를 제공하는 코드 소프트웨어 도구로서 오픈 소스 인프라입니다. Terraform은 클라우드 API를 선언적 구성 파일로 코드화합니다."

해시코프는 [HashiCorp 배우기](https://learn.hashicorp.com/terraform?utm_source=terraform_io&utm_content=terraform_io_hero)에서 모든 제품을 다루는 훌륭한 리소스를 제공하고 있으며, IaC로 무언가를 달성하려고 할 때 유용한 데모를 제공합니다.

모든 클라우드 제공업체와 온프레미스 플랫폼은 일반적으로 UI를 통해 리소스를 생성할 수 있는 관리 콘솔에 대한 액세스를 제공하며, 일반적으로 이러한 플랫폼은 동일한 리소스를 생성할 수 있는 CLI 또는 API 액세스도 제공하지만, API를 사용하면 빠르게 프로비저닝할 수 있습니다.형

IaC를 사용하면 이러한 API에 연결하여 원하는 상태로 리소스를 배포할 수 있습니다.

아래에 배타적이거나 완전한 도구가 아닌 다른 도구가 있습니다. 다른 도구가 있다면 PR을 통해 공유해 주세요.

| Cloud Specific                  | Cloud Agnostic |
| ------------------------------- | -------------- |
| AWS CloudFormation              | Terraform      |
| Azure Resource Manager          | Pulumi         |
| Google Cloud Deployment Manager |                |

데모뿐만 아니라 일반적으로 사용하고자 하는 클라우드와 플랫폼에 구애받지 않으려는 것이 바로 우리가 Terraform을 사용하는 또 다른 이유입니다.

## Terraform 개요

Terraform은 프로비저닝에 중점을 둔 도구로, 복잡한 인프라 환경을 프로비저닝할 수 있는 기능을 제공하는 CLI입니다. 로컬 또는 원격(클라우드)에 존재하는 복잡한 인프라 요구 사항을 정의할 수 있습니다. Terraform을 사용하면 처음에 구축할 수 있을 뿐만 아니라 해당 리소스의 수명 기간 동안 유지 관리 및 업데이트할 수 있습니다.

여기서는 개략적인 내용만 다루겠지만, 더 자세한 내용과 다양한 리소스를 보려면 [terraform.io](https://www.terraform.io/)를 방문하세요.

### 쓰기

Terraform을 사용하면 환경을 구축할 선언적 구성 파일을 만들 수 있습니다. 이 파일은 블록, 인수, 표현식을 사용하여 리소스를 간결하게 설명할 수 있는 해시코프 구성 언어(HCL)를 사용하여 작성됩니다. 물론 가상 머신, 컨테이너를 배포할 때와 쿠버네티스 내에서 이를 자세히 살펴볼 것입니다.

### 계획

위의 구성 파일들이 우리가 보고자 하는 것을 배포할 것인지 확인하기 위해 Terraform cli의 특정 기능을 사용하여 배포하거나 변경하기 전에 해당 계획을 테스트할 수 있는 기능입니다. Terraform은 인프라를 위한 지속적인 도구이므로 인프라의 측면을 변경하려면 코드에서 모두 캡처되도록 Terraform을 통해 변경해야 한다는 점을 기억하세요.

### 적용

만족스러우면 계속해서 이 구성을 Terraform 내에서 사용할 수 있는 많은 제공업체에 적용할 수 있습니다. [여기](https://registry.terraform.io/browse/providers)에서 사용 가능한 수많은 제공자를 확인할 수 있습니다.

또 한 가지 언급할 것은 사용 가능한 모듈도 있다는 것인데, 이러한 모듈은 공개적으로 생성 및 공유되어 있으므로 특정 인프라 리소스를 모든 곳에 동일한 방식으로 배포하는 모범 사례를 반복해서 생성할 필요가 없다는 점에서 컨테이너 이미지와 유사합니다. 사용 가능한 모듈은 [여기](https://registry.terraform.io/browse/modules)에서 찾을 수 있습니다.

Terraform workflow는 다음과 같습니다. (_Terraform 사이트에서 가져옴_)

![](/2022/Days/Images/Day57_IAC3.png)

### Terraform vs Vagrant

이번 챌린지에서는 개발 환경에 집중하는 또 다른 해시코프 오픈소스 도구인 Vagrant를 사용했습니다.

- Vagrant는 개발 환경 관리에 중점을 둔 도구입니다.

- Terraform은 인프라 구축을 위한 도구입니다.

두 도구에 대한 자세한 비교는 공식 [HashiCorp 사이트](https://www.vagrantup.com/intro/vs/terraform)에서 확인할 수 있습니다.

## Terraform 설치

Terraform을 설치하는 데에는 많은 것이 필요하지 않습니다.

Terraform은 크로스 플랫폼이며, 제 리눅스 머신에서 아래에서 CLI를 다운로드하고 설치하는 몇 가지 옵션을 볼 수 있습니다.

![](/2022/Days/Images/Day57_IAC2.png)

`arkade`를 사용하여 Terraform을 설치하면, 아케이드는 필요한 도구, 앱, 클리스를 시스템에 설치할 수 있는 편리한 작은 도구입니다. 간단한 `arkade get terraform`으로 Terraform을 업데이트할 수 있으며, Terraform이 있다면 같은 명령으로 Terraform CLI도 설치할 수 있습니다.

![](/2022/Days/Images/Day57_IAC1.png)

앞으로 HCL에 대해 좀 더 자세히 살펴본 다음 다양한 플랫폼에서 인프라 리소스를 생성하는 데 Terraform을 사용해 보도록 하겠습니다.

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

[Day 58](day58.md)에서 봐요!
