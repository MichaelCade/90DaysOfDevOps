---
title: '#90DaysOfDevOps - ELK Stack - Day 80'
published: false
description: 90DaysOfDevOps - ELK Stack
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048746
---

## ELK 스택

이 세션에서는 앞서 언급한 몇 가지 옵션에 대해 조금 더 실습해 보겠습니다.

ELK Stack은 세 가지 개별 도구의 조합입니다:

- [Elasticsearch](https://www.elastic.co/what-is/elasticsearch)는 텍스트, 숫자, 위치 기반 정보, 정형, 비정형 등 모든 유형의 데이터를 위한 분산형 무료 개방형 검색 및 분석 엔진입니다.

- [Logstash](https://www.elastic.co/logstash/)는 다양한 소스에서 데이터를 수집하고 변환한 다음 원하는 "stash(저장소)"로 전송하는 무료 개방형 서버 측 데이터 처리 파이프라인입니다.

- [Kibana](https://www.elastic.co/kibana/)는 무료 개방형 사용자 인터페이스로, Elasticsearch 데이터를 시각화하고 Elastic Stack을 탐색할 수 있게 해줍니다. 쿼리 로드 추적부터 앱에서 요청이 흐르는 방식 이해까지 무엇이든 할 수 있습니다.

ELK 스택을 사용하면 모든 소스에서 모든 형식의 데이터를 안정적이고 안전하게 가져온 다음 실시간으로 검색, 분석 및 시각화할 수 있습니다.

위에서 언급한 구성 요소 외에도 스택으로 전달하기 위해 다양한 유형의 데이터를 수집하기 위해 엣지 호스트에 설치되는 경량 에이전트인 Beats도 볼 수 있습니다.

- Log: 분석해야 하는 서버 로그가 식별됩니다.

- Logstash: 로그와 이벤트 데이터를 수집합니다. 심지어 데이터를 구문 분석하고 변환합니다.

- ElasticSearch: Logstash에서 변환된 데이터는 저장, 검색 및 색인됩니다.

- Kibana는 Elasticsearch DB를 사용해 탐색, 시각화, 공유합니다.

![](/2022/Days/Images/Day80_Monitoring8.png)

[Guru99에서 가져온 사진](https://www.guru99.com/elk-stack-tutorial.html)

이를 설명하는 좋은 리소스는 [ELK 스택에 대한 완전한 가이드](https://logz.io/learn/complete-guide-elk-stack/)입니다.

Beats가 추가되면서 ELK Stack은 이제 Elastic Stack이라고도 불립니다.

실습 시나리오에서는 Elastic Stack을 배포할 수 있는 많은 위치가 있지만, 여기서는 시스템에 로컬로 배포하기 위해 docker-compose를 사용하겠습니다.

[Docker Compose로 Elastic Stack 시작하기](https://www.elastic.co/guide/en/elastic-stack-get-started/current/get-started-stack-docker.html#get-started-docker-tls)

![](/2022/Days/Images/Day80_Monitoring1.png)

여기에서 제가 사용한 원본 파일과 연습을 찾을 수 있습니다. [deviantony/docker-elk](https://github.com/deviantony/docker-elk)

이제 `docker-compose up -d`를 실행할 수 있는데, 처음 실행할 때는 이미지를 가져와야 합니다.

![](/2022/Days/Images/Day80_Monitoring2.png)

이 리포지토리 또는 제가 사용한 리포지토리를 팔로우하면 "changeme"의 비밀번호를 갖게 되거나 제 리포지토리에서 "90DaysOfDevOps"의 비밀번호를 갖게 됩니다. 사용자 이름은 "elastic"입니다.

몇 분 후, Kibana server / Docker 컨테이너인 `http://localhost:5601/`로 이동할 수 있습니다.

![](/2022/Days/Images/Day80_Monitoring3.png)

초기 홈 화면은 다음과 같이 표시될 것입니다.

![](/2022/Days/Images/Day80_Monitoring4.png)

"Get started by adding integrations"라는 제목의 섹션 아래에 "Try sample data"가 있으며, 이를 클릭하면 아래 표시된 것 중 하나를 추가할 수 있습니다.

![](/2022/Days/Images/Day80_Monitoring5.png)

저는 "Sample web logs"를 선택하려고 하는데, 이것은 실제로 ELK 스택에 어떤 데이터 세트를 가져올 수 있는지 살펴보기 위한 것입니다.

"Add Data"를 선택하면 일부 데이터를 채우는 데 시간이 걸리고 "View Data" 옵션과 드롭다운에 해당 데이터를 볼 수 있는 사용 가능한 방법의 목록이 표시됩니다.

![](/2022/Days/Images/Day80_Monitoring6.png)

대시보드 보기에 명시된 대로:

**샘플 로그 데이터**

> 이 대시보드에는 사용해 볼 수 있는 샘플 데이터가 포함되어 있습니다. 이 데이터를 보고, 검색하고, 시각화와 상호 작용할 수 있습니다. Kibana에 대한 자세한 내용은 문서를 참조하세요.

![](/2022/Days/Images/Day80_Monitoring7.png)

이것은 Kibana를 사용하여 Logstash를 통해 ElasticSearch에 추가된 데이터를 시각화하는 것입니다. 이것이 유일한 옵션은 아니지만, 저는 이것을 배포하고 살펴보고 싶었습니다.

언젠가는 Grafana에 대해서도 다룰 예정이며, 이 둘 사이의 데이터 시각화 유사점을 보게 될 것이고, Prometheus도 보셨을 것입니다.

제가 Elastic Stack과 Prometheus + Grafana를 비교하면서 느낀 핵심은 Elastic Stack 또는 ELK Stack은 로그에 초점을 맞추고 있고 Prometheus는 메트릭에 초점을 맞추고 있다는 것입니다.

저는 서로 다른 제품에 대해 더 잘 이해하기 위해 MetricFire의 이 기사 [Prometheus vs. ELK](https://www.metricfire.com/blog/prometheus-vs-elk/)를 읽었습니다.

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

[Day 81](day81.md)에서 봐요!
