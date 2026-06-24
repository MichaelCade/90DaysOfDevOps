---
title: '#90DaysOfDevOps - Hands-On Monitoring Tools - Day 78'
published: false
description: 90DaysOfDevOps - Hands-On Monitoring Tools
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049056
---

## 실습용 모니터링 도구

지난 세션에서 모니터링의 큰 그림에 대해 이야기하면서 Nagios를 살펴본 이유는 두 가지였습니다. 첫 번째는 수년 동안 많이 들어본 소프트웨어이기 때문에 그 기능에 대해 조금 더 알고 싶었기 때문입니다.

오늘은 Prometheus에 대해 살펴보려고 하는데요, 클라우드 네이티브 환경에서 점점 더 많이 사용되고 있지만 Kubernetes 등 외부에서도 이러한 물리적 리소스를 관리하는 데 사용할 수 있습니다.

### Prometheus - 거의 모든 것을 모니터링합니다.

우선, Prometheus는 컨테이너와 마이크로서비스 기반 시스템뿐만 아니라 물리적, 가상 및 기타 서비스를 모니터링하는 데 도움이 되는 오픈 소스입니다. Prometheus 뒤에는 대규모 커뮤니티가 있습니다.

Prometheus에는 다양한 [통합 및 exporter](https://prometheus.io/docs/instrumenting/exporters/)가 있습니다. 핵심은 기존 메트릭을 Prometheus 메트릭으로 내보내는 것입니다. 또한 여러 프로그래밍 언어도 지원합니다.

pull 접근 방식 - 수천 개의 마이크로서비스 또는 시스템 및 서비스와 대화하는 경우 push 방식은 일반적으로 서비스가 모니터링 시스템으로 push되는 것을 볼 수 있는 곳입니다. 이 경우 네트워크 과부하, 높은 CPU 사용량, 단일 장애 지점 등의 문제가 발생합니다. pull 방식은 모든 서비스의 메트릭 엔드포인트에서 Prometheus가 끌어오는 훨씬 더 나은 경험을 제공합니다.

다시 한번 Prometheus의 구성을 위한 YAML을 살펴봅니다.

![](/2022/Days/Images/Day78_Monitoring7.png)

나중에 이것이 Kubernetes에 배포되었을 때 어떻게 보이는지, 특히 작업/exporter로부터 메트릭을 가져오는 **PushGateway**가 있는 것을 보게 될 것입니다.

알림을 push하는 **AlertManager**가 있으며, 여기에서 이메일, 슬랙 및 기타 도구와 같은 외부 서비스와 통합할 수 있습니다.

그런 다음 PushGateway에서 이러한 pull 메트릭의 검색을 관리하고 push 알림을 AlertManager로 전송하는 Prometheus Server가 있습니다. Prometheus Server는 또한 로컬 디스크에 데이터를 저장합니다. 원격 스토리지 솔루션을 활용할 수도 있습니다.

또한 메트릭과 상호 작용하는 데 사용되는 언어인 PromQL도 있는데, 이는 나중에 Prometheus 웹 UI에서 볼 수 있지만, 이 섹션의 뒷부분에서 Grafana와 같은 데이터 시각화 도구에서도 어떻게 사용되는지 볼 수 있습니다.

### Prometheus 배포 방법

Prometheus 설치 방법은 [다운로드 섹션](https://prometheus.io/download/) 도커 이미지에서도 다양하게 확인할 수 있습니다.

`docker run --name prometheus -d -p 127.0.0.1:9090:9090 prom/prometheus`

하지만 여기서는 Kubernetes에 배포하는 데 집중하겠습니다. 여기에도 몇 가지 옵션이 있습니다.

- 구성 YAML 파일 생성
- 오퍼레이터 사용(모든 Prometheus 구성 요소의 관리자)
- Helm 차트를 사용하여 오퍼레이터 배포

### Kubernetes에 배포하기

이 빠르고 간단한 설치를 위해 다시 로컬에서 Minikube 클러스터를 사용하겠습니다. Minikube와의 이전 접점과 마찬가지로, Helm을 사용하여 Prometheus Helm 차트를 배포할 것입니다.

`helm repo add prometheus-community https://prometheus-community.github.io/helm-charts`

![](/2022/Days/Images/Day78_Monitoring1.png)

위에서 볼 수 있듯이 Helm 리포지토리 업데이트도 실행했으므로, 이제 `helm install stable prometheus-community/prometheus` 명령을 사용하여 Minikube 환경에 Prometheus를 배포할 준비가 되었습니다.

![](/2022/Days/Images/Day78_Monitoring2.png)

몇 분 후, 몇 개의 새로운 pod가 나타나는데, 이 데모에서는 기본 네임스페이스에 배포했으며, 일반적으로는 이 pod를 해당 네임스페이스에 push합니다.

![](/2022/Days/Images/Day78_Monitoring3.png)

모든 pod가 실행되고 나면 Prometheus의 배포된 모든 측면도 살펴볼 수 있습니다.

![](/2022/Days/Images/Day78_Monitoring4.png)

이제 Prometheus Server UI에 액세스하기 위해 다음 명령어를 사용하여 포팅 포워드할 수 있습니다.

```Shell
export POD_NAME=$(kubectl get pods --namespace default -l "app=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace default port-forward $POD_NAME 9090
```

브라우저를 `http://localhost:9090`으로 처음 열면 다음과 같이 빈 화면이 표시됩니다.

![](/2022/Days/Images/Day78_Monitoring5.png)

Kubernetes 클러스터에 배포했기 때문에 Kubernetes API에서 자동으로 메트릭을 가져올 것이므로 일부 PromQL을 사용하여 최소한 `container_cpu_usage_seconds_total` 메트릭을 캡처하고 있는지 확인할 수 있습니다.

![](/2022/Days/Images/Day78_Monitoring6.png)

앞서 말씀드린 것처럼 메트릭을 확보하는 것도 좋지만 모니터링도 중요하지만, 모니터링 대상과 그 이유, 모니터링하지 않는 대상과 그 이유를 알아야 한다는 점에서 PromQL을 배우고 이를 실제로 적용하는 것은 매우 중요합니다!

다시 Prometheus로 돌아오고 싶지만, 지금은 로그 관리와 데이터 시각화에 대해 생각해 보고 나중에 다시 Prometheus로 돌아올 수 있도록 해야 할 것 같습니다.

## 자료

- [The Importance of Monitoring in DevOps](https://www.devopsonline.co.uk/the-importance-of-monitoring-in-devops/)
- [Understanding Continuous Monitoring in DevOps?](https://medium.com/devopscurry/understanding-continuous-monitoring-in-devops-f6695b004e3b)
- [DevOps Monitoring Tools](https://www.youtube.com/watch?v=Zu53QQuYqJ0)
- [Top 5 - DevOps Monitoring Tools](https://www.youtube.com/watch?v=4t71iv_9t_4)
- [How Prometheus Monitoring works](https://www.youtube.com/watch?v=h4Sl21AKiDg)
- [Introduction to Prometheus monitoring](https://www.youtube.com/watch?v=5o37CGlNLr8)
- [Promql cheat sheet with examples](https://www.containiq.com/post/promql-cheat-sheet-with-examples)

[Day 79](day79.md)에서 봐요!
