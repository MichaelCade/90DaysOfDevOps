---
title: '#90DaysOfDevOps - The Big Picture: Kubernetes - Day 49'
published: false
description: 90DaysOfDevOps - The Big Picture Kubernetes
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049049
---

## 큰 그림: Kubernetes

지난 섹션에서 컨테이너에 대해 살펴보았는데, 컨테이너는 scale과 오케스트레이션만으로는 부족합니다. 우리가 할 수 있는 최선은 docker-compose를 사용하여 여러 컨테이너를 함께 불러오는 것입니다. 컨테이너 오케스트레이터인 Kubernetes를 사용하면 애플리케이션과 서비스의 부하에 따라 자동화된 방식으로 확장 및 축소할 수 있습니다.

플랫폼으로서 Kubernetes는 요구사항과 원하는 상태에 따라 컨테이너를 오케스트레이션할 수 있는 기능을 제공합니다. 이 섹션에서는 차세대 인프라로 빠르게 성장하고 있는 Kubernetes에 대해 다룰 예정입니다. 또한 데브옵스 관점에서 볼 때 Kubernetes는 기본적인 이해가 필요한 하나의 플랫폼일 뿐이며, 베어메탈, 가상화 및 대부분의 클라우드 기반 서비스도 이해해야 합니다. Kubernetes는 애플리케이션을 실행하기 위한 또 다른 옵션일 뿐입니다.

### 컨테이너 오케스트레이션이란 무엇인가요?

앞서 Kubernetes와 컨테이너 오케스트레이션에 대해 언급했는데, Kubernetes는 기술인 반면 컨테이너 오케스트레이션은 기술 이면의 개념 또는 프로세스입니다. 컨테이너 오케스트레이션 플랫폼은 Kubernetes뿐만 아니라 Docker Swarm, HashiCorp Nomad 등도 있습니다. 하지만 Kubernetes가 점점 더 강세를 보이고 있기 때문에 Kubernetes를 다루고 싶지만, Kubernetes만이 유일한 것은 아니라는 점을 말씀드리고 싶었습니다.

### Kubernetes란 무엇인가요?

Kubernetes를 처음 접하는 경우 가장 먼저 읽어야 할 것은 공식 문서입니다. 1년 조금 전에 Kubernetes에 대해 깊이 파고든 제 경험에 따르면 학습 곡선이 가파르게 진행될 것입니다. 가상화 및 스토리지에 대한 배경지식이 있는 저는 이것이 얼마나 벅차게 느껴질지 생각했습니다.

하지만 커뮤니티, 무료 학습 리소스 및 문서는 놀랍습니다. [Kubernetes.io](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/)

Kubernetes는 컨테이너화된 워크로드와 서비스를 관리하기 위한 이식 가능하고 확장 가능한 오픈소스 플랫폼으로, 선언적 구성과 자동화를 모두 용이하게 합니다. 빠르게 성장하는 대규모 에코시스템이 있습니다. Kubernetes 서비스, 지원, 도구는 널리 이용 가능합니다.

위의 인용문에서 주목해야 할 중요한 점은 Kubernetes는 클라우드 네이티브 컴퓨팅 재단(CNCF)에 프로젝트를 기부한 Google로 거슬러 올라가는 풍부한 역사를 가진 오픈소스이며, 현재 오픈소스 커뮤니티와 대기업 벤더가 오늘날의 Kubernetes를 만드는 데 기여하면서 발전해 왔다는 점입니다.

위에서 컨테이너가 훌륭하다고 말씀드렸고 이전 섹션에서는 컨테이너와 컨테이너 이미지가 어떻게 클라우드 네이티브 시스템의 채택을 변화시키고 가속화했는지에 대해 이야기했습니다. 하지만 컨테이너만으로는 애플리케이션에 필요한 프로덕션 지원 환경을 제공할 수 없습니다. Kubernetes는 다음과 같은 이점을 제공합니다:

- **Services discovery 및 로드 밸런싱** Kubernetes는 DNS 이름 또는 IP 주소를 사용하여 컨테이너를 노출할 수 있습니다. 컨테이너에 대한 트래픽이 많을 경우, Kubernetes는 네트워크 트래픽을 로드 밸런싱하고 분산하여 배포가 안정적으로 이루어지도록 할 수 있습니다.

- **스토리지 오케스트레이션** Kubernetes를 사용하면 로컬 스토리지, 퍼블릭 클라우드 제공자 등 원하는 스토리지 시스템을 자동으로 마운트할 수 있습니다.

