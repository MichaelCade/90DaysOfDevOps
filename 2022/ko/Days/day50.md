---
title: '#90DaysOfDevOps - Choosing your Kubernetes platform - Day 50'
published: false
description: 90DaysOfDevOps - Choosing your Kubernetes platform
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049046
---

## Kubernetes 플랫폼 선택하기

이 세션을 통해 몇 가지 플랫폼 또는 배포판이라는 용어를 사용하는 것이 더 적합할 수도 있는데, Kubernetes 세계에서 도전 과제 중 하나는 복잡성을 제거하는 것입니다.

Kubernetes 어려운 길은 아무것도 없는 상태에서 완전한 기능을 갖춘 Kubernetes 클러스터로 구축하는 방법을 안내하지만, 적어도 제가 이야기하는 사람들은 점점 더 많은 사람들이 이러한 복잡성을 제거하고 관리형 Kubernetes 클러스터를 실행하기를 원하고 있습니다. 문제는 비용이 더 많이 들지만, 관리형 서비스를 사용하면 기본 노드 아키텍처와 컨트롤 플레인 노드 관점에서 무슨 일이 일어나고 있는지 알아야 하지만 일반적으로 이에 액세스할 수 없는 경우 이점을 누릴 수 있다는 것입니다.

그런 다음 시스템을 사용할 수 있는 로컬 개발 배포판이 있고, 개발자가 의도한 플랫폼에서 앱을 실행할 수 있는 완전한 작업 환경을 갖출 수 있도록 로컬 버전의 Kubernetes를 실행할 수 있습니다.

이 모든 개념의 일반적인 기본은 모두 Kubernetes의 한 종류이므로 요구 사항에 맞게 워크로드를 필요한 곳으로 자유롭게 마이그레이션하고 이동할 수 있어야 한다는 것입니다.

또한 어떤 투자가 이루어졌는지에 따라 많은 선택이 달라질 것입니다. 개발자 경험에 대해서도 언급했지만, 노트북에서 실행되는 로컬 Kubernetes 환경 중 일부는 비용을 들이지 않고도 기술을 익힐 수 있는 훌륭한 환경입니다.

### 베어 메탈 클러스터

많은 사람들이 클러스터를 생성하기 위해 여러 대의 물리적 서버에서 바로 Linux OS를 실행하는 옵션을 선택할 수 있으며, Windows일 수도 있지만 Windows, 컨테이너 및 Kubernetes와 관련된 채택률에 대해서는 많이 듣지 못했습니다. 만약 여러분이 기업이고 물리적 서버를 구매하기로 CAPEX 결정을 내렸다면, Kubernetes 클러스터를 구축할 때 이 방법을 사용할 수 있지만, 관리 및 관리자 측면에서는 여러분이 직접 구축하고 모든 것을 처음부터 관리해야 한다는 것을 의미합니다.

### 가상화

테스트 및 학습 환경이나 엔터프라이즈급 Kubernetes 클러스터에 관계없이 가상화는 일반적으로 가상 머신을 스핀업하여 노드 역할을 하도록 한 다음 함께 클러스터링할 수 있는 훌륭한 방법입니다. 가상화의 기본 아키텍처, 효율성 및 속도를 활용할 수 있을 뿐만 아니라 기존 지출을 활용할 수 있습니다. 예를 들어 VMware는 가상 머신과 Kubernetes 모두를 위한 훌륭한 솔루션을 다양한 형태로 제공합니다.

제가 처음으로 구축한 Kubernetes 클러스터는 몇 개의 VM을 노드로 실행할 수 있는 오래된 서버에서 Microsoft Hyper-V를 사용한 가상화를 기반으로 구축되었습니다.

### 로컬 데스크톱 옵션

데스크톱이나 노트북에서 로컬 Kubernetes 클러스터를 실행하는 데는 몇 가지 옵션이 있습니다. 앞서 말했듯이 개발자는 비용이 많이 들거나 복잡한 클러스터를 여러 개 보유하지 않고도 앱이 어떻게 보일지 확인할 수 있습니다. 개인적으로 저는 이 클러스터를 많이 사용해 왔으며 특히 Minikube를 사용해 왔습니다. 여기에는 무언가를 시작하고 실행하는 방식을 바꾸는 몇 가지 훌륭한 기능과 애드온이 있습니다.

