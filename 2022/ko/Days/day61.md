---
title: '#90DaysOfDevOps - Kubernetes & Multiple Environments - Day 61'
published: false
description: 90DaysOfDevOps - Kubernetes & Multiple Environments
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048743
---

## Kubernetes 및 다중 환경

지금까지 IaC에 대한 이 섹션에서는 가상 머신을 배포하는 방법을 살펴보았지만, 가상 머신의 모양을 코드에서 정의한 다음 배포한다는 전제는 실제로 동일합니다. Docker 컨테이너도 마찬가지이며, 이 세션에서는 Terraform을 사용하여 Kubernetes에서 지원하는 리소스와 상호 작용하는 방법을 살펴보겠습니다.

저는 데모 목적으로 3개의 주요 클라우드 제공업체에 Terraform을 사용하여 Kubernetes 클러스터를 배포해 왔으며, 리포지토리 [tf_k8deploy](https://github.com/MichaelCade/tf_k8deploy)를 찾을 수 있습니다.

그러나 Terraform을 사용하여 Kubernetes 클러스터 내의 객체와 상호 작용할 수도 있는데, 이는 [Kubernetes 공급자](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs)를 사용하거나 [Helm 공급자](https://registry.terraform.io/providers/hashicorp/helm/latest)를 사용하여 차트 배포를 관리할 수 있습니다.

이제 이전 섹션에서 살펴본 것처럼 `kubectl`을 사용할 수 있습니다. 하지만 Kubernetes 환경에서 Terraform을 사용하면 몇 가지 이점이 있습니다.

- 통합 workflow - 클러스터 배포에 Terraform을 사용했다면, 동일한 workflow와 도구를 사용하여 Kubernetes 클러스터 내에 배포할 수 있습니다.

- 라이프사이클 관리 - Terraform은 단순한 프로비저닝 도구가 아니라 변경, 업데이트, 삭제를 가능하게 합니다.

### 간단한 Kubernetes 데모

지난 세션에서 만든 데모와 마찬가지로 이제 Kubernetes 클러스터에 nginx를 배포할 수 있습니다. 여기서는 데모 목적으로 Minikube를 다시 사용하겠습니다. Kubernetes.tf 파일을 생성하고 [여기](2022/Days/IaC/Kubernetes/Kubernetes.tf)에서 이 파일을 찾을 수 있습니다.

이 파일에서 Kubernetes 공급자를 정의하고, kubeconfig 파일을 가리키고, nginx라는 네임스페이스를 생성한 다음, 2개의 복제본과 마지막으로 서비스를 포함하는 배포를 생성하겠습니다.

```
terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 2.0.0"
    }
  }
}
provider "kubernetes" {
  config_path = "~/.kube/config"
}
resource "kubernetes_namespace" "test" {
  metadata {
    name = "nginx"
  }
}
resource "kubernetes_deployment" "test" {
  metadata {
    name      = "nginx"
    namespace = kubernetes_namespace.test.metadata.0.name
  }
  spec {
    replicas = 2
    selector {
      match_labels = {
        app = "MyTestApp"
      }
    }
    template {
      metadata {
        labels = {
          app = "MyTestApp"
        }
      }
      spec {
        container {
          image = "nginx"
          name  = "nginx-container"
          port {
            container_port = 80
          }
        }
      }
    }
  }
}
resource "kubernetes_service" "test" {
  metadata {
    name      = "nginx"
    namespace = kubernetes_namespace.test.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment.test.spec.0.template.0.metadata.0.labels.app
    }
    type = "NodePort"
    port {
      node_port   = 30201
      port        = 80
      target_port = 80
    }
  }
}
```

새 프로젝트 폴더에서 가장 먼저 해야 할 일은 `terraform init` 명령을 실행하는 것입니다.

![](/2022/Days/Images/Day61_IAC1.png)

그리고 `terraform apply` 명령을 실행하기 전에 네임스페이스가 없다는 것을 보여드리겠습니다.

![](/2022/Days/Images/Day61_IAC2.png)

apply 명령을 실행하면 Kubernetes 클러스터 내에 3개의 새로운 리소스, 네임스페이스, 배포 및 서비스가 생성됩니다.

![](/2022/Days/Images/Day61_IAC3.png)

이제 클러스터 내에 배포된 리소스를 살펴볼 수 있습니다.

![](/2022/Days/Images/Day61_IAC4.png)

이전 섹션에서 보셨듯이 Minikube를 사용하고 있기 때문에 도커 네트워킹으로 인그레스를 시도할 때 한계가 있습니다. 하지만 `kubectl port-forward -n nginx svc/nginx 30201:80` 명령을 실행하고 `http://localhost:30201/`으로 브라우저를 열면 NGINX 페이지를 볼 수 있습니다.

![](/2022/Days/Images/Day61_IAC5.png)

Terraform과 Kubernetes에 대해 더 자세한 데모를 해보고 싶으시다면 [HashiCorp 학습 사이트](https://learn.hashicorp.com/tutorials/terraform/kubernetes-provider)를 방문해 보시기 바랍니다.

### 여러 환경

우리가 실행한 데모를 이제 특정 프로덕션, 스테이징 및 개발 환경이 동일하게 보이고 이 코드를 활용하기를 원한다면 Terraform을 사용하여 이를 달성하는 두 가지 접근 방식이 있습니다.

- `terraform workspaces` - 단일 백엔드 내의 여러 개의 명명된 섹션

- 파일 구조 - 디렉토리 레이아웃은 분리를 제공하고, 모듈은 재사용을 제공합니다.

하지만 위의 각 방법에는 장단점이 있습니다.

### Terraform workspaces

장점

- 쉬운 시작
- 편리한 terraform.workspace 표현식
- 코드 중복 최소화

단점

- 인적 오류가 발생하기 쉬움(TF를 사용하여 이를 제거하려고 노력 중임)
- 동일한 백엔드 내에 저장된 상태
- 코드베이스가 배포 구성을 명확하게 보여주지 않음.

### 파일 구조

장점

- 백엔드 격리
  - 보안 향상
  - 인적 오류 가능성 감소
- 배포된 상태를 완벽하게 나타내는 코드베이스

단점

- 프로비저닝 환경에 여러 Terraform 적용 필요
- 코드 중복이 더 많지만, 모듈을 사용하여 최소화할 수 있습니다.

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

[Day 62](day62.md)에서 봐요!
