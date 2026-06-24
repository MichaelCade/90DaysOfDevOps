---
title: '#90DaysOfDevOps - Data Visualisation - Grafana - Day 83'
published: false
description: 90DaysOfDevOps - Data Visualisation - Grafana
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048767
---

## 데이터 시각화 - Grafana

이 섹션에서는 통합 observability과 관련하여 많은 Kibana를 살펴봤습니다. 하지만 시간을 내어 Grafana에 대해서도 다뤄야 합니다. 하지만 이 둘은 동일하지도 않고 서로 완전히 경쟁하는 것도 아닙니다.

Kibana의 핵심 기능은 데이터 쿼리 및 분석입니다. 사용자는 다양한 방법을 사용해 근본 원인 분석과 진단을 위해 데이터 내에서 특정 이벤트나 문자열에 대해 Elasticsearch에서 색인된 데이터를 검색할 수 있습니다. 이러한 쿼리를 기반으로 사용자는 차트, 표, 지리적 지도 및 기타 유형의 시각화를 사용하여 다양한 방식으로 데이터를 시각화할 수 있는 Kibana의 시각화 기능을 사용할 수 있습니다.

Grafana는 Kibana의 folk로 시작되었으며, 당시 Kibana가 제공하지 않았던 메트릭, 즉 모니터링에 대한 지원을 제공하는 것이 목표였습니다.

Grafana는 무료 오픈 소스 데이터 시각화 도구입니다. 현장에서는 일반적으로 Prometheus와 Grafana를 함께 사용하는 것을 볼 수 있지만, Elasticsearch 및 Graphite와 함께 사용하는 경우도 있습니다.

두 도구의 주요 차이점은 로깅과 모니터링입니다. 이 섹션에서는 Nagios로 모니터링을 다루기 시작해 Prometheus를 살펴본 다음 로깅으로 이동해 ELK 및 EFK 스택을 다루었습니다.

Grafana는 시스템 CPU, 메모리, 디스크 및 I/O 사용률과 같은 메트릭을 분석하고 시각화하는 데 적합합니다. 이 플랫폼은 전체 텍스트 데이터 쿼리를 허용하지 않습니다. Kibana는 Elasticsearch 위에서 실행되며 주로 로그 메시지 분석에 사용됩니다.

Kibana는 배포가 매우 쉬울 뿐만 아니라 배포 위치를 선택할 수 있다는 것을 이미 확인했듯이, Grafana도 마찬가지입니다.

둘 다 Linux, Mac, Windows, Docker에 설치하거나 소스에서 빌드하는 것을 지원합니다.

의심할 여지 없이 다른 도구도 있지만, Grafana는 가상, 클라우드 및 클라우드 네이티브 플랫폼에 걸쳐서 제가 본 도구이므로 이 섹션에서 다루고 싶었습니다.

### Prometheus Operator + Grafana Deployment

이 섹션에서 이미 Prometheus에 대해 다루었지만, 이 두 가지를 함께 사용하는 경우가 너무 많아서 최소한 시각화에서 어떤 메트릭을 표시할 수 있는지 확인할 수 있는 환경을 만들어 보고 싶었습니다. 환경을 모니터링하는 것이 중요하다는 것은 알고 있지만, Prometheus나 다른 메트릭 도구에서 이러한 메트릭만 살펴보는 것은 번거롭고 확장성이 떨어집니다. 바로 이 지점에서 Grafana가 등장하여 Prometheus 데이터베이스에 수집 및 저장된 메트릭의 대화형 시각화를 제공합니다.

이 시각화를 통해 환경에 맞는 사용자 정의 차트, 그래프 및 알림을 만들 수 있습니다. 이 연습에서는 Minikube 클러스터를 사용하겠습니다.

먼저 이것을 로컬 시스템에 복제하는 것으로 시작하겠습니다. `git clone https://github.com/prometheus-operator/kube-prometheus.git` 및 `cd kube-prometheus`를 입력하세요.

![](/2022/Days/Images/Day83_Monitoring1.png)

첫 번째 작업은 Minikube 클러스터 내에서 네임스페이스를 생성하는 것입니다. `kubectl create -f manifests/setup` 이전 섹션에서 따라하지 않은 경우 `minikube start`를 사용하여 여기에 새 클러스터를 불러올 수 있습니다.

