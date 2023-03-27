---
title: '#90DaysOfDevOps - Deploying your first Kubernetes Cluster - Day 51'
published: false
description: 90DaysOfDevOps - Deploying your first Kubernetes Cluster
tags: 'DevOps, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048778
---

## 첫 번째 Kubernetes 클러스터 배포하기

이 글에서는 Minikube를 사용하여 로컬 머신에서 Kubernetes 클러스터를 시작하고 실행해 보겠습니다. 이렇게 하면 나머지 Kubernetes 섹션을 위한 기본 Kubernetes 클러스터가 제공되지만, 나중에 VirtualBox에서도 Kubernetes 클러스터를 배포하는 방법을 살펴볼 것입니다. 퍼블릭 클라우드에서 관리형 Kubernetes 클러스터를 스핀업하는 대신 이 방법을 선택한 이유는 무료 티어를 사용하더라도 비용이 들기 때문이며, 이전 섹션 [Day 50](day50.md)에서 해당 환경을 스핀업하려는 경우 몇 가지 블로그를 공유한 바 있습니다.

### Minikube란?

> "Minikube는 macOS, Linux, Windows에서 로컬 Kubernetes 클러스터를 빠르게 설정합니다. 우리는 애플리케이션 개발자와 새로운 Kubernetes 사용자를 돕는 데 주력하고 있습니다."

위에 해당되지 않을 수도 있지만, 저는 Minikube가 Kubernetes 방식으로 무언가를 테스트하고 싶을 때 앱을 쉽게 배포할 수 있고 몇 가지 놀라운 애드온이 있다는 것을 알게 되었으며, 이 글에서도 다룰 것입니다.

