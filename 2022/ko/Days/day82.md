---
title: '#90DaysOfDevOps - EFK Stack - Day 82'
published: false
description: 90DaysOfDevOps - EFK Stack
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049059
---

### EFK 스택

이전 섹션에서는 스택의 로그 수집기로 Logstash를 사용하는 ELK 스택에 대해 이야기했지만, EFK 스택에서는 이를 FluentD 또는 FluentBit으로 대체합니다.

이 섹션에서 우리의 임무는 EFK를 사용하여 Kubernetes 로그를 모니터링하는 것입니다.

### EFK 개요

우리는 다음을 Kubernetes 클러스터에 배포할 것입니다.

![](/2022/Days/Images/Day82_Monitoring1.png)

EFK 스택은 다음과 같은 3가지 소프트웨어가 함께 번들로 제공되는 모음입니다:

- Elasticsearch: NoSQL 데이터베이스는 데이터를 저장하는 데 사용되며 검색 및 로그 쿼리를 위한 인터페이스를 제공합니다.

- Fluentd: Fluentd는 통합 로깅 계층을 위한 오픈 소스 데이터 수집기입니다. Fluentd를 사용하면 데이터 수집과 소비를 통합하여 데이터를 더 잘 사용하고 이해할 수 있습니다.

- Kibana: 로그 관리 및 통계를 위한 인터페이스. elasticsearch에서 정보를 읽는 역할을 담당합니다.

### Minikube에 EFK 배포하기

우리는 신뢰할 수 있는 Minikube 클러스터를 사용해 EFK 스택을 배포할 것입니다. 우리 시스템에서 `minikube start`를 사용하여 클러스터를 시작하겠습니다. WSL2가 활성화된 Windows OS를 사용하고 있습니다.

![](/2022/Days/Images/Day82_Monitoring2.png)

EFK 스택을 클러스터에 배포하는 데 필요한 모든 것을 포함하는 [efk-stack.yaml](/2022/Days/Monitoring/EFK%20Stack/efk-stack.yaml)을 생성했으며, `kubectl create -f efk-stack.yaml` 명령을 사용하여 배포되는 모든 것을 확인할 수 있습니다.

![](/2022/Days/Images/Day82_Monitoring3.png)

시스템에 따라 다르지만, 이미 이 작업을 실행하여 이미지를 가져온 경우 다음 명령으로 진행 상황을 확인할 수 있습니다. `kubectl get pods -n kube-logging -w` 몇 분 정도 걸릴 수 있습니다.

![](/2022/Days/Images/Day82_Monitoring4.png)

위의 명령으로 상황을 계속 주시할 수 있지만, 모든 pod가 이제 실행 중인지 확인하기 위해 다음 `kubectl get pods -n kube-logging` 명령을 실행하여 모든 것이 정상임을 명확히 하고 싶습니다.

![](/2022/Days/Images/Day82_Monitoring5.png)

모든 pod가 실행되고 실행 중이고 이 단계에서는 다음을 볼 수 있어야 합니다.

- ElasticSearch와 연결된 3개의 pod
- Fluentd와 연결된 pod 1개
- Kibana와 연결된 pod 1개

또한 `kubectl get all -n kube-logging`을 사용하여 네임스페이스에 모두 표시할 수 있으며, 앞서 설명한 대로 fluentd는 데몬셋으로 배포되고, kibana는 deployment로, Elasticsearch는 statefulset으로 배포됩니다.

![](/2022/Days/Images/Day82_Monitoring6.png)

이제 모든 pod가 실행되고 실행 중이므로 이제 새 터미널에서 포트 포워드 명령을 실행하여 kibana 대시보드에 액세스할 수 있습니다. pod 이름은 여기에 표시된 명령과 다를 것입니다. `kubectl port-forward kibana-84cf7f59c-v2l8v 5601:5601 -n kube-logging`

![](/2022/Days/Images/Day82_Monitoring7.png)

이제 브라우저를 열고 `http://localhost:5601` 주소로 이동하면 아래와 같은 화면이 표시되거나 실제로 샘플 데이터 화면이 표시되거나 계속 진행하여 직접 구성할 수 있습니다. 어느 쪽이든 이전 세션에서 ELK 스택을 살펴볼 때 다루었던 테스트 데이터를 꼭 살펴보시기 바랍니다.

![](/2022/Days/Images/Day82_Monitoring8.png)

다음으로, 왼쪽 메뉴의 "discover" 탭을 누르고 인덱스 패턴에 "\*"를 추가해야 합니다. "Next step"를 눌러 다음 단계로 계속 진행합니다.

![](/2022/Days/Images/Day82_Monitoring9.png)

2단계 2번에서는 시간별로 데이터를 필터링하기 위해 드롭다운에서 @timestamp 옵션을 사용하겠습니다. 패턴 만들기를 누르면 완료하는 데 몇 초 정도 걸릴 수 있습니다.

![](/2022/Days/Images/Day82_Monitoring10.png)

이제 몇 초 후에 "discover" 탭으로 돌아가면 Kubernetes 클러스터에서 데이터가 들어오는 것을 볼 수 있을 것입니다.

![](/2022/Days/Images/Day82_Monitoring11.png)

이제 EFK 스택이 실행 중이고 Fluentd를 통해 Kubernetes 클러스터에서 로그를 수집하고 있으므로 왼쪽 상단의 Kibana 로고를 눌러 홈 화면으로 이동하면 처음 로그인할 때 보았던 것과 동일한 페이지가 표시됩니다.

다른 플러그인이나 소스로부터 APM, 로그 데이터, 메트릭 데이터, 보안 이벤트를 추가할 수 있습니다.

![](/2022/Days/Images/Day82_Monitoring12.png)

"Add log data"를 선택하면 아래에서 로그를 가져올 위치에 대한 많은 선택 사항이 있음을 알 수 있으며, ELK 스택의 일부인 Logstash가 언급되어 있는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day82_Monitoring13.png)

메트릭 데이터 아래에서 Prometheus 및 기타 여러 서비스에 대한 소스를 추가할 수 있음을 알 수 있습니다.

### APM(Application Performance Monitoring)

애플리케이션 내부에서 심층적인 성능 메트릭과 오류를 수집하는 APM(Application Performance Monitoring)을 수집하는 옵션도 있습니다. 이를 통해 수천 개의 애플리케이션의 성능을 실시간으로 모니터링할 수 있습니다.

여기서는 APM에 대해 자세히 설명하지 않겠지만, [Elastic 사이트](https://www.elastic.co/observability/application-performance-monitoring)에서 자세한 내용을 확인할 수 있습니다.

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

[Day 83](day83.md)에서 봐요!