![](/2022/Days/Images/Day83_Monitoring2.png)

다음으로, `kubectl create -f manifests/` 명령을 사용하여 데모에 필요한 모든 것을 배포할 것인데, 보시다시피 클러스터 내에 다양한 리소스가 배포될 것입니다.

![](/2022/Days/Images/Day83_Monitoring3.png)

이제 pod가 실행될 때까지 기다려야 하며, 실행 상태가 되면 `kubectl get pods -n monitoring -w` 명령어를 사용하여 pod를 계속 감시할 수 있습니다.

![](/2022/Days/Images/Day83_Monitoring4.png)

모든 것이 실행 중이면 `kubectl get pods -n monitoring` 명령으로 모든 pod가 실행 중이고 정상 상태인지 확인할 수 있습니다.

![](/2022/Days/Images/Day83_Monitoring5.png)

배포와 함께 데모 후반부에서 사용할 몇 가지 서비스를 배포했으며, `kubectl get svc -n monitoring` 명령으로 확인할 수 있습니다.

![](/2022/Days/Images/Day83_Monitoring6.png)

마지막으로 `kubectl get all -n monitoring` 명령으로 새 모니터링 네임스페이스에 배포된 모든 리소스를 확인해 봅시다.

![](/2022/Days/Images/Day83_Monitoring7.png)

새 터미널을 열면 이제 Grafana 도구에 액세스하여 몇 가지 메트릭을 수집하고 시각화할 준비가 되었으며, 사용할 명령은 `kubectl --namespace monitoring port-forward svc/grafana 3000`입니다.

![](/2022/Days/Images/Day83_Monitoring8.png)

브라우저를 열고 http://localhost:3000 로 이동하면 사용자 이름과 비밀번호를 입력하라는 메시지가 표시됩니다.

![](/2022/Days/Images/Day83_Monitoring9.png)
액세스하기 위한 기본 사용자 이름과 비밀번호는 다음과 같습니다.

```
Username: admin
Password: admin
```

하지만 처음 로그인할 때 새 비밀번호를 입력하라는 메시지가 표시됩니다. 초기 화면 또는 홈페이지에는 탐색할 수 있는 몇 가지 영역과 Grafana와 그 기능을 익히는 데 유용한 몇 가지 리소스가 표시됩니다. 나중에 사용하게 될 "Add your first data source" 및 "create your first dashboard" 위젯을 주목하세요.

![](/2022/Days/Images/Day83_Monitoring10.png)

Grafana 데이터 소스에 이미 prometheus 데이터 소스가 추가되어 있는 것을 볼 수 있지만, Minikube를 사용하고 있기 때문에 로컬 호스트에서 사용할 수 있도록 prometheus도 포팅해야 하므로 새 터미널을 열고 다음 명령을 실행할 수 있습니다. `kubectl --namespace monitoring port-forward svc/prometheus-k8s 9090` 이제 Grafana의 홈페이지에서 "Add your first data source"라는 위젯에 들어가서 여기에서 Prometheus를 선택합니다.

![](/2022/Days/Images/Day83_Monitoring11.png)

새 데이터 소스의 경우 http://localhost:9090 주소를 사용할 수 있으며, 아래에 강조 표시된 대로 드롭다운을 브라우저로 변경해야 합니다.

![](/2022/Days/Images/Day83_Monitoring12.png)

이제 페이지 하단에서 저장 및 테스트를 누르면 됩니다. 그러면 아래와 같이 Prometheus의 포트 포워딩이 작동하는 결과를 확인할 수 있습니다.

![](/2022/Days/Images/Day83_Monitoring13.png)

홈페이지로 돌아가서 "Create your first dashboard" 옵션을 찾아 "Add a new panel"를 선택합니다.

![](/2022/Days/Images/Day83_Monitoring14.png)

아래에서 이미 Grafana 데이터 원본에서 수집하고 있는 것을 볼 수 있지만, Prometheus 데이터 원본에서 메트릭을 수집하려면 데이터 원본 드롭다운을 선택하고 새로 만든 "Prometheus-1"을 선택합니다.

![](/2022/Days/Images/Day83_Monitoring15.png)