우선, 워크스테이션 OS에 관계없이 Minikube를 실행할 수 있습니다. 먼저 [프로젝트 페이지](https://minikube.sigs.k8s.io/docs/start/)로 이동합니다. 첫 번째 옵션은 설치 방법을 선택하는 것입니다. 저는 이 방법을 사용하지 않았지만, 여러분은 제가 사용하는 방법과 다른 방법을 선택할 수 있습니다(곧 소개할 예정입니다).

아래에 언급된 "Container or virtual machine managers, such as Docker, Hyper kit, Hyper-V, KVM, Parallels, Podman, VirtualBox, or VMware"가 있어야 한다고 명시되어 있는데, 이것이 Minikube가 실행되는 곳이고 쉬운 옵션이며 저장소에 명시되어 있지 않은 한 저는 Docker를 사용하고 있습니다. 그리고 [여기](https://docs.docker.com/get-docker/)에서 시스템에 Docker를 설치할 수 있습니다.

![](/2022/Days/Images/Day51_Kubernetes1.png)

### Minikube 및 기타 전제조건을 설치하는 방법...

저는 한동안 arkade를 사용하여 모든 Kubernetes 도구와 CLI를 설치해왔는데, arkade를 시작하기 위한 설치 단계는 이 [github 저장소](https://github.com/alexellis/arkade)에서 확인할 수 있습니다. 설치가 필요한 다른 블로그 게시물에서도 언급했습니다. arkade get을 누른 다음 툴이나 CLI를 사용할 수 있는지 확인하는 것만으로도 간단합니다. 리눅스 섹션에서 패키지 관리자와 소프트웨어를 얻는 프로세스에 대해 이야기했는데, arkade는 모든 앱과 Kubernetes용 CLI를 위한 마켓플레이스라고 생각하시면 됩니다. 시스템에서 사용할 수 있는 매우 편리한 작은 도구로, GO로 작성되어 크로스 플랫폼을 지원합니다.

![](/2022/Days/Images/Day51_Kubernetes2.png)

arkade 내에서 사용 가능한 긴 앱 목록의 일부로 Minikube도 그중 하나이므로 간단한 `arkade get minikube` 명령으로 바이너리를 다운로드하고 이제 바로 사용할 수 있습니다.

![](/2022/Days/Images/Day51_Kubernetes3.png)

또한 도구의 일부로 kubectl이 필요하므로 arkade를 통해서도 얻을 수 있으며, 위에서 언급한 curl 명령의 일부로 Minikube 문서에 나와 있다고 생각합니다. 이 포스트의 뒷부분에서 kubectl에 대해 더 자세히 다루겠습니다.

### Kubernetes 클러스터 시작 및 실행하기

이 특정 섹션에서는 로컬 머신에서 Kubernetes 클러스터를 시작하고 실행할 때 사용할 수 있는 옵션에 대해 다루고자 합니다. 다음 명령을 실행하기만 하면 사용할 수 있는 클러스터가 스핀업됩니다.

커맨드라인에 minikube가 사용되며, 모든 설치가 완료되면 `minikube start`를 실행하여 첫 번째 Kubernetes 클러스터를 배포할 수 있습니다. 아래에서 중첩된 가상화 노드를 실행할 위치에 대한 기본값이 Docker 드라이버인 것을 확인할 수 있습니다. 게시물의 시작 부분에서 사용 가능한 다른 옵션에 대해 언급했는데, 다른 옵션은 이 로컬 Kubernetes 클러스터의 모양을 확장하고자 할 때 도움이 됩니다.

이 인스턴스에서는 단일 Minikube 클러스터가 단일 docker 컨테이너로 구성되며, 이 컨테이너에는 컨트롤 플레인 노드와 워커 노드가 하나의 인스턴스에 포함됩니다. 일반적으로는 이러한 노드를 분리합니다. 다음 섹션에서는 아직 홈 랩 유형의 Kubernetes 환경이지만 프로덕션 아키텍처에 조금 더 가까운 환경을 살펴볼 것입니다.

![](/2022/Days/Images/Day51_Kubernetes4.png)

지금까지 몇 번 언급했지만, 저는 사용 가능한 애드온 때문에 Minikube를 좋아하는데, 처음부터 필요한 모든 애드온을 포함한 간단한 명령으로 클러스터를 배포할 수 있기 때문에 매번 동일한 설정을 배포하는 데 도움이 됩니다.

아래에서 이러한 애드온 목록을 볼 수 있는데, 저는 일반적으로 `CSI-host path-driver`와 `volumesnapshots` 애드온을 사용하지만, 아래에서 긴 목록을 볼 수 있습니다. 물론 이러한 애드온은 나중에 Kubernetes 섹션에서 다루겠지만 일반적으로 Helm을 사용하여 배포할 수 있지만 훨씬 더 간단해진다.

![](/2022/Days/Images/Day51_Kubernetes5.png)

또한 프로젝트에서 몇 가지 추가 구성을 정의하고 있는데, apiserver는 임의의 API 포트 대신 6433으로 설정하고 container runtime도 containerd로 정의하고 있지만 기본값은 docker이고 CRI-O도 사용할 수 있습니다. 또한 특정 Kubernetes 버전도 설정하고 있습니다.

![](/2022/Days/Images/Day51_Kubernetes6.png)

이제 Minikube를 사용하여 첫 번째 Kubernetes 클러스터를 배포할 준비가 되었습니다. 클러스터와 상호 작용하려면 `kubectl`도 필요하다고 앞서 언급했습니다. arkade를 사용하여 `arkade get kubectl` 명령으로 kubectl을 설치할 수 있습니다.

![](/2022/Days/Images/Day51_Kubernetes7.png)

또는 다음에서 크로스 플랫폼으로 다운로드할 수 있습니다.

- [리눅스](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux)
- [macOS](https://kubernetes.io/docs/tasks/tools/install-kubectl-macos)
- [윈도우](https://kubernetes.io/docs/tasks/tools/install-kubectl-windows)

kubectl을 설치했으면 `kubectl get nodes`와 같은 간단한 명령으로 클러스터와 상호 작용할 수 있습니다.

![](/2022/Days/Images/Day51_Kubernetes8.png)

### kubectl이란?

이제 Minikube와 Kubernetes 클러스터를 모두 설치하고 실행 중이며, Minikube에 대해서는 최소한 Minikube가 무엇을 하는지에 대해 설명했지만, kubectl이 무엇이고 어떤 역할을 하는지에 대해서는 설명하지 않았습니다.

kubectl은 Kubernetes 클러스터와 상호 작용하는 데 사용되거나 상호 작용할 수 있게 해주는 클리어로, 여기서는 Minikube 클러스터와 상호 작용하는 데 사용하고 있지만 퍼블릭 클라우드 전반의 엔터프라이즈 클러스터와 상호 작용하는 데도 kubectl을 사용할 수 있습니다.

우리는 애플리케이션을 배포하고 클러스터 리소스를 검사 및 관리하기 위해 kubectl을 사용합니다. 훨씬 더 자세한 개요는 Kubernetes [공식 문서](https://kubernetes.io/docs/reference/kubectl/overview/)에서 확인할 수 있습니다.

kubectl은 이전 포스트에서 간략하게 다룬 컨트롤 플레인 노드에 있는 API 서버와 상호작용합니다.

### kubectl 치트 시트

공식 문서와 함께, 필자는 kubectl 명령어를 찾을 때 [이 페이지](https://unofficial-kubernetes.readthedocs.io/en/latest/)를 항상 열어두는 것을 추천합니다.

| Listing Resources        |                                            |
| ------------------------ | ------------------------------------------ |
| kubectl get nodes        | 클러스터의 모든 노드 나열                  |
| kubectl get namespaces   | 클러스터의 모든 네임스페이스 나열          |
| kubectl get pods         | 기본 네임스페이스 클러스터에 모든 pod 나열 |
| kubectl get pods -n name | "이름" 네임스페이스에 모든 pod를 나열      |

| Creating Resources            |                                            |
| ----------------------------- | ------------------------------------------ |
| kubectl create namespace name | "name"이라는 네임스페이스를 생성           |
| kubectl create -f [filename]  | JSON 또는 YAML 파일에서 리소스를 다시 생성 |

| Editing Resources            |               |
| ---------------------------- | ------------- |
| kubectl edit svc/servicename | 서비스를 편집 |

| More detail on Resources |                                       |
| ------------------------ | ------------------------------------- |
| kubectl describe nodes   | 원하는 수의 리소스 상태를 자세히 표시 |

| Delete Resources   |                                                           |
| ------------------ | --------------------------------------------------------- |
| kubectl delete pod | 리소스를 제거할 수 있으며, 이는 stdin 또는 파일에서 제거. |

예를 들어 `-n`은 `namespace`의 줄임말로, 명령을 입력하기 쉬울 뿐만 아니라 스크립트를 작성할 때 훨씬 더 깔끔한 코드를 만들 수 있습니다.

| Short name | Full name                  |
| ---------- | -------------------------- |
| csr        | certificatesigningrequests |
| cs         | componentstatuses          |
| cm         | configmaps                 |
| ds         | daemonsets                 |
| deploy     | deployments                |
| ep         | endpoints                  |
| ev         | events                     |
| hpa        | horizontalpodautoscalers   |
| ing        | ingresses                  |
| limits     | limitranges                |
| ns         | namespaces                 |
| no         | nodes                      |
| pvc        | persistentvolumeclaims     |
| pv         | persistentvolumes          |
| po         | pods                       |
| pdb        | poddisruptionbudgets       |
| psp        | podsecuritypolicies        |
| rs         | replicasets                |
| rc         | replicationcontrollers     |
| quota      | resourcequotas             |
| sa         | serviceaccounts            |
| svc        | services                   |

마지막으로 추가하고 싶은 것은 데이터 서비스를 표시하기 위해 데모 환경을 빠르게 스핀업하고 Kasten K10으로 이러한 워크로드를 보호하기 위해 Minikube와 관련된 또 다른 프로젝트를 만들었습니다. [Project Pace](https://github.com/MichaelCade/project_pace)는 여기에서 찾을 수 있으며 여러분의 피드백이나 상호 작용을 원하며 Minikube 클러스터를 배포하고 다양한 데이터 서비스 애플리케이션을 만드는 몇 가지 자동화된 방법을 표시하거나 포함합니다.

다음에는 VirtualBox를 사용하여 여러 노드를 가상 머신에 배포하는 방법을 살펴보겠지만, Linux 섹션에서 vagrant를 사용하여 머신을 빠르게 스핀업하고 원하는 방식으로 소프트웨어를 배포했던 것처럼 여기에서도 쉽게 진행할 것입니다.

어제 포스트에 배포 중인 다양한 Kubernetes 클러스터에 대해 제가 수행한 워크스루 블로그인 이 목록을 추가했습니다.

- [Kubernetes playground – How to choose your platform](https://vzilla.co.uk/vzilla-blog/building-the-home-lab-kubernetes-playground-part-1)
- [Kubernetes playground – Setting up your cluster](https://vzilla.co.uk/vzilla-blog/building-the-home-lab-kubernetes-playground-part-2)
- [Getting started with Amazon Elastic Kubernetes Service (Amazon EKS)](https://vzilla.co.uk/vzilla-blog/getting-started-with-amazon-elastic-kubernetes-service-amazon-eks)
- [Getting started with Microsoft Azure Kubernetes Service (AKS)](https://vzilla.co.uk/vzilla-blog/getting-started-with-microsoft-azure-kubernetes-service-aks)
- [Getting Started with Microsoft AKS – Azure PowerShell Edition](https://vzilla.co.uk/vzilla-blog/getting-started-with-microsoft-aks-azure-powershell-edition)
- [Getting started with Google Kubernetes Service (GKE)](https://vzilla.co.uk/vzilla-blog/getting-started-with-google-kubernetes-service-gke)
- [Kubernetes, How to – AWS Bottlerocket + Amazon EKS](https://vzilla.co.uk/vzilla-blog/kubernetes-how-to-aws-bottlerocket-amazon-eks)
- [Getting started with CIVO Cloud](https://vzilla.co.uk/vzilla-blog/getting-started-with-civo-cloud)
- [Minikube - Kubernetes Demo Environment For Everyone](https://vzilla.co.uk/vzilla-blog/project_pace-kasten-k10-demo-environment-for-everyone)
- [Minikube - Deploy Minikube Using Vagrant and Ansible on VirtualBox](https://medium.com/techbeatly/deploy-minikube-using-vagrant-and-ansible-on-virtualbox-infrastructure-as-code-2baf98188847)

### Kubernetes 시리즈에서 다룰 내용

아래에 언급된 내용 중 일부를 다루기 시작했지만, 내일 두 번째 클러스터 배포를 통해 더 많은 실습을 한 후 클러스터에 애플리케이션 배포를 시작할 수 있습니다.

- Kubernetes 아키텍처
- Kubectl 커맨드
- Kubernetes YAML
- Kubernetes Ingress
- Kubernetes Services
- Helm 패키지 관리자
- 영속성 스토리지
- stateful 앱

## 자료

사용하신 무료 리소스가 있다면 리포지토리에 PR을 통해 여기에 추가해 주시면 기꺼이 포함시켜드리겠습니다.

- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [TechWorld with Nana - Kubernetes Tutorial for Beginners [FULL COURSE in 4 Hours]](https://www.youtube.com/watch?v=X48VuDVv0do)
- [TechWorld with Nana - Kubernetes Crash Course for Absolute Beginners](https://www.youtube.com/watch?v=s_o8dwzRlu4)
- [Kunal Kushwaha - Kubernetes Tutorial for Beginners | What is Kubernetes? Architecture Simplified!](https://www.youtube.com/watch?v=KVBON1lA9N8)
- [Techbeatly - Deploy Minikube Using Vagrant and Ansible on VirtualBox](https://www.youtube.com/watch?v=xPLQqHbp9BM&t=371s)

[Day 52](day52.md)에서 봐요!