- **자동화된 롤아웃 및 롤백** 배포된 컨테이너에 대해 원하는 상태를 설명할 수 있으며, 제어된 속도로 실제 상태를 원하는 상태로 변경할 수 있습니다. 예를 들어, 배포를 위한 새 컨테이너를 생성하고, 기존 컨테이너를 제거하고, 모든 리소스를 새 컨테이너에 적용하도록 Kubernetes를 자동화할 수 있습니다.

- **자동 bin 패킹** 컨테이너화된 작업을 실행하는 데 사용할 수 있는 노드 클러스터를 Kubernetes에 제공하고, 각 컨테이너에 필요한 CPU와 메모리(RAM)의 양을 Kubernetes에 알려줍니다. Kubernetes는 리소스를 최대한 활용하기 위해 컨테이너를 노드에 맞출 수 있습니다.

- **자가 복구** Kubernetes는 장애가 발생한 컨테이너를 다시 시작하고, 컨테이너를 교체하고, 사용자 정의 상태 확인에 응답하지 않는 컨테이너를 죽이고, 제공할 준비가 될 때까지 클라이언트에게 알리지 않습니다.

- **비밀 및 구성 관리** Kubernetes를 사용하면 비밀번호, OAuth 토큰, SSH 키와 같은 민감한 정보를 저장하고 관리할 수 있습니다. 컨테이너 이미지를 다시 빌드하거나 스택 구성에 시크릿을 노출하지 않고도 시크릿 및 애플리케이션 구성을 배포하고 업데이트할 수 있습니다.

Kubernetes는 분산 시스템을 탄력적으로 실행할 수 있는 프레임워크를 제공합니다.

컨테이너 오케스트레이션은 컨테이너의 배포, 배치 및 라이프사이클을 관리합니다.

또한 다른 많은 책임도 있습니다:

- 클러스터 관리는 호스트를 하나의 대상으로 묶습니다.

- 스케줄 관리는 스케줄러를 통해 컨테이너를 노드 간에 배포합니다.

- Services discovery은 컨테이너의 위치를 파악하고 클라이언트 요청을 컨테이너에 분산합니다.

- 복제는 요청된 워크로드에 적합한 수의 노드와 컨테이너를 사용할 수 있도록 보장합니다.

- 상태 관리는 건강하지 않은 컨테이너와 노드를 감지하고 교체합니다.

### 주요 Kubernetes 구성 요소

Kubernetes는 애플리케이션을 프로비저닝, 관리 및 확장하기 위한 컨테이너 오케스트레이터입니다. 이를 사용하여 VM 또는 물리적 머신과 같은 워커 머신의 모음인 노드 클러스터에서 컨테이너화된 앱의 라이프사이클을 관리할 수 있습니다.

앱을 실행하려면 데이터베이스 연결, 방화벽 백엔드와의 통신, 키 보안에 도움이 되는 볼륨, 네트워크, 시크릿 등 다른 많은 리소스가 필요할 수 있습니다. Kubernetes를 사용하면 이러한 리소스를 앱에 추가할 수 있습니다. 앱에 필요한 인프라 리소스는 선언적으로 관리됩니다.

Kubernetes의 핵심 패러다임은 선언적 모델입니다. 사용자가 원하는 상태를 제공하면 Kubernetes가 이를 실현합니다. 인스턴스 다섯 개가 필요한 경우, 사용자가 직접 다섯 개의 인스턴스를 시작하지 않습니다. 대신 인스턴스 5개가 필요하다고 Kubernetes에 알려주면 Kubernetes가 자동으로 상태를 조정합니다. 인스턴스 중 하나에 문제가 발생하여 실패하더라도 Kubernetes는 여전히 사용자가 원하는 상태를 파악하고 사용 가능한 노드에 인스턴스를 생성합니다.

### 노드

#### 컨트롤 플레인

모든 Kubernetes 클러스터에는 컨트롤 플레인 노드가 필요하며, 컨트롤 플레인의 구성 요소는 클러스터에 대한 전역 결정(예: 스케줄링)을 내리고 클러스터 이벤트를 감지 및 응답합니다.

![](/2022/Days/Images/Day49_Kubernetes1.png)

#### 워커 노드

Kubernetes 워크로드를 실행하는 워커 머신입니다. 물리적(베어메탈) 머신이거나 가상 머신(VM)일 수 있습니다. 각 노드는 하나 이상의 pod를 호스트할 수 있습니다. Kubernetes 노드는 컨트롤 플레인에 의해 관리됩니다.

![](/2022/Days/Images/Day49_Kubernetes2.png)

다른 노드 유형이 있지만 여기서는 다루지 않겠습니다.

#### kubelet

클러스터의 각 노드에서 실행되는 에이전트입니다. 컨테이너가 pod에서 실행되고 있는지 확인합니다.

