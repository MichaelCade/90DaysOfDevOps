---
title: '#90DaysOfDevOps - Testing, Tools & Alternatives - Day 62'
published: false
description: '90DaysOfDevOps - Testing, Tools & Alternatives'
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049053
---

## 테스트, 도구 및 대안

IaC에 대한 이 섹션을 마무리하면서 코드 테스트, 사용 가능한 다양한 도구, 그리고 이를 달성하기 위한 Terraform의 몇 가지 대안에 대해 언급하지 않을 수 없습니다. 이 섹션의 시작 부분에서 말씀드렸듯이 제가 Terraform에 초점을 맞춘 이유는 첫째로 무료이며 오픈 소스이고 둘째로 크로스 플랫폼이며 환경에 구애받지 않기 때문입니다. 하지만 고려해야 할 다른 대안도 있지만, 전반적인 목표는 이것이 인프라를 배포하는 방법이라는 것을 사람들에게 알리는 것입니다.

### Code Rot

이 세션에서 다루고자 하는 첫 번째 영역은 Code Rot으로, 애플리케이션 코드와 달리 코드로서의 인프라는 한 번 사용되었다가 오랫동안 사용하지 않을 수도 있습니다. Terraform을 사용하여 AWS에 VM 환경을 배포하려고 하는데, 처음에 완벽하게 작동하고 환경이 갖춰졌지만, 이 환경이 자주 변경되지 않아 코드가 중앙 위치에 저장되기를 바라지만 코드가 변경되지 않는 상태를 그대로 유지한다고 가정해 보겠습니다.

인프라에 변화가 생기면 어떻게 될까요? 하지만 대역 외에서 수행되거나 다른 환경이 변경되는 경우가 있습니다.

- 대역 외 변경 사항
- 고정되지 않은 버전
- 더 이상 사용되지 않는 종속성
- 적용되지 않은 변경 사항

### 테스트

코드 썩음과 일반적으로 뒤따르는 또 다른 큰 영역은 IaC를 테스트하고 모든 영역이 정상적으로 작동하는지 확인하는 기능입니다.

먼저 살펴볼 수 있는 몇 가지 기본 제공 테스트 명령어가 있습니다:

| Command              | Description                                                              |
| -------------------- | ------------------------------------------------------------------------ |
| `terraform fmt`      | Terraform 구성 파일을 표준 형식과 스타일로 다시 작성합니다.              |
| `terraform validate` | 디렉터리에 있는 구성 파일의 유효성을 검사하여 구성만 참조합니다.         |
| `terraform plan`     | 실행 계획을 생성하여 Terraform이 계획한 변경 사항을 미리 볼 수 있습니다. |
| Custom validation    | 입력 변수가 예상과 일치하는지 확인하기 위한 입력 변수 유효성 검사        |

Terraform 외부에서 사용할 수 있는 몇 가지 테스트 도구도 있습니다:

- [tflint](https://github.com/terraform-linters/tflint)

  - 가능한 오류를 찾습니다.
  - 더 이상 사용되지 않는 구문과 사용되지 않는 선언에 대해 경고합니다.
  - 모범 사례와 명명 규칙을 적용합니다.

검사 도구

- [checkov](https://www.checkov.io/) - 클라우드 인프라 구성을 스캔하여 배포하기 전에 잘못된 구성을 찾습니다.
- [tfsec](https://aquasecurity.github.io/tfsec/v1.4.2/) - Terraform 코드에 대한 정적 분석 보안 스캐너입니다.
- [terrascan](https://github.com/accurics/terrascan) - IaC를 위한 정적 코드 분석기입니다.
- [terraform-compliance](https://terraform-compliance.com/) - IaC에 대한 네거티브 테스트 기능을 지원하는 Terraform에 대한 경량의 보안 및 규정 준수 중심 테스트 프레임워크입니다.
- [snyk](https://docs.snyk.io/products/snyk-infrastructure-as-code/scan-terraform-files/scan-and-fix-security-issues-in-terraform-files) - Terraform 코드에서 잘못된 구성과 보안 문제를 스캔합니다.

관리형 클라우드 제품

- [Terraform Sentinel](https://www.terraform.io/cloud-docs/sentinel) - HashCorp 엔터프라이즈 제품과 통합된 임베디드 정책 기반 프레임워크로 세분화된 로직 기반 정책 결정을 가능하게 하며, 외부 소스의 정보를 사용하도록 확장할 수 있습니다.

### 자동화된 테스트

- [Terratest](https://terratest.gruntwork.io/) - Terratest는 인프라 테스트를 위한 패턴과 도우미 함수를 제공하는 Go 라이브러리입니다.
- Terratest는 인프라 코드에 대한 자동화된 테스트 작성을 용이하게 합니다. 일반적인 인프라 테스트를 위한 다양한 도우미 함수와 패턴을 제공합니다.
- 코드는 2022/Days/IaC/Terratest에서 찾을 수 있습니다.
- 이 애플리케이션을 실행하려면
  - git clone #repo_url# <br />
  - cd test <br />
  - go mod init "<MODULE_NAME>" <br />
    **MODULE_NAME은 github.com/<YOUR_USERNAME>/<YOUR_REPO_NAME>이 됩니다.** <br />
  - go mod init github.com/<FOLDER-PATH> <br/>
  - go run

---

go mod init "<MODULE_NAME>"은 테스트 폴더에 go.mod 파일을 생성합니다. <br />

- go.mod 파일은 GoLang에서 의존성 관리의 기초입니다.
- 프로젝트에서 필요하거나 사용할 모든 모듈이 go.mod 파일에 유지됩니다.
- 프로젝트에서 사용하거나 가져올 모든 패키지의 항목을 생성합니다.
- 수동으로 각 종속성을 가져오는 노력을 줄입니다.

처음 **go test**를 실행하면 go.sum 파일이 생성됩니다. <br />

- **go test** 또는 **go build**가 처음 실행될 때 go.sum 파일이 생성됩니다.
- 특정 버전(최신)으로 모든 패키지를 설치합니다.
- 우린 파일을 편집하거나 수정할 필요가 없습니다.

추가로 언급할 만한 것들

- [Terraform 클라우드](https://cloud.hashicorp.com/products/terraform) - Terraform 클라우드는 HashCorp의 관리형 서비스 제품입니다. 실무자, 팀 및 조직이 프로덕션에서 Terraform을 사용하기 위해 불필요한 도구와 문서가 필요하지 않습니다.

- [Terragrunt](https://terragrunt.gruntwork.io/) - Terragrunt는 구성을 건조하게 유지하고, 여러 Terraform 모듈로 작업하고, 원격 상태를 관리하기 위한 추가 도구를 제공하는 얇은 래퍼입니다.

- [Atlantis](https://www.runatlantis.io/) - Terraform 풀 리퀘스트 자동화

### 대안

이 섹션을 시작할 때 57일째에 몇 가지 대안이 있다고 언급했으며, 이번 챌린지에서도 이에 대해 계속 살펴볼 계획입니다.

| Cloud Specific                  | Cloud Agnostic |
| ------------------------------- | -------------- |
| AWS CloudFormation              | Terraform      |
| Azure Resource Manager          | Pulumi         |
| Google Cloud Deployment Manager |                |

저는 위의 목록 중 AWS CloudFormation을 가장 많이 사용했으며 AWS에서 기본으로 제공하지만, Terraform 이외의 다른 제품은 사용해 본 적이 없습니다. 클라우드별 버전은 특정 클라우드에서는 매우 훌륭하지만, 여러 클라우드 환경이 있는 경우 이러한 구성을 마이그레이션하는 데 어려움을 겪거나 IaC 작업을 위해 여러 관리 플레인을 사용해야 할 것입니다.

흥미로운 다음 단계는 시간을 내서 [Pulumi](https://www.pulumi.com/)에 대해 자세히 알아보는 것입니다.

Pulumi 사이트의 비교 부분에서

> "Terraform과 Pulumi 모두 코드가 원하는 인프라 상태를 나타내는 코드 모델로 원하는 상태 인프라를 제공하며, 배포 엔진은 이 원하는 상태를 스택의 현재 상태와 비교하여 어떤 리소스를 생성, 업데이트 또는 삭제해야 하는지를 결정합니다."

제가 볼 수 있는 가장 큰 차이점은 HCL(HashCorp 구성 언어)과 달리 Pulumi는 Python, TypeScript, JavaScript, Go, .NET과 같은 범용 언어를 허용한다는 점입니다.

간략한 개요 [Pulumi 소개: 코드형 최신 인프라](https://www.youtube.com/watch?v=QfJTJs24-JM) 저는 선택의 폭이 넓다는 점이 마음에 들어서 좀 더 자세히 알아보고 싶었습니다.

이것으로 IaC 섹션을 마무리하고, 다음에는 구성 관리와 약간 겹치는 부분, 특히 일부 작업 및 데모에 Ansible을 사용할 구성 관리의 큰 그림을 살펴볼 것입니다.

## 리소스

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
- [Pulumi - IaC in your favorite programming language!](https://www.youtube.com/watch?v=vIjeiDcsR3Q&t=51s)

[Day 63](day63.md)에서 봐요!