그런 다음 메트릭 브라우저를 선택하면 Minikube 클러스터와 관련된 Prometheus에서 수집되는 메트릭의 긴 목록이 표시됩니다.

![](/2022/Days/Images/Day83_Monitoring16.png)

데모에서는 시스템 리소스에 대한 몇 가지 출력을 제공하는 메트릭을 찾아보겠습니다. `cluster:node_cpu:ratio{}`는 클러스터의 노드에 대한 세부 정보를 제공하고 이 통합이 작동하고 있음을 증명합니다.

![](/2022/Days/Images/Day83_Monitoring17.png)

이 시각화가 마음에 들면 오른쪽 상단의 적용 버튼을 누르면 이 그래프를 대시보드에 추가할 수 있습니다. 계속해서 그래프와 다른 차트를 추가하여 필요한 시각화를 제공할 수 있습니다.

![](/2022/Days/Images/Day83_Monitoring18.png)

그러나 이전에 만든 수천 개의 대시보드를 활용할 수 있으므로 새로 만들 필요가 없습니다.

![](/2022/Days/Images/Day83_Monitoring19.png)

Kubernetes를 검색하면 선택할 수 있는 미리 빌드된 대시보드의 긴 목록을 볼 수 있습니다.

![](/2022/Days/Images/Day83_Monitoring20.png)

Kubernetes API 서버 대시보드를 선택하고 새로 추가된 Prometheus-1 데이터 소스에 맞게 데이터 소스를 변경하면 아래에 표시되는 몇 가지 메트릭을 볼 수 있습니다.

![](/2022/Days/Images/Day83_Monitoring21.png)

### 알림

또한 배포한 알림 관리자를 활용하여 슬랙이나 다른 통합으로 알림을 보낼 수도 있는데, 이렇게 하려면 아래 세부 정보를 사용하여 알림 관리자 서비스를 포트로 포워드해야 합니다.

`kubectl --namespace monitoring port-forward svc/alertmanager-main 9093`
`http://localhost:9093`

이것으로 통합 observability에 대한 모든 것에 대한 섹션을 마쳤습니다. 개인적으로 이 섹션은 이 주제가 얼마나 광범위한 주제인지 강조했지만, 마찬가지로 이것이 우리의 역할에 얼마나 중요한지, 특히 다른 섹션에서 이미 다룬 모든 자동화를 통해 매우 극적으로 변화할 수 있는 경우 메트릭, 로깅 또는 추적이든 앞으로 광범위한 환경에서 무슨 일이 일어나고 있는지 잘 파악해야 한다는 것을 알 수 있었습니다.

다음 섹션에서는 데이터 관리와 데이터 관리와 관련하여 데브옵스 원칙을 어떻게 고려해야 하는지에 대해 살펴보겠습니다.

## 자료

- [Understanding Logging: Containers & Microservices](https://www.youtube.com/watch?v=MMVdkzeQ848)
- [The Importance of Monitoring in DevOps](https://www.devopsonline.co.uk/the-importance-of-monitoring-in-devops/)
- [Understanding Continuous Monitoring in DevOps?](https://medium.com/devopscurry/understanding-continuous-monitoring-in-devops-f6695b004e3b)
- [DevOps Monitoring Tools](https://www.youtube.com/watch?v=Zu53QQuYqJ0)
- [Top 5 - DevOps Monitoring Tools](https://www.youtube.com/watch?v=4t71iv_9t_4)
- [How Prometheus Monitoring works](https://www.youtube.com/watch?v=h4Sl21AKiDg)
- [Introduction to Prometheus monitoring](https://www.youtube.com/watch?v=5o37CGlNLr8)
- [Promql cheat sheet with examples](https://www.containiq.com/post/promql-cheat-sheet-with-examples)
- [Log Management for DevOps | Manage application, server, and cloud logs with Site24x7](https://www.youtube.com/watch?v=J0csO_Shsj0)
- [Log Management what DevOps need to know](https://devops.com/log-management-what-devops-teams-need-to-know/)
- [What is ELK Stack?](https://www.youtube.com/watch?v=4X0WLg05ASw)
- [Fluentd simply explained](https://www.youtube.com/watch?v=5ofsNyHZwWE&t=14s)

[Day 84](day84.md)에서 봐요!