kubelet은 다양한 메커니즘을 통해 제공되는 일련의 PodSpec을 가져와서 해당 PodSpec에 설명된 컨테이너가 실행 중이고 정상인지 확인합니다. kubelet은 Kubernetes가 생성하지 않은 컨테이너는 관리하지 않습니다.

![](/2022/Days/Images/Day49_Kubernetes3.png)

#### kube-proxy

kube-proxy는 클러스터의 각 노드에서 실행되는 네트워크 프록시로, Kubernetes Services 개념의 일부를 구현합니다.

kube-proxy는 노드에서 네트워크 규칙을 유지 관리합니다. 이러한 네트워크 규칙은 클러스터 내부 또는 외부의 네트워크 세션에서 pod로의 네트워크 통신을 허용합니다.

운영 체제 패킷 필터링 계층이 있고 사용 가능한 경우, kube-proxy는 이를 사용합니다. 그렇지 않으면, kube-proxy는 트래픽 자체를 전달합니다.

![](/2022/Days/Images/Day49_Kubernetes4.png)

#### 컨테이너 런타임

컨테이너 런타임은 컨테이너 실행을 담당하는 소프트웨어입니다.

Kubernetes는 여러 컨테이너 런타임을 지원합니다: docker, containerd, CRI-O 그리고 Kubernetes CRI(Container Runtime Interface)의 모든 구현

![](/2022/Days/Images/Day49_Kubernetes5.png)

### 클러스터

클러스터는 노드의 그룹으로, 노드는 물리적 머신 또는 가상 머신이 될 수 있습니다. 각 노드에는 컨테이너 런타임(Docker)이 있으며 마스터 컨트롤러(나중에 자세히 설명)의 명령을 받는 에이전트인 kubelet 서비스와 다른 구성 요소(나중에 자세히 설명)에서 pod에 대한 연결을 프록시하는 데 사용되는 프록시도 실행됩니다.

고가용성으로 만들 수 있는 컨트롤 플레인에는 워커 노드와 비교하여 몇 가지 고유한 역할이 포함되며, 가장 중요한 것은 정보를 가져오거나 Kubernetes 클러스터로 정보를 푸시하기 위한 모든 통신이 이루어지는 곳인 kube API 서버입니다.

#### Kube API 서버

Kubernetes API 서버는 pod, 서비스, 응답 컨트롤러 등을 포함하는 API 오브젝트에 대한 데이터의 유효성을 검사하고 구성합니다. API 서버는 REST 작업을 수행하고 다른 모든 구성 요소가 상호 작용하는 클러스터의 공유 상태에 대한 프론트엔드를 제공합니다.

#### 스케줄러

Kubernetes 스케줄러는 노드에 pod를 할당하는 컨트롤 플레인 프로세스입니다. 스케줄러는 제약 조건과 사용 가능한 리소스에 따라 스케줄링 대기열에서 각 pod에 대해 유효한 배치가 되는 노드를 결정합니다. 그런 다음 스케줄러는 각 유효한 노드의 순위를 매기고 pod를 적합한 노드에 바인딩합니다.

#### 컨트롤러 매니저

Kubernetes 컨트롤러 매니저는 Kubernetes와 함께 제공되는 핵심 제어 루프를 임베드하는 daemon입니다. 로보틱스 및 자동화 애플리케이션에서 제어 루프는 시스템 상태를 조절하는 비종료 루프입니다. Kubernetes에서 컨트롤러는 API 서버를 통해 클러스터의 공유 상태를 감시하고 현재 상태를 원하는 상태로 이동시키기 위해 변경을 시도하는 제어 루프입니다.

#### etcd

모든 클러스터 데이터에 대한 Kubernetes의 백업 저장소로 사용되는 일관되고 가용성이 높은 키 값 저장소입니다.

![](/2022/Days/Images/Day49_Kubernetes6.png)

#### kubectl

CLI 관점에서 이를 관리하기 위해 kubectl이 있으며, kubectl은 API 서버와 상호 작용합니다.

Kubernetes 커맨드-라인 도구인 kubectl을 사용하면 Kubernetes 클러스터에 대해 명령을 실행할 수 있습니다. kubectl을 사용하여 애플리케이션을 배포하고, 클러스터 리소스를 검사 및 관리하고, 로그를 볼 수 있습니다.

![](/2022/Days/Images/Day49_Kubernetes7.png)

### Pods

