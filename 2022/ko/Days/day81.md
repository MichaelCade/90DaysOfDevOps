---
title: '#90DaysOfDevOps - Fluentd & FluentBit - Day 81'
published: false
description: 90DaysOfDevOps - Fluentd & FluentBit
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048716
---

## Fluentd & FluentBit

이 통합 observability 섹션의 일부로 살펴보고 싶었던 또 다른 데이터 수집기는 [Fluentd](https://docs.fluentd.org/)였습니다. 오픈소스 통합 로깅 레이어입니다.

Fluentd는 깨끗하고 안정적인 로깅 파이프라인을 구축하는 데 적합한 네 가지 주요 기능을 갖추고 있습니다:

JSON을 사용한 통합 로깅: Fluentd는 가능한 한 데이터를 JSON으로 구조화하려고 노력합니다. 이를 통해 여러 소스와 대상에 걸쳐 로그를 수집, 필터링, 버퍼링, 출력하는 등 로그 데이터 처리의 모든 측면을 통합할 수 있습니다. 엄격한 스키마를 강요하지 않고도 액세스할 수 있는 충분한 구조를 가지고 있기 때문에 JSON을 사용하면 다운스트림 데이터 처리가 훨씬 쉬워집니다.

플러그 가능한 아키텍처: Fluentd는 커뮤니티가 기능을 확장할 수 있는 유연한 플러그인 시스템을 갖추고 있습니다. 300개 이상의 커뮤니티 기여 플러그인이 수십 개의 데이터 소스를 수십 개의 데이터 출력에 연결하여 필요에 따라 데이터를 조작합니다. 플러그인을 사용하면 로그를 즉시 더 효과적으로 활용할 수 있습니다.

최소한의 리소스 필요: 데이터 수집기는 바쁜 컴퓨터에서도 편안하게 실행될 수 있도록 가벼워야 합니다. Fluentd는 C와 Ruby의 조합으로 작성되어 최소한의 시스템 리소스를 필요로 합니다. 바닐라 인스턴스는 30~40MB의 메모리에서 실행되며 초당 13,000개의 이벤트를 처리할 수 있습니다.

내장된 안정성: 데이터 손실은 절대 일어나지 않습니다. Fluentd는 메모리 및 파일 기반 버퍼링을 지원하여 노드 간 데이터 손실을 방지합니다. 또한, 강력한 페일오버를 지원하며 고가용성을 위해 설정할 수 있습니다.

[Fluentd 설치하기](https://docs.fluentd.org/quickstart#step-1-installing-fluentd)

### 앱은 데이터를 어떻게 기록하나요?

- 파일에 기록합니다. `.log` 파일(도구 없이는 대규모로 분석하기 어려움)
- 데이터베이스에 직접 기록(각 애플리케이션이 올바른 형식으로 구성되어 있어야 함)
- 타사 애플리케이션(NodeJS, NGINX, PostgreSQL)

이것이 바로 통합 로깅 레이어가 필요한 이유입니다.

FluentD는 위에 표시된 3가지 로깅 데이터 유형을 허용하고 이를 수집, 처리하여 대상(예: Elastic, MongoDB 또는 Kafka 데이터베이스)으로 로그를 전송할 수 있는 기능을 제공합니다.

모든 데이터, 모든 데이터 소스를 FluentD로 전송할 수 있으며, 이를 모든 대상으로 전송할 수 있습니다. FluentD는 특정 소스나 대상에 종속되지 않습니다.

Fluentd를 조사하는 과정에서 또 다른 옵션으로 Fluent bit를 계속 발견하게 되었는데, 서버뿐만 아니라 컨테이너에도 배포할 수 있지만 로깅 도구를 Kubernetes 환경에 배포하려는 경우 FluentBit가 해당 기능을 제공할 수 있을 것 같았습니다.

[Fluentd & FluentBit](https://docs.fluentbit.io/manual/about/fluentd-and-fluent-bit)

Fluentd와 FluentBit는 입력 플러그인을 사용하여 데이터를 FluentBit 형식으로 변환한 다음, 출력 플러그인을 사용하여 엘라스틱서치와 같은 출력 대상이 무엇이든 출력할 수 있습니다.

또한 구성 간에 태그와 일치를 사용할 수도 있습니다.

일부 아키텍처에서는 함께 사용할 수 있지만, FluentBit를 사용해야 할 좋은 이유를 찾을 수 없으며, FluentBit가 시작하기에 가장 좋은 방법인 것 같습니다.

### Kubernetes의 FluentBit

Kubernetes의 FluentBit는 데몬셋으로 배포되며, 이는 클러스터의 각 노드에서 실행된다는 것을 의미합니다. 그러면 각 노드의 각 Fluent Bit pod는 해당 노드의 각 컨테이너를 읽고 사용 가능한 모든 로그를 수집합니다. 또한 Kubernetes API 서버에서 메타데이터를 수집합니다.

Kubernetes 어노테이션은 애플리케이션의 구성 YAML 내에서 사용할 수 있습니다.

우선, Fluent Helm 리포지토리에서 배포할 수 있습니다. Helm 리포지토리에서 `helm repo add fluent https://fluent.github.io/helm-charts`를 실행한 다음 `helm install fluent-bit fluent/fluent-bit` 명령어를 사용하여 설치합니다.

![](/2022/Days/Images/Day81_Monitoring1.png)

내 클러스터에서는 (테스트 목적으로) 기본 네임스페이스에서 Prometheus를 실행 중이므로 fluent-bit pod가 실행 중인지 확인해야 합니다. `kubectl get all | grep fluent`를 사용하면 앞서 언급한 실행 중인 pod, 서비스 및 데몬셋을 확인할 수 있습니다.

![](/2022/Days/Images/Day81_Monitoring2.png)

FluentBit이 로그를 어디에서 가져올지 알 수 있도록 구성 파일이 있으며, 이 FluentBit의 Kubernetes 배포에는 구성 파일과 유사한 configmap이 있습니다.

![](/2022/Days/Images/Day81_Monitoring3.png)

이 configmap은 다음과 같이 보일 것입니다:

```
Name:         fluent-bit
Namespace:    default
Labels:       app.kubernetes.io/instance=fluent-bit
              app.kubernetes.io/managed-by=Helm
              app.kubernetes.io/name=fluent-bit
              app.kubernetes.io/version=1.8.14
              helm.sh/chart=fluent-bit-0.19.21
Annotations:  meta.helm.sh/release-name: fluent-bit
              meta.helm.sh/release-namespace: default

Data
====
custom_parsers.conf:
----
[PARSER]
    Name docker_no_time
    Format json
    Time_Keep Off
    Time_Key time
    Time_Format %Y-%m-%dT%H:%M:%S.%L

fluent-bit.conf:
----
[SERVICE]
    Daemon Off
    Flush 1
    Log_Level info
    Parsers_File parsers.conf
    Parsers_File custom_parsers.conf
    HTTP_Server On
    HTTP_Listen 0.0.0.0
    HTTP_Port 2020
    Health_Check On

[INPUT]
    Name tail
    Path /var/log/containers/*.log
    multiline.parser docker, cri
    Tag kube.*
    Mem_Buf_Limit 5MB
    Skip_Long_Lines On

[INPUT]
    Name systemd
    Tag host.*
    Systemd_Filter _SYSTEMD_UNIT=kubelet.service
    Read_From_Tail On

[FILTER]
    Name Kubernetes
    Match kube.*
    Merge_Log On
    Keep_Log Off
    K8S-Logging.Parser On
    K8S-Logging.Exclude On

[OUTPUT]
    Name es
    Match kube.*
    Host elasticsearch-master
    Logstash_Format On
    Retry_Limit False

[OUTPUT]
    Name es
    Match host.*
    Host elasticsearch-master
    Logstash_Format On
    Logstash_Prefix node
    Retry_Limit False

Events:  <none>
```

이제 pod를 로컬호스트로 포트 포워딩하여 연결성을 확보할 수 있습니다. 먼저 `kubectl get pods | grep fluent`로 pod의 이름을 가져온 다음 `kubectl port-forward fluent-bit-8kvl4 2020:2020`을 사용하여 `http://localhost:2020/`으로 웹 브라우저를 엽니다.

![](/2022/Days/Images/Day81_Monitoring4.png)

[FluentBit](https://medium.com/kubernetes-tutorials/exporting-kubernetes-logs-to-elasticsearch-using-fluent-bit-758e8de606af)에 대한 자세한 내용을 다루는 이 훌륭한 매체 기사도 발견했습니다.

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
- [Fluent Bit explained | Fluent Bit vs Fluentd](https://www.youtube.com/watch?v=B2IS-XS-cc0)

[Day 82](day82.md)에서 봐요!
