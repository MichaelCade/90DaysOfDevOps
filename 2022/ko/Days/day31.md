---
title: '#90DaysOfDevOps - Microsoft Azure Compute Models - Day 31'
published: false
description: 90DaysOfDevOps - Microsoft Azure Compute Models
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049040
---

## Microsoft Azure 컴퓨팅 모델

어제 Microsoft Azure의 보안 모델에 대한 기본 사항을 다룬 데 이어 오늘은 Azure에서 사용할 수 있는 다양한 컴퓨팅 서비스에 대해 살펴보겠습니다.

### 서비스 가용성 옵션

이 섹션은 데이터 관리에서 제 역할을 고려할 때 제 마음에 와닿는 부분입니다. 온-프레미스와 마찬가지로 서비스의 가용성을 보장하는 것이 중요합니다.

- 고가용성(지역 내 보호)
- 재해 복구(지역 간 보호)
- 백업(특정 시점으로부터의 복구)

Microsoft는 지정학적 경계 내에 여러 지역을 배포합니다.

서비스 가용성을 위한 Azure의 두 가지 개념엔 Sets와 영역이 있습니다.

가용성 Sets - 데이터 센터 내에서 복원력 제공

가용성 영역 - 지역 내 데이터 센터 간에 복원력을 제공합니다.

### 가상 머신

퍼블릭 클라우드의 시작점일 가능성이 높습니다.

- 다양한 기능을 갖춘 다양한 계열 및 [크기](https://docs.microsoft.com/en-us/azure/virtual-machines/sizes)의 가상 머신을 제공합니다(때로는 압도적일 수 있음).
- 고성능, 짧은 지연 시간, 높은 메모리 옵션의 VM에 이르기까지 다양한 옵션과 초점이 있는 VM을 제공합니다.
- 또한 B-Series에서 찾을 수 있는 burstable VM 유형도 있습니다. 이는 대부분의 경우 CPU 요구 사항이 낮지만, 한 달에 한 번 정도 성능 급증이 필요한 워크로드에 적합합니다.
- 가상 머신은 모든 네트워크에 연결을 제공할 수 있는 가상 네트워크에 배치됩니다.
- Windows 및 Linux 게스트 OS 지원.
- 특정 Linux 배포판의 경우 [Azure 튜닝된 커널](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/endorsed-distros#azure-tuned-kernels)도 있습니다.

### 템플릿

앞서 마이크로소프트 애저의 이면이나 밑에 있는 모든 것이 JSON이라고 언급했습니다.

리소스를 생성하는 데 사용할 수 있는 여러 가지 관리 포털과 콘솔이 있지만, 가장 선호하는 경로는 JSON 템플릿을 통한 것입니다.

증분 또는 전체 모드, 즉 반복 가능한 원하는 상태의 무동력 배포.

배포된 리소스 정의를 내보낼 수 있는 다양한 템플릿이 있습니다. 저는 이 템플릿 기능을 AWS CloudFormation과 같은 것에 적용하거나 멀티클라우드 옵션을 위한 Terraform이 될 수 있다고 생각하고 있습니다. Terraform에 대해서는 IaC 섹션에서 자세히 다루겠습니다.

### 스케일링

자동 확장은 사용하지 않는 리소스를 스핀다운 하거나 필요할 때 스핀업할 수 있는 퍼블릭 클라우드의 큰 기능입니다.

Azure에는 IaaS용 VMSS(Virtual Machine Scale Sets)라는 것이 있습니다. 이를 통해 일정 및 메트릭을 기반으로 골드 스탠다드 이미지를 자동으로 생성하고 확장할 수 있습니다.

이는 윈도우를 업데이트하는 데 이상적이므로 이미지를 업데이트하고 최소한의 영향을 미치면서 배포할 수 있습니다.

Azure 앱 서비스와 같은 다른 서비스에는 자동 확장 기능이 내장되어 있습니다.

### 컨테이너

컨테이너를 사용 사례로 다루지 않았고 DevOps 학습 여정에서 컨테이너가 필요한 이유와 방법을 다루지 않았지만, Azure에는 몇 가지 특정 컨테이너 중심 서비스가 있다는 점을 언급할 필요가 있습니다.

Azure Kubernetes Service(AKS) - 기반 클러스터 관리의 control plane이나 관리에 대해 걱정할 필요가 없는 관리형 Kubernetes 솔루션을 제공합니다. 나중에 Kubernetes에 대해 자세히 알아보세요.

Azure Container Instances - 초당 과금 방식의 서비스형 컨테이너. 컨테이너 오케스트레이션이 필요 없이 이미지를 실행하고 가상 네트워크와 통합할 수 있습니다.

Service Fabric - 많은 기능을 제공하지만, 컨테이너 인스턴스에 대한 오케스트레이션이 포함되어 있습니다.

Azure에는 Docker 이미지, Helm 차트, OCI 아티팩트 및 이미지에 대한 비공개 레지스트리를 제공하는 컨테이너 레지스트리도 있습니다. 이에 대해서는 컨테이너 섹션에서 다시 자세히 설명하겠습니다.

또한 많은 컨테이너 서비스가 실제로 컨테이너를 내부적으로 활용할 수도 있지만 이는 관리 요구 사항에서 추상화되어 있다는 점도 언급해야 합니다.

앞서 언급한 컨테이너 중심 서비스는 다른 모든 퍼블릭 클라우드에서도 비슷한 서비스를 찾을 수 있습니다.

### 애플리케이션 서비스

- Azure Application Services는 서비스를 쉽게 설정할 수 있는 애플리케이션 호스팅 솔루션을 제공합니다.
- 자동 배포 및 확장.
- Windows 및 Linux 기반 솔루션을 지원합니다.
- 서비스는 유형과 크기가 지정된 앱 서비스 플랜에서 실행됩니다.
- 웹 앱, API 앱, 모바일 앱 등 다양한 서비스를 지원합니다.
- 안정적인 테스트 및 홍보를 위한 배포 슬롯 지원.

### 서버리스 컴퓨팅

저에게 서버리스는 더 배우고 싶은 흥미진진한 다음 단계입니다.

서버리스의 목표는 함수의 런타임에 대해서만 비용을 지불하고 가상 머신이나 PaaS 애플리케이션을 항상 실행할 필요가 없다는 것입니다. 필요할 때 함수를 실행하기만 하면 사라집니다.

Azure Functions - 서버리스 코드를 제공합니다. 퍼블릭 클라우드에 대해 처음 살펴본 것을 기억한다면 관리의 추상화 계층을 기억할 수 있는데, 서버리스 함수를 사용하면 코드만 관리하게 됩니다.

대규모의 이벤트 기반은 나중에 실습을 하게 되면 무언가를 구축할 계획이 있습니다.

많은 Azure 및 타사 서비스에 대한 입력 및 출력 바인딩을 제공합니다.

다양한 프로그래밍 언어를 지원합니다. (C#, NodeJS, Python, PHP, batch, bash, Golang 및 Rust. 또는 모든 실행 파일)

Azure Event Grid를 사용하면 서비스 및 이벤트에서 로직을 트리거할 수 있습니다.

Azure Logic App은 그래픽 기반 workflow 및 통합을 제공합니다.

또한 일관된 관리 및 예약을 통해 Windows 및 Linux 노드 모두에서 대규모 작업을 실행할 수 있는 Azure Batch도 살펴볼 수 있습니다.

## 자료

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

[Day 32](day32.md)에서 봐요!