pod는 논리적 애플리케이션을 구성하는 컨테이너 그룹입니다. 예를 들어, NodeJS 컨테이너와 MySQL 컨테이너를 실행하는 웹 애플리케이션이 있는 경우, 이 두 컨테이너는 모두 단일 pod에 위치하게 됩니다. 또한 pod는 공통 데이터 볼륨을 공유할 수 있으며 동일한 네트워킹 네임스페이스도 공유합니다. pod는 임시적이며 마스터 컨트롤러에 의해 위아래로 이동될 수 있다는 것을 기억합시다. Kubernetes는 레이블(이름-값) 개념을 통해 pod를 식별하는 간단하지만, 효과적인 수단을 사용합니다.

- pod는 컨테이너의 볼륨, 시크릿 및 구성을 처리합니다.

- pod는 임시적입니다. 죽으면 자동으로 재시작되도록 설계되었습니다.

- pod는 ReplicationSet에 의해 앱이 수평으로 스케일 될 때 복사됩니다. 각 pod는 동일한 컨테이너 코드를 실행합니다.

- pod는 워커 노드에서 실행됩니다.

![](/2022/Days/Images/Day49_Kubernetes8.png)

### Deployments

- pod를 실행하기로 결정할 수 있지만 죽으면 끝입니다.

- Deployments를 사용하면 pod가 지속적으로 실행될 수 있습니다.

- Deployments를 사용하면 다운타임 없이 실행 중인 앱을 업데이트할 수 있습니다.

- 또한, Deployments는 pod가 죽었을 때 재시작하는 전략을 지정합니다.

![](/2022/Days/Images/Day49_Kubernetes9.png)

### ReplicaSets

- Deployments는 ReplicaSets을 생성할 수도 있습니다.

- ReplicaSets은 앱이 원하는 수의 pod를 갖도록 보장합니다.

- ReplicaSets은 Deployments에 따라 pod를 생성하고 확장합니다.

- Deployments, ReplicaSets, pod는 배타적이지는 않지만 그렇게 될 수 있습니다.

### StatefulSets

- 앱의 상태에 대한 정보를 유지해야 하나요?

- 데이터베이스에는 상태가 필요합니다.

- StatefulSets의 pod는 서로 호환되지 않습니다.

- 각 pod에는 컨트롤러가 모든 스케줄링에 대해 유지하는 고유하고 영구적인 식별자가 있습니다.

![](/2022/Days/Images/Day49_Kubernetes10.png)

### DaemonSets

- DaemonSets은 연속 프로세스를 위한 것입니다.

- 노드당 하나의 pod를 실행합니다.

- 클러스터에 새로운 노드가 추가될 때마다 pod가 시작됩니다.

- 모니터링 및 로그 수집과 같은 백그라운드 작업에 유용합니다.

- 각 pod에는 컨트롤러가 모든 스케줄링에 대해 유지하는 고유하고 영구적인 식별자가 있습니다.

![](/2022/Days/Images/Day49_Kubernetes11.png)

### Services

- pod에 액세스하기 위한 단일 엔드포인트입니다.

- 클러스터와 최종적으로 pod 목록으로 트래픽을 라우팅하는 통합된 방법입니다.

- Services를 사용하면 아무 영향 없이 pod를 올리고 내릴 수 있습니다.

이것은 Kubernetes의 기본 구성 요소에 대한 간략한 개요와 참고 사항일 뿐이며, 이 지식을 바탕으로 스토리지와 Ingress 관련 몇 가지 다른 영역을 추가하여 애플리케이션을 개선할 수 있지만, Kubernetes 클러스터가 실행되는 위치에 대한 선택의 폭도 넓어집니다. 다음 세션에서는 스토리지와 관련된 몇 가지 세부 사항을 살펴보면서 Kubernetes 클러스터를 실행할 수 있는 위치에 대한 이러한 옵션에 중점을 두겠습니다.

![](/2022/Days/Images/Day49_Kubernetes12.png)

### Kubernetes 시리즈에서 다룰 내용

- Kubernetes 아키텍처
- Kubectl 커맨드
- Kubernetes YAML
- Kubernetes Ingress
- Kubernetes Services
- Helm 패키지 관리자
- 영속성 스토리지
- stateful 앱

## 자료

- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [TechWorld with Nana - Kubernetes Tutorial for Beginners [FULL COURSE in 4 Hours]](https://www.youtube.com/watch?v=X48VuDVv0do)
- [TechWorld with Nana - Kubernetes Crash Course for Absolute Beginners](https://www.youtube.com/watch?v=s_o8dwzRlu4)
- [Kunal Kushwaha - Kubernetes Tutorial for Beginners | What is Kubernetes? Architecture Simplified!](https://www.youtube.com/watch?v=KVBON1lA9N8)

[Day 50](day50.md)에서 봐요!