### Kubernetes Managed Services

가상화에 대해 언급했는데, 이는 로컬에서 하이퍼바이저를 통해 달성할 수 있지만 이전 섹션에서 퍼블릭 클라우드의 가상 머신을 활용하여 노드 역할을 할 수도 있다는 것을 알고 있습니다. 여기서 말하는 Kubernetes 관리형 서비스란 대규모 하이퍼스케일러뿐만 아니라 최종 사용자로부터 관리 및 제어 계층을 제거하여 최종 사용자로부터 제어 플레인을 제거하는 MSP에서 제공하는 서비스를 말하며, 이는 Amazon EKS, Microsoft AKS 및 Google Kubernetes Engine에서 일어나는 일입니다. (GKE)

### 압도적인 선택

선택의 폭이 넓다는 것은 좋지만, 위에 나열된 각 카테고리의 모든 옵션에 대해 자세히 살펴본 것은 아닙니다. 위의 옵션 외에도 Red Hat의 OpenShift가 있으며, 이 옵션은 위의 모든 주요 클라우드 제공업체에서 실행할 수 있으며 클러스터가 배포된 위치에 관계없이 관리자에게 최고의 전반적인 사용성을 제공할 수 있습니다.

제가 가상화 경로로 시작했다고 말씀드렸지만, 이는 제가 목적에 맞게 사용할 수 있는 물리적 서버에 액세스할 수 있었기 때문에 감사하게도 그 이후로는 더 이상 이 옵션을 사용할 수 없었습니다.

지금 제가 드리고 싶은 조언은 Minikube를 첫 번째 옵션으로 사용하거나 Kind(Docker의 Kubernetes)를 사용하라는 것이지만, Minikube는 애드온을 사용하고 빠르게 구축한 다음 완료되면 날려버릴 수 있고, 여러 클러스터를 실행할 수 있으며, 거의 모든 곳에서 실행할 수 있고, 크로스 플랫폼 및 하드웨어에 구애받지 않기 때문에 복잡성을 거의 추상화할 수 있는 몇 가지 추가적인 이점을 제공합니다.

저는 Kubernetes에 대해 배우면서 약간의 여정을 거쳤기 때문에 플랫폼 선택과 구체적인 내용은 여기서는 플랫폼인 Kubernetes와 실행 가능한 위치에 대한 이해를 돕기 위해 시도했던 옵션들을 나열해 보겠습니다. 아래 블로그 포스팅을 다시 한번 살펴보고 블로그 게시물에 링크되어 있는 것보다 여기에 더 많이 소개할 수 있도록 하겠습니다.

- [Kubernetes playground – How to choose your platform](https://vzilla.co.uk/vzilla-blog/building-the-home-lab-kubernetes-playground-part-1)
- [Kubernetes playground – Setting up your cluster](https://vzilla.co.uk/vzilla-blog/building-the-home-lab-kubernetes-playground-part-2)
- [Getting started with Amazon Elastic Kubernetes Service (Amazon EKS)](https://vzilla.co.uk/vzilla-blog/getting-started-with-amazon-elastic-kubernetes-service-amazon-eks)
- [Getting started with Microsoft Azure Kubernetes Service (AKS)](https://vzilla.co.uk/vzilla-blog/getting-started-with-microsoft-azure-kubernetes-service-aks)
- [Getting Started with Microsoft AKS – Azure PowerShell Edition](https://vzilla.co.uk/vzilla-blog/getting-started-with-microsoft-aks-azure-powershell-edition)
- [Getting started with Google Kubernetes Service (GKE)](https://vzilla.co.uk/vzilla-blog/getting-started-with-google-kubernetes-service-gke)
- [Kubernetes, How to – AWS Bottlerocket + Amazon EKS](https://vzilla.co.uk/vzilla-blog/kubernetes-how-to-aws-bottlerocket-amazon-eks)
- [Getting started with CIVO Cloud](https://vzilla.co.uk/vzilla-blog/getting-started-with-civo-cloud)
- [Minikube - Kubernetes Demo Environment For Everyone](https://vzilla.co.uk/vzilla-blog/project_pace-kasten-k10-demo-environment-for-everyone)

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

[Day 51](day51.md)에서 봐요!
